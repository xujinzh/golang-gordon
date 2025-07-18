## init 函数
每一个源文件都可以包含一个 init 函数，该函数会在 main 函数执行前，被 golang 运行框架调用，也就是说 init 会在 main 函数前被调用。

```go
package main

import (
	"fmt"
	"xgo2/utils"
)

// init 函数，通常用于完成初始化工作
func init() {
	fmt.Println("init function called")
}

func main() {
	utils.Test(4)
	n := 1
	utils.Test2(&n)
	fmt.Println(n)
	fmt.Println(utils.Test3(2, 3))
}
```

init 函数的注意细节
1. 如果以源文件同时包含全局变量定义、init 函数和 main 函数，则执行的流程：全局变量定义->init函数->main函数；
```go
package main

import (
	"fmt"
	_ "xgo2/utils"
)

var name = testVarFirst()

// 为了看到全局变量先被初始化，先定义一个函数
func testVarFirst() string {
	fmt.Println("testVarFirst called")
	return "xgo2"
}

func init() {
	fmt.Println("init function called")
}

func main() {
	fmt.Println(name)
}
```
2. init 函数最主要的作用就是完成一些初始化工作，先执行导入包源文件中的 init 函数，再调用 main 函数所在源文件中的 init 函数；
```go
// utils.go
package utils

import (
	"fmt"
)

var Age int
var Name string

func init() {
	fmt.Println("utils init function called")
	Age = 18
	Name = "xgo2"
}

// main.go
package main

import (
	"fmt"
	"xgo2/utils"
)

var MainAge = testVarFirst()

func testVarFirst() int {
	fmt.Println("testVarFirst function called")
	return 20
}

func init() {
	fmt.Println("main init function called")
}

func main() {
	fmt.Println("main()", utils.Age)
	fmt.Println("main()", utils.Name)
	fmt.Println("main()", MainAge)
}

// 执行结果
utils init function called
testVarFirst function called
main init function called
main() 18
main() xgo2
main() 20
```
3. 执行顺序：先执行被引源文件中的全局变量定义->init函数，再执行 main 函数源文件中的全局变量定义->init函数->main函数。

## 匿名函数
golang 支持匿名函数（就是没有名字的函数），如果我们某个函数只希望使用一次，可以考虑使用匿名函数。匿名函数也可以实现多次调用。

1. 在定义匿名函数时就直接调用；
```go
// 使用匿名函数计算两个数的和
	res := func(a, b int) int {
		return a + b
	}(3, 5)
	fmt.Println("The sum is:", res)
```
2. 将匿名函数赋给一个变量（就称为函数变量），再通过该变量来调用匿名函数。可以打破在函数中定义函数；
```go
// 将匿名函数赋值给变量
sumFuncVar := func(a, b int) int {
    return a + b
}
fmt.Println("The sum using variable is:", sumFuncVar(10, 20))
```
3. 如果将匿名函数赋给一个全局变量，那么匿名函数称为一个全局匿名函数，可以在整个程序中有效。


## 闭包
闭包是一个函数和与其相关的引用环境组合的一个整体（实体）
```go
// 累加器
func AddUpper() func (int) int {
	var n int = 10
	return func (x int) int {
		n += x
		return n
	}
}
```
说明：
1. AddUpper() 是一个函数，返回的数据类型是函数 fun (int) int;
2. 变量 n 和 return 返回的匿名函数一起组成了一个整体，构成闭包；
3. 可以理解为闭包是类，函数是类的方法，n 是字段。函数和它使用到的 n 构成闭包；
4. 当我们反复调用函数时，因为 n 是初始化一次，因此每调用一次就进行累计；
5. 搞清楚闭包的关键，就是要分析出返回的函数它引用哪些变量，因为函数和它引用的变量共同构成闭包。
```go
package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 累加器
func AddUpper() func(int) (int, string) {
	// return 的函数和 n 构成了闭包
	var n int = 10
	var str string = "hello"
	return func(x int) (int, string) {
		n += x
		str += strconv.FormatInt(int64(x), 10)
		// fmt.Println("n:", n, "str:", str)
		return n, str
	}
}

// 请编写一个函数 makeSuffix(suffix string) 可以接收一个文件后缀名（如 ".txt"）
// 并返回一个闭包。调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀名，
// 那么返回该文件名加上后缀名；如果已经有了指定的后缀名，则直接返回该文件名。

func makeSuffix(suffix string) func(string) string {
	// return 返回的函数和 suffix 构成了闭包
	return func(filename string) string {
		if strings.HasSuffix(filename, suffix) {
			return filename
		}
		return filename + suffix
	}
}

func main() {
	// 使用累加器
	addUpper := AddUpper()
	for i := 0; i < 10; i++ {
		resultN, resultStr := addUpper(i)
		fmt.Println(resultN, resultStr)
	}

	// 使用文件后缀名闭包
	addSuffix := makeSuffix(".txt")
	fmt.Println(addSuffix("file1"))     // 输出: file1.txt
	fmt.Println(addSuffix("file2.txt")) // 输出: file2.txt
	fmt.Println(addSuffix("file3.doc")) // 输出: file3.doc.txt
	fmt.Println("main() finished")
}
```

闭包可以在一定程度上弥补 golang 中函数不能设置默认值的缺陷。

## defer
在函数中，经常需要创建资源（如数据库连接，文件句柄，锁等），为了在函数执行完毕后，及时的释放资源，golang 设计者提供了 defer（延时机制）。

