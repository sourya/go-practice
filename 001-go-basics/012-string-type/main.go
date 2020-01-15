package main

import "fmt"

func main() {

	s := `line 1
line 2`

	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	fmt.Println(s)

}
