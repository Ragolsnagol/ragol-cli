package flag

import "fmt"

type RequiredFlagError struct {
	Flag string
}

func (e *RequiredFlagError) Error() string {
	return fmt.Sprintf("Required flag %v not provided", e.Flag)
}

type InvalidFlagError struct {
	Flag string
}

func (e *InvalidFlagError) Error() string {
	return fmt.Sprintf("Invalid flag %v", e.Flag)
}
