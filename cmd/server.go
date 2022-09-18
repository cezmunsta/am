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

import "flag"

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

func (s *ServerCommand) PrintDefaults() {
	s.fs.PrintDefaults()
}

func (s *ServerCommand) Run() error {
	return nil
}

// NewServerCommand creates the ServerCommand FlagSet
func NewServerCommand() *ServerCommand {
	s := &ServerCommand{
		fs: flag.NewFlagSet("server", flag.ContinueOnError),
	}

	s.fs.Uint("port", 12345, "Set the port for the server")

	return s
}
