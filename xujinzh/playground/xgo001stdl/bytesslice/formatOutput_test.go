package bytesslice

import (
	"fmt"
	"math/big"
	"os"
	"runtime"
	"strings"
	"testing"
)

type User struct {
	Id int64
}

func TestFormatOutput(t *testing.T) {
	user := &User{Id: 1}
	fmt.Printf("%v\n", user)
	fmt.Printf("%+v\n", user)
	fmt.Printf("%#v\n", user)
	fmt.Printf("%T\n", user)
	fmt.Printf("%q\n", user)
	fmt.Printf("%%\n")
}

func TestBoolFormatOutput(t *testing.T) {
	fmt.Printf("%t\n", true)
}

func TestDigitFormatOutput(t *testing.T) {
	n := 180
	fmt.Printf("%b\n", n) // 二进制
	fmt.Printf("%c\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%x\n", n)
	fmt.Printf("%X\n", n)
	// fmt.Printf("%f\n", n)
	fmt.Printf("%U\n", n)
	a := 96
	fmt.Printf("%q\n", a)
	fmt.Printf("%q\n", 0x4e2d)
}

func TestFloatFormatOutput(t *testing.T) {
	/*
	   %b 无小数部分、二进制指数的科学计数法
	   %e 科学计数法
	   %E 科学计数法
	   %f 有小数部分但无指数部分
	   %F 等价于 %f
	   %g 根据实际情况采用 %e 或 %f 格式，以获得更简洁、准确的输出
	   %G 根据实际情况采用 %E 或 %F 格式，以获得更简洁、准确的输出
	*/
	m := 18.51
	fmt.Printf("%b\n", m)
	fmt.Printf("%e\n", m)
	fmt.Printf("%E\n", m)
	fmt.Printf("%f\n", m)
	fmt.Printf("%.3f\n", m)
	fmt.Printf("%F\n", m)
	fmt.Printf("%g\n", m)
	fmt.Printf("%G\n", m)
}

// 字符串和[]byte
func TestStringFormatOutput(t *testing.T) {
	/*
		%s 直接输出字符串或[]byte
		%q 该值对应的双引号括起来的go语法字符串字面值，必要时会采用安全的转义表示
		%x 每个字节用两字符十六进制数表示(a-f)
		%X 每个字节用两字符十六进制数表示(A-F)
	*/
	s := "原木纯品"
	b := []byte("清风十三式")
	fmt.Printf("%s\n", s)
	fmt.Printf("%s\n", b)
	fmt.Printf("%q\n", s)
	fmt.Printf("%x\n", s)
	fmt.Printf("%X\n", s)
	fmt.Printf("%s\n", []byte{0xe5, 0x8e, 0x9f, 0xe6, 0x9c, 0xa8})
	fmt.Printf("%#q\n", []byte{0xe5, 0x8e, 0x9f, 0xe6, 0x9c, 0xa8})
	fmt.Printf("%x\n", []byte{0xe5, 0x8e, 0x9f, 0xe6, 0x9c, 0xa8})
	fmt.Printf("%X\n", []byte{0xe5, 0x8e, 0x9f, 0xe6, 0x9c, 0xa8})
	bi := new(big.Int).SetBytes([]byte{0xe5, 0x8e, 0x9f, 0xe6, 0x9c, 0xa8})
	fmt.Printf("%v\n", bi)
	fmt.Printf("%q\n", bi)
	fmt.Printf("%#q\n", bi)

	s = "Hello, 世界"
	// %v: Hello, 世界
	fmt.Printf("%%v: %v\n", s)
	// %q: "Hello, 世界"
	fmt.Printf("%%q: %q\n", s)
	// %+q: "Hello, \u4e16\u754c" (同上，强制转义非ASCII)
	fmt.Printf("%%+q: %+q\n", s)
	// %#q: `Hello, 世界` (通常使用双引号，若有换行等会用反引号)
	fmt.Printf("%%#q: %#q\n", s)

	// 指针：%p 表示为十六进制，并加上前缀0x
	var i int = 3
	p := &i
	fmt.Printf("%%p: %p\n", p)
	// 打印当前运行时golang版本
	fmt.Println("golang version:", runtime.Version())
}

// 测试宽度标识符
// 宽度通过一个紧跟在百分号后面的十进制数指定，如果未指定宽度，则表示值时除必需之外不做填充
// 精度通过可选的宽度后跟点号跟的十进制数指定，如果未指定精度，会使用默认精度；如果点号后没有跟数字，表示精度为0
// %f 默认宽度，默认精度
// %10f 宽度为10，默认精度
// %.2f 默认宽度，精度为2
// %10.2f 宽度为10，精度为2
// %10.f 宽度为10，精度为0
func TestWidthFormatOutput(t *testing.T) {
	n := 13.14
	fmt.Printf("%f\n", n)
	fmt.Printf("%10f\n", n)
	fmt.Printf("%10s\n", "我是字符串")
	fmt.Printf("%s%5s\n", strings.Repeat("有", 10-5), "我是字符串")
	fmt.Printf("%.2f\n", n)
	fmt.Printf("%10.2f\n", n)
	fmt.Printf("%010.f\n", n)
	fmt.Printf("%s%.f\n", strings.Repeat("*", 10-2), n)
	fmt.Printf("\033[1;32m+%s+\033[0m\n", strings.Repeat("-", 38))
}

// 测试FLAG
// + 总是输出数值的正负号，对%+q会生成全部是ASCII字符的输出（通过转义）
// 空格 对数值，正数前加空格而负数前加符号，对字符串采用 "% x" 或 "% X" 会给各打印字节之间加空格
// - 在输出右边填充空白而不是默认的左边。即从默认的右对齐切换为左对齐
// # 八进制数（%#o）前加 0 ；十六进制数（%#x）前加 0x；指针（%#p）去掉前面的 0x；对 %#U 会输出空格和单引号括起来的go字面值
// 0 使用0而不是空格填充，对于数值类型会把填充的0放在正负号后面
func TestFlagFormatOutput(t *testing.T) {
	s := "我是字符串"
	fmt.Printf("% d\n", 10)
	fmt.Printf("% d\n", -10)
	fmt.Printf("%+d\n", 10)
	fmt.Printf("%s\n", s)
	fmt.Printf("%10s\n", s)
	fmt.Printf("%-10s\n", s)
	fmt.Printf("%10.2f\n", 10.14)
	fmt.Printf("%-10.2f\n", 10.14)
	fmt.Printf("%010s\n", s)

	fmt.Printf("\033[1;32m+%s+\033[0m\n", strings.Repeat("-", 33))
	n := 69
	fmt.Printf("%#x\n", n)
	fmt.Printf("%#X\n", n)
	fmt.Printf("%#o\n", n)
	fmt.Printf("%#b\n", n)
	fmt.Printf("%U\n", n)
	fmt.Printf("%#U\n", n)
	// fmt.Printf("%#d\n", n) // error
	fmt.Printf("%+d\n", n)
}

// 测试Fprint
func TestFprint(t *testing.T) {
	fmt.Fprint(os.Stdout, "打印到标准输出\n")
	fmt.Fprintln(os.Stdout, "打印到标准输出")
	fmt.Fprintf(os.Stdout, "%s 打印到标准输出", "golang")
}
