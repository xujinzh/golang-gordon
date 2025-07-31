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
