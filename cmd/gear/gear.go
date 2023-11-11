// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"goki.dev/gear/gear"
	"goki.dev/gi/v2/gi"
	"goki.dev/gi/v2/gimain"
	"goki.dev/grease"
)

type config struct { //gti:add
	// Command is the command to run gear on
	Command string `posarg:"0"`
}

func main() {
	opts := grease.DefaultOptions("gear", "Gear",
		"Gear provides the generation of GUIs and interactive CLIs for any existing command line tools.")
	grease.Run(opts, &config{}, &grease.Cmd[*config]{
		Func: run,
		Root: true,
	})
}

func run(c *config) error {
	var err error
	gimain.Run(func() {
		err = app(c)
	})
	return err
}

func app(c *config) error {
	sc := gi.NewScene("gear").SetTitle("Gear")
	cmd := gear.NewCmd(c.Command)
	err := cmd.Parse()
	if err != nil {
		return err
	}
	app := gear.NewApp(sc).SetCmd(cmd)
	gi.DefaultTopAppBar = app.TopAppBar
	gi.NewWindow(sc).Run().Wait()
	return nil
}