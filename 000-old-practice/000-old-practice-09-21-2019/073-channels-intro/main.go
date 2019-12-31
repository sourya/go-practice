package main

import "fmt"

// Does not run

// CHANNELS BLOCK
func main() {
	c := make(chan int)
	c <- 42
	fmt.Println(<-c)
}
