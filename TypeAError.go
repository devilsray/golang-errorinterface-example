package main

import (
	"fmt"
)

type TypeAError struct {
	ErrorNumber int
}

func (t *TypeAError) Error() string {
	return fmt.Sprintf("Error number %s", t.ErrorNumber)
}
