package main

import "fmt"

var y = 42

var z = "Russell & Ruby"

func main() {
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println(z)
	fmt.Printf("%T\n", z)

	y = 19
	fmt.Println(y)
}
