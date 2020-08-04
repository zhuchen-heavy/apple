/**
 * package main表示一个可独立执行的程序，每个 Go 应用程序都包含一个名为 main 的包。
 */
package main

import (
	"fmt"
	"runtime"
)

// 示例代码：https://www.runoob.com/go/go-program-structure.html
func main() {
	// 输出 hello world
	fmt.Println("hello world")
	// 输出go的版本，跟cmd中go version的输出结果一致
	fmt.Println(runtime.Version())
	// 字符串拼接
	fmt.Println("hello " + "hello")
	// 常量声明 const 变量名称,变量名称 ... ,数据类型 = 变量值,变量值 ...
	const name1 string = "tom"
	const name2 string = "tom"
	fmt.Println(name1 + name2)

	// var 变量名称,变量名称 ... ,数据类型 = 变量值,变量值 ...
	var age1 int = 31
	var age2 int = 32
	fmt.Println(age1, age2)

	// 变量如果没有被使用到，则会出现编译错误，编译器会表明删除变量
	var b, c int = 1, 2
	fmt.Println(b, c)

	// var f string = "roob" 等价于 f := "roob"
	var f string = "roob"
	fmt.Println(f)

}
