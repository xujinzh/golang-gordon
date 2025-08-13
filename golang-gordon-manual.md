# 基本介绍

## Golang 
Golang 是 Google 开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。 罗伯特·格瑞史莫、罗勃·派克及肯·汤普逊于 2007 年 9 月开始设计 Go，稍后伊恩·兰斯·泰勒、拉斯·考克斯加入项目。Golang 是基于 Inferno 操作系统所开发的。有时 Golang 也叫 Go 语言。

- 作业系统：Linux、macOS、FreeBSD、Windows
- 副档名：go
- 发行时间：2009年11月10日，​13年前
- 实作者：Google
- 当前版本(截止到2023年)：1.21.3 (2023 年 10 月 10 日；稳定版本);
- 编程范型：编译型，可平行化，结构化，指令式
- 设计者：罗伯特·格瑞史莫（英语：Robert Griesemer）; 罗勃·派克; 肯·汤普逊

## 撰写风格
Golang 有定义如下的撰写风格：
1. 每行程序结束后不需要撰写分号（`;`）。
2. 大括号 `{` 不能够换行放置，且必须成对出现。
   ```go
   // 错误示例
   func main()
   {
	// TODO
   }

   // 正确示例
   func main() {
    // TODO
   }
   ```
3. if 判断式和 for 循环不需要以小括号包覆起来。
   ```go
   if true {
	  fmt.Println("yes")
   }
   ```
4. 使用 tab 做排版。
5. 每个源文件必须属于一个包(package)。
6. 程序的执行入口是 main() 函数，且必须是在 main 包中。
7. Go 语言严格区分大小写。
8. Go 编译器是一行一行编译的，因此，不要多个语句写在同一行，会报错。
9. Go 语言代码文件中定义的变量或导入的包如果没有使用，代码编译不能通过。

使用内置 gofmt 工具后，会自动整理代码，使之符合规定的撰写风格。方法是：`gofmt -w main.go`，`-w` 表示自动将整理后的代码重写会源文件。更多参数可以使用 `gofmt --help` 查看.

## 目标
Golang 的主要目标是”兼具 Python 等动态语言的开发速度和 C/C++ 等编译型语言的性能与安全性“。

## 特性
- Go 语言的用途众多，可以进行网络编程、系统编程、并发编程、分布式编程。
- <font color=blue>Go 语言天然支持并发</font>，提出 goroutine 轻量级协程，可实现大并发处理，高效利用多核。
- 因为 Go 语言没有类和继承的概念，所以它和 Java 或 C++ 看起来并不相同。但是它通过接口（interface）的概念来实现多态性。Go 语言有一个清晰易懂的轻量级类型系统，在类型之间也没有层级之说。因此可以说 Go 语言是一门混合型的语言。
- Go 是编译型语言。在编译代码时，编译器检查错误、优化性能并输出可在不同平台上运行的二进制文件。
- Go 自带了编译器，因此无须单独安装编译器。
- Go 语言支持交叉编译，比如说你可以在运行 Linux 系统的计算机上开发可以在 Windows 上运行的应用程序。这是第一门完全支持 UTF-8 的编程语言，这不仅体现在它可以处理使用 UTF-8 编码的字符串，就连它的源码文件格式都是使用的 UTF-8 编码。Go 语言做到了真正的国际化！
- Go 语言有一个吉祥物，在会议、文档页面和博文中，大多会包含下图所示的 Go Gopher，这是才华横溢的插画家 Renee French 设计的，她也是 Go 设计者之一 Rob Pike 的妻子。

