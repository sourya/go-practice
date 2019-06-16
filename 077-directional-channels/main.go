package main

import "fmt"

func main() {
	c := make(chan int, 2)

	c <- 19
	c <- 42

	fmt.Println(<-c)
	fmt.Println(<-c)
	fmt.Println("-----")
	fmt.Printf("%T\n", c)
}
