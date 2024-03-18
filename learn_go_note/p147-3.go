package main

type data struct {
	x int
}

func main() {
	d := data{100}
	var t interface{} = &d

	t.(*data).x = 200
	println(t.(*data).x)
}
