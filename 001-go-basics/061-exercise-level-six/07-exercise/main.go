package main

import "fmt"

func main() {

	g := greeting
	g("Andrew")
}

func greeting(s string) {
	fmt.Println("Hello,", s)
}


