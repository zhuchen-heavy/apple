package main

import (
	"fmt"
)

/**
go的条件语句
*/

func main() {

	condition1()
	fmt.Println("--------------")

	condition2()
	fmt.Println("--------------")

	condition3()
	fmt.Println("--------------")

	condition4()
	fmt.Println("--------------")

	condition5()

}

/**
select语句：TODO
select {
    case communication clause  :
       statement(s);
    case communication clause  :
       statement(s);
    default :
       statement(s);
}
*/
func condition5() {
	var c1, c2, c3 chan int
	var i1, i2 int
	select {
	case i1 = <-c1:
		fmt.Printf("received ", i1, " from c1\n")
	case c2 <- i2:
		fmt.Printf("sent ", i2, " to c2\n")
	case i3, ok := (<-c3): // same as: i3, ok := <-c3
		if ok {
			fmt.Printf("received ", i3, " from c3\n")
		} else {
			fmt.Printf("c3 is closed\n")
		}
	default:
		fmt.Printf("no communication\n")
	}
}

/**
fallthrough：会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true。
--- fallthrough，类似于Java中的每个case都不加break，case匹配到的语句后的语句都会执行
*/
func condition4() {
	var grade = "B"
	var marks = 90
	switch marks {
	case 90:
		grade = "A"
		fallthrough
	case 80:
		grade = "B"
		fallthrough
	case 50, 60, 70:
		grade = "C"
	}
	// grade is  C
	fmt.Println("grade is ", grade)
}

/**
type...switch
switch x.(type){
    case type:
       statement(s);
    case type:
       statement(s);

	default:
		statement(s);
}
*/
func condition3() {
	var x interface{}
	// :=：声明并赋值，并且系统自动推断类型，不需要var关键字
	switch i := x.(type) {
	case nil:
		fmt.Println(" x 的类型 :%T", i)
	case int:
		fmt.Println("x 是 int 型")
	case float64:
		fmt.Println("x 是 float64 型")
	case func(int) float64:
		fmt.Println("x 是 func(int) 型")
	case bool, string:
		fmt.Println("x 是 bool 或 string 型")
	default:
		fmt.Println("未知型")
	}
}

/**
switch语句：不使用fallthrough，类似于Java中的每个case都加了break
switch var1 {
    case val1:
        ...
    case val2:
        ...
    default:
        ...
}
*/
func condition2() {

	var grade = "B"
	var marks = 90
	switch marks {
	case 90:
		grade = "A"
	case 80:
		grade = "B"
	case 50, 60, 70:
		grade = "C"
	}
	// grade is  A
	fmt.Println("grade is ", grade)

}

/**
if语法：
if 布尔表达式 {//在布尔表达式为 true 时执行 }
*/
func condition1() {

	var a int = 10

	if a < 20 {
		if a == 10 {
			fmt.Println("a等于10")
		} else {
			fmt.Println("a小于20.")
		}
	} else {
		fmt.Println("a不小于20.")
	}

}
