package logstl_test

import (
	"fmt"
	"io"
	"log"
	"os"
	"testing"
)

// log 标准库
// golang 内置了 log 包，实现简单的日志服务。
// 通过调用 log 包的函数，可以实现简单的日志打印功能。
// log 包定义了 Logger 类型，该类型提供了一些格式化输出的方法。
// log  包也提供了一个预定义的"标准" logger，可以通过调用函数 Print 系列（Print|Printf|Println）、Fatal 系列（Fatal|Fatalf|Fatalln）和 Panic 系列（Panic|Panicf|Panicln）来使用。

// log 包中有 3 个系列的日志打印函数：Print 系列、Panic 系列、Fatal 系列
// Print 系列（Print|Printf|Println）：单纯打印日志
// Panic 系列（Panic|Panicf|Panicln），抛出 panic 异常
// Fatal 系列（Fatal|Fatalf|Fatalln）：打印日志：打印日志，强制结束程序（os.Exit(1)），defer 函数不会执行

// Print 系列
func TestLogPrint(t *testing.T) {
	// 直接打印日志，并换行。与常规的 Print 不同，这里自带换行
	log.Print("this is a log")
	// 格式化打印日志。与常规的 Printf 不同，这里自带换行
	log.Printf("this is a log: %d", 100)
	// 拼接字符串，打印日志
	name := "zhangsan"
	age := 20
	log.Println(name, " ", age)
}

// Panic 系列
// 打印出日志并且抛出 panic 异常，在 panic 之后声明的代码将不会执行
func TestLogPanic(t *testing.T) {
	// 声明一个 defer 语句，看是否能够在 panic 出现后还会执行否
	// 可以执行，函数退出后执行
	defer fmt.Println("我在 panic 后可以执行吗？")
	log.Print("this is a log")
	log.Panic("this is a panic log")
	// 这里后的代码将不会被执行
	fmt.Println("运行结束...")
}

// Fatal 系列
// 将日志内容打印输出，接着调用系统的 `os.Exit(1)` 接口，强制退出程序并返回状态 1
// 但是需要注意的是，由于直接调用系统的 os 接口退出，defer 函数将不会被调用
func TestLogFatal(t *testing.T) {
	// 声明一个 defer 语句，看是否能够在 fatal 出现后还会执行否
	// 不可以执行，直接调用了 os.Exit
	defer fmt.Println("我在 fatal 后可以执行吗？")
	log.Print("this is a log")
	log.Fatal("this is a fatal log")
	// fatal 后的代码将不会被执行
	fmt.Println("运行结束...")
}

// 日志配置
// 默认情况下 log 只会打印出时间，但是实际情况下我们还需要获取文件名、行号等信息，log 包提供给我们定制的接口
// func Flags() int 返回标准 log 输出配置
// func SetFlags(flag int) 设置标准 log 输出配置
/*
const (
	// 控制输出日志信息的细节，不能控制输出的顺序和格式
	// 输出的日志在每一项后会有一个冒号分隔，例如
	// 2026/02/06 22:56:54.123123 /a/b/c/d.go:23: message
	Ldate    = 1 << iota  // 日期，2026/02/06
	Ltime                 // 时间，22:56:54
	Lmicroseconds         // 微秒级别的时间，22:56:54.123123，用于增强 Ltime 位
	Llongfile             // 文件全路径名+行号，/a/b/c/d.go:23
	Lshortfile            // 文件名+行号，d.go:23 会覆盖掉 Llongfile
	LUTC                  // 使用 UTC 时间
	LstdFlag = Ldate | Ltime // 标准 logger 的初始值
)
*/
func TestLogSetFlags(t *testing.T) {
	// 默认 flags
	i := log.Flags()
	fmt.Printf("i: %v\n", i)
	// 设置 flags
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	log.Print("this is a log")
	// 打印 flags
	fmt.Printf("log.Flags(): %v\n", log.Flags())
	// 设置 flags
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Llongfile)
	log.Print("this is a log")
	// 打印 flags
	fmt.Printf("log.Flags(): %v\n", log.Flags())
}

// 前缀配置
// func Prefix() string 返回日志的前缀配置
// func SetPrefix(prefix string) 设置日志前缀
func TestLogPrefix(t *testing.T) {
	// 打印默认前缀
	fmt.Printf("log.Prefix(): %v\n", log.Prefix())
	// 设置前缀
	log.SetPrefix("[MyLog] ")
	// 再次打印前缀
	fmt.Printf("log.Prefix(): %v\n", log.Prefix())
	// 打印日志，看前缀
	log.Print("this is a log")
}

// 输出到文件
// log 包提供了 func SetOutput(w io.Writer) 函数，将日志输出到文件中
func TestLogOutputFile(t *testing.T) {
	// 创建一个文件
	f, err := os.OpenFile("./resources/test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		t.Log(err)
	}
	defer f.Close()
	// 设置把日志输出到文件
	log.SetOutput(f)
	log.SetFlags(log.Ldate | log.Lmicroseconds | log.Lshortfile)
	log.SetPrefix("[MyLog] ")
	// 输出日志信息
	log.Print("this is a log")
}

// 自定义 Logger
// log 包中提供了 func New(out io.Writer, prefix string, flag int) *Logger 函数来实现自定义 logger
func TestLogCustom(t *testing.T) {
	// 声明一个 logger 对象
	var logger *log.Logger
	// 打开一个文件，记录日志信息
	f, err := os.OpenFile("./resources/test.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		t.Log(err)
	}
	defer f.Close()
	// 创建一个 logger 实例
	logger = log.New(f, "[MyCustomLog] ", log.Ldate|log.Ltime|log.Llongfile)
	// 打印日志
	logger.Print("this is a custom log")

	// 同时输出到控制台和文件中
	multiWriter := io.MultiWriter(os.Stdout, f)
	logger = log.New(multiWriter, "[MyCustomMultWriterLog] ", log.Ldate|log.Lmicroseconds|log.Lshortfile)
	logger.Print("this is a multi writer custom log...")
}
