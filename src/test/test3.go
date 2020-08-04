package main

import (
	"fmt"
	"unsafe"
)

/**
go语言常量
解决goland控制台中文乱码：https://www.cnblogs.com/motianlong/p/13427310.html
*/
func main() {

	//1：常量。语法：const identifier [type] = value
	// 显示声明类型
	const identifier string = "3"
	// 隐示声明类型
	//省略类型说明符 [type]，因为编译器可以根据变量的值来推断其类型。
	const identifier1 = "4"
	// 声明多个相同类型。语法：const c_name1, c_name2 = value1, value2
	const c_name1, c_name2 = "4", "5"

	// 2：常量做枚举
	const (
		unKnow = 0
		female = 1
		male   = 2
	)

	/**
	示例
	*/
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str"

	area = LENGTH * WIDTH

	// 面积为 : %d 50
	fmt.Println("面积为 : %d", area)
	fmt.Println("-----------------")
	// 1   false   str
	fmt.Println(a, " ", b, " ", c)
	fmt.Println("-----------------")

	test()
	fmt.Println("-----------------")

	test1()
	fmt.Println("-----------------")

}

/**
iota，特殊常量，可以认为是一个可以被编译器修改的常量。
iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。
*/
func test1() {
	const (
		a = iota
		b = iota
		c = iota
	)
	fmt.Println(a, " ", b, " ", c)
}

/**
常量示例
*/
func test() {
	const (
		a = "abc"
		b = len(a)
		c = unsafe.Sizeof(a)
	)
	// abc   3   16
	fmt.Println(a, " ", b, " ", c)
}
