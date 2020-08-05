package main

import "fmt"

/**
 * <p>
 * 接口。就是类似于Java的接口
 * </p>
 * @author: zhu.chen
 * @date: 2020/8/5
 * @version: v1.0.0
 */

/**
语法：
type interface_name interface {
   method_name1 [return_type]
   method_name2 [return_type]
   method_name3 [return_type]
   ...
   method_namen [return_type]
}

type struct_name struct {

}
// 实现接口方法
func (struct_name_variable struct_name) method_name1() [return_type] {
	// 方法实现
}
.....
func (struct_name_variable struct_name) method_namen() [return_type] {

}
*/
type Phone interface {
	call()
}
type NokiaPhone struct {
}

// 实现接口方法
func (nokiaPhone NokiaPhone) call() {
	// 方法实现
	fmt.Println("I am Nokia, I can call you!")
}

type IPhone struct {
}

func (iPhone IPhone) call() {
	fmt.Println("I am iPhone, I can call you!")
}

func main() {
	var phone Phone
	phone = new(NokiaPhone)
	phone.call()

	phone = new(IPhone)
	phone.call()
}
