package main

import "fmt"

const (
	a = iota + 2016
	b
	c
	d
)

func main() {

	// last four years
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}
