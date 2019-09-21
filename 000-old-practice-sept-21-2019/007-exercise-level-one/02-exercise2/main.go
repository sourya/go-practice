package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Printf("%T - type of variable x\n", x)
	fmt.Printf("%T - type of variable y\n", y)
	fmt.Printf("%T - type of variable z\n", z)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println("Compiler assigned these values: ZERO VALUE")
}
