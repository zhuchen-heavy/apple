package main

import (
	"fmt"
	"runtime"
	"strconv"
	"strings"
)

/**
声明全局变量
*/
var x, y int
var (
	a int
	b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

func main() {

	g, h := 123, "hello"
	fmt.Println(x, y, a, b, c, d, e, f, g, h)

	/**
	 	1：若变量声明未被使用；若包导入未被使用，则编译器会报错
		.\test2.go:3:8: imported and not used: "fmt"
		强制要求程序员删除。
	*/
	var i int = 3
	fmt.Println(i)
	go run("test")
	fmt.Println("go parent id is ", Goid())
}

func run(s string) {
	fmt.Println("go child id is ", Goid())
	fmt.Println(s)
}

/**
获取go的携程id
*/
func Goid() int {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic recover:panic info:%v", err)
		}
	}()

	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idField := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, err := strconv.Atoi(idField)
	if err != nil {
		panic(fmt.Sprintf("cannot get goroutine id: %v", err))
	}
	return id
}
