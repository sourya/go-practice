package main

import (
	"fmt"
	"time"
)

func main() {
	born := 1989
	age := 0
	currentTime := time.Now()
	currentYear := currentTime.Year()

	for born <= currentYear {

		fmt.Printf("Year: %v \t\t age: %v\n", born, age)
		born++
		age++
	}
}
