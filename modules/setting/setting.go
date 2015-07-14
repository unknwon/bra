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

package setting

import (
	"bufio"
	"os"
	"path"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/Unknwon/com"
	"github.com/Unknwon/log"
)

var (
	WorkDir    string
	ignoreDirs = []string{".git"}
)

var Cfg struct {
	Run struct {
		InitCmds   [][]string `toml:"init_cmds"`
		WatchAll   bool       `toml:"watch_all"`
		WatchDirs  []string   `toml:"watch_dirs"`
		WatchExts  []string   `toml:"watch_exts"`
		BuildDelay int        `toml:"build_delay"`
		Cmds       [][]string `toml:"cmds"`
	} `toml:"run"`
	Sync struct {
		ListenAddr string `toml:"listen_addr"`
		RemoteAddr string `toml:"remote_addr"`
	} `toml:"sync"`
}

// UnpackPath replaces special path variables and returns full path.
func UnpackPath(path string) string {
	path = strings.Replace(path, "$WORKDIR", WorkDir, 1)
	path = strings.Replace(path, "$GOPATH", com.GetGOPATHs()[0], 1)
	return path
}

// IgnoreDir determines whether specified dir must be ignored.
func IgnoreDir(dir string) bool {
	for _, s := range ignoreDirs {
		if strings.Contains(dir, s) {
			return true
		}
	}
	return false
}

func InitSetting() {
	var err error
	WorkDir, err = os.Getwd()
	if err != nil {
		log.Fatal("Fail to get work directory: %v", err)
	}

	confPath := path.Join(WorkDir, ".bra.toml")
	if !com.IsFile(confPath) {
		log.Fatal(".bra.toml not found in work directory")
	} else if _, err = toml.DecodeFile(confPath, &Cfg); err != nil {
		log.Fatal("Fail to decode .bra.toml: %v", err)
	}

	// init list of ignored dirs
	file, err := os.Open(path.Join(WorkDir, ".braignore"))
	if err == nil {
		defer file.Close()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := strings.TrimSpace(scanner.Text())
			if len(line) > 0 && !strings.HasPrefix(line, "#") {
				ignoreDirs = append(ignoreDirs, line)
			}
		}
	}
}
