package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	// 打开源文件
	f, err := os.Open(srcFileName)
	if err != nil {
		fmt.Println("open src file err=", err)
	}
	defer f.Close()
	// 获取 reader
	reader := bufio.NewReader(f)

	// 打开目标文件
	f_writer, err_writer := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err_writer != nil {
		fmt.Println("open dst file err=", err_writer)
	}
	defer f_writer.Close()
	// 获取writer
	writer := bufio.NewWriter(f_writer)

	return io.Copy(writer, reader)

}

func main() {
	// 将 ../data/flower.webp 拷贝到 ../data/flower-copy.webp
	srcFileName := "../data/flower.webp"
	dstFileName := "../data/flower-copy.webp"
	w, e := CopyFile(dstFileName, srcFileName)
	if e != nil {
		fmt.Println("err=", e)
	} else {
		fmt.Println("拷贝完成")
		fmt.Println(w)
	}

}
