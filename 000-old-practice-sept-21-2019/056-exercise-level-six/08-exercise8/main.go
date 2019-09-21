package main

import "fmt"

func main() {
	x := speak()

	y := x()

	fmt.Println(y)
}

func speak() func() string {
	return func() string {
		return "Hello from returned func!"
	}
}