<!-- ![](https://c.biancheng.net/uploads/allimg/180808/1-1PPQA9545W.jpg) -->
![](https://camo.githubusercontent.com/2b507540e2681c1a25698f246b9dca69c30548ed66a7323075b0224cbb1bf058/68747470733a2f2f676f6c616e672e6f72672f646f632f676f706865722f6669766579656172732e6a7067)
- 一种静态强类型、编译型、并发型、并具有垃圾回收功能的编程语言。
- Go 语言的一个文件都要归属于一个包，不能单独存在，即每个 `x.go` 文件上面要有一个 `package xxx`。包的概念用于组织程序结构。
- Go 语言支持返回多个值。
- 编写 Go 语言代码时，每一行编写一个语句，每个语句后面不需要增加 `;`，编译时会自动增加。
- Go 语言代码中未使用的导入包会报错；代码中未使用的变量定义会报错。


# 运行环境
## Golang 环境配置
访问 [Golang官网](https://go.dev/) 下载合适的版本。解压到喜欢的目录，并添加到环境变量（如 /etc/profile，生效需要注销当前登录用户）即可。类似于 java 和 nodejs。大概如下：
```bash
export GOROOT=/usr/local/go
export PATH=$GOROOT/bin:$PATH
export GOPATH=/Users/jinzhongxu/codes/go

# GOPATH 非必须，如果不配置会将 pkg（网络下载的包存放的位置）、bin（编译安装后的可执行文件存放的位置）自动存放到当前用户的家目录里的 go 文件夹下。建议配置
```

## 测试环境
### 查看运行版本
```bash
# 查看 Golang 版本信息
go version
```

### 编写简单程序
1. 创建文件夹：`go/src/learn/proj1/`；
2. 在该文件夹下创建文件：`main.go`，在该文件中增加如下内容：
	```go
	// 表示 main.go 文件归属于 main 包
	package main

	// 导入多个包
	import (
		"fmt"
		_ "unsafe"  // 暂时不用的包，前面增加 _ 表示忽略
	)

	//// 导入单个包
	//import "fmt"

	//// 声明多个变量
	// var (
	// 	i = 3
	// 	f = 3.14
	// 	b = true
	// )

	// func 是关键字，定义函数
	// main() 是主函数，程序的入口
	func main() {
		fmt.Println("hello, go!")
	}
	```
3. 在终端中执行：
	```bash
	# 直接执行。会后台编译，然后再执行。速度慢
	go run main.go

	# 先编译，再执行
	# 编译成同名的可执行程序
	# 注意变异后的程序一般会比较大，因为编译时会把依赖包一起打包进去，
	# 因为可执行程序中已有程序运行时的依赖包，
	# 所以移动到其他机器上仍然可以运行，而不需要配置 Go jdk
	go build main.go
	# 执行
	./main

	# 编译成指定的文件名可执行程序
	go build -o my_main main.go
	```

## 导入本地包
在 Golang 中，我们常常需要导入其他文件夹或项目的包，如导入本项目下子文件夹中某个 `.go` 文件中的变量或函数，导入其他项目（不在本项目目录下，可能同级目录或其他磁盘目录）下某个子文件夹中某个 `.go` 文件中的变量或函数。

如果一个 `.go` 文件中某个变量想要被其他 `.go` 文件中的代码访问，那么该变量的**命名要求首字母大写**。

### 导入本项目中的变量或函数
如果想要导入本项目（包括子文件夹）下的变量或函数，需要该项目进行包初始化，即 `go mod init`，具体地，假设项目目录如下：
```bash
proj1
├── go.mod
├── main.go
└── misc
    └── utils.go
```
其中，`proj1/misc/utils.go` 里面的内容如下：
```bash
package proj1

var Clock string = "石英"
```
我们想要在 `proj1/main.go` 中导入 `proj1/misc/utils.go` 里面的变量 `Clock`，方法如下：
1. 初始化包名:
	```bash
	cd proj1

	go mod init
	# or
	go mod init proj1

	# 此时会在 proj1 下生成 go.mod 文件
	# 查看该文件内容
	cat go.mod
	# 结果显示如下
	module proj1
	go 1.21.3
	```
2. 在 `proj1/main.go` 中导入需要的变量：
	```go
	package main

	import (
		"fmt"
		misc "proj1/misc" // 前面是包别名，后面是包路径名
	)

	func main() {
		fmt.Println(misc.Clock)
	}
	```


### 导入其他项目中的变量或函数
如果想要导入其他项目（包括子文件夹）下的变量或函数，需要将导入的项目进行包初始化，即 `go mod init`，具体地，假设待导入的项目目录和本项目的目录结构如下：
```bash
learn
├── proj1
│   ├── go.mod
│   ├── main.go
│   └── misc
│       └── utils.go
└── proj2
    ├── go.mod
    ├── misc
    │   └── hello.go
    └── web.go
```
`proj2/misc/hello.go` 内容如下：
```go
package proj2

var SayHello string = "hello"
```

我们需要将 `proj2/misc/hello.go` 里面的变量导入到 `proj1/main.go` 中。我们需要将 `proj1` 和 `proj2` 都进行包初始化: `go mod init`，然后，修改 `proj1/go.mod` 文件，最后在 `proj1/main.go` 中导入。具体如下：
1. 初始化包：
	```bash
	# 初始化两个包
	cd proj1
	go mod init
	# 此时在 proj1 下会产生 go.mod 文件

	cd proj2
	go mod init
	# 此时在 proj2 下会产生 go.mod 文件
	```
2. 修改 `proj1/go.mod` 文件：
	```bash
	cd proj1
	vim go.mod
	# 增加如下内容：
	require proj2 v0.0.0
	replace proj2 => ../proj2

	# 修改后，go.mod 内容大概如下
	module proj1

	go 1.21.3

	require proj2 v0.0.0
	replace proj2 => ../proj2

	# 即上面将项目 proj2 的路径写出来
	# 当 proj2 和 proj1 同目录，replace 项目路径就是 ../proj2
	# 当 proj2 在其他目录，那么 replace 填写正确路径即可
	```
3. 在 `proj1/main.go` 中导入：
	```go
	package main

	import (
		"fmt"
		proj2Misc "proj2/misc" // 前面是包别名，后面是包路径名
	)

	func main() {
		fmt.Println(proj2Misc.SayHello)
	}
	```
如果想要在 `proj1` 下的子文件夹中的其他 `.go` 中导入 `proj2` 中的变量或函数，上面同样适用。


# 转义字符（escape char）
常用的转义字符有：
- \t : 制表符；
- \n : 换行符；
- \\\\ : \ 符号；
- \\' : ' 符号；
- \\" : " 符号；
- \a : 响铃；
- \b : 退格；
- \f : 换页；
- \v : 垂直制表符；
- \r : 回车符；
   ```go
   fmt.Println("hello, \rgo!")

   // 打印结果不是
   hello,
   go!
   // 而是。因为不换行回车会在当前行打印
   go!lo, 
   ```
- `` : 将所有文本书写在 \`和\` 之间，将失去转义效果。如将代码写入其中，将会原封不动打印出来，并保持原结构。

# 注释
通过注释提高代码的可阅读性。

```go
// 行注释，参考 C++

/*
   块注释，参考 C
*/
```

Go 语言官方推荐尽量使用行注释。块注释中不要出现块注释。


# 变量
变量是程序的基本组成单位。Go 中变量的声明：
```go
// 方法一，显式声明变量的数据类型 int
var x int // 默认赋初值
var y int = 3

// 方法二，隐式声明变量的数据类型 int，根据赋值自动推断
var x = 3

// 方法三，常用. 省略 var 关键字，使用 := 表示声明新变量并赋值，隐式声明变量的数据类型 int
x := 3

// 多个变量同时声明。单引号'表示字符，双引号"表示字符串
var x, y, z int
var a, b, c = 1, '2', "3"
v1, v2, v3 := 1.1, true, "中"
var (
	x1 = 3
	x2 = "marry"
)

// 变量赋值
x = 33
```

注意：
- 声明的变量在前面程序中没有被声明过，否则编译会出错。
- 声明的变量在后面程序中必须使用，否则编译会出错。
- 声明的变量如果没有赋值，那么会隐式赋默认值，如 int 类型赋值 0，string 类型赋值空串（''）。
- 声明的变量只能在同一数据类型内变化，不能被其他数据类型的值重新赋值，如 `var x int`，那么 `x` 在后续程序使用中只能被改变为整数值，不能被赋值为字符串。
- 在一个作用域内，不能出现重名的变量。
- 变量代表了三重属性：变量名、值和数据类型。
- "+" 两边都是数值类型变量时，做数值加法（如 `2 + 3.1`）；当两边都是字符串变量时，则做字符串拼接（如 `"hello " + "world" ` ）。但不能是一边是数值类型变量，一边是字符串类型变量。

# 常量
1. 常量使用 const 修饰
2. 常量在定义的时候，必须初始化
3. 常量不能修改
4. 常量只能修饰 bool、数值类型（int, float）、string 类型
5. golang 中没有要求常量必须首字母大写
```go
// 语法
const 标识符 [type类型] = value值

// example
const name = "tom"

const tax float64 = 0.8

const a int  // error

const b = 9 / 3

var num = 9
const b = num / 3 // error

const c = getVal() // error

const (
	a = 1
	b = 2
)

// 专业写法
const (
	a = iota // 表示给 a 赋值 0
	b  // b 在 a 的基础上 +1
	c  // c 在 b 的基础上 +1
	// 结果是 a = 0, b = 1, c = 2
)

const (
	a = iota
	b = iota
	c, d = iota, iota
	// 结果是 a = 0, b = 1, c = d = 2
)
```


# 数据类型
Go 语言中每一种变量都要有对应的数据类型，每一种数据类型在内存中分配了不同大小的内存空间。Go 语言中分为两大类数据类型，一种是基本数据类型，一种是复杂数据类型（也叫派生数据类型）。按照变量存储的数据情况又分为值类型（变量存储数据值，内存通常在栈中分配）和引用类型（变量存储内存地址，该地址对应的空间才真正存储数据或值，内存通常在堆中分配。当没有任何变量引用这个地址时，该地址对应的数据空间就称为一个垃圾，**有 GC 来回收**）。值类型包括所有基本数据类型和复杂数据类型中的数组、结构体，引用类型包括复杂数据类型中的其他类型（指针、切片 slice、map、管道 channel、接口 interface）。

## 基本数据类型
基本数据类型包括有数值型（又分为整数类型和浮点类型）、布尔型和字符串型（官方把该类型作为基本数据类型）。

### 整数类型
整数类型包含有：`int8, int16, int32, int64, uint8, uint16, uint32, uint64, int, uint, rune, byte`。

`int8` ~ `int64` 表示有符号整数，范围是 $-2^{x-1}$ ~ $2^{x-1} - 1$，这里 $x$ 取值 `8, 16, 32, 64`，分别占 `1, 2, 3, 4` 个字节。

`uint8` ~ `uint64` 表示无符号整数（非负整数），范围是 $0$ ~ $2^x - 1$，这里 $x$ 取值 `8, 16, 32, 64`，分别占 `1, 2, 3, 4` 个字节。

除此之外，还有 4 种：`int, uint, rune, byte`:
|数据类型|有无符号|占用内存空间|表示范围|备注|
|---|---|---|---|---|
|int|有|64位系统占8个字节；32位系统占4个字节|64位系统: $-2^{64-1}$ ~ $2^{64-1} - 1$; 32位系统: $-2^{32-1}$ ~ $2^{32-1}-1$|跟操作系统有关|
|uint|无|64位系统占8个字节；32位系统占4个字节|64位系统: $0$ ~ $2^{64} - 1$; 32位系统: $0$ ~ $2^{32}-1$|跟操作系统有关|
|rune|有|4个字节|$-2^{32-1}$ ~ $2^{32}-1$|同 int32，**表示一个 Unicode 码**，处理中文比较合适|
|byte|无|1个字节|$0$ ~ $255$|同 uint8，常用于存储字符|

注意：
- 变量赋值不能超过声明的数据类型范围，否则编译错误。如 `var x byte = 256` 超过 255，会编译出错。
- 查看变量的数据类型：
	```go
	x := 3
	fmt.Printf("x 的数据类型: %T", x)
	```
- 整型变量的默认值是 0。
- 整型变量的默认数据类型为 `int`:
	```go
	var x = 3
	// 等价于
	var x int = 3
	```
- 查看变量的字节大小：
	```go
	import (
		"fmt"
		"unsafe"
	)

	func main() {
		x := 3
		fmt.Printf("x 占用内存字节大小：%d", unsafe.Sizeof(x))
	}
	```
- 声明变量时，当知道变量的取值（变化）范围时，应使用占用字节最小的数据类型，应尽量避免内存浪费：`var age byte = 100`。

#### 字符
Golang 中没有专门的字符类型。如果存储单个字符，那么可以使用 `byte` 或 `uint`，也就是说字符按照整数存储。Golang 中的字符串是由每个整数字节组成。
```go
var c1 byte = '0'
var c2 byte = 'a'
var c3 byte = 97
var c4 byte = '\n'
// 不能使用 byte 声明中文，因为会导致溢出
// Golang 使用 UTF-8 编码，一个英文字符占1个字节，一个中文字符占3个字节。
var c5 uint = '中'
```
打印输出字符，可以使用格式化输出，直接打印将会输出对应的整数编码值：
```go
var c = 'a'
fmt.Printf("c=%c", c)
```
因为字符在 Golang 中本质上是整数，所以它可以和数值类型变量进行运算：
```go
var x = 'a'
fmt.Println(x + 10)
// 结果是107
```
字符和码值的对应关系依据字符编码表（这里是 UTF-8）。

### 浮点类型
浮点类型就是用于存储小数的类型。分类有单精度 `float32` 和双精度 `float64`:
|数据类型|占用内存空间|表示范围|
|---|---|---|
|float32|4字节|-3.403E38 ~ 3.403E38|
|float64|8字节|-1.798E308 ~ 1.798E308|

注意：
- 浮点数存储时，包含有符号位、指数位、尾数位。所以浮点数都是有符号的。
- 浮点数的尾数位在存储中可能会造成精度损失
	```go
	var x float32 = 3.00000902
	fmt.Println(x)

	// 结果为 3.000009
	```
- 浮点数都是有符号的。
- 浮点数变量的默认类型是 `float64`。
- 浮点数变量的默认值是 0。
- 浮点数有两种表示形式：
	```go
	// 十进制形式
	var x = 5.21
	var x = 0.521
	var x = .521

	// 科学计数法
	var x = 5.21e2
	var x = 5.21e+2
	var x float32= 5.21e+02
	var x = 5.21E2
	var x = 5.21E-2
	var x float32 = 5.21E-02
	```
- 精度敏感时应使用 `float64`。


### 布尔类型（bool）
Golang 中布尔类型只能取 `true` 和 `false`，布尔类型值占用 1 个字节。

布尔类型常用于逻辑判断和流程控制。

布尔变量的默认值是 `false`。

打印布尔类型数据值：
```go
var b bool = true
fmt.Printf("b = %t", b)
```

### 字符串类型（string）
字符串是一串固定长度的字符连接起来的字符序列。使用双引号括起来，如：`"你好"`。

- Golang 的字符串是由单个整数字节连接起来的。统一使用 UTF-8 编码。
- 字符串一旦赋值，就不能被修改，即 Golang 中字符串是不可变的。
	```go
	var s = "abc"
	fmt.Println(s[0])
	s[0] = 'A' // 编译出错
	```
- 字符串表示由两种：
	```go
	// 方法1。会识别转义字符，需要注意。
	var s1 = "她说：\"天真好！\""

	// 方法2。以字符串原始形式输出，非常方便。可以自由书写换行、特殊字符，也可以输出源代码，并能防止攻击。
	var s2 = `
	package main

	import "fmt"

	func main() {
		fmt.Printf("Hello, World!")
	}
	`
	```
- 当字符串太长时，可以分行拼接，“+” 要防止上一行的末尾：
	```go
	var s = "我们今天去南京大排档吃了美食，" + 
			"那里的金陵烤鸭风味挺独特的，与北京烤鸭不同！"
	```
- 字符串变量的默认值是 `""`:
	```go
	var s string
	fmt.Printf("s=%v", s)
	// 输出结果是 ""
	```

### 基本数据类型的转换
Golang 在不同类型的变量之间赋值时需要**显示转换**。即 Golang 中数据类型不能自动转换。

#### 数值类型间的转换
```go
var x int32 = 100

// x convert datatype to float32
var xf float32 = float32(x)

// x convert datatype to int8
var xi8 int8 = int8(x)

// xf convert datatype to float64
var xf64 float64 = float64(xf)
fmt.Println(x, xf, xi8, xf64)
```
注意:
- 从高精度到低精度转换，结果可能会溢出。
	```go
	var x int32 = 1000
	var y int8 = int8(x)
	fmt.Println(y)
	// 结果是 -24
	```
- "+" 号类型转换：
	```go
	// 我这里机器是 64 位
	var i int32 = 12
	x := i + 10
	fmt.Printf("%T, %d\n", i, unsafe.Sizeof(i))
	fmt.Printf("%T, %d", x, unsafe.Sizeof(x))
	// result
	// int32, 4
	// int32, 4
	```
并不是所有数据类型间都能用这种方式转换：
```go
var b1 bool = true
var b2 bool = false
fmt.Println("b1 + b2 =", uint8(b1) + uint8(b2))  // error

// 可以自定义函数进行转换

func Bool2Int(b bool) uint8 {
	var i uint8
	if b {
		i = 1
	} else {
		i = 0
	}
	return i
}

fmt.Println("b1 + b2 =", misc.Bool2Int(b1) + misc.Bool2Int(b2))
```
#### 数值类型的转为字符串类型
1. fmt.Sprintf
	```go
	var i = 12
	var f float32 = 1.23
	var b bool = true
	var c byte = 'a'
	var s string
	s = fmt.Sprintf("%d", i)
	s = fmt.Sprintf("%f", f)
	s = fmt.Sprintf("%b", b)
	s = fmt.Sprintf("%c", c)
	```
2. strconv
	```go
	var i = 12
	var f float32 = 1.23
	var b bool = true
	// var c byte = 'a'
	var s string
	s = strconv.FormatInt(int64(i), 10)
	s = strconv.Itoa(int(i))
	// 'f' 表示输出格式，10 表示保留精度10位，64 表示 float64
	s = strconv.FormatFloat(float64(f), 'f', 10, 64)
	s = strconv.FormatBool(b)
	
	```

#### 字符串类型转数值类型
- strconv
	```go
	var s string = "true"
	var b bool
	b, _ = strconv.ParseBool(s)
	fmt.Println(b)

	var s2 string = "123"
	var i int64
	i, _ = strconv.ParseInt(s2, 10, 64)
	fmt.Println(i)

	var s3 string = "1.45"
	var f float64
	f, _ = strconv.ParseFloat(s3, 64)
	fmt.Println(f)
	```
如果将 "abc" 转成数值类型，Golang 会直接将该字符串转为 0（即，如果转换不成功，会转为目标类型的默认值）. 因此，使用 string 类型转数值类型时，要保证输入内容是一个有效数据。

## 复杂数据类型
复杂数据类型包含有指针、数组、结构体（struct）、管道（channel）、**函数**、切片（slice）、接口（interface）、map 等。

### 指针
关于指针，我们要知道：
- 获取变量的内存地址：
	```go
	var i = 10
	fmt.Println("变量 i 的地址是", &i)
	```
- 定义指针：`&ptr` 表示变量的内存地址，`ptr` 表示变量的值，`*ptr` 表示指针指向的变量的值。
	```go
	var i = 10
	// ptr 是指针，指向 int 变量的内存，
	// 里面存的值是 &i，即 i 的内存地址
	var ptr *int = &i
	fmt.Println("变量 i 的内存地址：", &i)
	fmt.Println("ptr 也是变量，请把 ptr 的地址打印出来：", &ptr)
	fmt.Println("ptr 指针变量里面存的值（另一个变量的内存地址）：", ptr)
	fmt.Println("ptr 指向的变量在内存中的值：", *ptr)

	// 结果
	// 变量 i 的内存地址： 0xc000112008
	// ptr 也是变量，请把 ptr 的地址打印出来： 0xc000108018
	// ptr 指针变量里面存的值（另一个变量的内存地址）： 0xc000112008
	// ptr 指向的变量在内存中的值： 10

	var f float32 = 3
	var ptr *float32 = &f
	*ptr = 3.33
	fmt.Println("变量 f 的值现在是：", f)

	// 结果
	// 变量 f 的值现在是： 3.33

	var i = 10
	fmt.Printf("i 的地址是 %p\n", &i)
	fmt.Printf("i 的地址是 %#p\n", &i)

	var p = &i 
	fmt.Printf("p = %p\n", p)
	fmt.Printf("p 指向的变量值：%d\n", *p)
	```
- 所有的值类型（基本数据类型：整型、浮点型、布尔型、字符串型，复杂数据类型：数组、结构体）都有对应的指针类型，形式为 "*数据类型"，如 `*int`, `*float32`, `*bool`, `*string` 等。

# 值类型和引用类型

## 值类型和引用类型分类
- 值类型：基本数据类型（int系列、float系统、bool、string），复杂数据类型（数组和结构体 struct）
- 引用类型：指针、slice切片、map、管道chan、接口interface等

## 值类型和引用类型特点
- 值类型：变量直接存储值，内存**通常**（排除逃逸情况）在**栈**中分配
- 应用类型：变量存储的是一个地址，这个地址对应的空间才真正存储数据（值），内存**通常**（排除逃逸情况）在**堆**上分配，当没有任何变量引用这个地址时，该地址对应的数据空间就成为一个垃圾，有GC来回收

*golang中对栈区和堆区区分的就不是特别明确，主要根据变量到底使用多长时间决定放在什么地方，如果使用时间很长就放在堆区。*

# 标识符
Golang 对各种变量、函数、方法等命名时使用的字符序列称为标识符。或者我们称在程序中凡是可以自己起名字的地方都叫做标识符。

## 标识符的命名
Golang 中标识符的命名规则：
- 标识符只能有 a-z, A-Z,  0-9, _ 组成；
- 数字不能作为标识符的开头；
- 标识符区分大小写；
- "_" 单独作为标识符时，只是空标识符，常用于接收函数某些不关心的返回值，它的值会被忽略；
- 标识符的命名应该避免关键字（如 `break`, `if` 等）； 

命名注意：
- 包(package)名建议和目录保持一致，尽量使用有意义的包名，应简短、有意义、且不能和标准库（如 `fmt, unsafe, strconv, net` 等）冲突；
- 变量名、函数名、常量名建议采用驼峰法（如 `addFloat`）；
- 如果变量名、函数名、常量名首字母大写（如 `AddFloat`），那么可以被其他包访问；反之，则只能在当前包中使用。即首字母大写是公开访问，首字母小写是私有内部访问。

## 保留关键字和预定义标识符
Golang 中保留关键字有 25 个，如下：
|关键字|关键字|关键字|关键字|关键字|
|---|---|---|---|---|
|break|default|func|interface|select|
|case|defer|go|map|struct|
|chan|else|goto|package|switch|
|const|fallthrough|if|range|type|
|continue|for|import|return|var|

Golang 中预定义标识符（包括基础数据类型和系统内嵌函数）有 36 个，如下：
|预定义|预定义|预定义|预定义|预定义|预定义|
|---|---|---|---|---|---|
|append|bool|byte|cap|closs|complex|
|complex64|complex128|uint16|copy|false|float32|
|float64|imag|int|int8|int16|uint32|
|int32|int64|iota|len|make|new|
|nil|panic|uint64|print|println|real|
|recover|string|true|uint|uint8|uintprt|


# 运算符
运算符是一种特殊的符号，用于表示数据的运算、赋值和关系等等。常用的运算符有：
1. 算术运算符；
2. 赋值运算符；
3. 关系运算符；
4. 逻辑运算符；
5. 位运算符；
6. 其它运算符；

## 算术运算符
算术运算符是对数值类型的变量进行运算，如数值的加减乘除、字符串的拼接等。 

|运算符|含义|示例|示例结果|
|---|---|---|---|
|+|正号|+1|1|
|-|负号|-1|-1|
|+|加|1+1|2|
|-|减|2-1|1|
|*|乘|2*3|6|
|/|除|3/2|1，默认表示整除。如果想得到浮点数，float64(3)/2=1.5|
|%|取余或取模|3%2|1，x%y=x-x/y*y|
|++|自增|x := 1, x++|2，只能独立使用|
|--|自减|x := 1, x--|0，只能独立使用|
|+|字符串拼接|"ab"+"c"|"abc"|

注意：
1. golang 的自增自减只能当做一个独立语言使用，不能这样使用
```go
b := a++
b := a--
if i++ > 0 {
	//
}
```
2. golang的 ++ 和 -- 只能写在变量的后面，不能写在变量的前面，不能这样使用
```go
++a
--a
```

## 赋值运算符
赋值运算符将右边运算的结果赋值左边的变量。

|运算符|含义|示例|
|---|---|---|
|=|直接赋值|x = 3|
|+=|相加后再赋值| x += 3|
|-=|相减后再赋值| x -= 3|
|*=|相乘后再赋值| x *= 3|
|/=|相除后再赋值| x /= 3|
|%=|取余后再赋值| x %= 3|
|<<=|左移后再赋值| x <<= 3|
|>>=|右移后再赋值| x >>= 3|
|&=|按位与后再赋值| x &= 3|
|^=|按位异或后再赋值| x ^= 3|
|\|=|按位或后再赋值| x \|= 3|

注意：
1. 运算顺序从右向左；
2. 赋值运算符的左边只能是变量，右边可以是变量、表达式、常量；
3. `x += 3` 等价于 `x = x + 3`；

交换变量 a 和 b 的值：
```go
var a int = 10
var b int = 30

a = a + b
b = a - b
a = a - b
```

## 关系运算符
关系运算符输出的结果是布尔值，也就是只能取 `true` 或者 `false`。关系运算符组成的表达式称为关系表达式。

|运算符|含义|示例|示例结果|
|---|---|---|---|
|==|相等|1 == 1|true|
|!=|不等|1 != 2|true|
|<|小于|3 < 4|true|
|>|大于|3 > 4|false|
|<=|小于等于|3 <= 4|true|
|>=|大于等于|3 >= 4|false|


## 位运算符
位运算符是双目运算符，参与运算的两个数对应位置的二进制位相互运算。

|运算符|描述|
|---|---|
|&|按位与运算。二进制位同为1时，结果才为1|
|\||按位或运算。二进制位同为0时，结果才为0|
|^|按位异或运算。二进制位不同时，结果为1，相同时结果为0|
|<<|左移运算。二进制位向左（高位）移动，高位丢弃，低位补0；左移$n$位就是乘以2的$n$次方；低位溢出，符号位不变，并用符号位补溢出的高位|
|>>|右移运算。二进制位向右（低位）移动，低位丢弃，高位补0；右移$n$位就是除以2的$n$次方；符号位不变，低位补0|


## 逻辑运算符

假设 `a=true`，`b=false`，那么

|运算符|含义|示例|示例结果|
|---|---|---|---|
|&&|逻辑与。只有运算符两边都为true时，结果才为true|a && b|false|
|\|\||逻辑或。只有运算符两边都为false时，结果才为false|a\|\|b|true|
|!|逻辑非。运算符右边为true时结果为false，false时结果为true|!a|false|

- 逻辑与会出现<font color=red>短路与</font>。当左边的条件为 false，则右边的条件不会判断，直接输出 false;
- 逻辑或会输出<font color=red>短路或</font>。当左边的条件为 true，则右边的条件不会判断，直接输出 true;


## 运算符优先级

Go语言运算符优先级和结合性一览表

|优先级|分类|运算符|结合性|
|---|---|---|---|
|1|逗号运算符|,|从左到右|
|2|赋值运算符|=、+=、-=、*=、/=、 %=、 >=、 <<=、&=、^=、\|=|<font color=orange>从右到左</font>|
|3|逻辑或|\|\||从左到右|
|4|逻辑与|&&|从左到右|
|5|按位或|\||从左到右|
|6|按位异或|^|从左到右|
|7|按位与|&|从左到右|
|8|相等/不等|==、!=|从左到右|
|9|关系运算符|<、<=、>、>=|从左到右|
|10|位移运算符|<<、>>|从左到右|
|11|加法/减法|+、-|从左到右|
|12|乘法/除法/取余|*（乘号）、/、%|从左到右|
|13|单目运算符|!、*（指针）、& 、++、--、+（正号）、-（负号）|<font color=orange>从右到左</font>|
|14|后缀运算符|( )、[ ]、->|从左到右|

注意
- 建议使用小括号 `()` 来分割表达式的运算顺序，小括号里面的运算优先执行;
- 优先级值越大，运算时越优先；
- 单目运算符、赋值运算符是从右向左的；


## 其他运算符
1. `&` 取变量地址
```go
var a int = 30
fmt.Println("a的内存地址是", &a)
```
2. `*` 取指针指向的变量的值
```go
var a int = 30
var ptr *int = &a
fmt.Println("a的值是", *ptr)
```

## 特别说明
golang 中不支持三元运算符，可以使用 if-else 来达成相同目的：
```go
var a int = 30
var b int = 40
var max int

if a > b {
	max = a
} else {
	max = b
}
```

# 接收用户输入
1. fmt.Scanln，接收一行数据
```go
var name string
var age byte
var salary float32
var isPass bool

fmt.Printf("name: ")
fmt.Scanln(&name)
fmt.Printf("age: ")
fmt.Scanln(&age)
fmt.Printf("salary: ")
fmt.Scanln(&salary)
fmt.Printf("isPass: ")
fmt.Scanln(&isPass)

fmt.Printf("name=%q, age=%d, salary=%f, isPass=%t\n", name, age, salary, isPass)
```
2. fmt.Scanf，按照格式接收数据
```go
var name string
var age byte
var salary float32
var isPass bool

fmt.Println("please input name, age, salary and isPass, use space split")
fmt.Scanf("%s %d %f %t", &name, &age, &salary, &isPass)
fmt.Printf("name=%q, age=%d, salary=%f, isPass=%t\n", name, age, salary, isPass)
```

# 进制
## 二进制
golang中，不能直接使用二进制来表示一个整数，这沿用了C的特点。
```go
var i int = 3
fmt.Printf("%b\n", i)
```

## 八进制
```go
011 == 9
```

## 十六进制
```go
0x11 == 17
```
十六进制不区分大小写（a-f, A-F）

# 原码、反码、补码
1. 二进制的最高位是符号位，0表示正数，1表示负数
2. 正数的原码、反码、补码一样
3. 负数的反码：它的原码符号位不变，其他位取反
4. 负数的补码：它的原码+1
5. 0的反码、补码都是0
6. 在计算机中，都是以补码的方式来运算的

# 流程控制
1. 顺序控制
2. 分支控制
3. 循环控制

## 顺序控制
golang代码按照代码的顺序执行。
变量要先定义再使用。

## 分支控制
```go
// 单分支
var age byte = 30
if age > 18 {
	//
}

// or

if age := 18; age + 3 > 20 {
	//
}

// 双分支
if age > 18 {
	//
} else {
	//
}

// 多分支
if age > 13 {
	//
} else if age > 18 {
	//
} else { // 可以没有
	//
}


// switch 
var key byte
fmt.Println("please input a char: a, b, c")
fmt.Scanf("%c", &key)

func addOne(key byte) byte {
	return key + 1
}

switch addOne(key) + 1 {
	case 'a', 'g':
		fmt.Println("you input a or g")
	case 'b':
		fmt.Println("you input b")
	case 'c':
		fmt.Println("you input c") 
	default:
		fmt.Println("error")
}


// switch 当做 if-else 使用
var a int = 5
switch {
case a == 5:
	fmt.Println("a is 5")
case a == 6:
	fmt.Println("a is 6")
default:
	fmt.Println("a is neither 5 nor 6")
}

// switch 中 case 表达式进行运算
var score int = 88
switch {
case score >= 90:
	fmt.Println("A+")
case score >= 80 && score < 90:
	fmt.Println("A")
default:
	fmt.Println("B")
}

// 在 switch 表达式中声明变量
// 不推荐
switch score := 88 {
case score >= 90:
	fmt.Println("A+")
case score >= 80 && score < 90:
	fmt.Println("A")
default:
	fmt.Println("B")
}


// switch 的穿透 fallthrough
var num int = 10
switch num {
	case 10:
		fmt.Println(10)
		fallthrough  // 当进入case 10时，也会穿透下面的case，直接打印20
	case 20:
		fmt.Println(20)
	case 30:
		fmt.Println(30)
	default:
		fmt.Println(0)
}

// type switch
var x interface{}
var y float64 = 10.0
x = y
switch i := x.(type) {
case nil:
	fmt.Printf("x type: %T, value: %v\n", i, i)
case int:
	fmt.Printf("x type: %T, value: %v\n", i, i)
case float64:
	fmt.Printf("x type: %T, value: %v\n", i, i)
case func(int) float64:
	fmt.Printf("x type: %T, value: %v\n", i, i)
case bool:
	fmt.Printf("x type: %T, value: %v\n", i, i)
default:
	fmt.Printf("unkown type: %T, value: %v\n", i, i)
}

```

注意：
1. 即使 if 语句里面的代码块只有一个语句，也要有大括号。
2. 赋值语句不能作为条件表达式:
```go
if b = false { // error
	//
}
```
3. switch 语句用于基于不同条件执行不同动作，每一个 case 分支都是唯一的，从上到下逐一测试，直到匹配为止；匹配项后面也不需要再加 break；
4. switch/case 后是一个表达式，可以是常量值、变量，也可以是一个有返回值的函数；
5. case 后的各个表达式的值的数据类型，必须和 switch 后的表达式数据类型一致；
6. case 后可以带多个表达式，使用逗号间隔；
7. case 后的表达式如果是常量值（字面量），则要求不能重复
```go
// 这里 5 重复了，错误，但是可以将重复的 5 替换为变量
switch key {
	case 5, 10:
	//
	case 6, 5:  // 可以将这里的5替换为取值为5的变量
	//
	default:
	//
}
```
8. case 后不需要有 break 和大括号，程序匹配到一个 case 后就会执行对应的代码块，然后退出 switch，如果一个都匹配不到，则执行 default；
9. default 语句不是必须的；
10. switch case中fallthrough能够穿透一层case；
11. switch 语句可以用于 type-switch 来判断某个 interface 变量中实际指向的变量类型。


switch 和 if 使用选择：
1. 如果判断的具体数值不多，而且符合整数、浮点数、字符、字符串这几种类型，建议使用 switch 语句，简洁高效；
2. 其他情况，如对区间判断和结果为 bool 类型的判断，使用 if，其使用范围更广。

## 循环控制
```go
for i := 1; i <=10; i++ {
	fmt.Println("i=", i)
}
```

注意：
1. 循环条件是返回一个布尔值的表达式；
2. for 循环的第二种方法：
```go
i := 1
for i <= 10 {
	fmt.Println("i=", i)
	i++
}
```
3. for 循环的第三种方法，通常配合 break 使用，避免死循环：
```go
// 死循环
for {
	fmt.Println("hello")
}

// 等价于
for ; ; {
	fmt.Println("hello")
}

// break
i := 1
for {
	if i <= 10 {
		fmt.Println("i=", i)
	} else {
		break // 跳出循环
	}
	i++
}
```
4. golang 提供 for-range 的方式，可以变量字符串和数组：
```go
var str string = "hello"
for i := 0; i < len(str); i++ { // 按字节遍历，字符串中出现中文（一个中文字符3个字节）时会有问题；可以使用 str2 = []rune(str)
	fmt.Printf("str[%d] = %c\n", i, str[i])
}

// for-range
for index, val := range str { // 按字符遍历，但是 index 按照字节计数
	fmt.Printf("str[%d] = %c\n", index, val)
}
```

在golang中没有 while 循环和 do-while 循环，不过可以使用 for 实现 while 循环的效果：
```go
// while 
var i int = 1
for {
	if i > 10 {
		break
	}
	fmt.Println("hello", i)
	i++
}

// do-while
var i int = 1
for {
	fmt.Println("hello", i)
	i++
	if i > 10 {
		break
	}
}
```

打印金字塔
```go
var totalLevel int = 5
for i := 1; i <= totalLevel; i++ {
	for k := 1; k <= totalLevel-i; k++ {
		fmt.Print(" ")
	}
	for j := 1; j <= 2*i-1; j++ {
		fmt.Print("*")
	}
	fmt.Println()
}
```
打印镂空金字塔
```go
var totalLevel int = 5
for i := 1; i <= totalLevel; i++ {
	for k := 1; k <= totalLevel-i; k++ {
		fmt.Print(" ")
	}
	for j := 1; j <= 2*i-1; j++ {
		if j == 1 || j == 2*i-1 || i == totalLevel {
			fmt.Print("*")
		} else {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
```

判断一个数是否是水仙花数
```go
// 水仙花数：一个3位数的水仙花数是指其各位数字的立方和等于它本身。
// 例如：153 = 1^3 + 5^3 + 3^3
// 下面的代码会输出所有的水仙花数。
var num uint32
fmt.Println("请输入一个三位数：")
fmt.Scanf("%d", &num)
numHundreds := num / 100
numTens := (num / 10) % 10
numOnes := num % 10
if num == uint32(math.Pow(float64(numHundreds), 3)+math.Pow(float64(numTens), 3)+math.Pow(float64(numOnes), 3)) {
	fmt.Printf("%d 是水仙花数\n", num)
} else {
	fmt.Printf("%d 不是水仙花数\n", num)
}
```

生成随机数
```go
import "math/rand"

//
fmt.Println(time.Now().Unix())
fmt.Println(time.Now().Weekday())

var targetNum int = 99
var count int = 0
for {
	n := rand.Intn(100) + 1
	count++
	if n == targetNum {
		break
	}
}
fmt.Printf("生成 %d 循环了 %d 次\n", targetNum, count)
```
golang中标签的使用
```go
label2:
	for i := 0; i < 4; i++ {
	label1:
		for j := 0; j < 5; j++ {
			if j == 2 {
				break label1
			}
			fmt.Println("j=", j)
		}
		if i == 3 {
			break label2
		}
	}
```
break 的注意事项：
1. break 语句出现在多层嵌套的语句块时，可以通过标签指明要终止的哪一层循环；
2. break 默认会跳出最近的循环；

continue 的注意事项：
1. continue 语句用于结束本次循环，继续执行下一次循环；
2. continue 语句出现在多层嵌套的循环语句块中，可以通过标签指明要跳过的是哪一层循环；

goto 的注意事项：
1. goto 和 if、标签搭配使用。不建议在代码中使用 goto;

return 的注意事项：
1. return 返回函数返回值；
2. 用于退出函数，或程序（main函数）；


# 函数
为完成某一功能的程序指令（语句）的集合，称为函数。golang 中，函数分为自定义函数和系统函数。

## 包
包的本质就是创建不同的文件夹来存放程序文件。

golang 的每一个文件都是属于一个包，golang 以包的形式来管理文件和项目目录结构。

包的作用：
1. 区分相同名字的函数、变量等标识符；
2. 当程序文件很多时，可以很好的管理项目；
3. 控制函数、变量等访问范围，即作用域。

```go
// create package
package packageName

// use package
// 从 GOPATH下 src 里面的文件夹开始直到源文件的文件夹
import "packageName Path"

// 调用函数
packageName.funcName
```

包的注意事项：
1. 在给一个文件打包时，该包对应一个文件夹。文件的包名通常与文件所在的文件夹名一致，一般为小写字母；
2. 当一个文件要使用其他包函数或变量时，需要先引入对应的包
```go
import "packageName"
// or
import "packageName Path"


import (
	"packageName1"
	"packageName2"
)
```
3. package 指令在文件的第一行，然后是 import 指令；
4. 在 import 时，路径从 $GOPATH 的 src 下文件夹开始，不用带 src；
5. 为了让其他包的文件，可以访问到本包的函数，需要该函数名首字母大写，类似于其他语言的 public，这样才能跨包使用；
6. 在访问其他包的函数时，语法是 包名.函数名；
7. 如果包名太长，golang 支持给包取别名，但是取别名后原来的包名就不能使用了；
8. 在同一个包下不能有相同的函数名、全局变量名等，否则报重复定义；
9. 如果要编译成一个可执行文件，那么需要将这个包声明为 main 包（package main）。这是一个语法规范。如果要写一个库，包名可以自定义。
10. 编译流程：
```bash
# 基于GOPATH的旧方式
# 先设置 GO111MODULE 为自动模式
go env -w GO111MODULE=auto
cd $GOPATH
# 编译时，编译路径不需要带 src，编译器会自动带
go build learn/proj3
# 编译时需要编译 main 包所在的文件夹
# 项目的目录结构要按照规范组织：$GOPATH/src/项目名/包名/源码文件，main.go 一般放在 main 文件夹（包名）下
# 可以自定编译后可执行文件名
go build -o bin/my.exe learn/proj3


# 基于 GO MOD 的新方式，在 GO 1.11 后引入
mkdir xgo
mkdir xgo/main
touch xgo/main/main.go // 也可以不用把 main.go 放到 main 文件夹，直接放到 xgo 项目文件夹下
mkdir xgo/utils
touch xgo/utils/util.go

go mod init xgo

cat > xgo/utils/util.go <<-"EOF"
package utils

var Name = "Hero"
EOF

cat > xgo/main/main.go <<-"EOF"
package main


import (
    "fmt"
    "xgo/utils"
)

func main() {
    fmt.Println(utils.Name)
}
EOF

go mod tidy
go mode verify

# 先设置 GO111MODULE 开启
go env -w GO111MODULE=auto

# 进入项目，项目不一定要放到 src 下
cd xgo

# 编译
go build -o xgo main/main.go
# or
cd main && go build -o xgo main.go
# or
cd main && go build -o xgo
# 不指定编译后的文件名，则使用默认main
cd main && go build

# 编译成 windows 可执行程序
cd main && GOOS=windows GOARCH=amd64 go build -o xgo.exe main.go
```

**交叉编译**
```bash
# linux 平台下编译生成 windows 可执行程序
GOOS=windows GOARCH=amd64 go build -o demo.exe

# windows 平台下编译生成 linux 可执行程序
GOOS=linux GOARCH=amd64 go build -o demo
```

## 函数调用
函数调用的细节：
1. 在调用一个函数时，会给该函数分配一个新的空间，编译器会通过自身的处理让这个新的空间和其他的栈空间区分开来；
2. 在每个函数对应的栈中，数据空间是分开的，不会混淆；
3. 当一个函数调用完毕（执行完毕）后，程序会销毁这个函数对应的栈空间。

### 递归调用
一个函数在函数体内又调用了本身，称为递归调用。

## 函数注意事项
1. 函数的形参列表可以时多个，返回值列表也可以是多个。当返回值列表是多个时，应该用小括号括起来：
```go
func MultiAdd(n1 int, n2 int) (int, int) {
	return n1 + n2, n1 * n2
}
```
2. 形参列表和返回值列表的数据类型可以是值类型和引用类型；
3. 函数的命名遵循标识符命名规范，首字母不能是数字，首字母大写时可以被本包和其他包文件调用；首字母小写时，只能被本包文件调用，其他包文件不能调用；
4. 函数中的变量是局部变量，函数外不能使用；
5. 基本数据类型和**数组**默认都是值传递，即值拷贝。在函数内修改，不会影响到原来的值；
6. 如果希望函数内的变量能够修改函数外的变量，可以传入变量的地址（如 `&x`），函数内以指针的方式操作变量。类似于引用；
7. golang 函数不支持函数重载；
8. golang 中，函数也是一种数据类型，可以赋值给一个变量，该变量就是一个函数类型的变量，通过该变量可以对函数调用；
9. 函数既然是一种数据类型，因此函数可以作为形参，并且可以被调用；
```go
func Add(n1 int, n2 int) int {
	return n1 + n2
}
func ComAdd(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}

func main() {
	res := ComAdd(Add, 3, 2)
}
```
10. 为了简化数据类型定义，golang 支持自定义数据类型
```go
基本语法：type 自定义数据类型名 数据类型 // 相当于一个别名
案例：type customInt int // 此时 customInt 等价于 int，但是 golang 认为它们不是同一种数据类型
// 例子
type customInt int
var n1 custiomInt = 30
var n2 int
n2 = n1 // error
n2 = int(n1) // right

案例：type customFunc func(int, int) int // 当函数作为形参时可以简化形参定义
// 例子
func ComAdd(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}

type customFunc func(int, int) int
func ComAdd(funvar customFunc, num1 int, num2 int) int {
	return funvar(num1, num2)
}
```
11. golang 支持对函数返回值命名
```go
func cal(n1 int, n2 int) (sum int, sub int) {
	sum = n1 + n2
	sub = n1 - n2
	return
}
```
12. 使用 _ 标识符，忽略返回值
```go
a, _ = MultiAdd(3, 2)
```
13. golang 支持可变参数：
```go
// 支持 0 到多个参数
func sum(args... int) sum int {
	//
}

// 支持 1 到多个参数
func sum(n1 int, args... int) sum int {
	//
}

// args 是切片 slice（动态数组），通过 args[index] 可以访问到各个值
// args 名字可以变，如 xargs
// 可变参数要放到形参列表的后面
func sum(n1 int, xargs... int) s int {
	s = n1
	for i := 0; i < len(xargs); i++ {
		s += xargs[i]
	}
	return
}
```
14. 形参类型一致时，可以合并：
```go
func add(n1, n2 int) int {
	//
}

// 等价于
func add(n1 int, n2 int) int {
	//
}
```


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

## 日期和时间函数
```go
package main

import (
	"fmt"
	"strings"
	"time"
)

func concat() {
	var b strings.Builder
	for i := 0; i < 10000000; i++ {
		b.WriteString("xjz")
	}
	fmt.Println("concat() 执行完毕，字符串长度:", b.Len())
}

func main() {
	// 获取当前时间
	now := time.Now()
	// 打印当前时间
	fmt.Println("当前时间:", now)

	// 获取年、月、日、时、分、秒
	fmt.Printf("年: %v\n", now.Year())
	fmt.Printf("月: %v\n", int(now.Month()))
	fmt.Printf("日: %v\n", now.Day())
	fmt.Printf("时: %v\n", now.Hour())
	fmt.Printf("分: %v\n", now.Minute())
	fmt.Printf("秒: %v\n", now.Second())

	// 格式化日期和时间
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n",
		now.Year(), int(now.Month()), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	dateStr := fmt.Sprintf("%d-%02d-%02d %02d:%02d:%02d\n",
		now.Year(), int(now.Month()), now.Day(),
		now.Hour(), now.Minute(), now.Second())
	fmt.Println("格式化后的时间:", dateStr)
	// 该时间 2006-01-02 15:04:05 是 Go 语言的时间格式化标准，不能改变
	fmt.Printf("格式化后的时间: %s\n", now.Format("2006-01-02 15:04:05"))
	fmt.Printf("格式化后的时间: %s\n", now.Format("01-02"))

	// 结合 sleep 了解时间常量
	// 必须用 100 乘以 time.Millisecond
	// 不能用 0.1 * time.Second
	// 这是 golang 设计者的一个约定
	// time.Hour, time.Minute, time.Second, time.Millisecond
	// time.Microsecond, time.Nanosecond
	time.Sleep(100 * time.Millisecond)
	fmt.Println("睡眠 0.1 秒后继续执行")

	// 获取当前时间的 Unix 时间戳，距离 1970 年 1 月 1 日的秒数
	unixTime := now.Unix()
	fmt.Println("当前时间的 Unix 时间戳:", unixTime)
	// 获取当前时间的 Unix 纳秒时间戳
	unixNanoTime := now.UnixNano()
	fmt.Println("当前时间的 Unix 纳秒时间戳:", unixNanoTime)

	// 统计程序执行的耗时
	start := time.Now().UnixNano()
	concat()
	end := time.Now().UnixNano()
	fmt.Println("concat() 执行耗时:", end-start, "纳秒")
	fmt.Printf("concat() 执行耗时: %.3f 秒\n", float64(end-start)/1000/1000/1000)
}
```

## golang内置函数
golang为了编程方便，提供了一些函数，这些函数可以直接使用，称为golang的内值函数。

常用内置函数：
1. len：用来求数据的长度，如string, array, slice, map, channel
2. new：用来分配内存，主要用来分配值类型，如int, float32, struct 等，**返回的是指针**
```go
// num 是指向一个 int 的指针
num := new(int)
// 默认整数空间存默认0，更改num指向的空间的值
*num = 1 
```
3. make：用来分配内存，主要用来分配应用类型，如channel, map, slice


# 异常处理
1. golang 中不支持 try...catch...finally, try...except
2. golang 中引入 defer, panic, recover
3. golang 中抛出一个 panic 异常，然后在 defer 中通过 recover 捕获这个异常，然后正常处理


## 自定义错误
在 golang 程序中，也支持自定义错误，使用 errors.New 和 panic 内置函数。
1. errors.New("错误说明") 会返回一个 error 类型的值，表示一个错误
2. panic 内置函数，接收一个空接口 interface() 类型的值（也就是任何值了）作为参数。接收 error 类型的变量，输出错误信息，并退出程序

```go
package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) int {
	// 使用 defer + recover 来处理除以零的异常
	defer func() {
		// recover 是内置函数
		err := recover() // golang 中的 recover 函数可以捕获 panic
		// 如果发生 panic，err 将不为 nil
		if err != nil {
			fmt.Println(err)
			b += 1 // 如果发生异常，b 加 1，避免除以零
			fmt.Println("Adjusted b to avoid division by zero:", b)
		}
	}()
	res := a / b
	return res
}

// 函数读取一个配置文件 config.ini 的信息
// 如果文件名不正常，返回一个自定义的错误
func readConfigFile(filename string) (err error) {
	if filename == "config.ini" {
		// read
		return nil
	} else {
		// 返回一个自定义的错误
		return errors.New("读取文件错误")
	}
}

func read() {
	err := readConfigFile("config.ini")
	if err != nil {
		panic(err) // 如果发生错误，触发 panic
	}
	fmt.Print("读取配置文件成功...\n")
}

func main() {
	a := 10
	b := 0 // 这里故意设置为0来触发异常
	result := divide(a, b)
	fmt.Println("The result of division is:", result)
	fmt.Println("done!")

	read() // 调用读取配置文件的函数
}
```

# 数组
数组可以存放多个同一数据类型的数据。数组也是一种数据类型。数组是值类型。

1. 定义完数组后，会有默认值，默认值的取决于数组定义的数据类型，如 var arr [3]int
2. 数组的地址可以通过数组名来获取，如 &arr
3. 数组的第一个元素的地址就是数组的首地址，即 &arr[0] = &arr
4. 数组的第二个元素的地址等于第一个元素的地址加上数据类型所占的字节数，如 &arr[1] = &arr[0] + 8，int 占8个字节

数组初始化
```go
var arr [3]int = [3]int{1, 2, 3}
var arr2 = [3]int{4, 5, 6}
var arr3 = [...]int{7, 8, 9} // 使用...自动推断长度
var arr4 = [3]int{1: 10, 2: 20} // 指定索引初始化
var arr5 = [...]int{1: 9, 2: 8, 3: 7} // 使用...自动推断长度，指定索引初始化
arr6 := [...]string{1: "a", 2: "b", 3: "c"} // 使用...自动推断长度，指定索引初始化
```

数组的遍历
```go
var arr = [...]int{1, 3, 5}
for i := 0; i < len(arr); i++ {
	fmt.Println(arr[i])
}

for index, value := range arr {
	fmt.Println(arr[index])
	fmt.Println(value)
}
```

数组注意事项
1. 数组是多个相同类型数据的组合，一个数据一旦声明或定义，其长度是固定的，不能改变
2. var arr []int 如果不指定数组长度，则是切片 slice
3. 数组中的元素可以是任何数据类型，包括值类型和引用类型。不能混用
4. 数组创建后，如果没有赋值，则使用默认值
5. 使用数组的步骤：1. 声明数组并开辟内存空间；2. 给数组各个元素赋值；3. 使用数组
6. 数组的下标是从0开始的
7. 数组下标必须在指定范围内使用，否则报 panic 数组越界
8. golang 的数组属于值类型，在默认情况下是值传递，因此会进行值拷贝，数组间不会相互影响
9. 如果想在其他函数中去修改原来的数组，可以使用引用传递（指针）。默认传给函数的数组是一种值拷贝，不会改变原来数组的值。
```go
var arr = [...]int{1, 2, 3}
func changeArr(arr *[3]int) {
	(*arr)[0] = 99
}

changeArr(&arr)
fmt.Println(arr)
```
10. 长度是数组类型的一部分，在传递函数参数时，需要考虑数组的长度。

# 切片 slice

基本介绍
1. 切片是数组的一个引用，因此切片是引用类型，在进行传递时，遵守引用传递的机制
2. 切片的使用和数组类似，遍历切片、访问切片的元素和求切片长度都一样
3. 切片的长度可以是变化的，因此切片是一个可动态变化的数组
4. 切片的基本语法：
```go
// var 变量名 []数据类型

var sli []int

// 数组
var arr = [...]int{1, 2, 3, 4, 5}
var sli []int = arr[1:3]
fmt.Println("slice 的容量", cap(sli))
```
切片(var sli []int = arr[1:3])的容量 cap(slice) 是动态变化的，根据切片元素的个数动态调整。切片从底层来说，其实就是一个数据结构 struct 结构体
```go
type sli struct {
	ptr *[2]int
	len int
	cap int
}
```
改变 sli 的值会改变 arr 的值，因为切片是引用。

切片使用的三种方式：
1. 定义一个切片，然后让切片去引用一个已经创建好的数组。直接引用一个数组，这个数值实现存在，开发人员可见。
2. 通过 make 来创建切片。make 也会创建一个数组，是由切片在底层进行维护，开发人员是不可见的：
```go
var 切片名 []type = make([]type, len, [cap])
// type 是数据类型
// len 表示切片大小
// cap 表示切片容量，可选，但不小于 len
var intSlice []int = make([]int, 4, 10)
```
3. 定义一个切片时，直接指定具体数组，原理类似于 make 的方式
```go
var strSlice []string = []string{"liubei", "guanyu", "zhangfei"}
```

切片注意事项
1. 切片初始化时 var slice = arr[startIndex:endIndex]，从 arr 数组小标为 startIndex， 取到下标 endIndex 的元数（不含 arr[endIndex]）
2. 切片初始化时，仍然不能越界，范围在 [0, len(arr)]之间，但是可以动态增长
3. var slice = arr[0:end] 可以简写 var slice = arr[:end]
4. var slice = arr[start:len(arr)] 可以简写 var slice = arr[start:]
5. var slice = arr[0:len(arr)] 可以简写 var slice = arr[:]
6. cap 是一个内置函数，用于统计切片的容量，即最大可以存放多少个元素
7. 切片定义后，还不能使用，因为本身是一个空的，需要让其引用到一个数组，或者 make 一个空间供切片来使用
8. 切片可以继续切片
9. 用 append 内置函数，可以对切片进行动态追加
```go
var slice []int = []int{100, 200, 300}
append(slice, 400, 500)
```
10. 切片 append 操作的本质就是对数组扩容
11. golang 底层会创建一个新的数组 newArr（按照扩容后大小），将切片原来包含的元素拷贝到新的数组 newArr，切片重新引用到 newArr，注意 newArr 是在底层来维护的，开发人员不可见。
12. 切片的拷贝操作，可以使用内置拷贝函数完成
```go
// 切片拷贝
var sli4 []int = []int{1, 2, 3, 4, 5}
var sli5 = make([]int, 10)
copy(sli5, sli4) // 拷贝 sli4 到 sli5, sli4 and sli5 在内存中有独立的空间
fmt.Println("Original slice:", sli4)
fmt.Println("Copied slice:", sli5)

// sli4 和 sli5 的数据空间是独立的，互相不影响，也就是说， 改变 sli4[0] = 99，不会影响 sli5[0] 的值
// 如果 sli5 的空间不是10，而是3，小于 sli4 的空间元素个数，那么程序也不会报错，只会将 sli4 的前3个元素拷贝到 sli5
```

## string and slice
1. string 底层是一个 byte 数组，因此 string 也可以进行切片处理
2. string 是不可变的，也就是说不能通过 str[0] = 'z' (var str = "abc") 的方式来修改字符串
3. 如果需要修改字符串，那么可以先将 string 转成切片（[]byte 或 []rune），然后再修改，之后重写成 string
```go
// 修改字符串的值
var str2 = "hello, world"
sli10 := []byte(str2) // 二进制的数据要使用这种方式
sli10[0] = 'z' // 修改切片的第一个字节
fmt.Println("Modified byte slice:", sli10)
str2 = string(sli10) // 转换回字符串
fmt.Println("String after modification:", str2)

// 转成 []byte 后，可以处理英文和数字，但不能处理中文
// 原因是 []byte 只处理单字节字符，而中文字符通常是3字节的，因此会出现乱码
// 解决方法是将 string 转成 []rune 即可，因为 []rune 可以处理按字符处理，兼容中文
str3 := "你好，世界"
sli11 := []rune(str3)
sli11[0] = '都' // 修改切片的第一个字符
fmt.Println("Modified rune slice:", sli11)
str3 = string(sli11) // 转换回字符串
fmt.Println("String after rune modification:", str3)
```

# 排序和查找
排序是将一组数据，依据指定的顺序进行排列的过程。排序的分类：
1. 内部排序：将需要处理的所有数据都加载到内部存储器中进行排序，如交换式排序法、选择式排序法和插入式排序法
2. 外部排序：数据量过大，无法全部加载到内存中，需要借助外部存储进行排序，如合并排序法、直接合并排序法

## 交换式排序
交换式排序属于内部排序法，是运用数据值比较厚，依据判断规则对数据位置进行交换，以达到排序的目的。交换式排序法包括两种：
1. 冒泡排序法（Bubble sort）
2. 快速排序法（Quick sort）

### 冒泡排序法
冒泡排序的基本思想是通过对待排序序列从后向前（从下标较大的元素开始），依次比较相邻元素的排序码，若发现逆序则交换，使排序码较小的元素逐渐从后部移向前部（从下标较大的单元移向下标较小的单元），将像水底下的气泡一样逐渐向上冒泡。

因为排序的过程中，各元素不断接近自己的位置，如果一趟比较下来没有进行过交换，就说明序列有序，因此要在排序过程中设置一个 flag 进行判断元素是否进行过交换，从而减少不必要的比较，加快程序执行速度。

```go
package main

import (
	"fmt"
)

// 冒泡排序
// Bubble Sort
func bubbleSort(arr *[6]int) {
	for i := 0; i < len(*arr)-1; i++ {
		flag := false // 优化标志
		// 每次遍历，假设没有交换
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				// 交换位置
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
				flag = true // 发生了交换
			}
		}
		if !flag {
			// 如果没有发生交换，说明数组已经有序，可以提前结束
			break
		}
	}
}

func main() {
	// 定义数组
	arr := [...]int{5, 2, 9, 1, 5, 6}
	fmt.Println("排序前", arr)
	bubbleSort(&arr)
	fmt.Println("排序后", arr)
}
```
## 查找
在 golang 中，常用的查找有两种：
1. 顺序查找
2. 二分查找

```go
package main

import (
	"fmt"
)

// 二分查找
func binarySearch(arr *[]int, leftIndex int, rightIndex int, target int) int {
	if leftIndex > rightIndex {
		return -1 // 没有找到目标元素，返回 -1
	}
	
	midIndex := (leftIndex + rightIndex) / 2 // 计算中间索引
	if (*arr)[midIndex] == target {
		return midIndex // 找到目标元素，返回索引
	} else if (*arr)[midIndex] > target {
		// 目标元素在左半部分
		return binarySearch(arr, leftIndex, midIndex-1, target)
	} else {
		// 目标元素在右半部分
		return binarySearch(arr, midIndex+1, rightIndex, target)
	}
}


func main() {
	arr2 := []int{1, 2, 3, 4, 5, 6}
	target := 4
	index := binarySearch(&arr2, 0, len(arr2)-1, target)
	fmt.Printf("目标元素 %d 在数组中的索引是: %d\n", target, index)
}
```

# 二维数组
二维数组是多维数组的一种。
```go
package main

import (
	"fmt"
)

func main() {
	// 定义或声明一个二维数组
	var arr [4][6]int
	// 初始化二维数组
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3

	// 打印二维数组
	fmt.Println("二维数组 arr:", arr)

	// 遍历二维数组
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}

	fmt.Println(len(arr))
	fmt.Println(len(arr[0]))
}
```
二维数组 arr [2][3]int 指向一个数组，里面存放2个指针（地址），每一个指针指向一个数组，里面存放3个整数值。所以 arr[0] 的地址和 arr[0][0] 的地址一样；arr[1] 的地址和 arr[1][0] 的地址一样。

二维数组的声明：
```go
// 定义或声明一个二维数组
var arr2 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
// 打印二维数组
fmt.Println("二维数组 arr2:", arr2)

var arr3 [2][3]int = [...][3]int{{1, 2, 3}, {4, 5, 6}}
var arr4 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
var arr5 = [...][3]int{{1, 2, 3}, {4, 5, 6}}
arr6 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
arr7 := [...][3]int{{1, 2, 3}, {4, 5, 6}}
fmt.Println("二维数组 arr3:", arr3)
fmt.Println("二维数组 arr4:", arr4)
fmt.Println("二维数组 arr5:", arr5)
fmt.Println("二维数组 arr6:", arr6)
fmt.Println("二维数组 arr7:", arr7)

// 遍历
for i := 0; i < len(arr3); i++ {
	for j :=0; j < len(arr3[i]); j++ {
		fmt.Printf("%v\t", arr3[i][j])
	}
	fmt.Println()
}

for i, v := range arr3 {
	for j, v2 := range v {
		fmt.Printf("arr3[%d][%d] = %d\t", i, j, v2)
	}
	fmt.Println()
}
```

# 映射 map

map 是 key-value 数据结构，又称为字段或者关联数组，类似其他编程语言的集合或字典。

基本语法：
var 变量名 map[keytype]valuetype

key 可以是很多种类型，如 bool, int, string, 指针，chann，结构，结构体，数组，**通常情况下为 int，string**。但是 slice，map，function 不可以，因为它们不能用 == 来判断。

value 的类型和key基本是一样的。通常为 int, float, string, map, struct

map 的声明：
var m map[string]string
var m map[string]int
var m map[int]string
var m map[string]map[string]string
注意：map 声明不会分配内存，只有初始化（make）后才分配内存，分配内存后才能赋值和使用。

map 的使用方式
```go
// 方式1
var cities map[string]string
cities = make(map[string]string, 10)
cities["cityeName"] = "上海"

// 方式2
var cities = make(map[string]string)
cities["cityeName"] = "上海"

// 方式3
var cities map[string]string = map[string]string {
	"cityName": "成都",

}
cities["cityeName"] = "上海"

// 方式4
cities := map[string]string {
	"cityName": "成都",
	
}
cities["cityeName"] = "上海"
```

map 嵌套
```go
// map 的嵌套
// 创建一个学生信息的嵌套 map
// 学生姓名作为键，值是另一个 map，包含年龄、城市和爱好等信息
// 例如: map[string]map[string]string
studentMap := make(map[string]map[string]string)
studentMap["xujinzh"] = make(map[string]string)
studentMap["xujinzh"]["age"] = "18"
studentMap["xujinzh"]["city"] = "成都"
studentMap["xujinzh"]["hobby"] = "编程"

studentMap["lisi"] = make(map[string]string)
studentMap["lisi"]["age"] = "20"
studentMap["lisi"]["city"] = "上海"
studentMap["lisi"]["hobby"] = "音乐"	
fmt.Println("学生信息:", studentMap)
for student, info := range studentMap {
	fmt.Println("学生:", student)
	for key, value := range info {
		fmt.Printf("  %s: %s\n", key, value)
	}
}
```

map 增删改查
```go
// map 的增删改查
cities := make(map[string]string)
cities["no1"] = "北京"
cities["no2"] = "上海"
cities["no3"] = "广州"	
fmt.Println(cities)

// 增加
cities["no4"] = "深圳"
fmt.Println("增加后:", cities)

// 修改
cities["no2"] = "杭州"
fmt.Println("修改后:", cities)

// 删除
delete(cities, "no3") // 内置函数
fmt.Println("删除后:", cities)

// 查询
if city, exists := cities["no1"]; exists {
	fmt.Println("查询到的城市:", city)
} else {
	fmt.Println("没有查询到城市")
}
```

map 删除细节
1. 如果要删除 map 的所有 key，没有专门的方法一次删除。只能遍历逐个删除
2. 或者 map = make(), make 一个新的，让原来的成为垃圾，被 GC 回收

map 遍历
```go
// map 的遍历
for key, value := range cities {
	fmt.Printf("key: %s, value: %s\n", key, value)
}

// 嵌套 map 的遍历
studentMap := make(map[string]map[string]string)
studentMap["xujinzh"] = make(map[string]string)
studentMap["xujinzh"]["age"] = "18"
studentMap["xujinzh"]["city"] = "成都"
studentMap["xujinzh"]["hobby"] = "编程"

studentMap["lisi"] = make(map[string]string)
studentMap["lisi"]["age"] = "20"
studentMap["lisi"]["city"] = "上海"
studentMap["lisi"]["hobby"] = "音乐"	
fmt.Println("学生信息:", studentMap)
for student, info := range studentMap {
	fmt.Println("学生:", student)
	for key, value := range info {
		fmt.Printf("  %s: %s\n", key, value)
	}
}
```

map 的长度
```go
var cities := make(map[string]string)
len(cities)
```

map 切片
切片的数据类型如果是 map，叫做 slice of map，map 切片中 map 个数可以动态变化。map 切片要初始化两次。

```go
// map 切片

// 声明一个 map 切片
var sliceOfMaps = make([]map[string]string, 2)
// 初始化 map 切片
if sliceOfMaps[0] == nil {
	sliceOfMaps[0] = make(map[string]string, 2)
	sliceOfMaps[0]["name"] = "xujinzh"
	sliceOfMaps[0]["age"] = "18"

}
if sliceOfMaps[1] == nil {
	sliceOfMaps[1] = make(map[string]string, 2)
	sliceOfMaps[1]["name"] = "lisi"
	sliceOfMaps[1]["age"] = "20"
}

// 使用切片 append 函数，动态增加 map
newMap := make(map[string]string, 2)
newMap["name"] = "wangwu"
newMap["age"] = "22"
sliceOfMaps = append(sliceOfMaps, newMap)
// 打印 map 切片
fmt.Println(sliceOfMaps)
```

map 排序
1. golang 中没有一个专门的方法针对 map 的 key 进行排序
2. golang 中的 map 默认是无序的，也不是按照添加的顺序存放的，每次遍历的输出可能也不一样
3. golang 中 map 的排序是先将 key 进行排序，然后根据 key 遍历输出

```go
//  map 排序
mapVar := make(map[int]int, 10)
mapVar[3] = 30
mapVar[1] = 10
mapVar[2] = 20
mapVar[5] = 50
mapVar[4] = 40
mapVar[6] = 60
fmt.Println("未排序的 map:", mapVar)
// 使用切片存储 map 的键
keys := make([]int, 0, len(mapVar))
for key := range mapVar {
	keys = append(keys, key)
}
// 对键进行排序
sort.Ints(keys)
// 或者使用冒泡排序（不推荐，因为效率低）
// for i := 0; i < len(keys)-1; i++ {
// 	for j := i + 1; j < len(keys); j++ {
// 		if keys[i] > keys[j] {
// 			keys[i], keys[j] = keys[j], keys[i]
// 		}
// 	}
// }
fmt.Println("排序后的键:", keys)
// 打印排序后的 map
fmt.Println("排序后的 map:")
for _, key := range keys {
	fmt.Printf("key: %d, value: %d\n", key, mapVar[key])
}
```

map 使用细节
1. map 是引用类型。在一个函数接收 map，修改后，会直接修改原来的 map
2. map 的容量达到后，再向 map 中增加元素，会自动扩容，并不会发生 panic，也就是说 map 能动态增长
3. map 的 value 也经常使用 struct 类型，更适合管理复杂的数据
```go
// 定义一个结构体
type Stu struct {
	Name    string
	Age     int
	Address string
}


// map 的值用结构体
students := make(map[string]Stu)
// 创建学生
stu1 := Stu{"tom", 18, "Beijing"}
stu2 := Stu{"lisa", 18, "Chongqing"}
students["no1"] = stu1
students["no2"] = stu2
fmt.Println(students)

// 遍历
for k, v := range students {
	fmt.Println(k)
	fmt.Println(v.Name)
	fmt.Println(v.Age)
	fmt.Println(v.Address)
}
```

判断 key 是否在 map 中
```go
if v, ok := map1[key]; ok {
	// exists
}

if map1[key] != nil {
	// exists
}
```

# 面向“对象”编程与结构体
golang 的面向对象编程跟 python, c++, java 不同，没有类的概念，用结构体替代类（class）的概念，因此结构体类似于 class。

golang 面向对象编程说明
1. golang 支持面向对象编程(OOP)，但与传统的面向对象编程有区别，不是纯粹的面向对象语言，所以 golang 支持面向对象编程特性更合适
2. golang 没有类 class，golang 的结构体和其他编程语言的类有相等的地位，可以理解golang是基于结构体来实现面向对象编程
3. golang 面向对象编程非常简洁，去掉了传统面向对象编程的继承、方法重载、构造函数和析构函数、隐藏的this指针、self等
4. golang 支持面向对象编程的继承、封装和多态的特性，只是实现方式和其他面向对象编程语言不同，比如继承：golang 没有 extends 关键字，继承是通过匿名字段来实现
5. golang 面向对象编程很优雅，OOP本身就是语言类型系统(type system)的一部分，通过接口关联，耦合性低，也非常灵活。因此，golang 面向接口编程时非常重要的特性。

## 结构体

结构体是值类型。结构体字段（field），也叫属性，字段可以是基本数据类型、数组，也可以是引用类型。

1. 结构体是自定义的数据类型，代表一类事物
2. 结构体变量（实例）是具体的，代表一个具体变量

字段的使用说明
1. 字段声明语法同变量
2. 字段的类型可以是基本类型、数组、引用类型
3. 创建一个结构体变量后，如果没有给字段赋值，都对应一个零值（默认值），规则同前面讲的一样

|类型|默认值|
|---|---|
|bool|false|
|int|0|
|string|""|
|array [2]int|[0, 0]|
|ptr *int| nil，表示没有分配内存空间|
|slice []int |nil，表示没有分配内存空间|
|map map[int]int |nil，表示没有分配内存空间|

指针、切片、映射的默认值是 nil，表示还没有分配内存空间，如果需要使用，那么须先 make。

4. 不同结构体变量的字段是独立的，互相不影响，一个结构体变量字段的更改，不影响另外一个。

结构体创建方式
```go
// 方式一
var cat1 Cat
cat1.Name = "xiaobai"
cat1.Age = 1
cat1.Color = "white"
fmt.Println(cat1)

// 方式二
cat2 := Cat{"xiaohua", 2, "color"}
fmt.Println(cat2)

// 方式三
var cat3 *Cat = new(Cat)
// golang 设计者为了程序员方便，底层进行了优化
// 也可以用 cat3.Name = "xiaoqing"
(*cat3).Name = "xiaoqing"
fmt.Println(*cat3)

// 方式四
var cat4 *Cat = &Cat{}
// var cat4 *Cat = &Cat{"xiaowang"}
(*cat4).Name = "xiaowang"
fmt.Println(*cat4)
```
第三种和第四种方式返回的是结构体指针。结构体指针访问字段的标注方式应该是 `(*结构体指针变量).字段名`，当 golang 为了开发人员方便，做了底层简化，可以使用 `结构体指针变量.字段名` 迎合开发人员。

结构体使用细节
0. 结构体中也可以不定义字段，即空的结构体
1. 结构体的所有字段在内存中是连续的
```go
type Point struct {
	x int
	y int
}

type Rect struct {
	leftUp    Point
	rightDown Point
}

type Rect2 struct {
	leftUp    *Point
	rightDown *Point
}

// struct 字段在内存中是连续的
rect := Rect{Point{1, 2}, Point{3, 4}}
// rect 的四个int，在内存中是连续分布的
// 0xc000018200 0xc000018208 0xc000018210 0xc000018218
fmt.Println(&rect.leftUp.x, &rect.leftUp.y, &rect.rightDown.x, &rect.rightDown.y)

rect2 := Rect2{&Point{10, 20}, &Point{30, 40}}
// 0xc0000140a0 0xc0000140a8
fmt.Println(&rect2.leftUp, &rect2.rightDown)
```
2. 结构体是用户单独定义的类型，和其他类型进行转换时，要求结构体的字段要完全一样（包括名字、个数和类型）
```go
type A struct {
	Num int
}

type B struct {
	Num int
}

// 结构体转换
var a A
var b B
a = A(b)
fmt.Println(a, b)
```
3. 结构体进行 type 重新定义（相当于取别名），golang 认为是新的数据类型，但是相互间可以强转
```go
type Student struct {
	Name string
	Age int
}

type Stu Student // 取别名

func main() {
	var stu1 Student
	var stu2 Stu
	stu2 = stu1 // error
	stu2 = Stu(stu1) // right
	fmt.Println(stu1, stu2)
}


type integer int

var i integer = 10
var j int = 20
j = i // error
j = int(i) // right
```
4. struct 的每个字段上，可以写上一个 tag，该 tag 可以通过反射机制获取，常见的使用场景就是序列号和反序列化
```go
type Dog struct {
	Name  string `json:"name"`  // struct tag
	Age   int    `json:"age"`
	Color string `json:"color"`
}

// 将结构体序列号为json字符串
dog1 := Dog{"xiaohei", 1, "gray"}
jsonStr, err := json.Marshal(dog1)
if err != nil {
	fmt.Println("error")
} else {
	fmt.Println(jsonStr)
	fmt.Printf("%q\n", string(jsonStr))
}

// "{\"name\":\"xiaohei\",\"age\":1,\"color\":\"gray\"}"
```
将大写的变量名使用 struct tag 反射为小写的字符串，方便传输。

## 方法 Method
在某些情况下，需要声明或定义方法。如 Person 结构体除了有字段年龄、姓名外， 还可能有一些行为，如跑步、学习等，这时就需要有方法来定义。

golang 中方法是作用在指定的数据类型上（即和指定的数据类型绑定），因此（type）自定义类型都可以有方法，而不仅仅是结构体。
```go
type Dog struct {
	Name  string `json:"name"` // struct tag
	Age   int    `json:"age"`
	Color string `json:"color"`
}

// 结构体 Dog 的方法
func (dog Dog) Eat() {
	fmt.Println(dog.Name, "啃骨头")
}

// 方法
dog2 := Dog{"阿黄", 2, "yellow"}
dog2.Eat()
```
1. Eat 方法和 Dog 类型绑定
2. Eat 方法只能通过 Dog 类型的变量来调用，而不能直接调用，也不能使用其他类型变量来调用
3. dog Dog 表示它属于哪个结构体

### 方法的调用和传参机制原理
方法的调用和传参机制，与函数基本一样，不一样的地方是方法调用时，也会将调用方法的变量，当做实参也传递给方法。

```go
func (instance Type) mathodName(参数列表) (返回值列表) {
	方法体
	return 返回值
}
```
这里的 Type 不一定是结构体。如果没有返回值列表，可以不用 return.

实例：

```go
type Circle struct {
	radius float64
}

func (circle Circle) Area() float64 {
	return math.Pi * math.Pow(circle.radius, 2)
}

// 计算圆的面积
circle := Circle{4.0}
area := circle.Area()
fmt.Printf("圆的面积：%.3f\n", area)
```

方法注意事项
1. 结构体类型是值类型，在方法调用中，遵守值类型的传递机制，是值拷贝传递方式
2. 如果希望在方法中修改结构体变量的值，可以通过结构体指针的方式来处理
```go
func (circle *Circle) Area() float64 {
	return math.Pi * math.Pow(circle.radius, 2) // 等价于下面，golang底层进行了优化
	// return math.Pi * math.Pow((*circle).radius, 2)
}

circle := Circle{4.0}
area := circle.Area() // 等价于下面，golang底层进行了优化
// area := (&circle).Area()
```
3. golang 中的方法作用在指定的数据类型上（和指定的数据类型绑定），因此，对于自定义的类型都可以有方法，而不仅仅是结构体，比如 int, float64 等都可以有方法
```go
type integer int

func (i integer) printi() {
	fmt.Println(i)
}

func (i *integer) changei() {
	i += 1
	// *i += 1
}

func main() {
	var i integer = 10
	i.printi()

	i.changei()
	// (&i).changei()
	fmt.Println(i) // i=11
}
```
4. 方法的访问范围的控制规则和函数一样。方法名首字母小写，只能在本包中访问，方法名首字母大写，可以在本包和其他包访问
5. 如果一个变量实现了 String() 这个方法，那么 fmt.Println 默认会调用这个变量的 String() 进行输出
```go
type Person struct {
	Name string
	Age  int
}

// 给 Person 实现 String() 方法
func (person *Person) String() string {
	str := fmt.Sprintf("Name=[%v], Age=[%v]\n", person.Name, person.Age)
	return str
}

// String() 方法改变 fmt.Println() 的行为
person := Person{
	Name: "xujinzh",
	Age:  18, // 最后一行也要有逗号
}
fmt.Println(&person) // Name=[xujinzh], Age=[18]
// 如果没有实现 String() 方法，则打印的是 person 的地址
```

### 函数和方法的区别
1. 调用方式不同
```go
// 函数调用：函数名（参数列表）
// 方法调用：变量.方法名（参数列表）
```
2. 对于函数，参数为值类型时，不能将指针类型的数据直接传递，反之依然
```go
type Person struct {
	Name string
}

// 函数
// 值传递
func test01(person Person) {
	fmt.Println(person.Name)
}

// 指针引用传递
func test02(person *Person) {
	fmt.Println(person.Name)
}

func main() {
	person := Person("tom")
	test01(person) // right
	// test01(&person) // error
	test02(&person) // right
	// test02(person) // error
}
```
3. 对于方法（以struct方法为例），参数列表为值类型时，可以用指针类型变量调用，反之依然
```go
type Person struct {
	Name string
}

func (person Person) test03() {
	fmt.Println(person.Name)
}

func (person *Person) test04() {
	fmt.Println(person.Name)
}

func main() {
	person := Person{"jack"}
	person.test03() // right
	(&person).test03() // right
	(&person).test04() // right
	person.test04() // right

	// 因为golang的底层优化，所以使用变量person和(&person)都可以，两者等价
	// 具体传给方法的参数是值类型还是（指针）引用类型，看方法定义中参数的类型
	// 如果方法定义中是值类型，那么传入值变量，即使(&person)也会底层转为person再传入
	// 如果方法定义中是引用类型，那么传入指针引用变了，即使person也会底层转为(&person)再传入
}
```

### 结构体初始化
```go
type Student struct {
	Name string
	Age int
}

var student1 = Student{"xiaoming", 18} // 初始化值的顺序必须和结构体参数对应上
student2 := Student{"xiaowang", 19} // 初始化值的顺序必须和结构体参数对应上
var student3 = Student{ // 初始化值的顺序可以不用和结构体参数对应上，因为已经指明
	Age : 20,
	Name : "xiaoai",
}

student4 := Student{ // 初始化值的顺序可以不用和结构体参数对应上，因为已经指明
	Age : 20,
	Name : "xiaohua",
}

// 返回结构体的指针类型
var student5 *Student = &Student{"xiaoxiong", 21}

student6 := &Student{
	Name : "xiaoma",
	Age : 22,
}
```

## 工厂模式
golang 的结构体没有构造函数，可以使用工程模式来解决这个问题。

```go
package model

type Student struct {
	Name string
}
```
结构体 Student 的首字母是大写，如果要在其他包（如 main 包）中 Student 实例，引入 model 包后，就可以直接创建 Student 结构体变量。

```go
package model

type student struct {
	Name string
}
```

但是，当首字母是小写（type student struct）时，就不可以使用了，此时可以采用工厂模式来解决。

即，既想用小写的结构体名，又想在其他包中使用该结构体，可以用工厂模式。

创建 model 包：
```go
package model

type student struct { // 结构体名是小写
	Name string
	score float64  // 字段名也是小写
}

func (stu *student) GetScore() float64 {
	return stu.score
}

// 因为 student 结构体首字母是小写，因此是只能在 model 包中使用
// 使用工厂模式来解决
func NewStudent(n string, s float64) *student {
	return &student{
		Name : n,
		score : s,
	}
}
```
在 main 包中调用：
```go
package main

import (
	"fmt"
)

func main() {
	stu := model.NewStudent("xiaoming", 88.8) // 返回的是指针
	fmt.Println(*stu)
	fmt.Println(stu.Name, stu.GetScore) // 底层已转化
}
```

## 抽象
前面定义结构体的时候，实际上就是把一类事物共有的属性（字段）和行为（方法）提取出来，形成一个物理模型（模板、结构体）。这种研究问题的方法称为**抽象**。

## 面向对象编程的三大特性
golang 仍然有面向对象编程的继承、封装和多态特性，只是实现方式有所不同。

### 封装
封装（encapsulation）是把抽象出的字段和对字段的操作封装在一起，数据被保护在内部，程序的其他包只有通过被授权的操作（方法）才能对字段进行操作。

封装的好处
1. 隐藏实现的细节
2. 可以对数据进行验证，保证安全合理

如何体现封装
1. 对结构体中的属性进行封装
2. 通过方法、包，来实现封装

封装的实现步骤
1. 将结构体、字段（属性）的首字母小写，这将使其不能被导出，即其他包不能使用，类似 private
2. 给结构体所在包提供一个工厂模式的函数，首字母大写，类似于一个构造函数
3. 提供一个首字母大写的 SetXXX 方法，用于对属性判断并赋值
```go
func (person *Person) SetAge(参数列表) (返回值列表) {
	// 数据验证的业务逻辑
	person.字段 = 参数新值
}
```
4. 提供一个首字母大写的 GetXXX 方法，用于获取属性的值
```go
func (person *Person) GetAge() {
	return person.Age
}
```

### 继承
继承（inheritance）可以解决代码复用。

当多个结构体存在相同的属性或字段、方法时，可以从这些结构体中抽象出结构体，在该结构体中定义这些相同的属性和方法。

其他结构体不需要重新定义这些属性和方法，只需嵌套一个匿名结构体即可。也就是说，**golang 中，如果一个 struct 嵌套了另一个匿名结构体，那么这个结构体可以直接访问匿名结构体的字段和方法，从而实现继承特性**。

```go
type Goods struct {
	Name string
	Price float64
}

type Book struct {
	Goods // 嵌套匿名结构体, 即只有类型 Goods 没有变量名，如 goods
	Writer string
}
```

示例
```go
package main

import (
	"fmt"
)

type Student struct {
	Name  string
	Age   int
	Score int
}

func (student *Student) ShowInfo() {
	fmt.Printf("name=%q, age=%v, score=%v\n",
		student.Name, student.Age, student.Score)
}

func (student *Student) SetScore(score int) {
	student.Score = score
}

type Pupil struct {
	Student // 匿名结构体
}

func (pupil *Pupil) Testing() {
	fmt.Println("Pupil testing")
}

type Graduate struct {
	Student // 匿名结构体
}

func (graduate *Graduate) Testing() {
	fmt.Println("Graduate testing")
}

func main() {
	pupil := &Pupil{}
	pupil.Student.Name = "xiaoming"
	pupil.Student.Age = 8
	pupil.Student.SetScore(99)
	pupil.Testing()
	pupil.Student.ShowInfo()

	pupil1 := &Pupil{}
	pupil1.Name = "xiaohua"
	pupil1.Age = 9
	pupil1.SetScore(100)
	pupil.Testing()
	pupil.Student.ShowInfo()

	graduate := &Graduate{}
	graduate.Student.Name = "yuz"
	graduate.Student.Age = 28
	graduate.Student.SetScore(100)
	graduate.Testing()
	graduate.Student.ShowInfo()
}
```

继承的好处
1. 代码的复用性提高了
2. 代码的扩展性和维护性提高了


继承的细节
1. 结构体可以使用嵌套匿名结构体使用的字段和方法，即首字母大写或小写的字段、方法都可以使用
2. 匿名结构体字段访问可以简化
3. 当结构体和匿名结构体有相同的字段或者方法时，编译器采用就近访问原则，先在本结构体中找字段、方法，没有再去匿名结构体中找。如希望访问匿名结构体的字段和方法，可以通过匿名结构体名来区分
```go
package main

import (
	"fmt"
)

type A struct {
	Name string
	age  int
}

func (a *A) SayOk() {
	fmt.Println("A SayOk", a.Name)
}

func (a *A) hello() {
	fmt.Println("A hello", a.Name)
}

type B struct {
	A
	Name string // B 结构体中也有该字段，如果创建B的实例，想要赋值A结构体中该Name字段的值，需要显示指定 A.Name，否则就近原则默认访问B.Name
}

func (b *B) SayOk() {
	fmt.Println("B sayok", b.Name)
}

func main() {
	// var b B
	// b.A.Name = "tome"
	// b.A.age = 10
	// b.A.SayOk()
	// b.A.hello()

	// // 下面简化写
	// b.Name = "smith"
	// b.age = 20
	// b.SayOk()
	// b.hello()

	var b B
	b.Name = "jack" // 使用B的字段
	b.age = 100     // 使用A的字段
	b.SayOk()       // 使用B的方法
	b.hello()       // 使用A的方法，但是没有给 b.A.Name 赋值，所有a.Name为默认值空字符串。可以显示赋值解决
	b.A.Name = "lucy"
	b.hello()
}
```
4. 结构体嵌套两个或多个匿名结构体时，如果匿名结构体有相同的字段、方法（同时结构体本身没有同名的字段、方法），在访问时，必须明确指定匿名结构体名字，否则编译报错
```go
package main

import (
	"fmt"
)

type A struct {
	Name string
	age  int
}

type B struct {
	Name  string
	Score float64
}

type C struct {
	// A和B同级，不分先后
	A
	B
}

func main() {
	fmt.Println("hello")
	var c C
	c.Name = "tom" // error，编译器不能确定到底是A还是B结构体里的字段
}
```
5. 如果一个结构体嵌套了一个有名结构体，这种模式叫做组合，如果是组合关系，那么在访问组合结构体的字段、方法时，必须带上结构体的名字
```go
package main

import (
	"fmt"
)

type A struct {
	Name string
	age  int
}

type B struct { // B 和 A 组合
	a A // 有名结构体
}

func main() {
	fmt.Println("hello")
	b := B{}
	// b.Name = "jack" // error，在有名结构体中，必须指定结构体名 b.A.Name，因为它不会往里再找，认为只有a
	b.a.Name = "jack"
	b.a.age = 18
	fmt.Println(b)
}
```
6. 嵌套匿名结构体后，可以在创建结构体变量时，直接指定各个匿名结构体字段的值。嵌套多个结构体也叫多重继承。
```go
package main

import (
	"fmt"
)

type Goods struct {
	Name  string
	Price float64
}

type Brand struct {
	Name    string
	Address string
}

type TV struct {
	Goods
	*Brand
}

func main() {
	fmt.Println()
	tv := TV{
		Goods{"电视机SN30239", 5000.9},
		&Brand{"海尔", "shandong"},
	}
	fmt.Println(tv.Goods)
	fmt.Println(*tv.Brand)
}
```

7. 匿名字段可以用基本数据类型，如int
```go
package main

import (
	"fmt"
)

type Monster struct {
	Name string
	Age  int
}

type E struct {
	Monster
	int // 匿名字段，基本数据类型；只能有一个同名的匿名字段
	n   int
}

func main() {
	fmt.Println()
	e := E{}
	e.Monster.Name = "hulijing"
	e.Monster.Age = 500
	e.int = 10
	e.n = 10
	fmt.Println(e)
}
```

# 接口
golang 中多态主要是通过接口来实现的。接口类型可以定义一组方法（但不能包含任何变量），但是这些方法不需要也不能在接口定义时实现。某个自定义类型（比如结构体）要使用这个接口的时候，需要把方法实现出来。

```go
type 接口名 interface {
	method1(参数列表) 返回值列表
	method2(参数列表) 返回值列表
}

// 接口实现，必须包含接口中所有方法的实现
func (实例名 自定义类型) method1(参数列表) 返回值列表 {
	// 方法具体实现
}

func (实例名 自定义类型) method2(参数列表) 返回值列表 {
	// 方法具体实现
}
```
- 接口里的所有方法都没有方法体，即没有具体实现细节。接口体现了程序设计的多态和高内聚低耦合的思想
- golang 中的接口，**不需要显式实现**。只要一个变量或实例含有接口类型的所有方法，那么这个变量或实例就是实现了这个接口。因此，golang 中没有 implement 这样的关键字

接口编程的基本例子
```go
package main

import (
	"fmt"
)

type USB interface {
	// 声明或定义了两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
	Name string
}

// 让 Phone 实现 USB 接口的（所有）方法
func (phone Phone) Start() {
	fmt.Printf("%v started\n", phone.Name)
}

func (phone Phone) Stop() {
	fmt.Printf("%v stoped\n", phone.Name)
}

type Camera struct {
	Weight float64
}

// 让 Camera 实现 USB 接口的（所有）方法
func (camera Camera) Start() {
	fmt.Printf("Camera weighing %v is started\n", camera.Weight)
}

func (camera Camera) Stop() {
	fmt.Printf("Camera weighing %v is stoped\n", camera.Weight)
}

type Computer struct {
}

// 编写一个方法 working 方法，接收一个 USB 接口类型的变量
// 只要是实现了 USB 接口（指实现了USB接口声明的所有方法）
func (computer Computer) Working(usb USB) {
	usb.Start()
	usb.Stop()
}

func main() {
	fmt.Println()
	// 创建接口体变量
	computer := Computer{}
	phone := Phone{"iPhone"}
	camera := Camera{2.3}

	computer.Working(phone)

	computer.Working(camera)
}
```

接口的细节：
- 接口本身不能创建实例（不像结构体），但是可以指向一个实现了该接口的自定义类型的实例或变量
```go
type Ainterface interface {
	Say()
}

type Stu struct {
	Name string
}

func (stu Stu) Say() {
	fmt.Println(stu.Name, "say something")
}

func main() {
	var stu Stu{"xiaoming"} // 该结构体实现了接口 Ainterface 的 Say() 方法
	var a Ainterface = stu
	a.Say()
}

```
- 接口中所有的方法都没有方法体，即没有具体实现方法
- 在 golang 中，一个自定义类型需要将某个接口的所有方法都实现，我们才说这个自定义类型实现了接口
- 一个自定义类型只有实现了某个接口，才能将该自定义类型的实例或变量赋给接口类型
- 只要是自定义数据类型，就可以实现接口，不仅仅是结构体类型
```go
type Ainterface interface {
	Say()
}

type integer int

func (i integer) Say() {
	fmt.Println(i, "say something")
}

func main() {
	var i integer = 10 // 该结构体实现了接口 Ainterface 的 Say() 方法
	var b Ainterface = i
	b.Say()
}
```
- 一个自定义类型可以实现多个接口
```go
type A interface {
	Say()
}

type B interface {
	Hello()
}

type Stu struct {
	Name string
}

// Stu 既实现了A接口也实现了B接口
func (stu Stu) Say() {
	fmt.Println(stu.Name, "say something")
}

func (stu Stu) Hello() {
	fmt.Println(stu.Name, "hello world")
}

func main() {
	var monster Monster
	var a A = monster
	var b B = monster
	a.Say() // right
	a.Hello() // error
	b.Hello() // right
	b.Say() // error
}
```
- golang 接口中不能有任何变量
```go
type A interface {
	Say() // right
	// Name string // error
}
```
- 一个接口A可以继承多个别的接口B、C，这时如果要实现接口A，也必须实现接口B、C的所有方法。但是B和C中不能定义相同的方法
```go
type B interface {
	test01()
}

type C interface {
	// test01() // error，编译不通过
	test02()
}

type A interface { // 接口继承
	B
	C
	test03()
}

type Stu struct {

}

func (stu Stu) test01() {

}

func (stu Stu) test02() {

}

func (stu Stu) test03() {

}

func main() {
	var stu Stu
	var a A = stu
	a.test01()
	a.test02()
	a.test03()
}
```
- 接口类型默认是一个指针(引用类型)，如果没有对接口初始化就使用，那么会输出 nil
- 空接口 `interface{}` 没有任何方法，所以所有类型都实现了空接口，即可以把任何一种数据类型变量赋给空接口
```go
type T interface {}

type Stu struct {
	Name string
}

func main() {
	var stu Stu
	var t T = stu // Stu 默认是实现了空接口
	var t2 interface{} = stu // right
	var num float64 = 8.8
	t2 = num // right
	t1 = num // right
}
```

- 实现接口时，自定义类型实现时如果是引用类型，那么调用时也需要用指针
```go
type Usb interface {
	Say()
}

type Stu struct {

}

func (stu *Stu) Say() {

}

func main() {
	var stu Stu = Stu{}
	var u Usb = stu // error，会报 Stu 类型没有实现 Usb 接口
	var u Usb = &stu // right
}
```

对接口体切片排序
```go
package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct {
	Name string
	Age  int
}

// hero 结构体切片类型
type HeroSlice []Hero

// 实现 Interface 接口的三个方法
func (heroslice HeroSlice) Len() int {
	return len(heroslice)
}

// 按 hero 的年龄从小到大排序
func (heroslice HeroSlice) Less(i, j int) bool {
	return heroslice[i].Age < heroslice[j].Age
	// return heroslice[i].Name < heroslice[j].Name
}

func (heroslice HeroSlice) Swap(i, j int) {
	heroslice[i], heroslice[j] = heroslice[j], heroslice[i]
}

func main() {
	var intSlice = []int{0, -1, 10, 7, 90}
	// 除了手撕冒泡排序等算法外，还可以用系统函数
	sort.Ints(intSlice) // 切片是引用类型
	fmt.Println(intSlice)

	// 结构体切片排序
	var heroSlices HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("hero-%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroSlices = append(heroSlices, hero)
	}

	// 排序前顺序
	for _, v := range heroSlices {
		fmt.Println(v)
	}
	// 排序
	sort.Sort(heroSlices)
	// 排序后顺序
	fmt.Println("排序后")
	for _, v := range heroSlices {
		fmt.Println(v)
	}
}
```

## 多态
 接口体现多态的两种形式
 ### 多态参数
 例子参见接口章节开始时的电话和相机的例子。在 USB 接口例子中，`func (computer Computer) Working(usb USB)` 既可以接收 Phone 结构体，又可以接收 Camera 结构体。

 ### 多态数组
 在接口数组中，存放不同类型的接口体，如 Phone 结构体，Camera 结构体
 ```go
package main

import (
	"fmt"
)

type USB interface {
	// 声明或定义了两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
	Name string
}

// 让 Phone 实现 USB 接口的（所有）方法
func (phone Phone) Start() {
	fmt.Printf("%v started\n", phone.Name)
}

func (phone Phone) Stop() {
	fmt.Printf("%v stoped\n", phone.Name)
}

type Camera struct {
	Weight float64
}

// 让 Camera 实现 USB 接口的（所有）方法
func (camera Camera) Start() {
	fmt.Printf("Camera weighing %v is started\n", camera.Weight)
}

func (camera Camera) Stop() {
	fmt.Printf("Camera weighing %v is stoped\n", camera.Weight)
}

func main() {
	fmt.Println()
	// 创建接口数组
	var usb [3]USB
	usb[0] = Phone{"苹果"}
	usb[1] = Phone{"小米"}
	usb[2] = Camera{3.14}
	fmt.Println(usb)
}
 ```

 ## 类型断言
由于接口时一般类型，不知道具体类型，如果要转成具体类型，就需要使用类型断言。

```go
package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func main() {
	var a interface{}
	var point Point = Point{1, 2}
	a = point
	var b Point
	b = a.(Point) // 类型断言
	// b = a // error
	fmt.Println(b)
	fmt.Println(a)

	// 其他案例
	var x interface{}
	var t float32 = 1.1
	x = t
	y := x.(float32)
	// y := x.(float64) // error
	fmt.Printf("%T, %v\n", y, y)
}
```
进行类型断言时，如果类型不匹配，就会报错 panic，因此进行类型断言时，要确保原来的空接口执行的就是断言的类型。

可以带上检测机制：
```go
// 断言检测
var m interface{}
var n float64 = 3.14
m = n
// k, ok := m.(float64) // 带检测 right
// k, ok := m.(float32) // 带检测 right
// if ok {
// 	fmt.Printf("%T, %v\n", k, k)
// } else {
// 	fmt.Println("convert failed")
// }

if k, ok := m.(float32); ok {
	fmt.Printf("%T, %v\n", k, k)
} else {
	fmt.Println("convert failed")
}
```

类型断言的应用案例
```go
package main

import (
	"fmt"
)

type USB interface {
	// 声明或定义了两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
	Name string
}

// 让 Phone 实现 USB 接口的（所有）方法
func (phone Phone) Start() {
	fmt.Printf("%v started\n", phone.Name)
}

func (phone Phone) Stop() {
	fmt.Printf("%v stoped\n", phone.Name)
}

func (phone Phone) Call() {
	fmt.Printf("%v calling \n", phone.Name)
}

type Camera struct {
	Weight float64
}

// 让 Camera 实现 USB 接口的（所有）方法
func (camera Camera) Start() {
	fmt.Printf("Camera weighing %v is started\n", camera.Weight)
}

func (camera Camera) Stop() {
	fmt.Printf("Camera weighing %v is stoped\n", camera.Weight)
}

type Computer struct {
}

// 编写一个方法 working 方法，接收一个 USB 接口类型的变量
// 只要是实现了 USB 接口（指实现了USB接口声明的所有方法）
func (computer Computer) Working(usb USB) {
	usb.Start()
	// 如果USB指向Phone结构体变量，则需要调用Call方法
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}

func main() {
	fmt.Println()
	// 创建接口体数组
	var usbArr [3]USB
	usbArr[0] = Phone{"小米"}
	usbArr[1] = Phone{"苹果"}
	usbArr[2] = Camera{3.14}

	computer := Computer{}
	for _, device := range usbArr {
		computer.Working(device)
	}
}
```
类型断言的应用2
```go
package main

import (
	"fmt"
)

type Student struct{}
type USB interface{}

func TypeJudge(items ...interface{}) {

	for index, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("第%v个参数是bool类型，值是%v\n", index, x)
		case float32:
			fmt.Printf("第%v个参数是float32类型，值是%v\n", index, x)
		case float64:
			fmt.Printf("第%v个参数是float64类型，值是%v\n", index, x)
		case int, int32, int64:
			fmt.Printf("第%v个参数是整数类型，值是%v\n", index, x)
		case string:
			fmt.Printf("第%v个参数是string类型，值是%v\n", index, x)
		default:
			fmt.Printf("第%v个参数 类型不确定，值是%v\n", index, x)
		}
	}
}

