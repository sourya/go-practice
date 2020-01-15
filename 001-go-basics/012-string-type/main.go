package main

import "fmt"

func main() {

	s := `line 1
line 2`

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	fmt.Println(s)

	for i := 0; i < len(s); i++ {
		fmt.Printf("%#U ", s[i])
	}

	fmt.Println()
	for i, v := range s {
		fmt.Printf("at index position %d we have hex %#x\n", i, v)
	}
}
