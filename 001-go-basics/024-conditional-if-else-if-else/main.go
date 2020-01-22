package main

import "fmt"

func main() {
	x := 88
	if x == 23 {
		fmt.Println("our value was 23")
	} else if x == 22 {
		fmt.Println("our value was 22")
	} else if x == 21 {
		fmt.Println("our value was 21")
	} else {
		fmt.Printf("our value was weird: %v", x)
	}
}
