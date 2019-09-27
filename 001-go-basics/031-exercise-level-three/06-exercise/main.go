package main

import (
	"fmt"
	"time"
)

func main() {
	cy := time.Now().Year()
	yb := 1989
	ca := cy - yb
	if ca > 21 {
		fmt.Println("You can drink!")
	} else {
		fmt.Println("No minors allowed!")
	}
}
