package main

func main() {
	var a func(int, string)
	var b func(string, int)

	b = a

	b("s", 1)
}
