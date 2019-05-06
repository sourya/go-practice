package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		fmt.Println("The outer loop: %d\t The inner loop: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Println("The outer loop: %d\t The inner loop: %d\n", i, j)
		}
	}
}
