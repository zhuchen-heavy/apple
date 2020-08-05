package main

import "fmt"

/**
 * <p>
 * 语言范围：range
 * Go 语言中 range 关键字用于 for 循环中迭代数组(array)、切片(slice)、通道(channel)或集合(map)的元素。
 * 在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对。
 * </p>
 * @author: zhu.chen
 * @date: 2020/8/5
 * @version: v1.0.0
 */

func main() {
	//这是我们使用range去求一个slice的和。使用数组跟这个很类似
	nums := []int{2, 3, 4}
	sum := 0
	// "_"：表示忽略这个变量
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum:", sum)

	// 在数组中range返回元素的索引和对应的值
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index:", i)
		}
	}

	// 在集合中range返回元素的key-value对
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k, v)
	}

	//range也可以用来枚举Unicode字符串。第一个参数是字符的索引，第二个是字符（Unicode的值）本身。
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
