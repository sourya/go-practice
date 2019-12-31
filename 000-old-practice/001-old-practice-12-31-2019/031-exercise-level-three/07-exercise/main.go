package main

import "fmt"

func main() {

	bouncer(21)
	bouncer(18)
	bouncer(16)
}

func bouncer(age int) {
	if age >= 21 {
		fmt.Println("You can drink!")
	} else if age >= 18 {
		fmt.Println("You can come in, but you're wearing \"X's\" on your hands.")
	} else {
		fmt.Println("No minors allowed!")
	}
}
