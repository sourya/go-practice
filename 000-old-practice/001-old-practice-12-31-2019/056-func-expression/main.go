package main

import "fmt"

func main() {

	f := func() {
		fmt.Println("my first func expression")
	}

	f2 := func(x int) {
		fmt.Println("Age:", x)
	}

	fmt.Printf("Type of 'f' variable: %T\n", f)

	f()
	f2(29)
}
