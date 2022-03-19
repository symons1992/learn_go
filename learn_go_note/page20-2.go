package main

func main() {
	x := 100
	println(&x)

	x := 200
	println(&x, x)
}
