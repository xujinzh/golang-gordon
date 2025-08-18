package main

import "fmt"

func QuickSort(left int, right int, array *[6]int) {
	l := left
	r := right
	// 中轴，支点
	pivot := array[(left+right)/2]
	temp := 0
	// 将比 pivot 小的数放到左边，大的数放到右边
	for l < r {
		// 从 pivot 左边找到大于等于 pivot 的值
		for array[l] < pivot {
			l++
		}
		// 从 pivot 的右边找小于等于 pivot 的值
		for array[r] > pivot {
			r--
		}
		// 如果本次分解完成
		if l >= r {
			break
		}
		temp = array[l]
		array[l] = array[r]
		array[r] = temp
		//
		if array[l] == pivot {
			r--
		}
		if array[r] == pivot {
			l++
		}
	}
	if l == r {
		l++
		r--
	}
	if left < r {
		QuickSort(left, r, array)
	}
	if right > l {
		QuickSort(l, right, array)
	}
}
func main() {
	arr := [6]int{-9, 78, 0, 23, -567, 70}
	QuickSort(0, len(arr)-1, &arr)
	fmt.Println(arr)
}
