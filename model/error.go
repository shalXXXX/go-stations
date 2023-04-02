package model

import "fmt"

type ErrNotFound struct {
	What string
}

func (e *ErrNotFound) Error() string {
	return fmt.Sprintf("%s not found", e.What)
}
