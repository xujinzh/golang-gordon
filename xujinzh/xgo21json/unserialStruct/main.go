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

func unserial() {
	// 项目开发中通过其他方式获取得到
	str := "{\"name\":\"牛魔王\",\"age\":5000,\"Birthday\":\"1-11-11\",\"Sal\":8000,\"Skill\":\"一板斧\"}"

	// 定义一个 monster实例
	var monster Monster
	err := json.Unmarshal([]byte(str), &monster)
	if err != nil {
		fmt.Println("反序列化失败")
	}
	fmt.Println(monster)
	fmt.Println(monster.Name)
}

func main() {
	fmt.Println()
	unserial()
}
