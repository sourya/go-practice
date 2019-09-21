package main

import "fmt"

func main() {
	// THIS CODE FAILS TO RUN
	// Only one buffer declared
	c := make(chan int, 1)

	// Not enough room with only one buffer
	// DEADLOCK!
	c <- 19
	c <- 42

	fmt.Println(<-c)
}
