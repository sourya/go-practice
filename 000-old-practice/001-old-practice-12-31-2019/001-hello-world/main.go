// every program needs a package main
package main

import "fmt"

// every program needs a func main
func main() {
	fmt.Println("Hello Ruby & Russell")

	foo()
	fmt.Println("Another print line")

	// iterative
	for i := 0; i < 100; i++ {
		// if i MOD 2 is zero
		// prints even numbers
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("print line from foo function")
}

// control flow:

// (1) sequence
// (2) loop, iterative
// (3) conditional
