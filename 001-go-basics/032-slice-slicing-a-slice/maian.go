package main

import "fmt"

func main() {
	x := []int{1, 2, 3, 4, 5}
	fmt.Println(x[1])

	fmt.Println(x[1:])
	fmt.Println(x[4:])
	fmt.Println(x[1:3])

	for i, v := range x {
		fmt.Println(i, v)
	}

	fmt.Println("##########")

	for i := 0; i < len(x); i++ {
		fmt.Println(x[i])
	}
}
