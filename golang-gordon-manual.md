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

# map

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
切片的数据类型如果是 map，叫做 slice of map，map 切片中map 个数可以动态变化。

