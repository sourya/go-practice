package main

import "fmt"

func main() {
	e := 19 // keep scope narrow
	// "\t" inserts tab
	// "%#x" is for hexadecimal the hash adds 0x in front
	fmt.Printf("%d\t%b\t%#x\a", e, e, e)
}
