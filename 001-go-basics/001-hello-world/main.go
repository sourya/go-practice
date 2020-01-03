package main

import (
	"fmt"
	foo2 "github.com/brudnak/go-practice/001-go-basics/001-hello-world/foo"
)

func main() {
	fmt.Println("Hello Russell & Ruby!")
	fmt.Println(foo2.HelloMsg())
	fmt.Println("Something else")

	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("I'm in foo")
}

// control flow:
// (1) sequence
// (2) look; iterative
// (3) conditional
