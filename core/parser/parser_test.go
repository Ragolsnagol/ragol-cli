package parser

import (
	"os"
	"ragol-cli/core/command"
	"ragol-cli/core/flag"
	"reflect"
	"testing"
)

func TestNewParser(t *testing.T) {
	parser := NewParser()
	if parser == nil {
		t.Error("Expected non-nil parser")
	}
}

func TestParseInput(t *testing.T) {
	// Save original args and restore after test
	originalArgs := os.Args
	defer func() { os.Args = originalArgs }()

	tests := []struct {
		name     string
		args     []string
		expected []string
	}{
		{
			name:     "empty args",
			args:     []string{"program"},
			expected: []string{},
		},
		{
			name:     "single arg",
			args:     []string{"program", "arg1"},
			expected: []string{"arg1"},
		},
		{
			name:     "multiple args",
			args:     []string{"program", "arg1", "arg2", "arg3"},
			expected: []string{"arg1", "arg2", "arg3"},
		},
	}

	parser := NewParser()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.args
			result := parser.ParseInput()
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ParseInput() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestParseFlags(t *testing.T) {
	tests := []struct {
		name        string
		input       []string
		command     command.BaseCommand
		expected    []flag.Flag
		expectError bool
	}{
		{
			name:  "no flags",
			input: []string{},
			command: command.BaseCommand{
				Name:  "test",
				Flags: []flag.Flag{},
			},
			expected:    nil,
			expectError: false,
		},
		{
			name:  "single flag without value",
			input: []string{"-f"},
			command: command.BaseCommand{
				Name: "test",
				Flags: []flag.Flag{
					{Name: "-f", UseValue: false},
				},
			},
			expected: []flag.Flag{
				{Name: "-f", UseValue: false},
			},
			expectError: false,
		},
		{
			name:  "flag with value",
			input: []string{"-f", "value"},
			command: command.BaseCommand{
				Name: "test",
				Flags: []flag.Flag{
					{Name: "-f", UseValue: true},
				},
			},
			expected: []flag.Flag{
				{Name: "-f", UseValue: true, Value: "value"},
			},
			expectError: false,
		},
		{
			name:  "missing required flag",
			input: []string{},
			command: command.BaseCommand{
				Name: "test",
				Flags: []flag.Flag{
					{Name: "-r", Required: true},
				},
			},
			expected:    nil,
			expectError: true,
		},
		{
			name:  "alias flag",
			input: []string{"-a"},
			command: command.BaseCommand{
				Name: "test",
				Flags: []flag.Flag{
					{Name: "-flag", Alias: "-a"},
				},
			},
			expected: []flag.Flag{
				{Name: "-flag", Alias: "-a"},
			},
			expectError: false,
		},
	}

	parser := NewParser()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := parser.ParseFlags(tt.input, tt.command)

			if tt.expectError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tt.expectError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if !tt.expectError && !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ParseFlags() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestParseCommand(t *testing.T) {
	commands := []command.BaseCommand{
		{Name: "command1"},
		{Name: "command2"},
		{Name: "help"}, // Last command is help
	}

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "existing command",
			input:    "command1",
			expected: "command1",
		},
		{
			name:     "case insensitive command",
			input:    "COMMAND2",
			expected: "command2",
		},
		{
			name:     "non-existent command defaults to help",
			input:    "invalid",
			expected: "help",
		},
	}

	parser := NewParser()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := parser.ParseCommand(tt.input, commands)
			if result.Name != tt.expected {
				t.Errorf("ParseCommand() = %v, want %v", result.Name, tt.expected)
			}
		})
	}
}
