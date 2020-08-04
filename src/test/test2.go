package main

import "fmt"

/**
声明全局变量
*/
var x, y int
var (
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

func main() {

	g, h := 123, "hello"
	fmt.Println(x, y, a, b, c, d, e, f, g, h)

	/**
	 	1：若变量声明未被使用；若包导入未被使用，则编译器会报错
		.\test2.go:3:8: imported and not used: "fmt"
		强制要求程序员删除。
	*/
	var i int = 3
	fmt.Println(i)

	go run("test")

}

func run(s string) {

	fmt.Println(s)
}
