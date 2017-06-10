package main

type TypeBError struct {
}

func (t *TypeBError) Error() string {
	return "Error type b"
}
