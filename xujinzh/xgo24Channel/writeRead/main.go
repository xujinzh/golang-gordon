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
