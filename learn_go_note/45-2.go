package main

func main() {
	x := 10

	if xinit(); x == 0 { // 优先执行xinit函数
		println("a")
	}

	if a, b := x+1, x+10; a < b {
		println(a)
	} else {
		println(b)
	}
}
