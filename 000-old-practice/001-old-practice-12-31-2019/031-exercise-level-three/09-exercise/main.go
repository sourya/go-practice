package main

import "fmt"

func main() {

	sportMessage("snowboarding")
	sportMessage("surfing")
	sportMessage("football")
	sportMessage("watching Netflix")
}

func sportMessage(favSport string) {
	switch favSport {
	case "snowboarding":
		fmt.Println("Go to the mountains!")
	case "surfing":
		fmt.Println("Go to the ocean!")
	case "football":
		fmt.Println("Go to the stadium!")
	default:
		fmt.Println("Go outside!")
	}
}
