package main

import (
	"encoding/json"
	"fmt"
)

func unserialSlice() {
	str := "[{\"address\":[\"北京\",\"南京\"],\"age\":7,\"name\":\"jack\"},{\"address\":\"纽约\",\"age\":13,\"name\":\"tom\"}]"

	// 定义一个切片
	var slice []map[string]interface{}

	err := json.Unmarshal([]byte(str), &slice)
	if err != nil {
		fmt.Println("反序列化失败")
		return
	}
	fmt.Println(slice)
}

func main() {
	fmt.Println()
	unserialSlice()
}
