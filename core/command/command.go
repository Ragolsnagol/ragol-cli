package command

import "ragol-cli/core/flag"

type Command interface {
	Run() error
	SetActiveFlags([]flag.Flag)
}

type BaseCommand struct {
	Name        string
	Action      func([]flag.Flag) error
	Flags       []flag.Flag
	ActiveFlags []flag.Flag
}

func NewCommand(name string, action func([]flag.Flag) error, flags []flag.Flag) *BaseCommand {
	return &BaseCommand{
		Name:   name,
		Action: action,
		Flags:  flags,
	}
}

func (c *BaseCommand) Run() error {
	return c.Action(c.ActiveFlags)
}

func (c *BaseCommand) SetActiveFlags(flags []flag.Flag) {
	c.ActiveFlags = flags
}
