package main

import "fmt"

func main() {
	// DOES NOT RUN
	// SEND ONLY CHANNEL
	c := make(chan <- int, 2)
	c <- 42
	c <- 43
	// TRYING to RECEIVE
	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("-----")
	fmt.Printf("%T\n", c)
}
