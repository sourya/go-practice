package main

import "fmt"

func main() {
	a := (19 == 19)
	b := (2019 <= 2020)
	c := (1989 >= 2019)
	d := (a != b)
	e := (1 > 0)
	f:= ("a" < "b")
	fmt.Println(a, b, c, d, e, f)
}