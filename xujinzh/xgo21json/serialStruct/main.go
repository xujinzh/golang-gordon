package main

import (
	"encoding/json"
	"fmt"
)

type Monster struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Birthday string
	Sal      float64
	Skill    string
}

func testStruct() {
	var monster = Monster{
		Name:     "牛魔王",
		Age:      5000,
		Birthday: "1-11-11",
		Sal:      8000.0,
		Skill:    "一板斧",
	}
	// 序列化
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Printf("序列化失败")
		return
	}
	// 输出序列化后的及格过
	fmt.Printf("monster序列化后的结果%q\n", string(data))
}

func main() {
	fmt.Println()
	testStruct()
}
