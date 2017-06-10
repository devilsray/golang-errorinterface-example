package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 100; i++ {
		err := GiveMeSomeError() // get an error type a or b
		if err != nil {
			switch t := err.(type) {
			case *TypeAError:
				fmt.Printf("%d: this is an error from type TypeAError\n", i)
			case *TypeBError:
				fmt.Printf("%d: and this is an error from type TypeBError\n", i)
			default:
				fmt.Printf("unknown error type: %s\n", t)
			}
		}
	}
}

// Return random error type
func GiveMeSomeError() error {
	randomSource := rand.NewSource(time.Now().UnixNano()) // create random source
	if rand.New(randomSource).Intn(2) < 1 {
		return &TypeAError{} // return error type a
	} else {
		return &TypeBError{} // return error type b
	}
}
