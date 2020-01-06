package main

import "fmt"

var y string

func main() {
	fmt.Println("Printing the value of y", y)
	fmt.Printf("%T\n", y)
	y = "Russell & Ruby"
	fmt.Println("Printing the value of y", y)
	fmt.Printf("%T\n", y)
}
