package main

import "fmt"

func SelectSort(arr *[5]int) {
	// (*arr)[0] = 600 // 等价于 arr[0] = 600
	for i := 0; i < len(*arr)-1; i++ {
		for j := i + 1; j < len(*arr); j++ {
			if (*arr)[i] < (*arr)[j] {
				(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
			}
		}
	}
}

func main() {
	// 定义一个数组，从大到小排序
	arr := [5]int{10, 34, 18, 99, 79}
	fmt.Println(arr)
	SelectSort(&arr)
	fmt.Println(arr)
}
