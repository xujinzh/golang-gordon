package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	// 打开一个存在的文件，读里面的内容并输出到终端，同时可以写入内容：are you ok?
	filepath := "../data/textWriterBuffer.txt"

	f, err := os.OpenFile(filepath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("err", err)
	}

	defer f.Close()

	// 先读取原来的内容，并显示在终端
	reader := bufio.NewReader(f)
	for {
		strRead, errRead := reader.ReadString('\n') // 一行的结束标志为 \n
		if errRead == io.EOF {                      // 读到文件的末尾
			break
		} else if errRead != nil && errRead != io.EOF {
			fmt.Println("err=", errRead)
		} else {
			fmt.Print(strRead)
		}

	}

	// 再写入
	strWrite := "are you ok?"
	writer := bufio.NewWriter(f)
	for i := 0; i < 5; i++ {
		writer.WriteString(strWrite)
		writer.WriteString("\r\n")
	}

	// flush 缓冲
	errFlush := writer.Flush()
	if errFlush != nil {
		fmt.Println("flush err=", errFlush)
	}

}
