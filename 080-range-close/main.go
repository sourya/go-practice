package main

import "fmt"

func main() {
	c := make(chan int)

	// send
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		// No deadlock once channel is closed
		close(c)
	}()

	// receive
	for v := range c {
		fmt.Println(v)
	}

	fmt.Println("about to exit")
}
