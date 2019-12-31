package main

import "fmt"

func main() {
	x := 83 / 40
	y := 83 % 40
	fmt.Println(x, y)

	fmt.Println("**********")

	z := 1
	for {
		z++
		if z > 100 {
			break
		}
		if z%2 != 0 {
			continue
		}

		fmt.Println(z)

	}
	fmt.Println("done with for statement")
}
