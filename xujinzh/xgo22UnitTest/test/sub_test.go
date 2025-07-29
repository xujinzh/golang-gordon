package test

import (
	"testing"
)

// 编写一个测试用例，去测试 addUpper是否正确
// 必须以 Test 开头，AddUpper 首字母必须大写
func TestGetSub(t *testing.T) {
	// 调用
	res := getSub(10, 3)
	if res != 7 {
		// fmt.Printf("addUpper 错误，返回值=%v 期望值=%v\n", res, 55)
		t.Fatalf("getSub(10, 3) 错误，返回值=%v 期望值=%v\n", res, 7)
	}
	// 如果正确，输出日志
	t.Logf("getSub(10, 3)执行正确...")
}
