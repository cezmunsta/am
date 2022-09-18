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

// AddCommand provides the options for adding items
type AddCommand struct {
	fs *flag.FlagSet

	name string
}

func (i *AddCommand) Init(args []string) error {
	return i.fs.Parse(args)
}

func (i *AddCommand) Name() string {
	return i.fs.Name()
}

func (i *AddCommand) PrintDefaults() {
	i.fs.PrintDefaults()
}

func (i *AddCommand) Run() error {
	return nil
}

func (i *AddCommand) SubCommands() []CommandRunner {
	return []CommandRunner{}
}

// NewAddCommand creates the AddCommand FlagSet
func NewAddCommand() *AddCommand {
	i := &AddCommand{
		fs: flag.NewFlagSet("add", flag.ContinueOnError),
	}

	i.fs.String("data", "", "Set input data")

	return i
}

// RemoveCommand provides the options for adding items
type RemoveCommand struct {
	fs *flag.FlagSet

	name string
}

func (i *RemoveCommand) Init(args []string) error {
	return i.fs.Parse(args)
}

func (i *RemoveCommand) Name() string {
	return i.fs.Name()
}

func (i *RemoveCommand) PrintDefaults() {
	i.fs.PrintDefaults()
}

func (i *RemoveCommand) Run() error {
	return nil
}

func (i *RemoveCommand) SubCommands() []CommandRunner {
	return []CommandRunner{}
}

// NewRemoveCommand creates the RemoveCommand FlagSet
func NewRemoveCommand() *RemoveCommand {
	i := &RemoveCommand{
		fs: flag.NewFlagSet("remove", flag.ContinueOnError),
	}

	return i
}

// UpdateCommand provides the options for adding items
type UpdateCommand struct {
	fs *flag.FlagSet

	name string
}

func (i *UpdateCommand) Init(args []string) error {
	return i.fs.Parse(args)
}

func (i *UpdateCommand) Name() string {
	return i.fs.Name()
}

func (i *UpdateCommand) PrintDefaults() {
	i.fs.PrintDefaults()
}

func (i *UpdateCommand) Run() error {
	return nil
}

func (i *UpdateCommand) SubCommands() []CommandRunner {
	return []CommandRunner{}
}

// NewUpdateCommand creates the UpdateCommand FlagSet
func NewUpdateCommand() *UpdateCommand {
	i := &UpdateCommand{
		fs: flag.NewFlagSet("update", flag.ContinueOnError),
	}

	return i
}

// ListCommand provides the options for adding items
type ListCommand struct {
	fs *flag.FlagSet

	name string
}

func (i *ListCommand) Init(args []string) error {
	return i.fs.Parse(args)
}

func (i *ListCommand) Name() string {
	return i.fs.Name()
}

func (i *ListCommand) PrintDefaults() {
	i.fs.PrintDefaults()
}

func (i *ListCommand) Run() error {
	return nil
}

func (i *ListCommand) SubCommands() []CommandRunner {
	return []CommandRunner{}
}

// NewListCommand creates the ListCommand FlagSet
func NewListCommand() *ListCommand {
	i := &ListCommand{
		fs: flag.NewFlagSet("list", flag.ContinueOnError),
	}

	return i
}

func getDefaultCommands() map[string]CommandRunner {
	subCmd := map[string]CommandRunner{
		"add":    NewAddCommand(),
		"remove": NewRemoveCommand(),
		"update": NewUpdateCommand(),
		"list":   NewListCommand(),
	}

	return subCmd
}