```go
package main

import "fmt"

func sum(n1, n2 int) int {

	// 当执行到 defer 时，暂不执行 defer 后面的代码，而是将其压入栈中，等到函数执行完毕后再执行
	// 执行的顺序是先进后出
	defer fmt.Println("ok1 n1=", n1) // 压入栈中该语句和变量值，只是推迟执行
	defer fmt.Println("ok2 n2=", n2) // 压入栈中该语句和变量值，只是推迟执行

	n1++
	n2++
	res := n1 + n2
	fmt.Println("ok3 res=", res)
	return res
}

func main() {
	res := sum(10, 20)
	fmt.Println("res=", res)

}
```
执行结果：
```bash
ok3 res= 32
ok2 n2= 20
ok1 n1= 10
res= 32
```

defer 的细节
1. 当 golang 执行到一个 defer 时，不会立即执行 defer 后面的语句，而是将 defer 后的语句压入到一个栈区，然后继续执行函数下一个语句；
2. 当函数执行完毕后，再从 defer 栈区依次取出执行语句（遵循先入后出的机制）；
3. 在 defer 将语句放入到栈时，也会将相关的值拷贝同时入栈。
4. defer 最主要的价值是在函数执行完毕后，可以及时的释放函数创建的资源。如关闭文件句柄、关闭数据库连接等
```go
// 伪代码
func test() {
	// 打开文件资源
	file = openfile(filename)
	// 保证函数执行完毕后能够关闭文件句柄
	defer file.close()
	// 其他代码
	file.read(3)
}

func test() {
	// 创建数据库连接
	connect = openDatabase("192.168.1.3", 8080)
	// 保证函数执行完毕后能够关闭数据库连接
	defer connect.close9)
	// 其他代码
	connect.sql(xxx)
}
```

## 函数参数传递总结
1. 函数传参有两种方式，一种是值传递（基本数据类型int、float、bool、string, 以及复杂数据类型中的数组和结构体 struct），一种是引用传递(指针、切片slice、map、管道chan、接口interface)；
2. 不管是值传递还是引用传递，传递给函数的都是变量的副本，不同的是，值传递的是值的拷贝，引用传递的是地址的拷贝。一般来说，地址拷贝效率高，因为数据量小，而值传递的效率取决于拷贝数据的大小，数据越大，效率越低；
3. 值类型默认是值传递，变量直接存储值，内存通常在栈中分配；
4. 引用类型默认是引用传递，变量存储的是一个地址，这个地址对应的空间才真正存储数据或值，内存通常在堆上分配，当没有任何变量引用这个地址时，该地址对应的数据空间就成为一个垃圾，有 GC 来回收。

## 变量作用域
1. 函数内部声明或定义的变量叫做局部变量，作用域仅限于函数内部；
2. 函数外部声明或定义的变量叫做全局变量，作用域在整个包都有效，如果其首字母为大写，则作用域在整个程序有效；
3. 如果变量是在一个代码块（如for、if中），那么这个变量的作用域就在该代码块。
4. 如下的全局变量定义需要注意
```go

var Age int = 20 // right
Name := "xjz" // error, equivalent to var Name string; Name = "xjz"
// 赋值语句在函数体外时错误的

func main() {
	//
}
```

## 字符串函数
1. 统计字符串的长度，按字节 len(str)
2. 字符串遍历，同时处理有中文的问题 r := []rune(str)
3. 字符串转整数 n, err := strconv.Atoi("12")
4. 整数转字符串 str = strconv.Itoa(1234)
5. 字符串转 []byte, var bytes = []byte("hello go")
6. []byte 转字符串，str = string([]byte{97, 98, 99})
7. 10 进制转2，8，16进制，str = strconv.FormatInt(123, 2), strconv.FormatInt(123, 8), strconv.FormatInt(123, 16)
8. 查找子串是否在指定的字符串中，strings.Contains("seafood", "foo")，返回布尔值
9. 统计一个字符串有几个指定的子串，strings.Count("ceheese", "e")，返回整数值
10. 不区分大小写的字符串比较（==是区分字母大小写的比较），strings.EqualFold("abc", "Abc")
11. 返回子串在字符串第一次出现的 index 值，strings.Index("ABC_abc", "abc"), 如果没有则返回-1
12. 将 string 数组的元素按照指定的字符串拼接起来，strings.Join(s, "|||")
13. 返回子串在字符串最后一次出现的 index 值，strings.LastIndex("go gopher", "go")，如果没有返回-1
14. 将指定的子串替换成另外一个子串，strings.Replace("oink oink oink", "k", "ky", n), n 可以指定你希望替换几个，如果 n=-1，表示全部替换
15. 按照指定的某个字符，为分割标识，将一个字符串拆分成字符串数组，strings.Split("hello,world,ok",",")
16. 将字符串的字符进行大小写转换，strings.ToLower("Go") 表示转为小写，strings.ToUpper("go") 表示转为大写， 返回字符串
17. 将字符串左右两边的空格去掉，strings.TrimSpace("  tn a lone gopher  ")， 返回字符串
18. 将字符串左右两边指定的字符去掉，strings.Trim("!hello !", "! ") 表示去掉字符串左右两边的字符'!'和' '， 返回字符串
19. 将字符串左边指定的字符去掉，strings.TrimLeft("!hello!", "! ") 表示去掉字符串左两边的字符'!'和' '， 返回字符串
20. 将字符串右边指定的字符去掉，strings.TrimRight("!hello!", "! ") 表示去掉字符串右两边的字符'!'和' '， 返回字符串
21. 判断字符串是否以指定的字符串开头，strings.HasPrefix("ftp://192.168.10.1", "ftp")，返回布尔值
22. 判断字符串是否以指定的字符串结束，strings.HasSuffix("nature.jpg", "jpg")，返回布尔值