func main() {
	fmt.Println()
	var n1 bool = true
	var n2 float32 = 2.2
	var n3 float64 = 3.3
	var n4 int = 30
	var n5 string = "xjz"
	n6 := "北京"
	n7 := Student{}
	n77 := &Student{}
	var n8 USB
	TypeJudge(n1, n2, n3, n4, n5, n6, n7, n77, n8)
}
```

# 项目练习-家庭收支记账软件
对于一个200-500万的项目，流程步骤如下：
1. 需求分析
2. 设计阶段
3. 实现阶段
4. 测试阶段
5. 实施阶段
6. 维护阶段

实现阶段和测试阶段会多轮迭代。

|条目|需求分析|设计阶段|实现阶段|测试阶段|实施阶段|维护阶段|
|---|---|---|---|---|---|---|
|时间|30%时间|20%时间|20%时间|共用30%|共用30%|共用30%|
|负责人|需求分析师|项目经理或架构师|软件工程师（码农）|软件测试工程师（如用友）|实施工程师|一般没有专业人员，对接人|
|职位要求|懂技术、懂业务|架构（开发语言、架构、数据库、操作系统）、选人|实现各个模块（golang, python, c++, java etc)|黑盒测试、白盒测试（写代码）、灰盒测试|上线服务器、配置||
|结果|需求分析报告|设计文档（类图、**数据库设计、界面原型**）|代码|测试大纲、测试用例、测试文档|||

## 需求说明
1. 模拟基于文本界面的《家庭记账软件》
2. 能够记录家庭的收入、支出，并能够打印收支明细表

## 项目界面

```html
-----------------家庭收支记账软件-----------------
                  1 收支明细
                  2 登记收入
                  3 登记支出
                  4 退出软件
