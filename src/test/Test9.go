package main

import "fmt"

/**
变量的作用域
1：Go 语言程序中全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑。
*/

var g int

func main() {
	// 声明局部变量
	var a, b, c int

	a = 10
	b = 20
	c = a + b
	g = a + b
	var g int = 10
	// %d：会进行类似于Java中log的替换
	fmt.Printf("结果： a = %d, b = %d and c = %d and g = %d\n", a, b, c, g)
}
