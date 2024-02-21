package internal

import (
	"fmt"
)

type ErrConnAccept struct {
	ID   uint32
	Name string
	Err  error
}

func (e ErrConnAccept) Error() string {
	return fmt.Sprintf("failed to accept %s server connection %d: %s", e.Name, e.ID, e.Err)
}

func (e ErrConnAccept) Unwrap() error {
	return e.Err
}

type ErrConnDial struct {
	ID   uint32
	Name string
	Err  error
}

func (e ErrConnDial) Error() string {
	return fmt.Sprintf("failed to dial %s client connection %d: %s", e.Name, e.ID, e.Err)
}

func (e ErrConnDial) Unwrap() error {
	return e.Err
}
