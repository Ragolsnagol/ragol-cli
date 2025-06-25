package core

import (
	"github.com/ragolsnagol/ragol-cli/core/command"
	"github.com/ragolsnagol/ragol-cli/core/context"
	"github.com/ragolsnagol/ragol-cli/core/parser"
)

type App struct {
	Name     string
	Version  string
	Parser   parser.Parser
	Commands []command.BaseCommand
	Context  context.Context
}

func NewApp(name string, version string, commands []command.BaseCommand) *App {
	h := command.CreateHelpCommand(commands)
	commands = append(commands, h)

	return &App{
		Name:     name,
		Version:  version,
		Parser:   *parser.NewParser(),
		Commands: commands,
		Context:  *context.NewContext(),
	}
}

func (a *App) Run() error {
	args := a.Parser.ParseInput()

	if len(args) == 0 {
		// Show the help command if nothing was supplied
		help := a.Commands[len(a.Commands)-1]
		return help.Run(a.Context)
	}

	c := a.Parser.ParseCommand(args[0], a.Commands)

	// Pass in all arguments after the command
	fs, err := a.Parser.ParseFlags(args[1:], c)
	if err != nil {
		return err
	}

	if len(fs) > 0 {
		a.Context.SetFlags(fs)
	}

	err = c.Run(a.Context)
	if err != nil {
		return err
	}

	return nil
}