请选择(1-4): 1
-----------------当前收支明细记录-----------------
当前没有收支明细...来一笔吧！

-----------------家庭收支记账软件-----------------
                  1 收支明细
                  2 登记收入
                  3 登记支出
                  4 退出软件
请选择(1-4): 2
本次收入金额: 100
本次收入说明: gongzi 

-----------------家庭收支记账软件-----------------
                  1 收支明细
                  2 登记收入
                  3 登记支出
                  4 退出软件
请选择(1-4): 1
-----------------当前收支明细记录-----------------
收支    账户余额        收支余额        说明
收入    10100           100             gongzi

-----------------家庭收支记账软件-----------------
                  1 收支明细
                  2 登记收入
                  3 登记支出
                  4 退出软件
请选择(1-4): 3
本次支出金额: 10
本次支出说明: chifan

-----------------家庭收支记账软件-----------------
                  1 收支明细
                  2 登记收入
                  3 登记支出
                  4 退出软件
请选择(1-4): 1
-----------------当前收支明细记录-----------------
收支    账户余额        收支余额        说明
收入    10100           100             gongzi
支出    10090           10              chifan

-----------------家庭收支记账软件-----------------
                  1 收支明细
                  2 登记收入
                  3 登记支出
                  4 退出软件
请选择(1-4): 4
你确定要退出吗？y/n
y
你退出了家庭记账软件
```
代码实现：
代码结构如下
```bash
.
├── go.mod
├── main.go
└── utils
    └── familyAccount.go

