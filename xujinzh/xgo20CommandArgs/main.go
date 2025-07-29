package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("命令行的参数有", len(os.Args))
	// 遍历os.Args切片，得到每一个命令行参数值
	for i, v := range os.Args {
		fmt.Printf("args[%v]=%q\n", i, v)
	}
}
