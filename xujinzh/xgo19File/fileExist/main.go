package main

import (
	"fmt"
	"os"
)

func IsFileExist(filepath string) (bool, error) {
	_, err := os.Stat(filepath)
	if err == nil {
		// 文件存在
		return true, nil
	} else if os.IsNotExist(err) {
		// 文件不存在
		return false, nil
	} else {
		// 不能判断文件是否存在
		return false, err
	}
}

func main() {
	fmt.Println()
	// 判断文件是否存在
	filepath := "../data/text.txt"
	b, _ := IsFileExist(filepath)
	if b {
		fmt.Printf("%q 存在\n", filepath)
	} else {
		fmt.Printf("%q 不存在\n", filepath)
	}

}
