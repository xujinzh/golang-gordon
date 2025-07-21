package main

import (
	"fmt"
)

func changeSlice(sli []int) {
	sli[0] = 100 // 修改切片的第一个元素
}

func fbn(n int) []uint64 {
	// 声明一个切片，切片大小为 n
	fbnSlice := make([]uint64, n)
	// 初始化前两个元素
	fbnSlice[0] = 1
	fbnSlice[1] = 1
	// 计算斐波那契数列
	for i := 2; i < n; i++ {
		fbnSlice[i] = fbnSlice[i-1] + fbnSlice[i-2]
	}
	return fbnSlice
}

func main() {
	var arr = [...]int{1, 2, 3, 4, 5}
	sli0 := arr[1:3]
	fmt.Println("Slice from array:", sli0)
	fmt.Println("Length of slice:", len(sli0))
	fmt.Println("Capacity of slice:", cap(sli0)) // 切片的容量是动态变化的

	for i := 0; i < len(sli0); i++ {
		fmt.Println("Element at index", i, ":", sli0[i])
	}

	for i, v := range sli0 {
		fmt.Println("Element at index", i, ":", v)
	}

	var sli1 []int = make([]int, 4, 10)
	fmt.Println(sli1)

	sli2 := make([]int, 3, 6)
	fmt.Println(sli2)

	var slice []int
	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Println("Slice contents:", slice)
	fmt.Println("Hello, World!")

	slice = append(slice, slice...)
	fmt.Println("Slice after appending itself:", slice)

	// 切片拷贝
	var sli4 []int = []int{1, 2, 3, 4, 5}
	var sli5 = make([]int, 10)
	copy(sli5, sli4) // 拷贝 sli4 到 sli5, sli4 and sli5 在内存中有独立的空间
	fmt.Println("Original slice:", sli4)
	fmt.Println("Copied slice:", sli5)

	var slice6 []int
	var arr3 [5]int = [...]int{1, 2, 3, 4, 5}
	slice6 = arr3[:]
	var slice7 = slice6
	slice7[0] = 100 // 修改 slice7 的第一个元素
	fmt.Println("Original array after modifying slice7:", arr3)
	fmt.Println("Slice after modifying slice6:", slice6)
	fmt.Println("Slice7 after modifying:", slice7)

	var sli8 = []int{1, 2, 3, 4, 5}
	fmt.Println("Before changeSlice:", sli8)
	changeSlice(sli8) // 传递切片到函数
	fmt.Println("After changeSlice:", sli8)

	str := "hello, world"
	sli9 := str[7:12] // 切片字符串
	fmt.Println("String slice:", sli9)

	// 修改字符串的值
	var str2 = "hello, world"
	sli10 := []byte(str2)
	sli10[0] = 'z' // 修改切片的第一个字节
	fmt.Println("Modified byte slice:", sli10)
	str2 = string(sli10) // 转换回字符串
	fmt.Println("String after modification:", str2)

	// 转成 []byte 后，可以处理英文和数字，但不能处理中文
	// 原因是 []byte 只处理单字节字符，而中文字符通常是3字节的，因此会出现乱码
	// 解决方法是将 string 转成 []rune 即可，因为 []rune 可以处理按字符处理，兼容中文
	str3 := "你好，世界"
	sli11 := []rune(str3)
	sli11[0] = '都' // 修改切片的第一个字符
	fmt.Println("Modified rune slice:", sli11)
	str3 = string(sli11) // 转换回字符串
	fmt.Println("String after rune modification:", str3)

	// 斐波那契数列
	n := 10
	fbnSlice := fbn(n)
	fmt.Println("Fibonacci slice:", fbnSlice)
}
