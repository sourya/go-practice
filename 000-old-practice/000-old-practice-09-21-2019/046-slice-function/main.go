package main

import "fmt"

func main() {
	xi := []int{2, 3, 4, 5, 6, 7, 8, 9}
	x := sum(xi...)
	fmt.Println("The total is", x)
}

func sum(x ...int) int {
	fmt.Println(x)
	// turns into a SLICE of INT
	fmt.Printf("%T\n", x)

	sum := 0
	for i, v := range x {
		sum += v
		fmt.Println("for item in index position", i, "we are now adding", v, "to the total which is now", sum)
	}
	fmt.Println("the total is", sum)
	return sum

}

// func (r receiver) identifier(parameters(s)) (returns(s)) { code }
