package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	First string
	Age   int
}

func main() {
	u1 := User{
		First: "Jessica",
		Age:   30,
	}

	u2 := User{
		First: "Russell",
		Age:   2,
	}

	u3 := User{
		First: "Ruby",
		Age:   1,
	}

	users := []User{u1, u2, u3}

	fmt.Println(users)

	// your code goes here

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(string(bs))
}
