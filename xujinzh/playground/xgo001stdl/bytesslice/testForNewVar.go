package bytesslice

import "fmt"

// 辅助函数
func add(a, b int) (int, error) {
	return a + b, nil
}

// 在for循环中重新定义一个新变量（n, err := add(n, n)）会影响之前同名的变量吗？测试下
func TestForNewVar() {
	// 计数变量，辅助停止FOR循环
	var flagCount int = 0
	n := 1
	for n != 2 {
		fmt.Printf("---n= %d, n addr: %d\n", n, &n)
		// n, _ := add(n, n) // 注意，采用 := 得到 n 是一个新地址空间的新变量，它和上面的 n 不是同一个变量。上面的 n 地址空间不会变，但是这里得到的 n 会重新产生，每次产生的地址空间大概率是不一样的
		// n, err := add(n, n)
		// if err != nil {

		// }
		// 如果想达到预定的效果，需要采用下面的方式
		n, _ = add(n, n)
		fmt.Printf("+++n= %d, n addr: %d\n", n, &n)
		if flagCount >= 3 {
			break
		} else {
			flagCount++
		}
	}
}
