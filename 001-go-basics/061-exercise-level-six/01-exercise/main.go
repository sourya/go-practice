package main

import "fmt"

func main() {
	fmt.Println(foo())

	x, s := bar()
	fmt.Printf("%v says everything is %v\n", s, x)
}

func foo() int {
	return 1989
}

func bar() (int, string) {
	return 19, "Andrew"
}
