package main

import "fmt"

/**
 * <p>
 * 类型转换
 * 语法：type_name(expression) ：type_name 为类型，expression 为表达式。
 * </p>
 * @author: zhu.chen
 * @date: 2020/8/5
 * @version: v1.0.0
 */

func main() {
	var sum int = 7
	var count int = 5
	var mean float32
	mean = float32(sum) / float32(count)
	fmt.Printf("sum 的值为: %f\n", float32(sum))
	fmt.Printf("mean 的值为: %f\n", mean)
}
