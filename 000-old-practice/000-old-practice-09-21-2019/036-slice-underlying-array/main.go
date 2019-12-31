package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(x)

	// new underlying array allocated
	y := append(x, 11)
	fmt.Println(y)

	// the same underlying array stores the value of the new slice
	z := append(x[:2], x[5:]...)
	fmt.Println(z)
}
