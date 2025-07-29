package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println()
	// 打开一个已经存在的文件，在原来的内容追加内容：hell0, moon
	filepath := "../data/textWriterBuffer.txt"

	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("打开文件失败", err)
	}

	defer f.Close()

	// 带缓冲写文件，先写入到缓冲区
	str := "hello, moon"
	writer := bufio.NewWriter(f)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
		writer.WriteString("\r\n")
	}

	// flush缓冲到磁盘
	err = writer.Flush()
	if err != nil {
		fmt.Println("flush err=", err)
	}
}
