package main

import "fmt"

func main() {
	// Have to declare the size of the array
	var x [5]int
	var y [6]int
	fmt.Println(x)
	x[3] = 19
	fmt.Println(x)
	fmt.Println(y)

}
