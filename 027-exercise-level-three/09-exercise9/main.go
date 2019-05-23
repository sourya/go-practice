package main

import "fmt"

func main() {
	favSport := "snowboarding"
	switch favSport {
	case "surfing":
		fmt.Print("go to the ocean!")
	case "swimming":
		fmt.Print("go to the pool!")
	case "snowboarding":
		fmt.Print("go to the mountains!")
	case "skydiving":
		fmt.Print("go to the sky!")

	}
}
