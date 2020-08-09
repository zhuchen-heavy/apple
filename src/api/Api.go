package api

import "fmt"

/**
 * <p>
 * Go使用首字母大小写来区别该方法的访问权限。
 * 方法首字母大写：表示该方法可以给外部包引用。---> 对应Java中的public
 * 方法首字母大写：表示该方法不可以给外部包引用。---> 对应Java中的private
 * </p>
 *
 * @author zhu.chen
 * @date 2020/8/9
 * @version 1.0
 */
func Api() {
	fmt.Println("this is a api.")
}
