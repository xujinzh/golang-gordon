package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) int {
	// 使用 defer + recover 来处理除以零的异常
	defer func() {
		// recover 是内置函数
		err := recover() // golang 中的 recover 函数可以捕获 panic
		// 如果发生 panic，err 将不为 nil
		if err != nil {
			fmt.Println(err)
			b += 1 // 如果发生异常，b 加 1，避免除以零
			fmt.Println("Adjusted b to avoid division by zero:", b)
		}
	}()
	res := a / b
	return res
}

// 函数读取一个配置文件 config.ini 的信息
// 如果文件名不正常，返回一个自定义的错误
func readConfigFile(filename string) (err error) {
	if filename == "config.ini" {
		// read
		return nil
	} else {
		// 返回一个自定义的错误
		return errors.New("读取文件错误")
	}
}

func read() {
	err := readConfigFile("config.ini")
	if err != nil {
		panic(err) // 如果发生错误，触发 panic
	}
	fmt.Print("读取配置文件成功...\n")
}

func main() {
	a := 10
	b := 0 // 这里故意设置为0来触发异常
	result := divide(a, b)
	fmt.Println("The result of division is:", result)
	fmt.Println("done!")

	read() // 调用读取配置文件的函数
}
