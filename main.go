package main

import (
	"fmt"
)

type TestObject interface {
	GiveMeSomeError(int int) error
}

func main() {
	testStruct := TestStruct{}
	for i := 0; i < 100; i++ {
		err := testStruct.GiveMeSomeError(i)
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
