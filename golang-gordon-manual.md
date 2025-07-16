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
// 先设置 GO111MODULE 为自动模式
go env -w GO111MODULE=auto
cd $GOPATH
// 编译时，编译路径不需要带 src，编译器会自动带
go build learn/proj3
// 编译时需要编译 main 包所在的文件夹
// 项目的目录结构要按照规范组织：$GOPATH/src/项目名/包名/源码文件，main.go 一般放在 main 文件夹（包名）下
// 可以自定编译后可执行文件名
go build -o bin/my.exe learn/proj3
```



