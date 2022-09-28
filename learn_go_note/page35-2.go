package main

import "fmt"

func main() {
	var a struct {
		x int    `x`
		s string `s`
	}

	var b struct {
		x int
		s string
	}

	b = a

	fmt.Println(b)
}
