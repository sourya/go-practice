package main

import "fmt"

func main() {
	ii := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(sum(5, 5, 5, 5, 5))
	fmt.Println(sum(10, 10, 10, 10, 10))
	fmt.Println(sum(ii...))
	fmt.Println(odd(sum, 3, 3), "ODD THINGS")
	fmt.Println(odd(sum, ii...))

	s2 := even(sum, ii...)
	fmt.Println(s2)

	fmt.Println()

	s := sum(ii...)
	fmt.Println("all numbers", s)

	t1 := even(sum, ii...)
	fmt.Println("even numbers", t1)

	u1 := odd(sum, ii...)
	fmt.Println("odd numbers", u1)
}

func sum(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func even(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 == 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}

func odd(f func(xi ...int) int, vi ...int) int {
	var yi []int
	for _, v := range vi {
		if v%2 != 0 {
			yi = append(yi, v)
		}
	}
	return f(yi...)
}
