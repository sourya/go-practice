package main

import "fmt"

func main() {
	x := 19
	if x == 21 {
		fmt.Println("our value was 21")
	} else if x == 22 {
		fmt.Println("our value was 22")
	} else {
		fmt.Println("our value was not 21 or 22")
	}
}
