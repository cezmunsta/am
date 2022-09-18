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

// InvoiceCommand provides the client interface to invoice management
type InvoiceCommand struct {
	fs     *flag.FlagSet
	subCmd map[string]CommandRunner

	name string
}

func (i *InvoiceCommand) Init(args []string) error {
	return i.fs.Parse(args)
}

func (i *InvoiceCommand) Name() string {
	return i.fs.Name()
}

func (i *InvoiceCommand) PrintDefaults() {
	i.fs.PrintDefaults()
}

func (i *InvoiceCommand) Run() error {
	return nil
}

func (i *InvoiceCommand) SubCommands() []CommandRunner {
	return []CommandRunner{}
}

// NewInvoiceCommand creates the InvoiceCommand FlagSet
func NewInvoiceCommand() *InvoiceCommand {
	i := &InvoiceCommand{
		fs:     flag.NewFlagSet("invoice", flag.ContinueOnError),
		subCmd: getDefaultCommands(),
	}

	return i
}
