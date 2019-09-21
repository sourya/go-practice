package main

import "fmt"

var houseSold bool

func main() {
	fmt.Println(houseSold)
	houseSold = true
	fmt.Println(houseSold)

	a := 19
	b := 2020
	fmt.Println(a == b)
}
