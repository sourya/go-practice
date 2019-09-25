package main

import "fmt"

func main() {

	// Slice ia a data structure
	// THREE parts to a SLICE

	// 1. POINTER to an underlying ARRAY
	// 2. LENGTH of the SLICE
	// 3. CAPACITY is length of the ARRAY

	x := []int{1, 2, 3, 4}
	fmt.Println(x)

	y := append(x, 5)
	fmt.Println(y)

}
