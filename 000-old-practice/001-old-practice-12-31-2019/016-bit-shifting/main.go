package main

import "fmt"

var x = 16

func main() {

	// bit shifting
	fmt.Printf("%d\t\t%b\n", x, x)
	y := x << 1
	fmt.Printf("%d\t\t%b\n\n", y, y)

	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)

}
