package main

import "fmt"

func main() {
	x := [5]int{1, 2, 3, 4, 5}
	fmt.Println(x)

	fmt.Printf("Type of the variable 'x': %T\n", x)
	for i, v := range x {

		fmt.Printf("Index: %v\t\t Value: %v\n", i, v)
	}
}
