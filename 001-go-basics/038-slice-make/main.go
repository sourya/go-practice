package main

import "fmt"

func main() {

	// USE a composite literal for CREATING a SLICE
	// can USE MAKE, but it's better to use above method

	x := make([]int, 4, 4)
	fmt.Println(x)
	fmt.Println("Length:", len(x))
	fmt.Println("Capacity:", cap(x))
	x[0] = 1
	x[1] = 2
	x[2] = 3
	x[3] = 4
	fmt.Println(x)
	// below will panic
	// panic: runtime error: index out of range
	// x[4] = 5

	// can append though
	x = append(x, 5)
	fmt.Println(x)
	fmt.Println("Length:", len(x))
	fmt.Println("Capacity:", cap(x))
	fmt.Println("When appending the underlying capacity of the array is doubled")
}
