package main

import "fmt"

func main() {
	x := 0
	fmt.Println("x before", &x)
	fmt.Println("x before", x)
	fmt.Println("")
	foo(&x)
	fmt.Println("")
	fmt.Println("x after", &x)
	fmt.Println("x after", x)
	fmt.Println("")
}

func foo(y *int) {
	fmt.Println("y before", y)
	fmt.Println("y before", *y)
	fmt.Println("")
	*y = 19
	fmt.Println("y after", y)
	fmt.Println("y after", *y)
}
