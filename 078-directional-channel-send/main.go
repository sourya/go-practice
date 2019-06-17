package main

import "fmt"

func main() {
	// SEND ONLY CHANNEL
	c := make(<- chan int, 2)

	// ERROR
	c <- 19
	// ERROR
	c <- 42

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("-----")
	fmt.Printf("%T\n", c)
}

