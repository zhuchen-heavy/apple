package main

import "fmt"

/**
 * <p>
 * 数组变量
 * Java，go，c等数组都是从0开始。
 * </p>
 * @author: zhu.chen
 * @date: 2020/8/5
 * @version: v1.0.0
 */
func main() {

	var n [10]int
	var i, j int
	for i = 0; i < 10; i++ {
		n[i] = i + 100
	}
	for j = 0; j < 10; j++ {
		fmt.Printf("Element[%d] = %d\n", j, n[j])
	}

	fmt.Println("-----------")

	array()

}

func array() {
	var a = [3][4]int{
		{0, 1, 2, 3},   /*  第一行索引为 0 */
		{4, 5, 6, 7},   /*  第二行索引为 1 */
		{8, 9, 10, 11}, /* 第三行索引为 2 */
	}
	var i, j int
	for i = 0; i < 3; i++ {
		for j = 0; j < 4; j++ {
			fmt.Printf("a[%d][%d] = %d\n", i, j, a[i][j])
		}
	}
}
