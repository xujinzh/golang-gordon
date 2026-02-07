package errstl_test

import (
	"errors"
	"fmt"
	"testing"
)

// errors 标准库

// errors 包实现了操作错误的函数。
// go 语言使用 error 类型来返回函数执行过程中遇到的错误
// 如果返回的 error 值为 nil，则表示未遇到错误，否则 error 会返回一个字符串，说明遇到了什么错误
/*
type error interface {
	Error() string
}
*/

// error 不一定表示一个错误，它可以表示任何信息，比如 io 包中就用 error 类型的 io.EOF 表示数据读取结束，而不是遇到了什么错误

func TestError(t *testing.T) {
	s := ""
	var err error
	if s == "" {
		err = errors.New("字符串不能为空")
	} else {
		err = nil
	}
	fmt.Println(err.Error())
	fmt.Println(err)
}
