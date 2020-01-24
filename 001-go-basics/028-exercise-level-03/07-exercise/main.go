package main

import "fmt"

func main() {
	for i := 0; i <= 75; i++ {
		if i % 3 == 0 {
			fmt.Println(i)
		} else if i % 2 == 0 {
			fmt.Println("yeet")
		} else {
			fmt.Println("meep")
		}
	}
}
