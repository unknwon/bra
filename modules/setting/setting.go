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

package setting

import (
	"os"
	"path"
	"strings"

	"github.com/BurntSushi/toml"
	"github.com/Unknwon/com"

	"github.com/Unknwon/bra/modules/log"
)

var (
	WorkDir string
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
}

// UnpackPath replaces special path variables and returns full path.
func UnpackPath(path string) string {
	path = strings.Replace(path, "$WORKDIR", WorkDir, 1)
	return path
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
}
