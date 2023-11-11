// Copyright (c) 2023, The GoKi Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package gear provides the generation of GUIs and interactive CLIs for any existing command line tools.
package gear

import (
	"goki.dev/glop/sentencecase"
)

// Cmd contains all of the data for a parsed command line command.
type Cmd struct {
	// Cmd is the actual name of the command (eg: "git", "go build")
	Cmd string
	// Name is the formatted name of the command (eg: "Git", "Go build")
	Name string
	// Doc is the documentation for the command (eg: "compile packages and dependencies")
	Doc string
	// Flags contains the flags for the command
	Flags []string
	// Cmds contains the subcommands of the command
	Cmds []*Cmd
}

// NewCmd makes a new [App] object from the given command name.
// It does not parse it; see [App.Parse].
func NewCmd(cmd string) *Cmd {
	return &Cmd{
		Cmd:  cmd,
		Name: sentencecase.Of(cmd),
	}
}