package main

type TypeAError struct {
}

func (t *TypeAError) Error() string {
	return "Error type a"
}
