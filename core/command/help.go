package command

import (
	"fmt"
	"ragol-cli/core/flag"
)

func CreateHelpCommand(c []BaseCommand) BaseCommand {
	return BaseCommand{
		Name:   "help",
		Action: createActionHelp(c),
	}
}

func createActionHelp(c []BaseCommand) func([]flag.Flag) error {
	return func(flags []flag.Flag) error {
		fmt.Println("Available commands:")
		for _, cmd := range c {
			fmt.Println(cmd.Name)
		}
		return nil
	}
}
