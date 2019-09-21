package main

import "fmt"

func main() {
	// BUFFER OF ONE ADDED
	// This code runs because there
	// is a buffer of one declared for it
	c := make(chan int, 1)

	c <- 19

	fmt.Println(<-c)
}
