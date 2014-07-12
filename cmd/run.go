// Copyright 2014 Unknown
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
	"path"
	"strings"
	"time"

	"github.com/Unknwon/com"
	"github.com/codegangsta/cli"
	"github.com/howeyc/fsnotify"

	"github.com/Unknwon/bra/modules/log"
	"github.com/Unknwon/bra/modules/setting"
)

var (
	lastBuild time.Time
	eventTime = make(map[string]int64)
)

var CmdRun = cli.Command{
	Name:   "run",
	Usage:  "start monitoring and notifying",
	Action: runRun,
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

func notify() {
	for _, cmd := range setting.Cfg.Run.Cmds {
		stdout, stderr, err := com.ExecCmd(cmd[0], cmd[1:]...)
		if err != nil {
			log.Error("Fail to execute %v", cmd)
			fmt.Println(stderr)
			return
		}
		if len(stdout) > 0 {
			fmt.Println(stdout)
		}
	}
	log.Info("Notify operations are done!")
}

func runRun(ctx *cli.Context) {
	watchPathes := append([]string{setting.WorkDir}, setting.Cfg.Run.WatchDirs...)
	if setting.Cfg.Run.WatchAll {
		subdirs := make([]string, 0, 10)
		for _, dir := range watchPathes {
			dirs, err := com.GetAllSubDirs(setting.UnpackPath(dir))
			if err != nil {
				log.Fatal("Fail to get sub-directories: %v", err)
			}
			for i := range dirs {
				dirs[i] = path.Join(dir, dirs[i])
			}
			subdirs = append(subdirs, dirs...)
		}
		watchPathes = append(watchPathes, subdirs...)
	}

	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal("Fail to create new watcher: %v", err)
	}

	log.Info("Following directories are monitored:")
	for i, p := range watchPathes {
		if err = watcher.Watch(setting.UnpackPath(p)); err != nil {
			log.Fatal("Fail to watch diretory(%s): %v", p, err)
		}
		if i > 0 {
			p = strings.Replace(p, setting.WorkDir, "\033[47;30m$WORKDIR\033[0m", 1)
			p = strings.Replace(p, "$WORKDIR", "\033[47;30m$WORKDIR\033[0m", 1)
		}
		fmt.Printf("-> %s\n", p)
	}

	go func() {
		for {
			select {
			case e := <-watcher.Event:
				needsNotify := true

				if isTmpFile(e.Name) || !hasWatchExt(e.Name) {
					continue
				}

				// Prevent duplicated builds.
				if lastBuild.Add(time.Duration(setting.Cfg.Run.BuildDelay) * time.Millisecond).
					After(time.Now()) {
					continue
				}
				lastBuild = time.Now()

				showName := strings.Replace(e.String(), setting.WorkDir, "\033[47;30m$WORKDIR\033[0m", 1)

				if !e.IsDelete() {
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
					notify()
				}
			}
		}
	}()
	select {}
}
