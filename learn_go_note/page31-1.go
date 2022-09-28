package main

func add(x, y int) int {
	return x + y
}

func main() {
	var x int = 100
	var y int64 = x

	add(x, y)
}
