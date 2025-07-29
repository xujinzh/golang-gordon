package test

import (
	"testing"
)

// 编写一个测试用例，去测试 addUpper是否正确
// 必须以 Test 开头，AddUpper 首字母必须大写
func TestAddUpper(t *testing.T) {
	// 调用
	res := addUpper(10)
	if res != 55 {
		// fmt.Printf("addUpper 错误，返回值=%v 期望值=%v\n", res, 55)
		t.Fatalf("addUpper 错误，返回值=%v 期望值=%v\n", res, 55)
	}
	// 如果正确，输出日志
	t.Logf("addUpper(10)执行正确...")
}

func TestHello(t *testing.T) {
	t.Logf("TestHello被调用...")
}
