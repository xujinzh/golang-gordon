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
