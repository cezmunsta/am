/*
Copyright (c) 2022 Ceri Williams. All rights reserved.

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>
*/

package cmd

import (
	"flag"
	"fmt"
)

// HelpCommand provides the interface to the help information
type HelpCommand struct {
	fs *flag.FlagSet

	cmdList []CommandRunner
	name    string
}

func (h *HelpCommand) Init(args []string) error {
	return h.fs.Parse(args)
}

func (h *HelpCommand) Name() string {
	return h.fs.Name()
}

func (h *HelpCommand) PrintDefaults() {
	h.fs.PrintDefaults()
}

func (h *HelpCommand) Run() error {
	fmt.Print("Usage of am")

	sCmd := "all"
	msg := ":"

	if h.fs.NArg() > 0 {
		sCmd = h.fs.Arg(0)
		msg = fmt.Sprintf(" %s%s", sCmd, msg)
	}
	fmt.Printf("%s\n", msg)

	for _, cmd := range h.cmdList {
		if sCmd == "all" {
			fmt.Println(cmd.Name())
		}
		if sCmd == "all" || sCmd == cmd.Name() {
			cmd.PrintDefaults()
			fmt.Println()
		}
	}

	return nil
}

// NewHelpCommand creates the HelpCommand FlagSet
func NewHelpCommand() *HelpCommand {
	h := &HelpCommand{
		fs: flag.NewFlagSet("am", flag.ExitOnError),
	}

	return h
}
