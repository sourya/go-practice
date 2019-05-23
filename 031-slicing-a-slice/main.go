package main

import "fmt"

func main() {
	x := []int{4, 8, 15, 16, 23, 42}
	fmt.Println(x[1])
	fmt.Println(x)
	fmt.Println(x[1:])
	fmt.Println(x[1:3])

	for i, v := range x {
		fmt.Println(i, v)
	}

	for i := 0; i <= 5; i++ {
		fmt.Println(i, x[i])
	}
}
