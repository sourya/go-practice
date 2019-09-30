package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	c := make(chan int)
	go func() {
		c <- 42
	}()
	fmt.Println("Goroutines:", runtime.NumGoroutine())
	fmt.Println(<-c)
}
