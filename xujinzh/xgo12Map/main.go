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

	// map 的嵌套
	// 创建一个学生信息的嵌套 map
	// 学生姓名作为键，值是另一个 map，包含年龄、城市和爱好等信息
	// 例如: map[string]map[string]string
	studentMap := make(map[string]map[string]string)
	studentMap["xujinzh"] = make(map[string]string)
	studentMap["xujinzh"]["age"] = "18"
	studentMap["xujinzh"]["city"] = "成都"
	studentMap["xujinzh"]["hobby"] = "编程"

	studentMap["lisi"] = make(map[string]string)
	studentMap["lisi"]["age"] = "20"
	studentMap["lisi"]["city"] = "上海"
	studentMap["lisi"]["hobby"] = "音乐"	
	fmt.Println("学生信息:", studentMap)
	for student, info := range studentMap {
		fmt.Println("学生:", student)
		for key, value := range info {
			fmt.Printf("  %s: %s\n", key, value)
		}
	}
	
	// map 的增删改查
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "上海"
	cities["no3"] = "广州"	
	fmt.Println(cities)

	// 增加
	cities["no4"] = "深圳"
	fmt.Println("增加后:", cities)

	// 修改
	cities["no2"] = "杭州"
	fmt.Println("修改后:", cities)
	
	// 删除
	delete(cities, "no3") // 内置函数
	fmt.Println("删除后:", cities)

	// 查询
	if city, exists := cities["no1"]; exists {
		fmt.Println("查询到的城市:", city)
	} else {
		fmt.Println("没有查询到城市")
	}

	// map 的遍历
	for key, value := range cities {
		fmt.Printf("key: %s, value: %s\n", key, value)
	}
}
