package main

import "fmt"

var x int

type year int // creating own type

var y year

func main() {
	x = 19
	y = 2020
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(y)
	fmt.Printf("%T\n", y)

	// no casting in Go

	// Go has conversion & type assertions
	x = int(y)
	fmt.Println(x)
	fmt.Printf("%T\n", x)

}
