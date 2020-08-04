package main

import "fmt"

/**
go的函数，function_name：函数名。[parameter list]：参数列表。[return_types]：返回值类型
func function_name( [parameter list] ) [return_types] {
   函数体
}
1：默认情况下，Go 语言使用的是值传递，即在调用过程中不会影响到实际参数。
*/

func main() {

	fmt.Println("1和2的max值为 : ", max(1, 2))
	fmt.Println("------------")

	a, b := swap("google", "facebook")
	fmt.Println(a, b)
	fmt.Println("------------")

	nextNumber := getSequence()
	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println("------------")

	var c1 Circle
	c1.radius = 10.00
	fmt.Println("圆的面积 : ", c1.getArea())
}

/**
定义结构体，类似于Java中的实体类，也就是c++中的struct
*/
type Circle struct {
	radius float64
}

//  method 属于 Circle 类型对象中的方法
// 可以通过 "." 来调用struct中的方法
func (c Circle) getArea() float64 {
	return 3.14 * c.radius * c.radius
}

/**
创建了函数 getSequence() ，返回另外一个函数。该函数的目的是在闭包中递增 i 变量
*/
func getSequence() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

/**
比较交换
*/
func swap(x, y string) (string, string) {
	return y, x
}

/**
求最大值
*/
func max(num1, num2 int) int {
	var result int
	if num1 > num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}
