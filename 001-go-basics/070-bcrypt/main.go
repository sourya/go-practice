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

}
