package main

import "fmt"

type year int

var x year
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 2020
	fmt.Println(x)
	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
}
