package main

import (
	"fmt"
)

func main() {
	// map 的声明
	var m map[string]string
	// map 的初始化
	m = make(map[string]string)
	// map 的赋值
	m["name"] = "xujinzh"
	m["age"] = "18"
	fmt.Println("map 的内容:", m) // 输出: map 的内容: map[age:18 name:xujinzh], map 是无序的
	// map 的遍历
	m["age"] = "20" // 更新 age 的值
	for key, value := range m {
		fmt.Println("key:", key, "value:", value)
	}
	// map 的删除
	delete(m, "age")
	// map 的长度
	fmt.Println("map 的长度:", len(m))
	// map 的查询
	if value, ok := m["name"]; ok {
		fmt.Println("查询到的值:", value)
	} else {
		fmt.Println("没有查询到值")
	}
	// map 的 nil 判断
	if m == nil {
		fmt.Println("map 是 nil")
	} else {
		fmt.Println("map 不是 nil")
	}
	// map 的 nil 初始化
	m = nil
	if m == nil {
		fmt.Println("map 现在是 nil")
	} else {
		fmt.Println("map 仍然不是 nil")
	}
}
