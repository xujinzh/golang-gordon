package main

import (
	"encoding/json"
	"fmt"
)

func testSlice() {
	// 声明切片
	var slice []map[string]interface{}
	var m1 map[string]interface{}
	// 使用前先MAKE
	m1 = make(map[string]interface{})
	m1["name"] = "jack"
	m1["age"] = 7
	m1["address"] = [...]string{"北京", "南京"}
	// 添加到切片
	slice = append(slice, m1)
	// 先MAKE
	m2 := make(map[string]interface{})
	m2["name"] = "tom"
	m2["age"] = 13
	m2["address"] = "纽约"
	// 添加到切片
	slice = append(slice, m2)

	// 切片序列化
	data, err := json.Marshal(slice)
	if err != nil {
		fmt.Println("序列化错误")
		return
	}
	fmt.Printf("%q\n", string(data))

}

func main() {
	fmt.Println()
	testSlice()
}
