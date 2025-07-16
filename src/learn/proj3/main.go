package main

import (
	// "crypto/rand"
	"fmt"
	_ "math"
)

// 冒泡排序函数
func BubbleSort(arr []int) {
	n := len(arr)
	// 外层循环控制排序轮数
	for i := 0; i < n-1; i++ {
		swapped := false // 优化标志位
		// 内层循环进行相邻元素比较和交换
		for j := 0; j < n-i-1; j++ {
			if arr[j] > arr[j+1] {
				// 交换相邻元素
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true // 发生交换则置为true
			}
		}
		// 如果某轮未发生交换，说明已有序，提前终止
		if !swapped {
			break
		}
	}
}

func main() {

	// 测试用例
	arr := []int{5, 3, 8, 6, 7, 2}
	fmt.Println("原始数组:", arr)
	BubbleSort(arr)
	fmt.Println("排序后数组:", arr)

	var f float64 = 3.1415926
	fmt.Printf("f = %9.2f\n", f)

	// label2:
	// 	for i := 0; i < 4; i++ {
	// 	label1:
	// 		for j := 0; j < 5; j++ {
	// 			if j == 2 {
	// 				break label1
	// 			}
	// 			fmt.Println("j=", j)
	// 		}
	// 		if i == 3 {
	// 			break label2
	// 		}
	// 	}

	// fmt.Println(time.Now().Unix())
	// fmt.Println(time.Now().Weekday())

	// var targetNum int = 99
	// var count int = 0
	// for {
	// 	n := rand.Intn(100) + 1
	// 	count++
	// 	if n == targetNum {
	// 		break
	// 	}
	// }
	// fmt.Printf("生成 %d 循环了 %d 次\n", targetNum, count)

	// // 水仙花数：一个3位数的水仙花数是指其各位数字的立方和等于它本身。
	// // 例如：153 = 1^3 + 5^3 + 3^3
	// // 下面的代码会输出所有的水仙花数。
	// var num uint32
	// fmt.Println("请输入一个三位数：")
	// fmt.Scanf("%d", &num)
	// numHundreds := num / 100
	// numTens := (num / 10) % 10
	// numOnes := num % 10
	// if num == uint32(math.Pow(float64(numHundreds), 3)+math.Pow(float64(numTens), 3)+math.Pow(float64(numOnes), 3)) {
	// 	fmt.Printf("%d 是水仙花数\n", num)
	// } else {
	// 	fmt.Printf("%d 不是水仙花数\n", num)
	// }

	// // 九九乘法表
	// for i := 1; i <= 9; i++ {
	// 	for j := 1; j <= i; j++ {
	// 		fmt.Printf("%d * %d = %2d \t", j, i, i*j)
	// 	}
	// 	fmt.Println()
	// }

	// // 镂空金字塔
	// var totalLevel int = 5
	// for i := 1; i <= totalLevel; i++ {
	// 	for k := 1; k <= totalLevel-i; k++ {
	// 		fmt.Print(" ")
	// 	}
	// 	for j := 1; j <= 2*i-1; j++ {
	// 		if j == 1 || j == 2*i-1 || i == totalLevel {
	// 			fmt.Print("*")
	// 		} else {
	// 			fmt.Print(" ")
	// 		}
	// 	}
	// 	fmt.Println()
	// }
	// var str string = "hello，中国"
	// for i := 0; i < len(str); i++ {
	// 	fmt.Printf("str[%d] = %c\n", i, str[i])
	// }

	// for index, val := range str {
	// 	fmt.Printf("str[%d] = %c\n", index, val)
	// }
	// for i := 0; i < 10; i++ {
	// 	fmt.Println("i=", i)
	// }

	// var x interface{}
	// var y float64 = 10.0
	// x = y
	// switch i := x.(type) {
	// case nil:
	// 	fmt.Printf("x type: %T, value: %v\n", i, i)
	// case int:
	// 	fmt.Printf("x type: %T, value: %v\n", i, i)
	// case float64:
	// 	fmt.Printf("x type: %T, value: %v\n", i, i)
	// case func(int) float64:
	// 	fmt.Printf("x type: %T, value: %v\n", i, i)
	// case bool:
	// 	fmt.Printf("x type: %T, value: %v\n", i, i)
	// default:
	// 	fmt.Printf("unkown type: %T, value: %v\n", i, i)
	// }

	// var score int = 88
	// switch {
	// case score >= 90:
	// 	fmt.Println("A+")
	// case score >= 80 && score < 90:
	// 	fmt.Println("A")
	// default:
	// 	fmt.Println("B")
	// }

	// var key byte
	// fmt.Println("please input a char: a, b, c")
	// fmt.Scanf("%c", &key)
	// switch key {
	// case 'a':
	// 	fmt.Println("you input a")
	// case 'b':
	// 	fmt.Println("you input b")
	// case 'c':
	// 	fmt.Println("you input c")
	// default:
	// 	fmt.Println("error")
	// }

	// if age := 17; age >= 18 {
	// 	fmt.Println("成年")
	// } else if age >= 16 {
	// 	fmt.Println("未成年")
	// }

	// var i int = 3
	// fmt.Printf("%b\n", i)
	// fmt.Printf("%o\n", 9)
	// fmt.Println(011 == 9)
	// fmt.Printf("%x\n", 17)
	// fmt.Println(0x11 == 17)

	// var name string
	// var age byte
	// var salary float32
	// var isPass bool

	// fmt.Println("please input name, age, salary and isPass, use space split")
	// fmt.Scanf("%s %d %f %t", &name, &age, &salary, &isPass)
	// fmt.Printf("name=%q, age=%d, salary=%f, isPass=%t\n", name, age, salary, isPass)

	// var name string
	// var age byte
	// var salary float32
	// var isPass bool
	// fmt.Printf("name: ")
	// fmt.Scanln(&name)

	// fmt.Printf("age: ")
	// fmt.Scanln(&age)
	// fmt.Printf("salary: ")
	// fmt.Scanln(&salary)
	// fmt.Printf("isPass: ")
	// fmt.Scanln(&isPass)
	// fmt.Printf("name=%q, age=%d, salary=%f, isPass=%t\n", name, age, salary, isPass)

	// var a int = 30
	// var b int = 40
	// var c int = 50
	// var max int
	// if a > b {
	// 	max = a
	// } else {
	// 	max = b
	// }

	// if max < c {
	// 	max = c
	// }
	// fmt.Printf("max(%v, %v, %v) = %v\n", a, b, c, max)

	// var a int = 30
	// var b int = 40
	// var max int
	// if a > b {
	// 	max = a
	// } else {
	// 	max = b
	// }

	// fmt.Printf("max = %v\n", max)

	// fmt.Println(10 / 4)
	// fmt.Println(10 / 4.0)

	// fmt.Println(-10 % 3) // a % b = a - a / b * b
	// fmt.Println(10 % -3)
	// fmt.Println(-10 % -3)

}
