package main

import (
	"fmt"
	"github.com/brudnak/go-practice/001-go-basics/099-exercise-level-thirteen/01-exercise/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog.Years(10),
	}
	fmt.Println(fido)
}