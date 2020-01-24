package main

import (
	"fmt"
	"time"
)

func main() {
	bd := 1989
	yr, _, _ := time.Now().Date()

	for {
		if bd > yr {
			break
		}
		fmt.Println(bd)
		bd++
	}

}
