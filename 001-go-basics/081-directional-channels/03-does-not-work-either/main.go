package main

import "fmt"

func main() {
	// DOES NOT RUN
	// RECEIVE ONLY CHANNEL
	c := make(<-chan int, 2)
	// TRYING to SEND fails
	c <- 42
	c <- 43
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("-----")
	fmt.Printf("%T\n", c)
}
