package flag

import (
	"errors"
	"testing"
)

func TestNewFlag(t *testing.T) {
	tests := []struct {
		name     string
		flagName string
		alias    string
		required bool
		useValue bool
		wantErr  bool
		errType  string
	}{
		{
			name:     "valid flag with name and alias",
			flagName: "--test",
			alias:    "-t",
			wantErr:  false,
		},
		{
			name:     "valid flag with required and useValue",
			flagName: "--verbose",
			alias:    "-v",
			required: true,
			useValue: true,
			wantErr:  false,
		},
		{
			name:     "invalid short name without prefix",
			flagName: "t",
			alias:    "-v",
			wantErr:  true,
			errType:  "InvalidFlagError",
		},
		{
			name:     "invalid name too short",
			flagName: "--",
			alias:    "-t",
			wantErr:  true,
			errType:  "InvalidFlagError",
		},
		{
			name:     "invalid alias format",
			flagName: "--test",
			alias:    "invalid",
			wantErr:  true,
			errType:  "InvalidFlagError",
		},
		{
			name:     "single character alias without prefix",
			flagName: "--test",
			alias:    "t",
			wantErr:  true,
			errType:  "InvalidFlagError",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			flag, err := NewFlag(tt.flagName, tt.alias, tt.required, tt.useValue)

			if tt.wantErr {
				if err == nil {
					t.Errorf("NewFlag() error = nil, wantErr %v", tt.wantErr)
					return
				}
				var invalidFlagError *InvalidFlagError
				if !errors.As(err, &invalidFlagError) {
					t.Errorf("NewFlag() error type = %T, want %s", err, tt.errType)
				}
				return
			}

			if err != nil {
				t.Errorf("NewFlag() unexpected error = %v", err)
				return
			}

			if flag.Name != tt.flagName {
				t.Errorf("flag.Name = %v, want %v", flag.Name, tt.flagName)
			}
			if flag.Alias != tt.alias {
				t.Errorf("flag.Alias = %v, want %v", flag.Alias, tt.alias)
			}
			if flag.Required != tt.required {
				t.Errorf("flag.Required = %v, want %v", flag.Required, tt.required)
			}
			if flag.UseValue != tt.useValue {
				t.Errorf("flag.UseValue = %v, want %v", flag.UseValue, tt.useValue)
			}
		})
	}
}

func TestFlag_SetValue(t *testing.T) {
	flag := &Flag{
		Name:     "--test",
		Alias:    "-t",
		Required: true,
		UseValue: true,
	}

	testValue := "test-value"
	flag.SetValue(testValue)

	if flag.Value != testValue {
		t.Errorf("SetValue() = %v, want %v", flag.Value, testValue)
	}
}
