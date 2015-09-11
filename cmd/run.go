// Copyright 2014 Unknwon
//
// Licensed under the Apache License, Version 2.0 (the "License"): you may
// not use this file except in compliance with the License. You may obtain
// a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
// WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the
// License for the specific language governing permissions and limitations
// under the License.

package cmd

import (
	"fmt"
	"os"
	"os/exec"
	"os/signal"
	"path"
	"runtime"
	"strings"
	"sync"
	"syscall"
	"time"

	"github.com/Unknwon/com"
	"github.com/Unknwon/log"
	"github.com/codegangsta/cli"
	"gopkg.in/fsnotify.v1"

	"github.com/Unknwon/bra/modules/setting"
)

var (
	lastBuild time.Time
	eventTime = make(map[string]int64)

	runningCmd  *exec.Cmd
	runningLock = &sync.Mutex{}
	shutdown    = make(chan bool)
)

var CmdRun = cli.Command{
	Name:   "run",
	Usage:  "start monitoring and notifying",
	Action: runRun,
	Flags:  []cli.Flag{},
}

// isTmpFile returns true if the event was for temporary files.
func isTmpFile(name string) bool {
	if strings.HasSuffix(strings.ToLower(name), ".tmp") {
		return true
	}
	return false
}

// hasWatchExt returns true if the file name has watched extension.
func hasWatchExt(name string) bool {
	for _, ext := range setting.Cfg.Run.WatchExts {
		if strings.HasSuffix(name, ext) {
			return true
		}
	}
	return false
}

func notify(cmds [][]string) {
	runningLock.Lock()
	defer func() {
		runningCmd = nil
		runningLock.Unlock()
	}()

	for _, cmd := range cmds {
		command := exec.Command(cmd[0], cmd[1:]...)
		command.Stdout = os.Stdout
		command.Stderr = os.Stderr
		if err := command.Start(); err != nil {
			log.Error("Fail to start command %v - %v", cmd, err)
			fmt.Print("\x07")
			return
		}

		log.Debug("Running %v", cmd)
		runningCmd = command
		done := make(chan error)
		go func() {
			done <- command.Wait()
		}()

		isShutdown := false
		select {
		case err := <-done:
			if isShutdown {
				return
			} else if err != nil {
				log.Warn("Fail to execute command %v - %v", cmd, err)
				fmt.Print("\x07")
				return
			}
		case <-shutdown:
			isShutdown = true
			gracefulKill()
			return
		}
	}
	log.Info("Notify operations are done!")
}

func gracefulKill() {
	// Given process a chance to exit itself.
	runningCmd.Process.Signal(os.Interrupt)

	// Wait for timeout, and force kill after that.
	for i := 0; i < setting.Cfg.Run.InterruptTimeout; i++ {
		time.Sleep(1 * time.Second)

		if runningCmd.ProcessState == nil || runningCmd.ProcessState.Exited() {
			return
		}
	}

	log.Info("Fail to graceful kill, force killing...")
	runningCmd.Process.Kill()
}

func catchSignals() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGTERM)
	<-sigs

	if runningCmd != nil {
		shutdown <- true
	}
	os.Exit(0)
}

func runRun(ctx *cli.Context) {
	setup(ctx)

	go catchSignals()
	go notify(setting.Cfg.Run.InitCmds)

	watchPathes := append([]string{setting.WorkDir}, setting.Cfg.Run.WatchDirs...)
	if setting.Cfg.Run.WatchAll {
		subdirs := make([]string, 0, 10)
		for _, dir := range watchPathes[1:] {
			dirs, err := com.GetAllSubDirs(setting.UnpackPath(dir))
			if err != nil {
				log.Fatal("Fail to get sub-directories: %v", err)
			}
			for i := range dirs {
				if !setting.IgnoreDir(dirs[i]) {
					subdirs = append(subdirs, path.Join(dir, dirs[i]))
				}
			}
		}
		watchPathes = append(watchPathes, subdirs...)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("Fail to create new watcher: %v", err)
	}
	defer watcher.Close()

	go func() {
		for {
			select {
			case e := <-watcher.Events:
				needsNotify := true

				if isTmpFile(e.Name) || !hasWatchExt(e.Name) || setting.IgnoreFile(e.Name) {
					continue
				}

				// Prevent duplicated builds.
				if lastBuild.Add(time.Duration(setting.Cfg.Run.BuildDelay) * time.Millisecond).
					After(time.Now()) {
					continue
				}
				lastBuild = time.Now()

				showName := e.String()
				if !log.NonColor {
					showName = strings.Replace(showName, setting.WorkDir, "\033[47;30m$WORKDIR\033[0m", 1)
				}

				if e.Op&fsnotify.Remove != fsnotify.Remove {
					mt, err := com.FileMTime(e.Name)
					if err != nil {
						log.Error("Fail to get file modify time: %v", err)
						continue
					}
					if eventTime[e.Name] == mt {
						log.Debug("Skipped %s", showName)
						needsNotify = false
					}
					eventTime[e.Name] = mt
				}

				if needsNotify {
					log.Info(showName)
					if runningCmd != nil && runningCmd.Process != nil {
						if runningCmd.Args[0] == "sudo" && runtime.GOOS == "linux" {
							// 给父进程发送一个TERM信号，试图杀死它和它的子进程
							rootCmd := exec.Command("sudo", "kill", "-TERM", com.ToStr(runningCmd.Process.Pid))
							rootCmd.Stdout = os.Stdout
							rootCmd.Stderr = os.Stderr
							if err := rootCmd.Run(); err != nil {
								log.Error("Fail to start rootCmd %s", err.Error())
								fmt.Print("\x07")
							}
						} else {
							shutdown <- true
						}
					}
					go notify(setting.Cfg.Run.Cmds)
				}
			}
		}
	}()

	log.Info("Following directories are monitored:")
	for i, p := range watchPathes {
		if err = watcher.Add(setting.UnpackPath(p)); err != nil {
			log.Fatal("Fail to watch diretory(%s): %v", p, err)
		}
		if i > 0 && !log.NonColor {
			p = strings.Replace(p, setting.WorkDir, "\033[47;30m$WORKDIR\033[0m", 1)
			p = strings.Replace(p, "$WORKDIR", "\033[47;30m$WORKDIR\033[0m", 1)
		}
		fmt.Printf("-> %s\n", p)
	}
	select {}
}
