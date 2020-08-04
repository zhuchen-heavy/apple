package main

import "fmt"

/**
go中循环语句
*/

func main() {
	cycle()
	fmt.Println("------------")

	result := cycle1()
	fmt.Println("result is ", result)
	fmt.Println("------------")

	cycle2()

}

/**
For-each range 循环
*/
func cycle2() {
	strings := []string{"google", "facebook"}
	for i, s := range strings {
		/**
		0 google
		1 facebook
		*/
		fmt.Println(i, s)
	}

	numbers := [6]int{1, 2, 3, 5}
	for i, x := range numbers {
		/**
		未赋值的为0
		第 0 位 x 的值 = 1
		第 1 位 x 的值 = 2
		第 2 位 x 的值 = 3
		第 3 位 x 的值 = 5
		第 4 位 x 的值 = 0
		第 5 位 x 的值 = 0
		*/
		fmt.Printf("第 %d 位 x 的值 = %d\n", i, x)
	}

}

/*
for condition { }
*/
func cycle1() int {
	sum := 1
	for sum <= 10 {
		sum += sum
	}
	// 1+1 -> 2+2 -> 4+4 -> 8+8
	return sum
}

/**
for init; condition; post { }
*/
func cycle() {
	sum := 0
	for i := 0; i <= 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
