package main

import "fmt"

func main() {
	xi := []int{1, 2, 3, 4, 5}
	fmt.Println(foo(xi...))
	fmt.Println(bar(xi))
}

func foo(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func bar(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
