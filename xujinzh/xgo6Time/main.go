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
