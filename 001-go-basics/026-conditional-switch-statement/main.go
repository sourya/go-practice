package main

import (
	"fmt"
)

func main() {
	switch {
	case 1 == 2:
		fmt.Println("this should not print")
	default:
		fmt.Println("this is default")
	}
}
