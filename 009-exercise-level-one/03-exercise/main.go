package main

import "fmt"

var x = 29
var y = "Andrew Brudnak"
var z = true

func main() {
	s := fmt.Sprintf("%v is %v years old - has children: %v\n", y, x, z)
	fmt.Println(s)
	fmt.Printf("s is of type: %T\n", s)
}
