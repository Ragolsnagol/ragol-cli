package core

import (
	"ragol-cli/core/command"
	"ragol-cli/core/parser"
)

type App struct {
	Name     string
	Version  string
	Parser   parser.Parser
	Commands []command.BaseCommand
}

func NewApp(name string, version string, commands []command.BaseCommand) *App {
	return &App{
		Name:     name,
		Version:  version,
		Parser:   *parser.NewParser(),
		Commands: commands,
	}
}

func (a *App) Run() {
	args := a.Parser.ParseInput()

	if len(args) == 0 {
		// TODO: Default to the help command
	}

	c := a.Parser.ParseCommand(args[0], a.Commands)

	// Pass in all arguments after the command
	fs := a.Parser.ParseFlags(args[1:], c)

	if len(fs) > 0 {
		c.SetActiveFlags(fs)
	}

	err := c.Run()
	if err != nil {
		panic(err)
	}
}
