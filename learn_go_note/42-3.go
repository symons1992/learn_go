package main

func main() {
	x := 10
	p := &x

	p++
	var p2 *int = p + 1

	p2 = &x
	println(p == p2)
}
