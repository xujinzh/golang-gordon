package exercise

import (
	"testing"
)

// 测试用例，测试 Store 方法
func TestStore(t *testing.T) {
	// 先创建一个 monster 实例
	monster := Monster{
		Name:  "红孩儿",
		Age:   20,
		Skill: "三昧真火",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("monster Store() 错误，期望是%v， 实际是%v", true, res)
	}
	t.Logf("monster Store() 测试成功")
}

func TestReStore(t *testing.T) {
	// 先创建一个 monster 实例，先不需要指定字段的值
	var monster Monster
	res := monster.ReStore()
	if !res {
		t.Fatalf("monster.ReStore() 错误，期望是%v，实际是%v", true, res)
	}
	// 进一步判断，反序列后的值是正确的
	if monster.Name != "红孩儿" {
		t.Fatalf("monster.ReStore() 错误，期望是%v，实际是%v", "红孩儿", monster.Name)
	}

	t.Logf("monster.Store() 测试成功！")
}
