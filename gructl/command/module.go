package command

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/dnaeon/gru/module"
)

// NewModuleCommand creates a new sub-command for
// displaying the list of available modules
func NewModuleCommand() cli.Command {
	cmd := cli.Command{
		Name:   "module",
		Usage:  "display available modules",
		Action: execModuleCommand,
	}

	return cmd
}

// Executes the "module" command
func execModuleCommand(c *cli.Context) {
	path := c.GlobalString("modulepath")
	if path == "" {
		displayError(errInvalidModulePath, 64)
	}

	registry, err := module.Discover(path)
	if err != nil {
		displayError(err, 1)
	}

	for name := range registry {
		fmt.Println(name)
	}
}