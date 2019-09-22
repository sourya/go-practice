package main

import (
	"fmt"
)

func main() {
	s := "Ruby Genevieve Brudnak"
	fmt.Println(s)
	fmt.Printf("%T\n", s)

	// Slice of bytes
	bs := []byte(s)
	fmt.Println(bs)
	// output:
	// [82 117 98 121 32 71 101 110 101 118 105 101 118 101 32 66 114 117 100 110 97 107]
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U", s[i])
	}

	fmt.Println("")

	for i, v := range s {
		fmt.Printf("at index %d\t we have hex %#x\n", i, v)

	}
}
