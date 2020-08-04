package main

import (
	"fmt"
)

/**
go运算符的示例
大多数都是跟Java一样的
*/

func main() {

	math()

	relation()

	position()

	special()
}

/**
特殊运算符
&：返回变量存储地址。示例：&a：将给出变量的实际地址。
*：指针变量。 示例：*a：是一个指针变量。
*/
func special() {
	var a int = 4
	var b int32
	var c float32
	var ptr *int

	/* 运算符实例 */
	fmt.Printf("第 1 行 - a 变量类型为 = %T\n", a)
	fmt.Printf("第 2 行 - b 变量类型为 = %T\n", b)
	fmt.Printf("第 3 行 - c 变量类型为 = %T\n", c)

	/*  & 和 * 运算符实例 */
	ptr = &a /* 'ptr' 包含了 'a' 变量的地址 */
	fmt.Printf("a 的值为  %d\n", a)
	fmt.Printf("*ptr 为 %d\n", *ptr)
}

/**
位运算符
*/
func position() {

	var A = 00111100

	var B = 00001101

	fmt.Println("A&B 为 ", A&B)

}

/**
关系运算符
*/
func relation() {
	var a bool = true
	var b bool = false

	if a == b {
		fmt.Printf("第一行 - 条件为 true\n")
	} else {
		fmt.Printf("第一行 - 条件为 false\n")
	}

	if a || b {
		fmt.Printf("第二行 - 条件为 true\n")
	} else {
		fmt.Printf("第二行 - 条件为 false\n")
	}

	/* 修改 a 和 b 的值 */
	a = false
	b = true
	if a && b {
		fmt.Printf("第三行 - 条件为 true\n")
	} else {
		fmt.Printf("第三行 - 条件为 false\n")
	}
	if !(a && b) {
		fmt.Printf("第四行 - 条件为 true\n")
	}

}

/**
算术运算符
*/
func math() {
	var a = 21
	var b = 10
	var c int

	c = a + b
	fmt.Printf("第一行 - c 的值为 %d\n", c)
	c = a - b
	fmt.Printf("第二行 - c 的值为 %d\n", c)
	c = a * b
	fmt.Printf("第三行 - c 的值为 %d\n", c)
	c = a / b
	fmt.Printf("第四行 - c 的值为 %d\n", c)
	c = a % b
	fmt.Printf("第五行 - c 的值为 %d\n", c)
	a++
	fmt.Printf("第六行 - a 的值为 %d\n", a)
	a = 21 // 为了方便测试，a 这里重新赋值为 21
	a--
	fmt.Printf("第七行 - a 的值为 %d\n", a)
}
