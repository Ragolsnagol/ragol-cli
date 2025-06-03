package flag

import "fmt"

type RequiredFlagError struct {
	Flag string
}

func (e *RequiredFlagError) Error() string {
	return fmt.Sprintf("Required flag %v not provided", e.Flag)
}
