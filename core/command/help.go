package command

import (
	"fmt"
	"ragol-cli/core/flag"
)

func CreateHelpCommand(c []BaseCommand) BaseCommand {
	return BaseCommand{
		Name:   "help",
		Action: createHelpAction(c),
	}
}

func createHelpAction(c []BaseCommand) func([]flag.Flag) error {
	return func(flags []flag.Flag) error {
		fmt.Println("Available commands:")
		for _, cmd := range c {
			fmt.Printf("%v - %v\n", cmd.Name, cmd.HelpString)
		}
		return nil
	}
}
