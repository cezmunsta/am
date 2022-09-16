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
	"flag"
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
	Run() error
	Name() string
}

// HelpCommand provides the interface to the help information
type HelpCommand struct {
	fs *flag.FlagSet

	name string
}

func (h *HelpCommand) Init(args []string) error {
	return h.fs.Parse(args)
}

func (h *HelpCommand) Name() string {
	return h.fs.Name()
}

func (h *HelpCommand) Run() error {
	return nil
}

// NewHelpCommand creates the HelpCommand FlagSet
func NewHelpCommand() *HelpCommand {
	h := &HelpCommand{
		fs: flag.NewFlagSet("help", flag.ExitOnError),
	}

	return h
}

// ServerCommand provides the interface to the server functionality
type ServerCommand struct {
	fs *flag.FlagSet

	name string
}

func (s *ServerCommand) Init(args []string) error {
	return s.fs.Parse(args)
}

func (s *ServerCommand) Name() string {
	return s.fs.Name()
}

func (s *ServerCommand) Run() error {
	return nil
}

// NewServerCommand creates the ServerCommand FlagSet
func NewServerCommand() *ServerCommand {
	s := &ServerCommand{
		fs: flag.NewFlagSet("server", flag.ContinueOnError),
	}

	return s
}

func Execute(args []string) error {
	if len(args) == 0 {
		return errors.New(errNoArgs)
	}

	commands := []CommandRunner{
		NewServerCommand(),
	}

	subcommand := os.Args[1]

	for _, cmd := range commands {
		if cmd.Name() == subcommand {
			cmd.Init(os.Args[2:])
			return cmd.Run()
		}
	}

	return fmt.Errorf(errUnkSub, subcommand)
}
