package main

import (
	"fmt"
)

func main() {
	// 定义或声明一个二维数组
	var arr [4][6]int
	// 初始化二维数组
	arr[1][2] = 1
	arr[2][1] = 2
	arr[2][3] = 3

	// 打印二维数组
	fmt.Println("二维数组 arr:", arr)

	// 遍历二维数组
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[i]); j++ {
			fmt.Print(arr[i][j], " ")
		}
		fmt.Println()
	}

	fmt.Println(len(arr))
	fmt.Println(len(arr[0]))

	// 定义或声明一个二维数组
	var arr2 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	// 打印二维数组
	fmt.Println("二维数组 arr2:", arr2)

	var arr3 [2][3]int = [...][3]int{{1, 2, 3}, {4, 5, 6}}
	var arr4 = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	var arr5 = [...][3]int{{1, 2, 3}, {4, 5, 6}}
	arr6 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	arr7 := [...][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println("二维数组 arr3:", arr3)
	fmt.Println("二维数组 arr4:", arr4)
	fmt.Println("二维数组 arr5:", arr5)
	fmt.Println("二维数组 arr6:", arr6)
	fmt.Println("二维数组 arr7:", arr7)

	// 遍历
	for i := 0; i < len(arr3); i++ {
		for j := 0; j < len(arr3[i]); j++ {
			fmt.Printf("%v\t", arr3[i][j])
		}
		fmt.Println()
	}

	for i, v := range arr3 {
		for j, v2 := range v {
			fmt.Printf("arr3[%d][%d] = %d\t", i, j, v2)
		}
		fmt.Println()
	}

}
