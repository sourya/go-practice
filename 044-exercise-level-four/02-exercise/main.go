package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("Type of the 'xi' variable: %T\n", xi)

	for i, v := range xi {
		fmt.Printf("Index: %v\t\t Value: %v\n", i, v)
	}
}
