package flag

type Flag struct {
	Name     string
	Alias    string
	Required bool
}

func NewFlag(name string, alias string, required bool) *Flag {
	return &Flag{
		Name:     name,
		Alias:    alias,
		Required: required,
	}
}
