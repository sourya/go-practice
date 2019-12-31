package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Fprintln(os.Stdout, "Testing")
	io.WriteString(os.Stdout, "Russell & Ruby")
}
