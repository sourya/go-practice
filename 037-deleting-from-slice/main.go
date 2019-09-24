package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4}
	fmt.Println(x)
	x = append(x, 5, 6, 7, 8)
	fmt.Println(x)

	y := []int{9, 10, 11, 12}
	x = append(x, y...)
	fmt.Println(x)

	// DELETE from SLICE

	x = append(x[:2], x[10:]...)
	fmt.Println(x)
}
