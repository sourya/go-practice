package main

import (
	"fmt"
	dog2 "github.com/brudnak/go-practice/000-old-practice/001-old-practice-12-31-2019/099-exercise-level-thirteen/01-exercise/dog"
)

type canine struct {
	name string
	age  int
}

func main() {
	fido := canine{
		name: "Fido",
		age:  dog2.Years(10),
	}
	fmt.Println(fido)
	fmt.Println(dog2.YearsTwo(20))
}
