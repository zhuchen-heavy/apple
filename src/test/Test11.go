package main

import "fmt"

/**
 * <p>
 * 语言指针：https://zhuanlan.zhihu.com/p/70431669
 * # 变量是一种使用方便的占位符，用于引用计算机内存地址。
 * # Go 语言的取地址符是 &，放到一个变量前使用就会返回相应变量的内存地址。
 * Go中的指针就是Java里面的引用传递，Java所有的引用类型都是引用传递，也就都是指针
 * </p>
 * @author: zhu.chen
 * @date: 2020/8/5
 * @version: v1.0.0
 */

func main() {
	a := 10
	// 10 ---> %d 进行标准的十进制格式化。
	fmt.Printf("变量的值: %d\n", a)
	// 变量的地址: c00000a0e8 ---> %x 提供十六进制编码。
	fmt.Printf("变量的地址: %x\n", &a)
	fmt.Println("--------")

	test11()
	fmt.Println("--------")

	test12()
	fmt.Println("--------")

	test13()
}

/**
语法：var var_name *var-type
var-type 为指针类型，var_name 为指针变量名，* 号用于指定变量是作为一个指针
*/
func test11() {
	var a int = 20 /* 声明实际变量 */
	var ip *int    /* 声明指针变量 */

	ip = &a /* 指针变量的存储地址 */

	fmt.Printf("a 变量的地址是: %x\n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)
}

/**
当一个指针被定义后没有分配到任何变量时，它的值为 nil。
nil 指针也称为空指针。
nil在概念上和其它语言的null、None、nil、NULL一样，都指代零值或空值。
一个指针变量通常缩写为 ptr。
*/
func test12() {
	var ptr *int
	fmt.Printf("ptr 的值为 : %x\n", ptr)
}

/**
指针数组
*/
const MAX int = 3

func test13() {
	a := []int{10, 100, 200}
	var i int
	for i = 0; i < MAX; i++ {
		fmt.Printf("a[%d] = %d\n", i, a[i])
	}
}
