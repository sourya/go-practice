package main

import "fmt"

func main() {
	foo()

	func() {
		fmt.Println("Anonymous func ran")
	}()

	func(x int) {
		fmt.Println("Everything is:", x)
	}(19)
}

func foo() {
	fmt.Println("foo ran")
}
