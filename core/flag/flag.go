package flag

import "strings"

type Flag struct {
	Name     string
	Alias    string
	Required bool
	UseValue bool
	Value    interface{}
}

func NewFlag(name string, alias string, required bool, useValue bool) (*Flag, error) {
	if len(name) <= 2 || !strings.HasPrefix(name, "--") {
		return nil, &InvalidFlagError{Flag: name}
	}
	if (len(alias) > 1 && len(alias) != 2) || !strings.HasPrefix(alias, "-") {
		return nil, &InvalidFlagError{Flag: alias}
	}

	return &Flag{
		Name:     name,
		Alias:    alias,
		Required: required,
		UseValue: useValue,
	}, nil
}

func (f *Flag) SetValue(value interface{}) {
	f.Value = value
}
