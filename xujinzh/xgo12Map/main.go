package main

import (
	"fmt"
	"sort"
)

func modifyMap(m map[string]string) {
	// 修改 map 中的值
	m["name"] = "modifiedName"
	m["age"] = "30"
}

// 定义一个结构体
type Stu struct {
	Name    string
	Age     int
	Address string
}

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

	// map 切片

	// 声明一个 map 切片
	var sliceOfMaps = make([]map[string]string, 2)
	// 初始化 map 切片
	if sliceOfMaps[0] == nil {
		sliceOfMaps[0] = make(map[string]string, 2)
		sliceOfMaps[0]["name"] = "xujinzh"
		sliceOfMaps[0]["age"] = "18"

	}
	if sliceOfMaps[1] == nil {
		sliceOfMaps[1] = make(map[string]string, 2)
		sliceOfMaps[1]["name"] = "lisi"
		sliceOfMaps[1]["age"] = "20"
	}

	// 使用切片 append 函数，动态增加 map
	newMap := make(map[string]string, 2)
	newMap["name"] = "wangwu"
	newMap["age"] = "22"
	sliceOfMaps = append(sliceOfMaps, newMap)
	// 打印 map 切片
	fmt.Println(sliceOfMaps)

	//  map 排序
	mapVar := make(map[int]int, 10)
	mapVar[3] = 30
	mapVar[1] = 10
	mapVar[2] = 20
	mapVar[5] = 50
	mapVar[4] = 40
	mapVar[6] = 60
	fmt.Println("未排序的 map:", mapVar)
	// 使用切片存储 map 的键
	keys := make([]int, 0, len(mapVar))
	for key := range mapVar {
		keys = append(keys, key)
	}
	// 对键进行排序
	sort.Ints(keys)
	// 或者使用冒泡排序（不推荐，因为效率低）
	// for i := 0; i < len(keys)-1; i++ {
	// 	for j := i + 1; j < len(keys); j++ {
	// 		if keys[i] > keys[j] {
	// 			keys[i], keys[j] = keys[j], keys[i]
	// 		}
	// 	}
	// }
	fmt.Println("排序后的键:", keys)
	// 打印排序后的 map
	fmt.Println("排序后的 map:")
	for _, key := range keys {
		fmt.Printf("key: %d, value: %d\n", key, mapVar[key])
	}

	// 修改 map 的函数
	m1 := make(map[string]string)
	m1["name"] = "xujinzh"
	m1["age"] = "18"
	fmt.Println(m1)
	modifyMap(m1) // 说明 map 是引用类型
	fmt.Println(m1)

	// map 的值用结构体
	students := make(map[string]Stu)
	// 创建学生
	stu1 := Stu{"tom", 18, "Beijing"}
	stu2 := Stu{"lisa", 18, "Chongqing"}
	students["no1"] = stu1
	students["no2"] = stu2
	fmt.Println(students)

	// 遍历
	for k, v := range students {
		fmt.Println(k)
		fmt.Println(v.Name)
		fmt.Println(v.Age)
		fmt.Println(v.Address)
	}
}
