package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	EngCount   int // 英文的个数
	NumCount   int // 数字的个数
	SpaceCount int // 空格的个数
	OtherCount int // 其他的个数
}

func main() {
	fmt.Println()
	// 统计一个文件中有多少个英文、数字、空格和其他字符
	// 思路：每读取一行统计一次，然后将结果保存到一个结构体
	filename := "../data/text.txt"
	f, err := os.Open(filename)
	if err != nil {
		fmt.Println("err=", err)
	}
	defer f.Close()

	// CharCount实例
	var count CharCount
	// 创建一个reader
	reader := bufio.NewReader(f)
	// 循环读取内容
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}

		// str = []rune(str) // 可以处理中文字符
		// 遍历 str，进行统计
		for _, v := range str {
			// fmt.Println(v)
			switch {
			case v >= 'a' && v <= 'z':
				fallthrough // 穿透
			case v >= 'A' && v <= 'Z':
				count.EngCount++
			case v >= '0' && v <= '9':
				count.NumCount++
			case v == ' ':
				count.SpaceCount++
			default:
				count.OtherCount++
			}

		}
	}
	// 输出统计结果
	fmt.Printf("英文字符的个数%v，数字的个数%v，空格的个数%v，其他字符个数%v\n", count.EngCount, count.NumCount, count.SpaceCount, count.OtherCount)
}
