package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Println(s1)
	fmt.Println(foo(s1...))
	s2 := []int{10, 20}
	fmt.Println(s2)
	fmt.Println(bar(s2))
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
