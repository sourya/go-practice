package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	fmt.Println(x[1])
	fmt.Println(x)
	fmt.Println(x[1:])
	fmt.Println(x[0:2])
	fmt.Println(len(x))
	for i := 0; i <= len(x)-1; i++ {
		fmt.Println("Printing from the slice:", x[i])
	}
}
