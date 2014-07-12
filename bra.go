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

// Bra(Brilliant Ridiculous Assistant) is a command line utility tool for Unknown.
package main

import (
	"os"
	"runtime"

	"github.com/codegangsta/cli"

	"github.com/Unknwon/bra/cmd"
	"github.com/Unknwon/bra/modules/log"
)

const APP_VER = "0.0.1.0711"

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	app := cli.NewApp()
	app.Name = "bra"
	app.Usage = "Brilliant Ridiculous Assistant"
	app.Version = APP_VER
	app.Commands = []cli.Command{
		cmd.CmdRun,
	}
	app.Flags = append(app.Flags, []cli.Flag{}...)
	log.Info("App Version: %s", APP_VER)
	app.Run(os.Args)
}
