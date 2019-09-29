package main

import (
	"fmt"
	"golang.org/x/crypto/bcrypt"
)

func main() {
	s := "password123"

	bs, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.MinCost)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println(s)
	fmt.Println(bs)

	wrongPword1 := "qwerty"

	err2 := bcrypt.CompareHashAndPassword(bs, []byte(wrongPword1))
	if err2 != nil {
		fmt.Println("YOU CAN'T LOGIN")
	} else {
		fmt.Println("You're logged in")
	}
}
