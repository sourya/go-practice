package main

import (
	"encoding/json"
	"fmt"
)

type user struct {
	First string
	Age   int
}

func main() {
	u1 := user{
		First: "Andrew",
		Age:   29,
	}

	u2 := user{
		First: "Jessie",
		Age:   29,
	}

	u3 := user{
		First: "Russell",
		Age:   2,
	}

	u4 := user{
		First: "Ruby",
		Age:   1,
	}

	users := []user{u1, u2, u3, u4}

	fmt.Println(users)

	bs, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(bs))
}
