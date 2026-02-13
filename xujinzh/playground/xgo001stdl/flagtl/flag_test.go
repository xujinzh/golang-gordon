package flagtl

import (
	"fmt"
	"os"
	"testing"
)

/*
flag标准库

go语言内置的flag包实现了命令行参数的解析，flag包使得开发命令行工具更为简单。

1. os.Args
如果你只是简单的想要获取命令行参数，可以使用os.Args来获取命令行参数。

os.Args是一个存储命令行参数的字符串切片，它的第一个元素使执行文件的名称。
*/

func TestOsArgs(t *testing.T) {
	// 获取命令行参数
	// os.Args: []string
	if len(os.Args) > 0 {
		for i, v := range os.Args {
			fmt.Println(i, v)
		}
	}
}

/*
2.1 flag参数类型

flag包支持的命令行参数类型有bool, int, int64, uint, uint64, float, float64, string, duration.

- 字符串flag		---		合法字符串
- 整数flag			---		1234， 0664， 0x1234等类型，也可以是负数
- 浮点数flag		---		合法浮点数
- bool类型flag		---		1, 0, t, f, T, F, true, false, TRUE, FALSE, True, False.
- 时间段flag		---		任何合法的时间段字符串，如"300ms", "-1.5h", "2h45m"。合法的单位有"ns", "us"or"μs", "ms", "s", "m", "h"

2.2 flag.Type()

基本格式如下：
flag.Type(flag名，默认值，帮助信息) *Type

例如，通过命令行获取三个参数：姓名、年龄，婚否，定义方式如下：

func main() {
	// flag.Type()
	name := flag.String("name", "张三", "年龄")
	age := flag.Int("age", 18, "年龄")
	married := flag.Bool("married", false, "是否已婚")
	delay := flag.Duration("d", 0, "时间间隔")

	// parse flag
	flag.Parse()

	// print
	fmt.Println(*name, *age, *married, *delay)
}

命令行运行：
./main --help

./main -age 19 -d 10s -married true -name lisi
*/

/*
2.3 flag.TypeVar(Type指针，flag名，默认值，帮助信息)

例如

func main() {
	var name string
	var age uint
	var married bool
	var d time.Duration

	flag.StringVar(&name, "name", "张三", "姓名")
	flag.UintVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "m", false, "是否已婚")
	flag.DurationVar(&d, "duration", 0, "时间间隔")

	flag.Parse()

	fmt.Println(name, age, married, d)
}
*/

/*
2.4 flag.Parse()

通过以上两种方法定义好命令行flag参数后，需要通过调用flag.Parse()来对命令行参数进行解析。
支持的命令行参数格式有以下几种：

1. -flag xxx 使用空格，一个"-"符号
2. --flag xxx 使用空格，两个"-"符号
3. -flag=xxx 使用等号，一个"-"符号
4. --flag=xxx 使用等号，两个"-"符号

flag解析在第一个非flag参数（单个"-"不是flag参数）之前停止，或者在终止符"-"之后停止。
*/

/*
2.5 其他函数

- flag.Argss() 返回命令行参数后的其他参数，以[]string类型
- flag.NArg() 返回命令函参数后的其他参数个数
- flag.NFlag() 返回使用的命令行参数个数

func main() {
	var name string
	var age uint
	var married bool
	var d time.Duration

	flag.StringVar(&name, "name", "张三", "姓名")
	flag.UintVar(&age, "age", 18, "年龄")
	flag.BoolVar(&married, "m", false, "是否已婚")
	flag.DurationVar(&d, "duration", 0, "时间间隔")

	flag.Parse()

	fmt.Println(name, age, married, d)

	fmt.Println(flag.Args())
	fmt.Println(flag.NArg())
	fmt.Println(flag.NFlag)
}
*/
