package itunes_search

import (
	"fmt"
	"errors"
)

type notFoundError struct {
	ID   string
	Type string
}

func (self notFoundError) Error() string {
	return fmt.Sprintf("id %s of type %s not found.", self.ID, self.Type)
}

// IDNotFound generate a error from given id value & type
func IDNotFound(id, tp string) error {
	return notFoundError{id, tp}
}

var (
	ErrNotFound = errors.New("not found")
)
