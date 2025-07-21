package main

import (
	"fmt"
)

// 冒泡排序
// Bubble Sort
func bubbleSort(arr *[6]int) {
	for i := 0; i < len(*arr)-1; i++ {
		flag := false // 优化标志
		// 每次遍历，假设没有交换
		for j := 0; j < len(*arr)-1-i; j++ {
			if (*arr)[j] > (*arr)[j+1] {
				// 交换位置
				(*arr)[j], (*arr)[j+1] = (*arr)[j+1], (*arr)[j]
				flag = true // 发生了交换
			}
		}
		if !flag {
			// 如果没有发生交换，说明数组已经有序，可以提前结束
			break
		}
	}
}

// 二分查找
func binarySearch(arr *[]int, leftIndex int, rightIndex int, target int) int {
	if leftIndex > rightIndex {
		return -1 // 没有找到目标元素，返回 -1
	}
	
	midIndex := (leftIndex + rightIndex) / 2 // 计算中间索引
	if (*arr)[midIndex] == target {
		return midIndex // 找到目标元素，返回索引
	} else if (*arr)[midIndex] > target {
		// 目标元素在左半部分
		return binarySearch(arr, leftIndex, midIndex-1, target)
	} else {
		// 目标元素在右半部分
		return binarySearch(arr, midIndex+1, rightIndex, target)
	}
}

func main() {
	fmt.Println("切片的容量 cap(slice) 是动态变化的，根据切片元素的个数动态调整。")
	// 定义数组
	arr := [...]int{5, 2, 9, 1, 5, 6}
	fmt.Println("排序前", arr)
	bubbleSort(&arr)
	fmt.Println("排序后", arr)

	arr2 := []int{1, 2, 3, 4, 5, 6}
	target := 4
	index := binarySearch(&arr2, 0, len(arr2)-1, target)
	fmt.Printf("目标元素 %d 在数组中的索引是: %d\n", target, index)
}
