package main

import "fmt"

func main() {
	for i := 0; i <= 75; i++ {
		if i%3 == 0 {
			fmt.Println(i)
		}
	}
}