1 directory, 3 files
```

`go.mod` 如下
```go
module xgo17/oop

go 1.23.11
```

`familyAccount.go` 如下
```go
package utils

import (
	"fmt"
)

type FamilyAccount struct {
	// 声明一个字段，保存用户输入的选项
	key string

	// 声明一个字段，用于控制退出主循环
	loop bool

	// 定义账户余额
	balance float64
	// 每次收支的余额
	money float64
	// 每次收支的说明
	note string
	// 定义变量，记录是否有收支行为
	flag bool
	// 收支的详情
	details string
}

// 工厂模式
func FactoryFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收支\t账户余额\t收支余额\t说明",
	}
}

// 给该结构体绑定方法
// 显示明细
func (FamilyAccount *FamilyAccount) ShowDetails() {
	fmt.Println("-----------------当前收支明细记录-----------------")
	if FamilyAccount.flag {
		fmt.Println(FamilyAccount.details)
	} else {
		fmt.Println("当前没有收支明细...来一笔吧！")
	}
}

// 收入
func (FamilyAccount *FamilyAccount) Income() {
	fmt.Print("本次收入金额: ")
	fmt.Scanln(&FamilyAccount.money)
	fmt.Print("本次收入说明: ")
	fmt.Scanln(&FamilyAccount.note)
	FamilyAccount.balance += FamilyAccount.money // 修改账户余额
	FamilyAccount.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", FamilyAccount.balance, FamilyAccount.money, FamilyAccount.note)
	FamilyAccount.flag = true
}

