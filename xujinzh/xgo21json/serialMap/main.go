package main

import (
	"encoding/json"
	"fmt"
)

func testMap() {
	// 定义一个map
	var a map[string]interface{}
	// 使用map前，需要先make
	a = make(map[string]interface{})
	a["name"] = "红孩儿"
	a["age"] = 30
	a["address"] = "芭蕉洞"
	// 序列化map
	data, err := json.Marshal(a) // data 是 []byte
	if err != nil {
		fmt.Println("序列化失败")
		return
	}
	fmt.Printf("a map 序列化后=%q\n", string(data))
}

func main() {
	fmt.Println()
	testMap()
}
