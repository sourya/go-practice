package main

import "fmt"

type year int

var x year // underlying type is int

func main() {
	fmt.Println(x)
	fmt.Printf("%T is the TYPE of variable x\n", x)
	x = 2020
	fmt.Println(x)
}
