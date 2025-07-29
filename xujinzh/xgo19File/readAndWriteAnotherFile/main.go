package main

import (
	"fmt"
	"os"
)

func main() {
	// 将 ../data/textWriterBuffer.txt 读取并将内容写入到 ../data/textReadWrite.txt
	filepathRead := "../data/textWriterBuffer.txt"
	filepathWrite := "../data/textReadWrite.txt"
	// 读取文件，隐式打开，不需要显示关闭
	content, err := os.ReadFile(filepathRead)
	if err != nil {
		fmt.Println("err=", err)
	}
	// 写入文件
	errWrite := os.WriteFile(filepathWrite, content, 0666)
	if errWrite != nil {
		fmt.Println("write err=", err)
	}
}
