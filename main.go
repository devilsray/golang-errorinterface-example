package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	for i := 0; i < 100; i++ {
		err := GiveMeSomeError(i)
		if err != nil {
			switch t := err.(type) {
			case *TypeAError:
				fmt.Printf("%d: this is an error from type TypeAError\n", t.ErrorNumber)
			case *TypeBError:
				fmt.Printf("%d: and this is an error from type TypeBError\n", t.ErrorNumber)
			default:
				fmt.Println("not a model missing error")
			}
		}
	}
}

func GiveMeSomeError(number int) error {
	randomSource := rand.NewSource(time.Now().UnixNano())
	if rand.New(randomSource).Intn(2) < 1 {
		return &TypeAError{ErrorNumber: number}
	} else {
		return &TypeBError{ErrorNumber: number}
	}
}