package flag

type Flag struct {
	Name     string
	Alias    string
	Required bool
	UseValue bool
	Value    interface{}
}

func NewFlag(name string, alias string, required bool, useValue bool) *Flag {
	return &Flag{
		Name:     name,
		Alias:    alias,
		Required: required,
		UseValue: useValue,
	}
}

func (f *Flag) SetValue(value interface{}) {
	f.Value = value
}
