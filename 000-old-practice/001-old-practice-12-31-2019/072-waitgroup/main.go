package main

import (
	"fmt"
	"runtime"
	"sync"
)

func init() {
	fmt.Println("I'll run once before func main")
}

var wg sync.WaitGroup

func main() {
	fmt.Println("OS\t\t\t\t", runtime.GOOS)
	fmt.Println("ARCH\t\t\t", runtime.GOARCH)
	fmt.Println("CPUs\t\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t\t", runtime.NumGoroutine())

	wg.Add(1)
	go foo()

	bar()

	fmt.Println("CPUs\t\t\t", runtime.NumCPU())
	fmt.Println("Goroutines\t\t", runtime.NumGoroutine())
	wg.Wait()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("foo:", i)
	}
	wg.Done()
}

func bar() {
	for i := 0; i < 10; i++ {
		fmt.Println("bar:", i)
	}
}
