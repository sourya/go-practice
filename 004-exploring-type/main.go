package main

import "fmt"

// GO is all about TYPE

// STATIC programming language

// not a DYNAMIC language like JavaScript

// variable y is TYPE int
var y = 19

// variable z is TYPE string
var z = "Phoenix"

var msg = `Using grave accents
Which keep everything inside of them

like line returns, etc, etc..        ^_^v
`

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
	fmt.Println()
	fmt.Println(msg)
}
