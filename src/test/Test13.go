package main

import "fmt"

/**
 * <p>
 * 结构体
 * 类似于Java等面向对象语言中的对象
 *
 * 结构体语法：
 *	type struct_variable_type struct {
 *	   member definition
 *	   member definition
 *	   ...
 *	   member definition
 *	}
 * </p>
 * @author: zhu.chen
 * @date: 2020/8/5
 * @version: v1.0.0
 */

type Books struct {
	title   string
	author  string
	subject string
	bookId  int
}

func main() {
	// 创建一个新的结构体
	fmt.Println(Books{"go", "www.baidu.com", "go教程", 111})

	// 也可以使用 key => value 格式
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com", subject: "Go 语言教程", bookId: 6495407})

	// 忽略的字段为 0 或 空   基本类型为0，引用类型为nil
	fmt.Println(Books{title: "Go 语言", author: "www.runoob.com"})

	fmt.Println("----------")

	structTest1()
	fmt.Println("----------")

	structTest2()
}

/**
访问结构体成员
*/
func structTest1() {
	var book1 Books
	//var book2 Books
	book1.title = "Go语言"
	book1.author = "xxx"
	book1.subject = "go教程"
	book1.bookId = 111
	/* 打印 Book1 信息 */
	fmt.Printf("Book 1 title : %s\n", book1.title)
	fmt.Printf("Book 1 author : %s\n", book1.author)
	fmt.Printf("Book 1 subject : %s\n", book1.subject)
	fmt.Printf("Book 1 book_id : %d\n", book1.bookId)
}

/**
&：返回变量存储地址。示例：&a：将给出变量的实际地址。
*：指针变量。 示例：*a：是一个指针变量。
*/
func structTest2() {
	var book1 Books
	book1.title = "book1"
	book1.author = "zuozhe"
	book1.bookId = 1
	// 传递变量在内存中的地址
	changeBook(&book1)
	// fmt.Println是值传递，需要将book1的值打印出来
	fmt.Println(book1)
}

// 函数体接收指针
func changeBook(book *Books) {
	book.title = "book1_change"
}

// 直接传值是无法别赋值的
//func changeBook(book Books) {
//	book.title = "book1_change"
//}
