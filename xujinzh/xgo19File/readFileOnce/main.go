package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println()
	// ioutil.ReadFile一次性将文件读取到位
	filepath := "../data/text.txt"
	content, err := os.ReadFile(filepath) // 因为没有显示的 Open 该文件，因此也不需要显示的 Close 该文件
	if err != nil {
		fmt.Println(err)
	}
	for _, c := range content {
		fmt.Println(c)
		fmt.Printf("%c\n", c)
	}
	fmt.Printf("%v", string(content))
}
