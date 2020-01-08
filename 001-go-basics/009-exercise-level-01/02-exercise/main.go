package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Println("The zero value is", x)
	fmt.Printf("%T\n", x)
	fmt.Println("The zero value is", y)
	fmt.Printf("%T\n", y)
	fmt.Println("The zero value is", z)
	fmt.Printf("%T\n", z)
}
