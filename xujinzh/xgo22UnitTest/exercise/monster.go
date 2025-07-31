package exercise

import (
	"encoding/json"
	"fmt"
	"os"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

// 序列化对象
func (m *Monster) Store() bool {
	mData, err := json.Marshal(m)
	if err != nil {
		fmt.Println("序列化失败", err)
		return false
	}

	// 保存序列化字符串到文件
	filepath := "./data/monster.serial"
	errWrite := os.WriteFile(filepath, mData, 0666)
	if errWrite != nil {
		fmt.Println("write file err=", err)
		return false
	}
	return true
}

// 反序列化
func (m *Monster) ReStore() bool {
	// 从文件中读取序列化的字符串
	filepath := "./data/monster.serial"
	data, errRead := os.ReadFile(filepath)
	if errRead != nil {
		fmt.Println("read file err=", errRead)
		return false
	}
	err := json.Unmarshal(data, m)
	if err != nil {
		fmt.Println("反序列化失败", err)
		return false
	}
	return true
}
