package main

import "fmt"

func main() {

	favSport := "flying"

	switch favSport {
	case "snowboarding":
		fmt.Println("go to the mountains!")
	case "swimming":
		fmt.Println("get to the pool!")
	case "flying":
		fmt.Println("get to the chopper!")
	}
}
