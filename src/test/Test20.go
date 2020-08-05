package main

import (
	"fmt"
	"time"
)

/**
 * <p>
 * 并发处理：只需要通过 go 关键字来开启 goroutine 即可。
 * goroutine 是轻量级线程，goroutine 的调度是由 Golang 运行时进行管理的。
 * Go 允许使用 go 语句开启一个新的运行期线程， 即 goroutine，以一个不同的、新创建的 goroutine 来执行一个函数。 同一个程序中的所有 goroutine 共享同一个地址空间。
 * 语法：go 函数名( 参数列表 )
 * e.g
 * f(x, y, z) --开启协程处理--> go f(x, y, z)
 * </p>
 * @author: zhu.chen
 * @date: 2020/8/5
 * @version: v1.0.0
 */

func main() {
	go say("world")
	say("hello")
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}
