package main

type data struct {
	x int
}

func main() {
	d := data{100}
	var t interface{} = d

	p := &t.(data)
	t.(data).x = 200
}
