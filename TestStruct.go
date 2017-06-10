package main

import (
	"math/rand"
	"time"
)

type TestStruct struct {
}

func (a TestStruct) GiveMeSomeError(number int) error {
	randomSource := rand.NewSource(time.Now().UnixNano())
	if rand.New(randomSource).Intn(2) < 1 {
		return &TypeAError{ErrorNumber: number}
	} else {
		return &TypeBError{ErrorNumber: number}
	}
}
