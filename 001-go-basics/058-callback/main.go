package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(sum(5, 5, 5, 5, 5))
	fmt.Println(sum(10, 10, 10, 10, 10))
	fmt.Println(sum(ii...))
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
