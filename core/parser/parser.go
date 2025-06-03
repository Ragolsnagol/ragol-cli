package parser

import (
	"os"
	"ragol-cli/core/command"
	"ragol-cli/core/flag"
	"strings"
)

type Parse interface {
	ParseInput() []string
	ParseFlags([]string, command.BaseCommand) ([]flag.Flag, error)
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

func (p *Parser) ParseFlags(input []string, command command.BaseCommand) ([]flag.Flag, error) {
	fs := getFlags(input, command)

	err := checkRequiredFlags(fs, command.Flags)
	if err != nil {
		return nil, err
	}

	return fs, nil
}

func getFlags(input []string, command command.BaseCommand) []flag.Flag {
	var fs []flag.Flag

	// Handle dash for flags
	for i := 0; i < len(input); i++ {
		s := input[i]

		for _, f := range command.Flags {
			if strings.EqualFold(f.Name, s) || strings.EqualFold(f.Alias, s) {
				if f.UseValue {
					f.SetValue(input[i+1])
					i++
				}
				fs = append(fs, f)
			}
		}
	}

	return fs
}

func checkRequiredFlags(flags []flag.Flag, cmdFlags []flag.Flag) error {
	// Check if all required flags are present
	for _, cmdFlag := range cmdFlags {
		if !cmdFlag.Required {
			continue
		}

		found := false
		for _, f := range flags {
			if strings.EqualFold(cmdFlag.Name, f.Name) || strings.EqualFold(cmdFlag.Alias, f.Name) {
				found = true
				break
			}
		}

		if !found {
			return &flag.RequiredFlagError{Flag: cmdFlag.Name}
		}
	}

	return nil
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
