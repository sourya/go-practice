package main

import "fmt"

func main() {

	fmt.Println(foo())

	x := bar()

	fmt.Println(x())
	// wow
	fmt.Println(bar()())
}

func foo() string {
	return "Hello from foo"

}

func bar() func() int {
	return func() int {
		return 2020
	}
}
