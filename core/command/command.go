package command

import (
	"ragol-cli/core/action"
	"ragol-cli/core/context"
	"ragol-cli/core/flag"
)

type Command interface {
	Run() error
	SetActiveFlags([]flag.Flag)
}

type BaseCommand struct {
	Name        string
	HelpString  string
	Action      action.Action
	Flags       []flag.Flag
	ActiveFlags []flag.Flag
}

func NewCommand(name string, helpString string, action action.Action, flags []flag.Flag) *BaseCommand {
	return &BaseCommand{
		Name:       name,
		HelpString: helpString,
		Action:     action,
		Flags:      flags,
	}
}

func (c *BaseCommand) Run(ctx context.Context) error {
	return c.Action.Action(ctx)
}
