package main

import (
	"fmt"
	"xgo/utils"
)

func main() {
	fmt.Println(utils.Name)

	var num int = 456
	str := fmt.Sprintf("%x", num)
	fmt.Printf("转换后的字符串: %s\n", str) // 输出: 转换后的字符串: 456
}
