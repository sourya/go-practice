package main

import "fmt"

// GO ROUTINE 1
func main() {
	c := make(chan int)

	// GO ROUTINE 2
	go func() {
		c <- 19
	}()

	// TRANSFER DATA AT SAME TIME
	fmt.Println(<-c)
}
