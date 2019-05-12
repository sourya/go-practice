package main

import "fmt"

func main() {
	sum := 0
	bd := 1989
	for bd <= 2019{
		fmt.Printf("Years Alive: %d\n", bd)
		sum ++
		bd ++
	}
	sum -= 1
	fmt.Printf("Age / total years alive: %d", sum)
}
