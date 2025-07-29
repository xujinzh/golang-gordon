package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 打开一个已经存在的文件，将原来的内容覆盖成新内容：你好，月球
	filepath := "../data/textWriterBuffer.txt"
	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Printf("open file err=%v", err)
	}
	// 及时关闭文件句柄
	defer f.Close()
	// 准备写入的内容
	str := "你好，月球"
	// 使用带缓冲的写
	writer := bufio.NewWriter(f)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
		writer.WriteString("\r\n")
	}
	// flush缓冲
	err = writer.Flush()
	if err != nil {
		fmt.Println(err)
	}
}
