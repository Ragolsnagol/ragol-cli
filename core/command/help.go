package command

import (
	"fmt"
	"github.com/ragolsnagol/ragol-cli/core/action"
	"github.com/ragolsnagol/ragol-cli/core/context"
)

func CreateHelpCommand(c []BaseCommand) BaseCommand {
	return BaseCommand{
		Name:   "help",
		Action: action.NewAction(createHelpAction(c)),
	}
}

func createHelpAction(c []BaseCommand) func(context.Context) error {
	return func(ctx context.Context) error {
		fmt.Println("Available commands:")
		for _, cmd := range c {
			fmt.Printf("%v - %v\n", cmd.Name, cmd.HelpString)
		}
		return nil
	}
}