// 支出
func (FamilyAccount *FamilyAccount) Pay() {
	fmt.Print("本次支出金额: ")
	fmt.Scanln(&FamilyAccount.money)
	fmt.Print("本次支出说明: ")
	fmt.Scanln(&FamilyAccount.note)
	if FamilyAccount.money > FamilyAccount.balance {
		fmt.Println("余额不足")
	}
	FamilyAccount.balance -= FamilyAccount.money // 修改账户余额
	FamilyAccount.details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", FamilyAccount.balance, FamilyAccount.money, FamilyAccount.note)
	FamilyAccount.flag = true
}

// 退出系统
func (FamilyAccount *FamilyAccount) Exit() {
	fmt.Println("你确定要退出吗？y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入 y/n")
	}
	if choice == "y" {
		FamilyAccount.loop = false
	}
}

// 显示主菜单
func (FamilyAccount *FamilyAccount) MainMenu() {
	// 显示主菜单
	for {
		fmt.Println()
		fmt.Println("-----------------家庭收支记账软件-----------------")
		fmt.Println("                  1 收支明细")
		fmt.Println("                  2 登记收入")
		fmt.Println("                  3 登记支出")
		fmt.Println("                  4 退出软件")
		fmt.Print("请选择(1-4): ")

		// fmt.Scanln(&key)
		fmt.Scanf("%s", &FamilyAccount.key)

		switch FamilyAccount.key {
		case "1":
			FamilyAccount.ShowDetails()
		case "2":
			FamilyAccount.Income()
		case "3":
			FamilyAccount.Pay()
		case "4":
			FamilyAccount.Exit()
		default:
			fmt.Println("请输入正确的选项...")
		}
		if !FamilyAccount.loop {
			break
		}
	}
	fmt.Println("你退出了家庭记账软件")
}
```

`main.go` 如下
```go
package main

import (
	"fmt"
	"xgo17/oop/utils"
)

func main() {
	fmt.Println("OOP")
	utils.FactoryFamilyAccount().MainMenu()
}
```

# 项目练习-客户信息管理系统
## 需求说明
1. 模拟实现基于文本界面的《客户信息管理系统》
2. 该软件能够实现对客户对象的插入、修改和删除（用切片实现），并能够打印客户明细表
3. 项目采用分级菜单方式
## 项目界面
```html

---------------客户信息管理软件---------------
               1 添加客户
               2 修改客户
               3 删除客户
               4 客户列表
               5 退    出
请选择(1-5): 4
-----------------------客户列表-----------------------
编号    姓名    性别    年龄    电话    邮箱
1       张三    男      20      112     zs@gmail.com
-----------------------客户列表完成-------------------

---------------客户信息管理软件---------------
               1 添加客户
               2 修改客户
               3 删除客户
               4 客户列表
               5 退    出
请选择(1-5): 1
-----------------------添加客户-----------------------
姓名：
张三
性别：
男
年龄：
18
电话：
131
邮箱：
zs@qq.com
-----------------------添加完成-----------------------
编号    姓名    性别    年龄    电话    邮箱
---------------客户信息管理软件---------------
               1 添加客户
               2 修改客户
               3 删除客户
               4 客户列表
               5 退    出
请选择(1-5): 4
-----------------------客户列表-----------------------
编号    姓名    性别    年龄    电话    邮箱
1       张三    男      20      112     zs@gmail.com
2       张三    男      18      131     zs@qq.com
-----------------------客户列表完成-------------------

---------------客户信息管理软件---------------
               1 添加客户
               2 修改客户
               3 删除客户
               4 客户列表
               5 退    出
请选择(1-5): 3
---------------------删除客户---------------------
请选择待删除客户编号(-1退出): 
1
确认是否删除(y/n): 
y
---------------------删除成功---------------------
---------------客户信息管理软件---------------
               1 添加客户
               2 修改客户
               3 删除客户
               4 客户列表
               5 退    出
请选择(1-5): 5
确认是否退出(y/n): 
y
你退出了系统
```

## 编写代码
代码组织结构
```bash
xgo18CustomManager
├── go.mod
├── model
│   └── customer.go
├── service
│   └── customerService.go
└── view
    └── customerView.go
```

`go.mod` 内容
```go
module xgo18

go 1.23.11

```

`customerView.go` 内容
```go
package main

import (
	"fmt"
	"xgo18/model"
	"xgo18/service"
)

type customerView struct {
	key             string // 用户输入
	loop            bool   // 表示是否循环显示菜单
	customerService *service.CustomerService
}

// 显示所有客户信息
func (cv *customerView) list() {
	customers := cv.customerService.List()
	// 显示
	fmt.Println("-----------------------客户列表-----------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")

	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("-----------------------客户列表完成-------------------")
	fmt.Println()
}

func (cv *customerView) add() {
	fmt.Println("-----------------------添加客户-----------------------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	email := ""
	fmt.Scanln(&email)

	// id 号没有让用户输入，是唯一的，需要系统分配
	customer := model.FactoryCustomer2(name, gender, age, phone, email)
	if cv.customerService.Add(customer) {
		fmt.Println("-----------------------添加完成-----------------------")
	} else {
		fmt.Println("-----------------------添加失败-----------------------")
	}

	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")

}

// 得到用户的输入，删除ID对应的客户
func (cv *customerView) delete() {
	fmt.Println("---------------------删除客户---------------------")
	fmt.Println("请选择待删除客户编号(-1退出): ")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return // 放弃删除
	}
	fmt.Println("确认是否删除(y/n): ")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" {
		// 调用 customerService 的 Delete 方法
		if cv.customerService.Delete(id) {
			fmt.Println("---------------------删除成功---------------------")
		} else {
			fmt.Println("-----------------删除失败，ID不存在----------------")
		}
	}

}

// 退出软件
func (cv *customerView) exit() {
	fmt.Println("确认是否退出(y/n): ")
	for {
		fmt.Scanln(&cv.key)
		if cv.key == "y" || cv.key == "n" {
			break
		}
		fmt.Println("你输入有误，确认是否退出(y/n): ")
	}
	if cv.key == "y" {
		cv.loop = false
	}

}

func (cv *customerView) mainMenu() {
	for {
		fmt.Println("---------------客户信息管理软件---------------")
		fmt.Println("               1 添加客户")
		fmt.Println("               2 修改客户")
		fmt.Println("               3 删除客户")
		fmt.Println("               4 客户列表")
		fmt.Println("               5 退    出")
		fmt.Print("请选择(1-5): ")

		fmt.Scanln(&cv.key)
		switch cv.key {
		case "1":
			// fmt.Println("添加客户")
			cv.add()
		case "2":
			fmt.Println("修改客户")
		case "3":
			// fmt.Println("删除客户")
			cv.delete()
		case "4":
			// fmt.Println("客户列表")
			cv.list()
		case "5":
			cv.exit()
		default:
			fmt.Println("输入有误，请重新输入")
		}

		if !cv.loop {
			break
		}
	}
	fmt.Println("你退出了系统")
}

func main() {
	fmt.Println("ok")
	cv := customerView{
		key:             "",
		loop:            true,
		customerService: service.FactoryCustomerService(),
	}
	cv.mainMenu()
}

```

`customerService.go` 内容
```go
package service

import (
	"xgo18/model"
)

// 对 Customer 的操作，包括增删改查
type CustomerService struct {
	customers []model.Customer
	// 当前切片含有多少个客户
	customerNum int
}

func FactoryCustomerService() *CustomerService {
	customerService := &CustomerService{}
	customerService.customerNum = 1
	customer := model.FactoryCustomer(1, "张三", "男", 20, "112", "zs@gmail.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

// 返回客户切片
func (cs *CustomerService) List() []model.Customer {
	return cs.customers
}

// 添加客户到 customers 切片
func (cs *CustomerService) Add(customer model.Customer) bool {
	cs.customerNum += 1
	customer.Id = cs.customerNum
	cs.customers = append(cs.customers, customer)
	return true
}

// 根据ID查找客户在切片中的对应下标，如果没有该客户，返回-1
func (cs *CustomerService) FindById(id int) int {
	index := -1
	for i := 0; i < len(cs.customers); i++ {
		if cs.customers[i].Id == id {
			index = i
		}
	}
	return index
}

// 根据ID删除客户信息
func (cs *CustomerService) Delete(id int) bool {
	index := cs.FindById(id)
	if index == -1 { // 说明没有这个客户
		return false
	}
	// 如何从切片中删除元素
	cs.customers = append(cs.customers[:index], cs.customers[index+1:]...)
	return true
}

```

`customer.go` 内容
```go
package model

import (
	"fmt"
)

// 客户信息
type Customer struct {
	Id     int
	Name   string
	Gender string
	Age    int
	Phone  string
	Email  string
}

// 工厂模式，返回 Customer 实例
func FactoryCustomer(id int, name string, gender string,
	age int, phone string, email string) Customer {
	return Customer{
		Id:     id,
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

func FactoryCustomer2(name string, gender string,
	age int, phone string, email string) Customer {
	return Customer{
		Name:   name,
		Gender: gender,
		Age:    age,
		Phone:  phone,
		Email:  email,
	}
}

// 返回客户信息
func (customer Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v",
		customer.Id, customer.Name, customer.Gender,
		customer.Age, customer.Phone, customer.Email)
	return info
}

```

# 文件操作
文件是数据源（保存数据的地方）的一种，最主要的作用是保存数据，常见的如 word, txt, mp4, aac 等，可以保存文字、视频和声音等。

文件在程序中是以流的形式来操作的。以 golang 程序为中心，流是数据在数据源（文件）和 golang 程序（内存）之间经历的路径，输入流是数据从数据源到程序的路径，输出流是数据从程序到数据源的路径。

## 打开文件关闭文件
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("ok")
	f, err := os.Open("./data/text.txt")
	if err != nil {
		fmt.Println("打开失败")
	} else {
		fmt.Println("打开成功")
		fmt.Println(f)
	}

	err = f.Close()

	if err != nil {
		fmt.Println("关闭失败")
	} else {
		fmt.Println("关闭成功")
	}
}

```

## 读取文件

### 带缓冲的读取
```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("../data/text.txt")
	if err != nil {
		fmt.Println("文件打开失败")
	}
	// 函数退出时，及时关闭文件句柄
	defer f.Close()

	// 创建 Reader，默认缓冲 4096
	reader := bufio.NewReader(f)
	// 循环读取文件内容
	for {
		str, err := reader.ReadString('\n') // 读到字符'\n'就结束
		if err == io.EOF {                  // io.EOF代表文件末尾
			break
		} else if err != nil {
			fmt.Println("读取错误", err)
		} else {
			// 输出内容
			fmt.Print(str) // str 带 '\n'
		}

	}
	fmt.Println("文件读取结束...")
}
```

### 读取整个文件
使用 `ioutil.ReadFile`（go 1.16前）、`os.ReadFile`（go 1.16后） 一次将整个文件读入到内存中，这种方式适用于文件不大的情况。
```go
func ReadFile(filename string) ([]byte, error)
```
ReadFile 从 filename 指定的文件中读取数据并返回文件的内容。成功的调用返回的 err 为 nil 而非 EOF. 因为函数定义为读取整个文件，它不会将读取返回的 EOF 视为应报告的错误。

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println()
	// ioutil.ReadFile一次性将文件读取到位
	filepath := "../data/text.txt"
	content, err := os.ReadFile(filepath) // 因为没有显示的 Open 该文件，因此也不需要显示的 Close 该文件
	if err != nil {
		fmt.Println(err)
	}
	for _, c := range content {
		fmt.Println(c)
		fmt.Printf("%c\n", c)
	}
	fmt.Printf("%v", string(content))
}

```

## 写入文件
### 带缓冲写文件
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println()
	// 创建一个新文件 "../data/textWriterBuffer.txt"，并写入5句 hello, gordon
	filepath := "../data/textWriterBuffer.txt"
	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}

	// 及时关闭文件，防止内存泄露
	defer f.Close()

	// 写入5句 hello, gordon
	str := "hello, Gordon"
	// 使用带缓冲的 writer
	writer := bufio.NewWriter(f)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
		writer.WriteString("\r\n")
	}
	// 因为wirter是带缓冲，因此在调用 writerstring 方法时，其实内容是写入到缓冲的，不是写入到磁盘
	// 还需要调用 flush 方法写入到磁盘，否则文件中没有数据
	err = writer.Flush()
	if err != nil {
		fmt.Println("flush失败")
	}
}

```

### 覆盖写文件
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 打开一个已经存在的文件，将原来的内容覆盖成新内容：你好，月球
	filepath := "../data/textWriterBuffer.txt"
	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("open file err=%v", err)
	}
	// 及时关闭文件句柄
	defer f.Close()

	// 准备写入的内容
	str := "你好，月球"
	// 使用带缓冲的写
	writer := bufio.NewWriter(f)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
		writer.WriteString("\r\n")
	}
	// flush缓冲
	err = writer.Flush()
	if err != nil {
		fmt.Println(err)
	}
}

