package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {

	fmt.Println(runtime.NumCPU())

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		fmt.Println("hello from 1st routine")
		wg.Done()
	}()

	go func() {
		fmt.Println("hello from 2nd routine")
		wg.Done()

	}()

	wg.Wait()

	fmt.Println("about to exit")
}
