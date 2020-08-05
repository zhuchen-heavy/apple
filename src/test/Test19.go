package main

import "fmt"

/**
 * <p>
 * 错误处理
 * </p>
 * @author: zhu.chen
 * @date: 2020/8/5
 * @version: v1.0.0
 */

/**
语法：error类型是一个接口类型，这是它的定义：（该定义是Go中的定义，应用无需声明）
type error interface {
    Error() string
}
*/

type DivideError struct {
	dividee int
	divider int
}

func (de *DivideError) Error() string {
	strFormat := `
    Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
`
	return fmt.Sprintf(strFormat, de.dividee)
}
