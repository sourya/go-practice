package main

import "fmt"

func main() {
	x := 935

	//if x == 10 {
	//	fmt.Println("the value was 19")
	//} else {
	//	fmt.Println("the value was not 19")
	//}

	if x == 19 {
		fmt.Println("the value was 19")
	} else if x == 20 {
		fmt.Println("the value was 20")
	} else if x == 21 {
		fmt.Println("the value was 21")
	} else if x == 22 {
		fmt.Println("the value was 22")
	} else if x == 23 {
		fmt.Println("the value was 23")
	} else {
		fmt.Println("the value was not 19, 20, 21, 22, or 23")
	}
}
