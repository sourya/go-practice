package main

import "fmt"

func main() {
	// defer keyword
	defer foo()
	bar()
} // foo() runs here now before/as program closes

func foo() {
	fmt.Println("foo")
}

func bar() {
	fmt.Println("bar")
}
