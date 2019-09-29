package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("Hello from line 7 of func main")
	fmt.Println("Hello from line 8 of func main")
	fmt.Println("Hello from line 9 of func main")
}

func foo() {
	fmt.Println("Hello from line 6, called last thanks to the DEFER KEYWORD")
}
