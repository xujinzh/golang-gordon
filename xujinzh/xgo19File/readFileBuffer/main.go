package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("../data/text.txt")
	if err != nil {
		fmt.Println("文件打开失败")
	}
	// 函数退出时，及时关闭文件句柄
	defer f.Close()

	// 创建 Reader，默认缓冲 4096
	reader := bufio.NewReader(f)
	// 循环读取文件内容
	for {
		str, err := reader.ReadString('\n') // 读到字符'\n'就结束
		if err == io.EOF {                  // io.EOF代表文件末尾
			break
		} else if err != nil {
			fmt.Println("读取错误", err)
		} else {
			// 输出内容
			fmt.Print(str) // str 带 '\n'
		}

	}
	fmt.Println("文件读取结束...")
}
