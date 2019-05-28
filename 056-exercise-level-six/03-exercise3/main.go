package main

import "fmt"

func main() {
	defer foo()
	bar()
	greeting()
	woo()
}

func foo() {
	defer func() {
		fmt.Println("I was deferred...")
	}()
	fmt.Println("1st - Hello from foo!")
}

func bar() {
	fmt.Println("2nd - Hello from bar!")
}

func greeting() {
	fmt.Println("3rd - Hello from greeting!")
}

func woo() {
	fmt.Println("4th - Hello from woo!")
}
