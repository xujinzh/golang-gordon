# 基本介绍

## Golang 
Golang 是 Google 开发的一种静态强类型、编译型、并发型，并具有垃圾回收功能的编程语言。 罗伯特·格瑞史莫、罗勃·派克及肯·汤普逊于 2007 年 9 月开始设计 Go，稍后伊恩·兰斯·泰勒、拉斯·考克斯加入项目。Golang 是基于 Inferno 操作系统所开发的。有时 Golang 也叫 Go 语言。

- 作业系统： Linux、macOS、FreeBSD、Windows
- 副档名： go
- 发行时间： 2009年11月10日，​13年前
- 实作者： Google
- 当前版本： : 1.21.3 (2023 年 10 月 10 日；稳定版本);
- 编程范型： 编译型，可平行化，结构化，指令式
- 设计者： 罗伯特·格瑞史莫（英语：Robert Griesemer）; 罗勃·派克; 肯·汤普逊

## 撰写风格
Golang 有定义如下的撰写风格：
1. 每行程序结束后不需要撰写分号;。
2. 大括号 { 不能够换行放置。
3. if判断式和 for 循环不需要以小括号包覆起来。
4. 使用 tab 做排版
除了第二点外（换行会产生编译错误），在不符合上述规定时，仍旧可以编译，但使用了内置 gofmt 工具后，会自动整理代码，使之符合规定的撰写风格。方法是：`gofmt -w main.go`

## 目标
Golang 的主要目标是“兼具 Python 等动态语言的开发速度和 C/C++ 等编译型语言的性能与安全性”

## 特性
- Go 语言的用途众多，可以进行网络编程、系统编程、并发编程、分布式编程。
- 因为Go语言没有类和继承的概念，所以它和 Java 或 C++ 看起来并不相同。但是它通过接口（interface）的概念来实现多态性。Go语言有一个清晰易懂的轻量级类型系统，在类型之间也没有层级之说。因此可以说Go语言是一门混合型的语言。
- Go 是编译型语言。在编译代码时，编译器检查错误、优化性能并输出可在不同平台上运行的二进制文件。
- Go 自带了编译器，因此无须单独安装编译器。
- Go语言支持交叉编译，比如说你可以在运行 Linux 系统的计算机上开发可以在 Windows 上运行的应用程序。这是第一门完全支持 UTF-8 的编程语言，这不仅体现在它可以处理使用 UTF-8 编码的字符串，就连它的源码文件格式都是使用的 UTF-8 编码。Go语言做到了真正的国际化！
- Go语言有一个吉祥物，在会议、文档页面和博文中，大多会包含下图所示的 Go Gopher，这是才华横溢的插画家 Renee French 设计的，她也是 Go 设计者之一 Rob Pike 的妻子。

<!-- ![](https://c.biancheng.net/uploads/allimg/180808/1-1PPQA9545W.jpg) -->
![](https://camo.githubusercontent.com/2b507540e2681c1a25698f246b9dca69c30548ed66a7323075b0224cbb1bf058/68747470733a2f2f676f6c616e672e6f72672f646f632f676f706865722f6669766579656172732e6a7067)
- 一种静态强类型、编译型、并发型、并具有垃圾回收功能的编程语言。


# 环境
## Golang 环境配置
访问 [Golang官网](https://go.dev/) 下载合适的版本。解压到喜欢的目录，并添加到环境变量即可。类似于 java 和 nodejs。大概如下：
```bash
export GOROOT=/usr/local/go
export PATH=$GO_HOME/bin:$PATH
export GOPATH=/Users/jinzhongxu/codes/go
```


## 测试环境
```bash
# 查看 Golang 版本信息
go version
```

## 编写简单程序
1. 创建文件夹：`go/learn/proj1/`；
2. 创建文件：`main.go`，在该文件中增加如下内容：
```go
package main

import "fmt"

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
go build main.go
# 执行
./main

# 编译成指定的文件名可执行程序
go build -o my_main main.go
```

