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
	"errors"
	"fmt"
	"os"
)

const (
	errNoArgs = "no arguments were specified"
	errUnkSub = "unknown subcommand: %s"
)

// CommandRunner accepts a set of commands that form the CLI interface
type CommandRunner interface {
	Init([]string) error
	PrintDefaults()
	Run() error
	Name() string
}

func needsHelp(s string) bool {
	switch s {
	case "help", "-h", "--help", "-help":
		return true
	}

	return false
}

func getCmdArgs() []string {
	return os.Args[2:]
}

func Execute(args []string) error {
	if len(args) == 0 {
		return errors.New(errNoArgs)
	}

	help := NewHelpCommand()
	commands := []CommandRunner{
		NewInvoiceCommand(),
		NewServerCommand(),
	}

	help.cmdList = commands
	commands = append(commands, help)
	subcommand := os.Args[1]

	if needsHelp(subcommand) {
		help.Init(getCmdArgs())
		return help.Run()
	}

	for _, cmd := range commands {
		if cmd.Name() == subcommand {
			cmd.Init(getCmdArgs())
			return cmd.Run()
		}
	}

	return fmt.Errorf(errUnkSub, subcommand)
}
