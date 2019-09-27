package main

import "fmt"

func main() {

	//xi := []int{1, 2, 3, 4, 5}

	// SAME SLICE USED ON LINE 15
	//s := sum(xi...)
	s := sum()
	fmt.Println("The total is,", s)
}

// RECEIVING the SLICE from line 10 UNCHANGED, it's the SAME SLICE
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
