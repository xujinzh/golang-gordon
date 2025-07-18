package main

import (
	"fmt"
	"strconv"
	"strings"
)

// 累加器
func AddUpper() func(int) (int, string) {
	var n int = 10
	var str string = "hello"
	return func(x int) (int, string) {
		n += x
		str += strconv.FormatInt(int64(x), 10)
		// fmt.Println("n:", n, "str:", str)
		return n, str
	}
}

// 请编写一个函数 makeSuffix(suffix string) 可以接收一个文件后缀名（如 ".txt"）
// 并返回一个闭包。调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀名，
// 那么返回该文件名加上后缀名；如果已经有了指定的后缀名，则直接返回该文件名。

func makeSuffix(suffix string) func(string) string {
	return func(filename string) string {
		if strings.HasSuffix(filename, suffix) {
			return filename
		}
		return filename + suffix
	}
}

func main() {
	// 使用累加器
	addUpper := AddUpper()
	for i := 0; i < 10; i++ {
		resultN, resultStr := addUpper(i)
		fmt.Println(resultN, resultStr)
	}

	// 使用文件后缀名闭包
	addSuffix := makeSuffix(".txt")
	fmt.Println(addSuffix("file1"))     // 输出: file1.txt
	fmt.Println(addSuffix("file2.txt")) // 输出: file2.txt
	fmt.Println(addSuffix("file3.doc")) // 输出: file3.doc.txt
	fmt.Println("main() finished")
}
