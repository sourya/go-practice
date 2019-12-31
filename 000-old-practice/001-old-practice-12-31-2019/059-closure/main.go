package main

import "fmt"

// SCOPE of x is ENTIRE package

var x int

func main() {
	fmt.Println(x)
	x++
	fmt.Println(x)
	foo()
	fmt.Println(x)

	a := increment()
	b := increment()

	fmt.Println("From increment 'a':", a())
	fmt.Println("From increment 'a':", a())
	fmt.Println("From increment 'a':", a())
	fmt.Println()
	fmt.Println("From increment 'b':", b())
	fmt.Println("From increment 'b':", b())
	fmt.Println("From increment 'b':", b())
}

func foo() {
	fmt.Println("hello")
	x++
}

func increment() func() int {
	var x int
	return func() int {
		x++
		return x
	}
}
