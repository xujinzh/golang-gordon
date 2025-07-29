package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("ok")
	f, err := os.Open("./data/text.txt")
	// f 叫文件指针、文件对象、文件句柄
	if err != nil {
		fmt.Println("打开失败")
	} else {
		fmt.Println("打开成功")
		fmt.Println(f)
	}

	err = f.Close()

	if err != nil {
		fmt.Println("关闭失败")
	} else {
		fmt.Println("关闭成功")
	}
}
