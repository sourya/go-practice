package main

import "fmt"

func main() {
	count := 1
	for i := 65; i <= 90; i++ {
		fmt.Println(count)
		count++
		for j := 1; j <= 3; j++ {
			fmt.Printf("\t%#U\n", i)
		}
	}
}
