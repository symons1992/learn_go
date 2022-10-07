package main

func main() {
	a := 1
	++a    // 语法错误：unexpected ++ （不能前置）

	if (a++) > 1 {  // 语法错误：unexpected ++， expecting ） （语句不能作为表达式使用）
	}

	p := &a
	*p++
	println(a)
}
