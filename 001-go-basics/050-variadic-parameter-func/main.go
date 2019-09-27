package main

import "fmt"

func main() {
	s := sum(1, 2, 3, 4, 5)
	fmt.Println("The total is,", s)
}

func sum(x ...int) int {
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index postion,", i, "we are now adding,", v, "to the total which is now", sum)
	}

	return sum
}
