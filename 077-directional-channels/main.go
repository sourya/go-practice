package main

import "fmt"

func main() {
	// SEND ONLY CHANNEL
	// DOES NOT WORK
	c := make(chan <- int, 2)

	c <- 19
	c <- 42

	// ERROR
	fmt.Println(<-c)
	// ERROR
	fmt.Println(<-c)
	fmt.Println("-----")
	fmt.Printf("%T\n", c)
}
