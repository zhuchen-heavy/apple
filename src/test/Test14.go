package main

import "fmt"

/**
 * <p>
 * 语言切片：参考：https://www.runoob.com/go/go-slice.html
 * Go 语言切片是对数组的抽象。 ---> Go中的数组是定长数组，语言切片是非定长数组，跟Redis的自定义的结构体一致
 * Go 数组的长度不可改变，在特定场景中这样的集合就不太适用，Go中提供了一种灵活，功能强悍的内置类型切片("动态数组"),
 * 与数组相比切片的长度是不固定的，可以追加元素，在追加时可能使切片的容量增大。
 * </p>
 * @author: zhu.chen
 * @date: 2020/8/5
 * @version: v1.0.0
 */

/**
语法
var slice1 []type = make([]type, len)  ==== slice1 := make([]type, len)
make([]T, length, capacity) ： 切片的容量和长度
*/
func main() {
	// len() 和 cap() 函数
	var numbers = make([]int, 3, 5)
	printSlice(numbers)
	if numbers != nil {
		fmt.Printf("切片是非空的")
	}
	fmt.Println("------------")

	// 空切片
	var numbers1 []int
	printSlice(numbers1)
	if numbers1 == nil {
		fmt.Printf("切片是空的")
	}
	fmt.Println("------------")

	// 切片截取
	/* 创建切片 */
	numbers3 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	printSlice(numbers3)
	/* 打印原始切片 */
	fmt.Println("numbers3 ==", numbers3)

	/* 打印子切片从索引1(包含) 到索引4(不包含)*/
	fmt.Println("numbers3[1:4] ==", numbers3[1:4])

	/* 默认下限为 0*/
	fmt.Println("numbers3[:3] ==", numbers3[:3])

	/* 默认上限为 len(s)*/
	fmt.Println("numbers3[4:] ==", numbers3[4:])

	numbers4 := make([]int, 0, 5)
	printSlice(numbers4)

	/* 打印子切片从索引  0(包含) 到索引 2(不包含) */
	number5 := numbers[:2]
	printSlice(number5)

	/* 打印子切片从索引 2(包含) 到索引 5(不包含) */
	number6 := numbers[2:5]
	printSlice(number6)

	fmt.Println("---------------")
	slice()

}

func printSlice(x []int) {
	fmt.Printf("len=%d cap=%d slice=%v\n", len(x), cap(x), x)
}

func slice() {
	var numbers []int
	printSlice(numbers)

	/* 允许追加空切片 */
	numbers = append(numbers, 0)
	printSlice(numbers)

	/* 向切片添加一个元素 */
	numbers = append(numbers, 1)
	printSlice(numbers)

	/* 同时添加多个元素 */
	numbers = append(numbers, 2, 3, 4)
	printSlice(numbers)

	/* 创建切片 numbers1 是之前切片的两倍容量*/
	numbers1 := make([]int, len(numbers), (cap(numbers))*2)

	/* 拷贝 numbers 的内容到 numbers1 */
	copy(numbers1, numbers)
	printSlice(numbers1)
}
