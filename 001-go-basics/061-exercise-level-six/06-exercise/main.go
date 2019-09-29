package main

import "fmt"

func main() {

	func(x ...int) {
		sum := 0
		for _, v := range x {
			sum += v
		}
		fmt.Println("The total sum is:", sum)
	}(2, 2, 2)
}
