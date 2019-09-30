package main

import "fmt"

func main() {
	// DOES NOT RUN
	c := make(chan int, 1)
	c <- 42
	c <- 43
	fmt.Println(<-c)
}