```

### 追加写文件
```go
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println()
	// 打开一个已经存在的文件，在原来的内容追加内容：hell0, moon
	filepath := "../data/textWriterBuffer.txt"

	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开文件失败", err)
	}

	defer f.Close()

	// 带缓冲写文件，先写入到缓冲区
	str := "hello, moon"
	writer := bufio.NewWriter(f)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
		writer.WriteString("\r\n")
	}

	// flush缓冲到磁盘
	err = writer.Flush()
	if err != nil {
		fmt.Println("flush err=", err)
	}
}

```

## 同时读写
```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开一个存在的文件，读里面的内容并输出到终端，同时可以写入内容：are you ok?
	filepath := "../data/textWriterBuffer.txt"

	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("err", err)
	}

	defer f.Close()

	// 先读取原来的内容，并显示在终端
	reader := bufio.NewReader(f)
	for {
		strRead, errRead := reader.ReadString('\n') // 一行的结束标志为 \n
		if errRead == io.EOF {                      // 读到文件的末尾
			break
		} else if errRead != nil && errRead != io.EOF {
			fmt.Println("err=", errRead)
		} else {
			fmt.Print(strRead)
		}

	}

	// 再写入
	strWrite := "are you ok?"
	writer := bufio.NewWriter(f)
	for i := 0; i < 5; i++ {
		writer.WriteString(strWrite)
		writer.WriteString("\r\n")
	}

	// flush 缓冲
	errFlush := writer.Flush()
	if errFlush != nil {
		fmt.Println("flush err=", errFlush)
	}

}

```

## 读一个文件写入到另一个文件
```go
package main

import (
	"fmt"
	"os"
)

func main() {
	// 将 ../data/textWriterBuffer.txt 读取并将内容写入到 ../data/textReadWrite.txt
	filepathRead := "../data/textWriterBuffer.txt"
	filepathWrite := "../data/textReadWrite.txt"
	// 读取文件，隐式打开，不需要显示关闭
	content, err := os.ReadFile(filepathRead)
	if err != nil {
		fmt.Println("err=", err)
	}
	// 写入文件
	errWrite := os.WriteFile(filepathWrite, content, 0666)
	if errWrite != nil {
		fmt.Println("write err=", err)
	}
}

```

## 判断文件存在性
```go
package main

import (
	"fmt"
	"os"
)

func IsFileExist(filepath string) (bool, error) {
	_, err := os.Stat(filepath)
	if err == nil {
		// 文件存在
		return true, nil
	} else if os.IsNotExist(err) {
		// 文件不存在
		return false, nil
	} else {
		// 不能判断文件是否存在
		return false, err
	}
}

func main() {
	fmt.Println()
	// 判断文件是否存在
	filepath := "../data/text.txt"
	b, _ := IsFileExist(filepath)
	if b {
		fmt.Printf("%q 存在\n", filepath)
	} else {
		fmt.Printf("%q 不存在\n", filepath)
	}

}

```

## 拷贝文件
```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	// 打开源文件
	f, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("open src file err=", err)
	}
	defer f.Close()
	// 获取 reader
	reader := bufio.NewReader(f)

	// 打开目标文件
	f_writer, err_writer := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err_writer != nil {
		fmt.Println("open dst file err=", err_writer)
	}
	defer f_writer.Close()
	// 获取writer
	writer := bufio.NewWriter(f_writer)

	return io.Copy(writer, reader)

}

func main() {
	// 将 ../data/flower.webp 拷贝到 ../data/flower-copy.webp
	srcFileName := "../data/flower.webp"
	dstFileName := "../data/flower-copy.webp"
	w, e := CopyFile(dstFileName, srcFileName)
	if e != nil {
		fmt.Println("err=", e)
	} else {
		fmt.Println("拷贝完成")
		fmt.Println(w)
	}

}

```

## 文件字符统计
```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	EngCount   int // 英文的个数
	NumCount   int // 数字的个数
	SpaceCount int // 空格的个数
	OtherCount int // 其他的个数
}

func main() {
	fmt.Println()
	// 统计一个文件中有多少个英文、数字、空格和其他字符
	// 思路：每读取一行统计一次，然后将结果保存到一个结构体
	filename := "../data/text.txt"
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("err=", err)
	}
	defer f.Close()

	// CharCount实例
	var count CharCount
	// 创建一个reader
	reader := bufio.NewReader(f)
	// 循环读取内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		// str = []rune(str) // 可以处理中文字符
		// 遍历 str，进行统计
		for _, v := range str {
			// fmt.Println(v)
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough // 穿透
			case v >= 'A' && v <= 'Z':
				count.EngCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			case v == ' ':
				count.SpaceCount++
			default:
				count.OtherCount++
			}

		}
	}
	// 输出统计结果
	fmt.Printf("英文字符的个数%v，数字的个数%v，空格的个数%v，其他字符个数%v\n", count.EngCount, count.NumCount, count.SpaceCount, count.OtherCount)
}

```

# 命令行参数

## 按照参数顺序
`os.Args` 是一个 string 切片，用来存储所有的命令行参数，第一个参数是源文件

```go
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("命令行的参数有", len(os.Args))
	// 遍历os.Args切片，得到每一个命令行参数值
	for i, v := range os.Args {
		fmt.Printf("args[%v]=%q\n", i, v)
	}
}

```

运行命令
```bash
go run main.go tom /disk0/codes/golang-gordon/data/gordon.jpg 999

```
解析结果
```bash
命令行的参数有 4
args[0]="/tmp/go-build4143543334/b001/exe/main"
args[1]="tom"
args[2]="/disk0/codes/golang-gordon/data/gordon.jpg"
args[3]="999"
```

## 指定参数名
flag 包解析命令行参数，可以根据参数
```go
package main

import (
	"flag"
	"fmt"
)

func main() {
	// 定义几个变量，用于接收命令行参数
	var user string
	var passwd string
	var host string
	var port int

	flag.StringVar(&user, "u", "", "用户名，默认为空")
	flag.StringVar(&passwd, "p", "", "密码，默认为空")
	flag.StringVar(&host, "h", "localhost", "主机名，默认为localhost")
	flag.IntVar(&port, "port", 3306, "端口号，默认为3306")

	// 必须调用转换
	flag.Parse()

	// 输出结果
	fmt.Printf("user=%v\npassword=%v\nhost=%v\nport=%v\n", user, passwd, host, port)
}

```

运行
```bash
go build main.go
./main -u root -p 123456 -h 127.0.0.1 -port 3306
```
结果
```bash
user=root
password=123456
host=127.0.0.1
port=3306
```

# JSON 基本介绍
JSON(JavaScript Object Notation)是一种轻量级的数据交换格式，易于阅读和编写，也易于机器解析和生成。

JSON是在2001年开始推广使用的数据格式，目前已经成为主流的数据格式。

JSON易于机器解析和生成，并有效地提升网络传输效率，通常程序在网络传输时会先将数据（结构体、map、切片等）序列化成JSON字符串，当接收方接收到JSON字符串时，再反序列化恢复得到原来的数据类型（结构体、map、切片等）。

## JSON序列化
JSON序列化指将KEY-VALUE结构的数据类型序列化成JSON字符串的操作。

### 序列化结构体
```go
package main

import (
	"encoding/json"
	"fmt"
)

// 结构体序列化时，可以指定JSON标签，如下面 Name, Age，这里使用了 golang 的反射机制
// 如果字段名不是大写，那么json.marshal序列化会丢掉该字段
type Monster struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Birthday string
	Sal      float64
	Skill    string
}

func testStruct() {
	var monster = Monster{
		Name:     "牛魔王",
		Age:      5000,
		Birthday: "1-11-11",
		Sal:      8000.0,
		Skill:    "一板斧",
	}
	// 序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化失败")
		return
	}
	// 输出序列化后的及格过
	fmt.Printf("monster序列化后的结果%q\n", string(data))
}

func main() {
	fmt.Println()
	testStruct()
}

```

### 序列化map
```go
package main

import (
	"encoding/json"
	"fmt"
)

func testMap() {
	// 定义一个map
	var a map[string]interface{}
	// 使用map前，需要先make
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "芭蕉洞"
	// 序列化map
	data, err := json.Marshal(a) // data 是 []byte
	if err != nil {
		fmt.Println("序列化失败")
		return
	}
	fmt.Printf("a map 序列化后=%v\n", string(data))
}

func main() {
	fmt.Println()
	testMap()
}

```
### 序列化切片

```go
package main

import (
	"encoding/json"
	"fmt"
)

func testSlice() {
	// 声明切片
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	// 使用前先MAKE
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = 7
	m1["address"] = [...]string{"北京", "南京"}
	// 添加到切片
	slice = append(slice, m1)
	// 先MAKE
	m2 := make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = 13
	m2["address"] = "纽约"
	// 添加到切片
	slice = append(slice, m2)

	// 切片序列化
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("序列化错误")
		return
	}
	fmt.Println(string(data))

}

func main() {
	fmt.Println()
	testSlice()
}

```

### 序列化其他类型
```go
package main

import (
	"encoding/json"
	"fmt"
)

// 基本数据类型序列化
func testFloat64() {
	// 对基本数据类型序列化意义不大
	var num1 float64 = 2124.8
	// 对 num1 序列化
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Println("序列化失败")
		return
	}
	fmt.Printf("%q\n", string(data)) // 把浮点数转为字符串
}

func main() {
	fmt.Println()
	testFloat64()
}

```

## JSON反序列化
JSON反序列化指将JSON字符串反序列化成对应的数据类型。

### 反序列化结构体
需要保证字段一致

```go
package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Birthday string
	Sal      float64
	Skill    string
}

func unserial() {
	// 项目开发中通过其他方式获取得到
	str := "{\"name\":\"牛魔王\",\"age\":5000,\"Birthday\":\"1-11-11\",\"Sal\":8000,\"Skill\":\"一板斧\"}"

	// 定义一个 monster实例
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println("反序列化失败")
	}
	fmt.Println(monster)
	fmt.Println(monster.Name)
}

func main() {
	fmt.Println()
	unserial()
}

```
### 反序列化map
```go
package main

import (
	"encoding/json"
	"fmt"
)

func unserialMap() {
	str := "{\"address\":\"芭蕉洞\",\"age\":30,\"name\":\"红孩儿\"}"
	// 定义一个map
	var a map[string]interface{}

	// a = make(map[string]interface{})

	// 不需要再make，因为反序列化底层会自动make
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("反序列化失败")
		return
	}
	fmt.Println(a)
}

func main() {
	fmt.Println()
	unserialMap()
}

```

### 反序列化切片
```go
package main

import (
	"encoding/json"
	"fmt"
)

func unserialSlice() {
	str := "[{\"address\":[\"北京\",\"南京\"],\"age\":7,\"name\":\"jack\"},{\"address\":\"纽约\",\"age\":13,\"name\":\"tom\"}]"

	// 定义一个切片
	var slice []map[string]interface{}

	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Println("反序列化失败")
		return
	}
	fmt.Println(slice)
}

func main() {
	fmt.Println()
	unserialSlice()
}

```

# 单元测试
在项目开发中，会遇到确认一个函数或者模块是否正确（运行结果正确）。

传统测试示例
```go
package main

import (
	"fmt"
)

// 被测函数
func addUpper(n int) int {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	return res
}

func main() {
	fmt.Println()
	// 传统测试方法
	res := addUpper(10)
	if res != 55 {
		fmt.Printf("addUpper 错误，返回值=%v 期望值=%v\n", res, 55)
	} else {
		fmt.Printf("addUpper 正确，返回值=%v 期望值=%v\n", res, 55)
	}
}

```

golang 语义自带一个轻量级的测试框架 testing 和自带的 go test 命令来实现单元测试和性能测试，testing 框架和其他语言中的测试框架类似，可以基于这个框架写针对相应函数的测试用例，也可以基于该框架写相应的压力测试用例。通过单元测试，可以解决如下问题：
1. 确保每个函数是可运行的，并且结果是正确的
2. 确保写出来的代码性能是好的
3. 单元测试能及时发现程序设计或实现中逻辑错误，使问题及早暴露，便于问题的定位解决，而性能测试的重点在于发现程序设计上的一些问题，让程序能够在高并发的情况下还能保存稳定

单元测试示例

代码结构
```bash
.
├── go.mod
├── main.go
└── test
    ├── cal.go
    ├── cal_test.go
    └── sub_test.go

1 directory, 5 files
```

`cal.go`
```go
package test

// 待测试函数

func addUpper(n int) int {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	return res
}

func getSub(n1, n2 int) int {
	return n1 - n2
}

```

`cal_test.go`
```go
package test

import (
	"testing"
)

// 编写一个测试用例，去测试 addUpper是否正确
// 必须以 Test 开头，AddUpper 首字母必须大写
func TestAddUpper(t *testing.T) {
	// 调用
	res := addUpper(10)
	if res != 55 {
		// fmt.Printf("addUpper 错误，返回值=%v 期望值=%v\n", res, 55)
		t.Fatalf("addUpper 错误，返回值=%v 期望值=%v\n", res, 55)
	}
	// 如果正确，输出日志
	t.Logf("addUpper(10)执行正确...")
}

func TestHello(t *testing.T) {
	t.Logf("TestHello被调用...")
}

```

`sub_test.go`

```go
package test

import (
	"testing"
)

// 编写一个测试用例，去测试 addUpper是否正确
// 必须以 Test 开头，AddUpper 首字母必须大写
func TestGetSub(t *testing.T) {
	// 调用
	res := getSub(10, 3)
	if res != 7 {
		// fmt.Printf("addUpper 错误，返回值=%v 期望值=%v\n", res, 55)
		t.Fatalf("getSub(10, 3) 错误，返回值=%v 期望值=%v\n", res, 7)
	}
	// 如果正确，输出日志
	t.Logf("getSub(10, 3)执行正确...")
}

```

单元测试总结
1. 测试用例文件名必须以 xxx_test.go 命名，比如 cal_test.go, sub_test.go 
2. 测试用例函数必须以 Test 开头，一般来说是 Test+被测试函数名，比如 TestAddUpper()
3. TestAddUpper(t *testing.T) 的形参必须是 *testing.T
4. 一个测试用例文件中可以有多个测试用例函数
5. 运行测试用例指令：`go test -v`
6. 当出现错误时，可以使用 t.Fataf 来格式化输出错误信息，并退出程序
7. t.Logf方法可以输出相应的日志
8. 测试用例函数并没有放在 main 函数中，但是却能够正确执行，这就是测试用例的方便支出
9. PASS 表示测试用例运行成功，FAIL表示测试用例运行失败
10. 只测试单个文件，可以加上被测试的源文件：`go test -v cal_test.go cal.go`
11. 测试单个方法：`go test -v -test.run TestAddUpper`


综合测试例子

项目结构如下：
```bash
exercise
├── data
│   └── monster.serial
├── monster.go
└── monster_test.go

1 directory, 3 files
```

`monster.go`
```go
package exercise

import (
	"encoding/json"
	"fmt"
	"os"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

// 序列化对象
func (m *Monster) Store() bool {
	mData, err := json.Marshal(m)
	if err != nil {
		fmt.Println("序列化失败", err)
		return false
	}

	// 保存序列化字符串到文件
	filepath := "./data/monster.serial"
	errWrite := os.WriteFile(filepath, mData, 0666)
	if errWrite != nil {
		fmt.Println("write file err=", err)
		return false
	}
	return true
}

// 反序列化
func (m *Monster) ReStore() bool {
	// 从文件中读取序列化的字符串
	filepath := "./data/monster.serial"
	data, errRead := os.ReadFile(filepath)
	if errRead != nil {
		fmt.Println("read file err=", errRead)
		return false
	}
	err := json.Unmarshal(data, m)
	if err != nil {
		fmt.Println("反序列化失败", err)
		return false
	}
	return true
}

```

`monster_test.go`
```go
package exercise

import (
	"testing"
)

// 测试用例，测试 Store 方法
func TestStore(t *testing.T) {
	// 先创建一个 monster 实例
	monster := Monster{
		Name:  "红孩儿",
		Age:   20,
		Skill: "三昧真火",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("monster Store() 错误，期望是%v， 实际是%v", true, res)
	}
	t.Logf("monster Store() 测试成功")
}

func TestReStore(t *testing.T) {
	// 先创建一个 monster 实例，先不需要指定字段的值
	var monster Monster
	res := monster.ReStore()
	if !res {
		t.Fatalf("monster.ReStore() 错误，期望是%v，实际是%v", true, res)
	}
	// 进一步判断，反序列后的值是正确的
	if monster.Name != "红孩儿" {
		t.Fatalf("monster.ReStore() 错误，期望是%v，实际是%v", "红孩儿", monster.Name)
	}

	t.Logf("monster.Store() 测试成功！")
}

```

运行测试代码：
```bash

# 测试单个函数
go test -v -test.run TestStore

go test -v -test.run TestReStore

# 测试整个源文件
go test -v
```

# 协程（goroutine）和管道（channel）

## 协程的需求分析

当一个循环非常耗时时（如统计1到999999999999中素数的个数），可以使用并发或并行的方式，将统计素数的任务分配给多个 goroutine 去完成。


## 进程和线程
1. 进程是程序在操作系统中一次执行过程，是系统进行资源分配和调度的基本单位
2. 线程是进程的一个执行实例，是程序执行的最小单元，它是比进程更小的能独立运行的基本单位
3. 一个进程可以创建和销毁多个线程，同一个进程中的多个线程可以并发执行
4. 一个程序至少有一个进程，一个进程至少有一个线程

## 并发和并行
1. 多线程程序在单核上运行，称为并发
2. 多线程程序在多核上运行，称为并行

更详细的解释：
- 并发：因为是在一个CPU上，假设程序有10个线程，每个线程执行10毫秒（进行轮询操作），从人的角度看，好像这10个线程都在运行，但从微观上看，某个时刻，其实只有一个线程在执行。
- 并行：因为是在多个CPU上，假设有10个CPU，程序有10个线程，每个线程执行10毫秒（各自在不同CPU上执行），从人的角度看，这10个线程都在运行，从微观上看，在某个时刻，同时有10个线程执行。


## golang主线程和协程
1. golang主线程（有时称为线程，也可以理解为进程）：一个golang线程上，可以起多个协程，可以理解为 协程是轻量级的线程（编译器做优化）
2. golang协程的特点：
   - 有独立的栈空间
   - 共享程序堆空间
   - 调度有用户控制
   - 协程是轻量级的线程

简单的入门例子：
```go
package main

import (
	"fmt"
	"strconv"
	"time"
)

// 在主线程（或进程）中，开启一个 goroutine 协程，每隔 1 秒输出一个 "hello, world"
// 在主线程，也每隔 1 秒输出一个 "hello, golang"，输出 10 次后，退出程序
// 要求主线程和 goroutine 协程同时执行

// 每隔 1 秒输出一个 "hello, world"
func test() {
	for i := 0; i < 10; i++ {
		fmt.Printf("goroutine test() %v: hello, world\n", strconv.Itoa(i))
		time.Sleep(time.Second * 1)

	}
}

func main() {
	go test() // 开启协程
	go test() // 开启协程

	// 主线程
	for i := 0; i < 10; i++ {
		fmt.Printf("main() %v: hello, golang\n", strconv.Itoa(i))
		time.Sleep(time.Second * 1)
	}
}
```

1. **主线程执行完就会退出程序，无论协程是否执行完成。**
2. 主线程是一个物理线程，直接作用在CPU上，是重要级的，非常耗费CPU资源
3. 协程从主线程开启，是轻量级的线程，是逻辑态，对资源耗费相对小
4. golang 的协程机制是重要的特点，可以轻松的开启上万个协程，其他编程语言的开发机制是一般基于线程的，开启过多的线程资源耗费大

## goroutine调度模型MPG模式
1. M 表示操作系统的主线程（是物理线程）
2. P 表示协程执行需要的上下文
3. G 表示协程


## runtime
```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println()
	fmt.Println(runtime.NumCPU()) // 服务器逻辑CPU个数

	runtime.GOMAXPROCS(runtime.NumCPU() - 1) // golang 1.8后不用显示设置，默认开启
}
```

## 全局互斥锁
```go
package main

