package main

import (
	"fmt"
	_ "unsafe"
	_ "strconv"
)

func main() {
	var f float32 = 3
	var ptr *float32 = &f
	*ptr = 3.33
	fmt.Println("变量 f 的值现在是：", f)
}