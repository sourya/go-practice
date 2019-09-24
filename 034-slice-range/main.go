package main

import "fmt"

func main() {
	x := []int{20,30,40,50,60}
	fmt.Println(len(x))
	fmt.Println(x)
	fmt.Println(x[0])

	for i, v := range x {
		fmt.Printf("The index is: %v\t The value is: %v\n", i, v)
	}
}
