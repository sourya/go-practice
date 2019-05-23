package main

import "fmt"

func main() {
	x := []int{4, 8, 15, 16, 23, 42}
	fmt.Println(len(x))
	fmt.Println(x)
	fmt.Println(x[0])
	fmt.Println(x[1])
	fmt.Println(x[2])
	fmt.Println(x[3])
	fmt.Println(x[4])
	fmt.Println(x[5])

	// i = INDEX
	// v = VALUE
	// Could be named ANYTHING
	for i, v := range x {
		fmt.Println(i, v)
	}
	// Ease of programming
}


