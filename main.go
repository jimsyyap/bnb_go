package main

import (
	"fmt"
)

func main() {
	saySomething("Hello, World!")
	// fmt.Println(saySomething("Hello, World!"))
}

func saySomething(s string) string {
	fmt.Println(s)
	return s
}
