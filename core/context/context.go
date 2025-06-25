package context

import "github.com/ragolsnagol/ragol-cli/core/flag"

type Context struct {
	Flags []flag.Flag
}

func NewContext() *Context {
	return &Context{
		Flags: nil,
	}
}

func (c *Context) SetFlags(flags []flag.Flag) {
	c.Flags = flags
}

func (c *Context) GetFlag(name string) (*flag.Flag, error) {
	for _, f := range c.Flags {
		if f.Name == name || f.Alias == name {
			return &f, nil
		}
	}

	return nil, nil
}
