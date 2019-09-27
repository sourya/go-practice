package main

import "fmt"

func main() {

	// DEFER starts where SET
	defer foo()
	bar()

	// once main func exits, DEFER RUNS
}

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
