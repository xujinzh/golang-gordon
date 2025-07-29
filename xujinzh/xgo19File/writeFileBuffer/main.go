package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println()
	// 创建一个新文件 "../data/textWriterBuffer.txt"，并写入5句 hello, gordon
	filepath := "../data/textWriterBuffer.txt"
	f, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}

	// 及时关闭文件，防止内存泄露
	defer f.Close()

	// 写入5句 hello, gordon
	str := "hello, Gordon"
	// 使用带缓冲的 writer
	writer := bufio.NewWriter(f)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
		writer.WriteString("\r\n")
	}
	// 因为wirter是带缓冲，因此在调用 writerstring 方法时，其实内容是写入到缓冲的，不是写入到磁盘
	// 还需要调用 flush 方法写入到磁盘，否则文件中没有数据
	err = writer.Flush()
	if err != nil {
		fmt.Println("flush失败")
	}
}
