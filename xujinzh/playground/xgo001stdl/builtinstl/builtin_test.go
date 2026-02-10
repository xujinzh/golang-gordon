package builtinstl_test

import (
	"fmt"
	"testing"
)

/*
builtin 标准库

builtin包提供了一些错误声明、变量和常量声明，还有一些便利函数，这个包不需要导入，这些变量和函数就可以直接使用。
*/

/*
1. 常用函数

1.1 append

func append(slice []Type, elems ...Type) []Type

- slice = append(slice, elem1, elem2): 直接在slice后面添加单个元素 //（错误：要求能转化成slice元素类型的值）添加元素类型可以和slice相同，也可以不同
- slice = append(slice, anotherSlice...): 直接将另一个slice添加到slice后面，但其本质上是将anotherSlice中的元素一个一个添加到slice中和上一种方式类似。
*/

func TestAppend(t *testing.T) {
	// 添加元素
	s1 := []int{1, 2, 3}
	s2 := append(s1, 'a', 4)

	fmt.Printf("s2: %v\n", s2)

	// 添加另一个切片
	s3 := []int{7, 8, 9}
	s4 := append(s1, s3...)

	fmt.Printf("s4: %v\n", s4)

}

/*
1.2 len

返回：数组、切片、字符串、通道的长度
*/

func TestLen(t *testing.T) {
	// 字符串的长度
	s := "hello world"
	fmt.Printf("len(s): %v\n", len(s))
	// 切片的长度
	s2 := []int{1, 2, 3}
	fmt.Printf("len(s2): %v\n", len(s2))

}

/*
1.3 print, println

打印输出到控制台
*/

func TestPrint(t *testing.T) {

	// 定义变量并初始化
	name := "zhangsan"
	age := 20

	// 打印变量。手动增加换行符
	print(name, " ", age, "\n")

	fmt.Println("---------")

	// 打印变量并换行
	println(name, " ", age)

}

/*
2. 重点常用函数

2.1 panic

抛出一个panic异常
*/

func TestPanic(t *testing.T) {
	defer fmt.Println("panic 异常后执行...")
	panic("panic 错误...")
	fmt.Println("unreachable code. end...")

}

/*
2.2 new 和 make

new和make的区别：
- make只能用来分配及初始化类型为slice，map，chan数据，new可以分配任意类型的数据；
- new分配返回的指针，即类型*T；make返回引用，即T；
- new分配的空间被清零，make分配后，会进行初始化。

2.2.1 new
*/

func TestNew(t *testing.T) {
	b := new(bool)
	fmt.Println(*b)

	i := new(int)
	fmt.Println(*i)

	s := new(string)
	fmt.Printf("*s: %q\n", *s)
}

/*
2.2.2 make

例如：make([]int, 10, 100)
说明：分配一个有100个int的数组，然后创建一个长度为10，容量为100的slice。
该slice引用包含前10个元素的数组。对应的，new([]int)返回一个指向新分配的，被置为零的slice结构体的指针，即指向值为nil的slice的指针

内建函数make(T, args)与new(T)的用途不一样。
它只用来创建slice, map, channel，并且返回一个初始化的（而不是置零），类型为T的值（而不是*T)。
之所以不同，是因为这三个类型的背后引用了使用前必须初始化的数据结构。
例如，slice是一个三元描述符，包含一个指向数据（在数组中）的指针，长度，以及容量，在这些项被初始化前，slice都是nil的。
对于slice，map和channel，make初始化这些内部数据结构，并准备好可用的值。
*/
func TestMake(t *testing.T) {

	var p *[]int = new([]int)
	var v []int = make([]int, 10)

	fmt.Printf("p: %v\n", p)
	fmt.Printf("v: %v\n", v)

	var p1 *[]int = new([]int)
	*p1 = make([]int, 5, 10)

	// idiomatic 习惯做法
	v1 := make([]int, 10)

	fmt.Printf("p1: %v\n", p1)
	fmt.Printf("v1: %v\n", v1)

}
