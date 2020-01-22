package main

import "fmt"

func main() {
	for i := 0; i < 100; i++ {
		if i%16 == 0 {
			fmt.Println(i)
		}
	}
}
