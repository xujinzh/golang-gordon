package bytesslice

import (
	"bytes"
	"fmt"
)

func TestSlicesByteSplit() {
	// 打印分割符，方便查看本段函数的运行结果
	fmt.Println("\033[1;32m---------------------------------------------\033[0m")
	// 重复拼接
	var b = []byte{0xaa, 0xbb}
	result := bytes.Repeat(b, 3)
	fmt.Printf("result: %#q \n", result)
	// 转换为中文，方便计数
	b1 := []byte("你好中国")
	b2 := bytes.Runes(b1)
	fmt.Printf("转换前计数：%d\n", len(b1))
	fmt.Printf("转换后计数：%d\n", len(b2))

}
