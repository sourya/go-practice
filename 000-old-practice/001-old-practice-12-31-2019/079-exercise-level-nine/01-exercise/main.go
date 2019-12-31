package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println("Goroutines:\t", runtime.NumGoroutine())

	fmt.Println("Hello from main func")
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("Hello from 1st go func")
		wg.Done()
	}()

	go func() {
		fmt.Println("Hello from 2nd go func")
		wg.Done()
	}()

	fmt.Println("Goroutines:\t", runtime.NumGoroutine())

	wg.Wait()
}
