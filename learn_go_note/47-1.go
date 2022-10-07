package main

import "log"

func main() {
	x := 10

	if err := check(x); err != nil {
		log.Fatalln(err)
	}

	x++
	println(x)
}
