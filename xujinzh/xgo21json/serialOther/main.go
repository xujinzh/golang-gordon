package main

import (
	"encoding/json"
	"fmt"
)

// 基本数据类型序列化
func testFloat64() {
	// 对基本数据类型序列化意义不大
	var num1 float64 = 2124.8
	// 对 num1 序列化
	data, err := json.Marshal(num1)
	if err != nil {
		fmt.Println("序列化失败")
		return
	}
	fmt.Printf("%q\n", string(data)) // 把浮点数转为字符串
}

func main() {
	fmt.Println()
	testFloat64()
}
