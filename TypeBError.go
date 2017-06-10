package main

import (
	"fmt"
)

type TypeBError struct {
	ErrorNumber int
}

func (t *TypeBError) Error() string {
	return fmt.Sprintf("Error number %s", t.ErrorNumber)
}
