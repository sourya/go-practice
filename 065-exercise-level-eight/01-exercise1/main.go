package main

import (
	"fmt"
)

type user struct {
	first string
	age   int
}

func main() {
	u1 := user{
		first: "Andrew",
		age:   29,
	}

	u2 := user{
		first: "Jessie",
		age:   29,
	}

	u3 := user{
		first: "Russell",
		age:   2,
	}

	users := []user{u1, u2, u3}

	fmt.Println(users)

	// your code goes here
}
