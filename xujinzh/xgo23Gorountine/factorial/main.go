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
