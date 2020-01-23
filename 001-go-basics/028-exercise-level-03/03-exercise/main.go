package main

import (
	"fmt"
	"time"
)

func main() {
	yr, _, _ := time.Now().Date()
	count := 0
	for i := 1989; i <= yr; i++ {
		fmt.Printf("age %v in the year %v\n", count, i)
		count++
	}
}
