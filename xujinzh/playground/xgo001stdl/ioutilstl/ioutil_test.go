package ioutilstl_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

// ioutil 标准库

/*
ioutil包封装了一些I/O实用程序函数

ReadAll 读取数据，返回读到的字节slice
ReadDir	读取一个目录，返回目录入口数组[]os.FileInfo
ReadFile	读一个文件，返回读到的文件内容字节slice
WriteFile	根据文件路径，写入字节slice
TempDir	在一个目录中创建指定前缀名的临时目录，返回新临时目录的路径
TempFile	在一个目录中创建指定前缀名的临时文件，返回os.File
NopCloser	用一个无操作的Close方法包装文件并返回一个ReadCloser接口
*/

// 1 NopCloser

/*
func NopCloser(r io.Reader) io.ReadCloser

就是将一个不带Close的Reader封装成ReadCloser
*/

func TestNopCloser(t *testing.T) {
	f, err := os.Open("./resources/a.txt")
	if err != nil {
		t.Log(err)
	}
	readCloser := ioutil.NopCloser(f)
	fmt.Printf("readCloser: %v\n", readCloser)
}
