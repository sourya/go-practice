package main

import "fmt"

var y = 7 // package level scope
// cannot use short declaration outside of a function

// DECLARE the VARIABLE with IDENTIFIER "z"
var z int // zero value declared: z equals 0

func main() {
	// short declaration operator
	// DECLARE a variable and ASSIGN a VALUE (of a certain type)
	x := 2019

	fmt.Println(x)

	fmt.Println(y)

	foo()
} // program exit point

func foo() {
	fmt.Println("Printing variable y again:", y)
}