import (
	"fmt"
	"sync"
	"time"
)

// 利用goroutine计算1-50的阶乘，并把结果存到一个map里
var (
	mapFactorial = make(map[int]uint64) // 存放结果
	// 声明一个全局互斥锁
	// lock 是一个全局互斥锁
	// sync 是包，synchornized 同步
	// mutex 互斥
	locker sync.Mutex // 结构体
)

// 计算一个数的阶乘
func Factorial(n int) {
	// 计算阶乘
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}

	locker.Lock() // 加锁
	mapFactorial[n] = uint64(res)
	locker.Unlock() // 解锁
}

func main() {
	// 使用协程
	for i := 1; i <= 50; i++ {
		go Factorial(i)
	}
	// 让主线程等待协程执行完毕
	time.Sleep(time.Second * 3)
	// 这里仍然需要全局锁
	locker.Lock()
	for i, v := range mapFactorial {
		fmt.Printf("mapFactorial[%d]=%d\n", i, v)
	}
	locker.Unlock()
}

```

## 管道channel
前面使用全局变量锁来解决goroutine的通讯，有如下缺点：
1. 主线程在等待所有协程全部完成的时间很难确定，设置的时间3秒，很难估算
2. 如果主线程休眠时间过长，会导致等待时间过长，如果设置时间较短，可能会导致部分协程处于工作状态，但是随着主线程的退出而销毁
3. 通过全局变量锁同步，也没有利用多个协程对全局变量的读写操作

管道的介绍
1. 管道的本质是一个数据结构，即队列
2. 数据是先进先出（FIFO,First in first out）（与栈相反，它是先进后出）
3. 线程安全，多协程访问时，不需要加全局锁。也就是说，管道本身就是线程安全的
4. 管道是有类型的，一个string的管道只能存放string类型的数据

### 声明管道
```bash
var 变量名 chan 数据类型

example:
var intChan chan int    // intChan 用于存放 int 数据
var mapChan chan map[int]string     // mapChan 用于存放 map[int]string 类型
var perChan chan Person 
var perChan2 chan *Person
```
1. 管道是引用类型
2. 管道必须初始化才能写入数据，即 make 后才能使用
3. 管道是有类型要求的，intChan 只能写入整数
4. 想在管道中放置多种数据类型数据，可以声明空接口`var kongChan chan interface{}`

```go
package main

import (
	"fmt"
)

func main() {
	// 创建一个管道
	var intChan = make(chan int, 3)

	// 查看
	fmt.Printf("intChan 的值=%v intChan本身的地址=%p\n", intChan, &intChan) // 管道是引用类型

	// 写入数据
	intChan <- 10
	num := 211
	intChan <- num

	// 管道的长度和容量cap，管道的容量不像map能够自动增长
	fmt.Printf("channel len = %v, cap = %v \n", len(intChan), cap(intChan))

	// 虽然管道存放的数据不能超过管道的容量，但是，可以先读取后空出位置后再增加数据，先进先出
	num2 := <-intChan
	fmt.Println(num2)
	intChan <- num2
	fmt.Printf("channel len = %v, cap = %v \n", len(intChan), cap(intChan))

	// 在没有使用协程的情况下，如果我们的管道数据已经全部取出，再取就会报死锁 deadlock
	num3 := <-intChan
	num4 := <-intChan
	fmt.Println("num3=", num3, "num4=", num4)

	fmt.Printf("channel len = %v, cap = %v \n", len(intChan), cap(intChan))

	// num5 := <-intChan // 管道中已经没有数据，再取会报死锁
	// fmt.Println("num5=", num5)
	
	intChan <- 5
	<-intChan // 取出管道的数据不要了
}
```
管道使用的注意事项：
1. 管道中只能存放指定的数据类型
2. 管道的数据放满后，就不能再放入了
3. 如果从管道取出数据后，那么可以继续放入，因为已经腾出位置了，遵循先进先出的原则
4. 在没有使用协程的情况下，如果管道中的数据取完了，再取会报死锁 dead lock
5. `<-intChan` 取出管道里的数据可以扔掉，不赋值给别的变量

### 案例
1. 创建一个 intChan，最多可以存放 3 个 int
```go
package main

import (
	"fmt"
)

func main() {
	var intChan chan int
	intChan = make(chan int, 3)
	intChan <- 10
	intChan <- 20
	intChan <- 10
	// 因为 intChan 的容量是3，再存放数据会报告 deadlock
	num1 := <- intChan
	num2 := <- intChan
	num3 := <- intChan
	// 因为 intChan 中的数据已经取完，再取会报告 deadlock
	fmt.Println(num1, num2, num3)
}

```
2. 创建一个 mapChan，最多可以存放 10 个 map[string]string 的 key-value
```go
package main

import (
	"fmt"
)

func main() {
	var mapChan chan map[string]string
	mapChan = make(chan map[string]string, 10)

	m1 := make(map[string]string, 20)
	m1["city1"] = "beijing"
	m2["city2"] = "shanghai"

	m2 := make(map[string]string, 20)
	m2["fruit1"] = "apple"
	m2["fruit2"] = "orange"

	mapChan <- m1
	mapChan <- m2

	fmt.Println(m1, m2)
}
```
3. 创建一个 catChan，最多可以存放 10 个 Cat 结构体变量
```go
package main

import (
	"fmt"
)

var Cat struct {
	Name string
	Age int
}

func main() {
	catChan := make(chan Cat, 10)
	cat1 := Cat{
		Name: "tom",
		Age: 1,
	}
	cat2 := Cat{
		Name: "jack",
		Age: 2,
	}

	catChan <- cat1
	catChan <- cat2

	cat11 := <- catChan
	cat22 := <- catChan

	fmt.Pintln(cat11, cat22)
}
```

4. 创建一个 chtChan2，最多可以存放 10 个 *Cat 变量
```go
package main

import (
	"fmt"
)

var Cat struct {
	Name string
	Age int
}

func main() {
	var catChan chan *Cat
	catChan = make(chan *Cat, 10)

	cat1 := Cat{"Tom", 1}
	cat2 := Cat{"Lucy", 2}

	catChan2 <- &cat1
	catChan2 <- &cat2

	cat11 := <- catChan2
	cat22 := <- catChan2

	fmt.Println(cat11, cat22)
}
```
5. 创建一个 allChan，最多可以存放 10 个任意数据类型的变量
```go
package main

import (
	"fmt"
)

func main() {
	var allChan chan interface{}

	allChan = make(chan interface{}, 10)

	cat1 := Cat{"tom", 1}
	cat2 := Cat{"tomm", 11}

	allChan <- cat1
	allChan <- cat2
	allChan <- 10
	allChan <- "jack"
	// 取出
	cat11 := <- allChan
	cat22 := <- allChan
	v1 := <- allChan
	v2 := <- allChan

	fmt.Println(cat11, cat22, v1, v2)

	// cat11 是空接口类型，需要先用类型断言转成 Cat 结构体类型
	cat111 := cat11.type(Cat)
	fmt.Println(cat111.Name)
}
```

经典案例
```go
package main

/*
goroutine 和 channel 协同工作的案例
1. 开启一个 writeData 协程，向管道 intChan 中写入 50 个整数；
2. 开启一个 readData 协程，从管道 intChan 中读取 writeData 写入的数据；
3. 主线程需要等待 writeData 和 readData 协程都完成工作才能退出
*/

import (
	"fmt"
)

func writeData(intChan chan int) {
	for i := 1; i <= 50; i++ {
		// 放入数据
		intChan <- i
		fmt.Println("writeData 写入数据", i)
	}
	close(intChan) // 关闭管道
}

func readData(intChan chan int, exitChan chan bool) {
	for {
		// 对于关闭的管道，如果数据取完了，那么会返回 false 给 ok
		v, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println("readData 读到数据", v)
	}

	// readData 读取完数据后，任务完成
	exitChan <- true
	close(exitChan)
}

func main() {
	fmt.Println()
	// 创建两个管道
	intChan := make(chan int, 50)
	exitChan := make(chan bool, 1)

	// 创建协程
	go writeData(intChan)
	go readData(intChan, exitChan)

	// 主程序等待协程完成
	for {
		_, ok := <-exitChan
		if !ok { // 如果 exitChan 中的数据读完了，那么关闭
			break
		}
	}
}
```

### 管道阻塞

1. 如果只是想管道中写数据，而不进行读，那么就会出现阻塞而 deaklock， 如上面的程序，将 main() 中的 go readData(intChan, exitChan) 注销 ；
2. 如果读管道的数据的频率，和写管道的数据的频率不一样，编译器会分析，不会进行阻塞；


### 应用案例

使用多协程计算素数
```go
package main

import (
	"fmt"
	"time"
)

func putNum(intChan chan int) {
	for i := 1; i <= 800000; i++ {
		intChan <- i
	}
	close(intChan)
}

func primeNum(intChan chan int, primeChan chan int, exitChan chan bool) {
	// 使用 for 循环
	var flag bool
	for {
		num, ok := <-intChan

		if !ok { // 如果 intChan 中的数据取完了
			break
		}
		flag = true
		// 判断 num 是不是素数
		for i := 2; i < num; i++ {
			if num%i == 0 { // 说明 num 不是素数
				flag = false
				break

			}
		}
		if flag { // 如果是素数，就放入到管道 primeChan
			primeChan <- num
		}

	}
	// fmt.Println("有一个 primeNum goroutine 因为取不到数据，退出")
	exitChan <- true
}

func main() {
	start := time.Now().UnixMilli()
	intChan := make(chan int, 10000) // 存放 1 - 8000
	primeChan := make(chan int, 800) // 存放素数

	// 标识协程退出
	var threads int = 30
	exitChan := make(chan bool, threads)

	// 开启一个 goroutine，向 intChan 放入 1- 8000 个整数
	go putNum(intChan)
	// 开启四个 goroutine，从 intChan 中取数，判断是否为素数，如果是，就放入到 primeChan
	for i := 0; i < threads; i++ {
		go primeNum(intChan, primeChan, exitChan)
	}

	// 如果从 exitChan 中取到了四个 true，那么说明 goroutin 都处理完了，那么关闭 primeChan
	go func() {
		for i := 0; i < threads; i++ {
			<-exitChan // 会等待
		}
		close(primeChan)
	}()

	// 遍历管道 primeChan，把结果取出
	for {
		_, ok := <-primeChan
		if !ok {
			break

		}
		// fmt.Printf("素数=%v\n", res)
	}

	end := time.Now().UnixMilli()
	fmt.Printf("主线程退出，耗时：%d 毫秒\n", end-start)
}
```

channel 使用细节和注意事项
1. channel 可以声明为只读`var onlyWriteChan <-chan int`，或者只写`var onlyReadChan chan<- int`
2. 默认情况下，管道是双向的，可读可写`var myChan chan int`

```go
package main

import (
	"fmt"
)

// 只写
func send(ch chan<- int, exitChan chan struct{}) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
	var a struct {}
	exitChan <- a
}

// 只读
func recv(ch <-chan int, exitChan chan struct{}) {
	for {
		v, ok := <-ch
		if !ok {
			break
		}
		fmt.Println(v)
	}
	var a struct{}
	exitChan <- a
}



func main() {
	var ch chan int
	ch = make(chan int, 10)
	exitChan := make(chan struct{}, 2)

	go send(ch, exitChan)
	go recv(ch, exitChan)

	var total = 0
	for _ = range exitChan {
		total++
		if total == 2 {
			break
		}
	}
	fmt.Println("done!")
}

```
3. 某个 gorountine 的 panic 异常会导致整个程序退出，可以使用 defer ... recover 捕获异常，使主线程不受影响，可以继续执行
```go
package main

import (
	"fmt"
	"time"
)

func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("Hello, World")
	}
}

func test() {
	// 使用 defer + recover 捕获异常
	defer func() {
		// 捕获函数抛出的 panic 异常
		if err := recover(); err != nil {
			fmt.Println("test() 发送错误", err)
		}
	}()

	var myMap map[int]string
	myMap[0] = "golang" // 使用前没有先 make，会报 panic 异常
}

func main() {
	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("main() ok=", i)
		time.Sleep(time.Second)
	}
}
```

### select 语句
channel 当不关闭遍历取数据时，会出现阻塞导致 dead lock，但是有时候无法确定什么时候需要关闭 channel 又想便利 channel，可以使用 select
```go
package main

import (
	"fmt"
)

func main() {
	intChan := make(chan int, 10)
	stringChan := make(chan string, 5)

	labelfor :
	for {
		select {
			case v := <- intChan:
			fmt.Println(v)
			case v := <- stringChan:
			fmt.Println(v)
			default:
			fmt.Println("都取不到")
			// break // 只会跳出 select 循环
			break labelfor // 会跳出外层的 for 循环
		}
	}
}
```

# 反射 reflect

反射使用的场景：
1. 结构体序列化和反序列化标签的使用，可以回归前面json章节；
2. 适配器函数，桥连接
```go
test1 := func(v1 int, v2 int) {
	t.Log(v1, v2)
}

test2 := func(v1 int, v2 int, s string) {
	t.Log(v1, v2, s)
}

// 定义一个适配器函数用作统一处理接口
bridge := func(call interface{}, args... interface{}) {
	// todo
}

// 调用 test1
bridge(test1, 1, 2)
// 调用 test2
bridge(test2, 1, 2, "test2")
```

1. 反射可以在运行时动态获取变量的各种信息，比如变量的类型(type)， 类别(kind)。普通变量的 type 和 kind 是一致的，但是结构体的不同；
2. 如果是结构体变量，可以获取结构体本身的信息（包括结构体的字段、方法等）；
3. 通过反射，可以修改变量的值，可以调用关联的方法；
4. 使用反射，需要 import "reflect"
```go
package main

import (
	"fmt"
	"reflect"
)

func reflectTest(b interface{}) {
	// 通过反射获取传入的变量的 type，kind 和值
	// 1. 先获取 reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("reflect type =", rType)

	// 2. 获取 reflect.Value
	rValue := reflect.ValueOf(b)
	fmt.Printf("reflect value = %v, reflect value type = %T\n", rValue, rValue)

	fmt.Println("2 + rValue =", 2+rValue.Int())

	// 3. 将 rValue 转成 interface{}
	iValue := rValue.Interface()
	// 将 interface{} 通过断言转成需要的类型
	num2 := iValue.(int)
	fmt.Println(num2)
}

func reflectTest02(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Printf("iType = %v, iType Type = %T\n", rType, rType)

	rValue := reflect.ValueOf(b)

	iValue := rValue.Interface()
	fmt.Printf("iValue = %v, iValue Type = %T\n", iValue, iValue)

	// 这里必须使用类型断言，虽然从运行结果看 iValue 的类型是 Student 结构体类型，但是那是运行时，编译时无法判断，会报错
	// 将 interface() 通过类型断言转成需要的类型，如果类型不确定，可以使用 switch
	student, ok := iValue.(Student)
	if ok {
		fmt.Printf("student.Name = %v\n", student.Name)
	}

	// 获取变量对应的 Kind, 1. rValue.Kind(); 2. rType.Kind()
	kind1 := rValue.Kind()
	kind2 := rType.Kind()
	fmt.Printf("kind = %v, kind = %v\n", kind1, kind2)

}

func reflectTest03(b interface{}) {
	rValue := reflect.ValueOf(b)
	fmt.Printf("rValue kind = %v\n", rValue.Kind())
	rValue.Elem().SetInt(2)
}

type Student struct {
	Name string
	Age  int
}

func main() {
	// 基本数据类型、空接口interface{}、reflect.Value 进行反射基本操作
	// 先定义一个 int
	var num int = 100
	reflectTest(num)

	// 定义一个 Student 实例
	student := Student{
		Name: "tom",
		Age:  10,
	}
	reflectTest02(student)

	// 使用反射修改变量的值
	var n int = 1
	reflectTest03(&n)
	fmt.Println("n =", n)
}
```

## 放射注意事项和细节
1. reflect.ValueOf.Kind 获取变量的类别，返回一个常量
2. Type 是类型，Kind 是类别
```go
var num int = 10 // num Type is int, Kind is int
var student Student // student Type is package.Student, Kind is struct
```
3. 通过反射可以让变量在 interface{}, reflect.Value 之间互相转换
4. 使用反射获取变量的值，必须保证数据类型匹配，否则报 panic
```go
var x int
reflect.ValueOf(x).Int() // right
reflect.ValueOf(x).Float() // panic
```
5. 通过反射修改变量的值，需要传入变了的地址，即指针类型来完成。同时，需要使用 reflect.ValueOf.Elem().SetXXX()

```go
package main

import (
	"fmt"
	"reflect"
)

// define a Monster struct
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster-age"`
	Score float32
	Sex   string
}

// method, show struct info
func (monster Monster) Print() {
	fmt.Println(monster)
}

// method, get sum
func (monster Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

// method, set monster field value
func (monster Monster) Set(name string, age int, score float32, sex string) {
	monster.Name = name
	monster.Age = age
	monster.Score = score
	monster.Sex = sex
}

func TestStruct(any interface{}) {
	// get reflect.Type
	rType := reflect.TypeOf(any)
	// get reflect.Value
	rValue := reflect.ValueOf(any)
	// judge is struct
	anyKind := rValue.Kind()
	if anyKind != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	// get num of field of struct
	numField := rValue.NumField()
	fmt.Printf("struct has %d fields\n", numField)

	// traversal all field
	for i := 0; i < numField; i++ {
		fmt.Printf("Field %d: value = %v\n", i, rValue.Field(i))
		// get tag of struct
		tagValue := rType.Field(i).Tag.Get("json")
		if tagValue != "" {
			fmt.Printf("Field %d: tag = %v\n", i, tagValue)
		}
	}

	// get num of method
	numMethod := rValue.NumMethod()
	fmt.Printf("struct has %d methods\n", numMethod)

	// call struct method, input arg is [] reflect.Value, here is null
	rValue.Method(1).Call(nil)

	// call struct method with arguments, input arg is []reflect.Value, and output arg is also []reflect.Value
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := rValue.Method(0).Call(params)
	fmt.Println("res =", res[0].Int())
}

func main() {
	// define a Monster instance
	monster := Monster{
		Name:  "monkey",
		Age:   500,
		Score: 88,
	}
	TestStruct(monster)
}

```

