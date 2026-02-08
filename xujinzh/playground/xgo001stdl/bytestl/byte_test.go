package bytestl_test

import (
	"bytes"
	"fmt"
	"testing"
)

// bytes标准库

// bytes包提供了对字节切片进行读写操作的一系列函数
// 包括基本处理函数、比较函数、后缀检查函数、索引函数、分割函数、大小写处理函数和子切片处理函数等

// 常用函数
// 转换
// func ToUpper(s []byte) []byte 将s中的所有字符修改为大写格式返回
// func ToLower(s []byte) []byte 将s中的所有字符修改为小写格式返回
// func ToTitle(s []byte) []byte 将s中的所有字符修改为标题格式返回
// func ToUpperSpecial(_case unicode.SpecialCase, s []byte) []byte 使用指定的映射表将s中的所有字符修改为大写格式返回
// func ToLowerSpecial(_case unicode.SpecialCase, s []byte) []byte 使用指定的映射表将s中的所有字符修改为小写格式返回
// func ToTitleSpecial(_case unicode.SpecialCase, s []byte) []byte 使用指定的映射表将s中的所有字符修改为标题格式返回
// func Title(s []byte) []byte 将s中的所有单词的首字符修改为Title格式返回（不能很好的处理以Unicode标点符号分隔的单词）

// 字符串转换
func TestConvert(t *testing.T) {
	// 字符串小写转大写
	b := []byte("hainan")
	a := bytes.ToUpper(b)
	fmt.Printf("b: %s\n", b)
	fmt.Printf("a: %s\n", a)
	// 字符串大写转小写
	c := []byte("SHANGHAI")
	d := bytes.ToLower(c)
	fmt.Printf("c: %s\n", c)
	fmt.Printf("d: %s\n", d)
	// 字符串转为标题
	e := []byte("attention is all you need")
	f := bytes.ToTitle(e)
	fmt.Printf("e: %q\n", e)
	fmt.Printf("f: %q\n", f)
}

// 字符串比较
// func Compare(a, b []byte) int 比较两个[]byte, nil参数相当于空[]byte
// a < b 返回 -1；a == b 返回 0； a > b 返回1
// func Equal(a, b []byte) bool 判断 a, b 是否相等，nil 参数相当于空 []byte
// func EqualFold(s, t []byte) bool 判断 s, t 是否相似，忽略大写、小写、标题三种格式的区别

// 测试比较两个字符串
func TestCompare(t *testing.T) {
	a := []byte("beijing")
	b := []byte("BeiJing")
	// 比较字符串大小
	n := bytes.Compare(a, b)
	fmt.Printf("n: %v\n", n)
	// 比较字符串是否相等
	bt := bytes.Equal(a, b)
	fmt.Printf("bt: %v\n", bt)
	// 不区分大小写比较字符串是否相等
	btFold := bytes.EqualFold(a, b)
	fmt.Printf("btFold: %v\n", btFold)
}