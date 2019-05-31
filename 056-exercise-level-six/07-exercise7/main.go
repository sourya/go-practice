package main

import "fmt"

func main() {
	x := sum(5, 5)
	fmt.Println("Returned value:", x)
}

var sum = func(x, y float64) float64 {
	return x + y
}
