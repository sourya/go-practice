package main

import "fmt"

func main() {

	fmt.Println(foo()())
	f := foo()
	fmt.Println(f())
}

func foo() func() int {
	return func() int {
		return 42
	}
}
