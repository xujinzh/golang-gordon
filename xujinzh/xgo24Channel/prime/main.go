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
