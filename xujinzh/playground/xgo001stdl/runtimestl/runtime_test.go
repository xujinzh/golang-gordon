package runtimestl

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

/*
runtime标准库

runtime包提供和go运行时环境的互操作，如控制goroutine函数。
它也包括用于reflect包的低层次类型信息。

1. 环境变量
- 环境变量GOGC设置最初的垃圾收集目标百分比。当新申请的数据和前次垃圾收集剩下的存活数据的比率达到该百分比时，就会触发垃圾收集。
默认GOGC=100，设置GOGC=off会完全关闭垃圾收集。
runtime/debug包的SetGCPercent函数运行在运行时修改该百分比。
- 环境变量GODEBUG控制运行时的debug输出。GODEBUG的指是逗号分隔的name=val对。
支持的name如下：
	- allocfreetrace
		设置为1时，会导致每次分配都会记录每一个对象的分配、释放及其堆栈踪迹。
	- efence
		设置为1时，会导致分配器运行模式为每个对象申请在独立的页和地址，且永不循环利用。
	- gctrace
		设置为1时，会导致每次垃圾回收器触发一行日志，包含内存回收的概要信息和暂停时间。设置为2时，不但会写入同样的概述还会重复收集。
	- gcdead
		设置为1时，会导致垃圾收集器摧毁任何它认为已经死掉的执行堆栈。
	- schedtrace
		设置为x时，会导致调度程序每隔x毫秒输出单行信息到标准错误输出。
	- scheddetail
		设置为schedtrace为x并设置scheddetail为1时，会导致调度程序每隔x毫秒输出详细的多行信息，描述调度、进程、线程和go协程的状态。
- 环境变量GOMAXPROCS限制可以同时运行用户层次的go代码的操作系统进程数。没有对代表go代码的、可以在系统调用中阻塞的go协程数的限制；
那些阻塞的go协程不与GOMAXPROCS限制冲突。本包的GOMAXPROCS函数可以查询和修改该限制。
- 环境变量GOTRACEBACK控制当go程序因为不能恢复的panic或不期望的运行时情况失败时的输出。
失败的程序默认会打印所有现存go协程的堆栈踪迹（省略运行时系统中的函数），然后以状态码2退出。
如果GOTRACEBACK为0，会完全忽略所有go协程的堆栈踪迹。
如果GOTRACEBACK为1，会采用默认行为。
如果GOTRACEBACK为2，会打印所有现存GO协程包括运行时函数的堆栈踪迹。
如果GOTRACEBACK为crash，会打印所有现存go协程包括运行时函数的堆栈踪迹，并且如果可能会采用操作系统特定的方式崩溃，而不是退出。例如，在Unix系统里，程序会释放SIGABRT信号以触发核心信息转储。
- 环境变量GOARCH, GOOS, GOPATH, GOROOT构成完整的go环境变量集合。
它们影响go程序的构建，GOARCH, GOOS, GOROOT在编译时 被记录并可用本包的常量和函数获取，但它们不会影响运行时环境。


gctrace每一行打印的日志格式如下：

gc {0} @{1}s {2}%: {3}+...+{4} ms clock, {5}+...+{6} ms cpu, {7}->{8}->{9} MB, {10} MB goal, {11} P

{0}: gc运行次数
{1}: 程序已运行的时间
{2}: gc占用的CPU百分比
{3}: 执行时间，包括程序延迟和资源等待
{4}: 也是执行时间，一般看这个
{5}: CPU clock
{6}: CPU clock
{7}: GC 启动前的堆内存
{8}: GC 运行后的堆内存
{9}: 当前堆内存
{10}: GC目标
{11}: 进程数


下面是一个存在内存泄漏的代码：

`
package main

import (
	"os"
	"os/signal"
)

func main() {
	go func() {
		m := make(map[int]int)
		for i := 0; ; i++ {
			m[i] = i
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig)

	<-sig
}
`
执行：GODEBUG=gctrace=1 go run main.go，查看运行时的内存情况。
可以看到程序运行过程中，每次GC，堆内存都在不断增大，这是一个很明显的内存泄漏场景。
*/

/*
2. runtime.Gosched

让出CPU时间片，重新等待安排任务。

`
package main

import (
	"fmt"
	"runtime"
)

func main() {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
			runtime.Gosched() // 让出CPU时间片，重新等待安排任务
		}
	}("world")

	for i := 0; i < 2; i++ {
		fmt.Println("hello")
		runtime.Gosched() // 让出CPU时间片，重新等待安排任务
	}
}
`
*/

func TestGosched(t *testing.T) {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
			runtime.Gosched() // 让出CPU时间片，重新等待安排任务
		}
	}("world")

	for i := 0; i < 2; i++ {
		fmt.Println("hello")
		runtime.Gosched() // 让出CPU时间片，重新等待安排任务
	}

}

/*
3. runtime.Goexit

退出当前协程
*/

func TestGoexit(t *testing.T) {
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("A.defer")

		func() {
			defer fmt.Println("B.defer")
			// 结束协程
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	// 主goroutine等待子gorountine结束，主协程再结束
	wg.Wait()
}

/*
4. runtime.GOMAXPROCS

go运行时调度器使用runtime.GOMAXPROCS参数来确定需要使用多少个os线程来同时执行go代码，默认值是机器上的CPU核心梳理，
例如一个8核心的机器上，调度器会把go代码同时调度到8个os线程上。

go可以通过runtime.GOMAXPROCS()函数来设置当前程序并发时占用的CPU逻辑核心数。
go1.5版本之前默认使用的时单核心运行，1.5之后默认使用全部的CPU逻辑核心数。

我们可以通过将任务分配到不同的CPU逻辑核心上，从而实现并行效果。


5. runtime.NumCPU

返回当前系统的CPU核心数。

6. runtime.GOOS
目标操作系统。
*/

var wg sync.WaitGroup

func a() {
	defer wg.Done()
	for i := 0; i < 100000000; i++ {
		// fmt.Println("A", i)
	}
}

func b() {
	defer wg.Done()
	for i := 0; i < 100000000; i++ {
		// fmt.Println("B", i)

	}
}

func TestGoMaxProcs(t *testing.T) {
	startTime := time.Now()
	// 设置go运行时的os线程数
	// 设置为1时执行时间要比设置为4时用时更长
	// 如果不设置，会默认设置机器所有CPU逻辑核心数
	// runtime.GOMAXPROCS(1)
	runtime.GOMAXPROCS(4)
	wg.Add(1)
	go a()
	wg.Add(1)
	go b()
	wg.Add(1)
	go a()
	wg.Add(1)
	go b()

	wg.Wait()

	fmt.Printf("time.Now().Sub(startTime): %v\n", time.Since(startTime))

	fmt.Printf("runtime.NumCPU(): %v\n", runtime.NumCPU())

	fmt.Printf("runtime.GOOS: %v\n", runtime.GOOS)
}
