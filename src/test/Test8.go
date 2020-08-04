package main

import "fmt"

/**
go的数组
语法：var variable_name [SIZE] variable_type
示例：定义一维数组：var balance [10] float32
初始化数组：
	var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}
	 var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0} // go会根据元素的数量来设置大小
赋值： balance[4] = 50.0

已经看到：
https://www.runoob.com/go/go-arrays.html
*/

func main() {

	var n [10]int
	var i, j int
	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}

	for j = 0; j < 10; j++ {
		fmt.Println(j, n[j])
	}

}
