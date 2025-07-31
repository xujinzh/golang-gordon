package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println()
	fmt.Println(runtime.NumCPU()) // 服务器逻辑CPU个数

	runtime.GOMAXPROCS(runtime.NumCPU() - 1) // golang 1.8后不用显示设置，默认开启
}
