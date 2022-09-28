package main

func main() {
	type data int
	var d data = 10

	var x int = d
	println(x)

	println(d == x)
}
