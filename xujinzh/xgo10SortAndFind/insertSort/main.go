package main

import "fmt"

func InsertSort(arr *[7]int) {
	for i := 1; i < len(arr); i++ {
		insertVal := arr[i]
		insertIndex := i - 1
		// 从大到小排序
		for insertIndex >= 0 && arr[insertIndex] < insertVal {
			arr[insertIndex+1] = arr[insertIndex] // 数据后移
			insertIndex--
		}
		// 插入
		if insertIndex+1 != i {
			arr[insertIndex+1] = insertVal
		}
		fmt.Printf("第%d次插入的结果%v\n", i, *arr)
	}

}
func main() {
	arr := [7]int{23, 0, 12, 56, 34, 100, 3}
	InsertSort(&arr)
	fmt.Println(arr)

}
