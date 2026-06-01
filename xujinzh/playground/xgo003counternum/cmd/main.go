/*
计算一个整数的位数
用户从控制台输入一个整数，程序计算并输出该整数的位数
*/
package main

import (
	"fmt"
	"strings"

	"github.com/xujinzh/counternum/internal/counter"
)

func main() {
	counter.CounterNum()

	var input string

	for {
		fmt.Print("如果继续请输入 y，否则请输入 n 退出:")

		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("读取失败：", err)
			return
		}
		input = strings.TrimSpace(strings.ToLower(input))
		if input == "n" {
			break
		}
		counter.CounterNum()

	}
}
