package main

import "fmt"

type age int

var x age

func main() {
	fmt.Println("zero value of type age:", x)
	fmt.Printf("The type of variable 'x': %T\n", x)
	x = 19
	fmt.Println(x)
}
