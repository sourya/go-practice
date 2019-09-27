package main

import "fmt"

// USE var if DECLARING outside of function body
// best to limit scope of variables

var y = 100

func main() {
	// short declaration operator
	// DECLARE a variable and ASSIGN a VALUE (of a certain TYPE)
	x := 19
	fmt.Println(x)
	fmt.Println(y)
	// SHORT DECLARATION works fine in function body
	// WILL NOT work OUTSIDE of a function body
	foo()
}

func foo() {
	fmt.Println("again", y)
	fmt.Println("I can't access variable \"x\" from the main function")
}
