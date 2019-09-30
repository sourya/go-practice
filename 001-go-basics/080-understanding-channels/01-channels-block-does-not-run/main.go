package main

import "fmt"

func main() {
	// does not run
	c := make(chan int)
	c <- 42
	fmt.Println(<-c)
}
