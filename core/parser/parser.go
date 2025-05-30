package parser

import (
	"os"
	"ragol-cli/core/command"
	"ragol-cli/core/flag"
	"strings"
)

type Parse interface {
	ParseInput() []string
	ParseFlags([]string, command.BaseCommand) []flag.Flag
	ParseCommand(string, []command.BaseCommand) command.BaseCommand
}

type Parser struct {
}

func NewParser() *Parser {
	return &Parser{}
}

func (p *Parser) ParseInput() []string {
	// Ignore the first argument since that's just the name of the program
	return os.Args[1:]
}

func (p *Parser) ParseFlags(input []string, command command.BaseCommand) []flag.Flag {
	var fs []flag.Flag

	// TODO: Check alias as well
	for _, s := range input {
		for _, f := range command.Flags {
			if strings.EqualFold(f.Name, s) {
				fs = append(fs, f)
			}
		}
	}

	return fs
}

func (p *Parser) ParseCommand(input string, commands []command.BaseCommand) command.BaseCommand {
	for _, c := range commands {
		if strings.EqualFold(c.Name, input) {
			return c
		}
	}

	// Default to the help command (which is the last command)
	return commands[len(commands)-1]
}
