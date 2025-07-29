package main

import (
	"encoding/json"
	"fmt"
)

func unserialMap() {
	str := "{\"address\":\"芭蕉洞\",\"age\":30,\"name\":\"红孩儿\"}"
	// 定义一个map
	var a map[string]interface{}

	// a = make(map[string]interface{})

	// 不需要再make，因为反序列化底层会自动make
	err := json.Unmarshal([]byte(str), &a)
	if err != nil {
		fmt.Printf("反序列化失败")
		return
	}
	fmt.Println(a)
}

func main() {
	fmt.Println()
	unserialMap()
}
