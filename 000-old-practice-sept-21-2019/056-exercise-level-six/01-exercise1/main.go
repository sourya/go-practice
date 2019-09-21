package main

import "fmt"

func main() {
	fmt.Println("The int from foo:", foo())
	ib, sb := bar()
	fmt.Printf("%v %v", sb, ib)
}

func foo() int {
	return 19
}

func bar() (int, string) {
	return 1989, "Andrew"
}
