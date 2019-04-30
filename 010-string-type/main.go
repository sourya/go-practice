package main

import "fmt"

func main() {
	s := "If you get up in the morning and think the future is going to be better, it is a bright day."
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	// slice of byte
	bs := []byte(s) // conversion string to slice of byte
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	// byte is an alias for unit8
	// returns ASCII coding scheme

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	} // prints utf-8 characters of string

	fmt.Println("")

	for i, v := range s {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}
}
