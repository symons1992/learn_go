package main

import "fmt"

func main() {
	const (
		x = iota // 0
		y
		z
	)

	const (
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
	)

	const (
		_, _ = iota, iota * 10
		a, b
		c, d
	)

	fmt.Printf("KB = %v\n", KB)
	fmt.Printf("MB = %v\n", MB)
	fmt.Printf("GB = %v\n", GB)
}
