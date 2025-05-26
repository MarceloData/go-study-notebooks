package main

import (
	"fmt"

	"github.com/google/uuid"
)

// Package with basic concepts of go lang

func main() {
	fmt.Println("This is Go basics!")
	fmt.Println("UUID:", uuid.New().String())
	variables()
	pointers()
}
