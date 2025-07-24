package main

import (
	"fmt"
	"math/rand"
	"sort"
)

type Hero struct {
	Name string
	Age  int
}

// hero 结构体切片类型
type HeroSlice []Hero

// 实现 Interface 接口的三个方法
func (heroslice HeroSlice) Len() int {
	return len(heroslice)
}

// 按 hero 的年龄从小到大排序
func (heroslice HeroSlice) Less(i, j int) bool {
	return heroslice[i].Age < heroslice[j].Age
	// return heroslice[i].Name < heroslice[j].Name
}

func (heroslice HeroSlice) Swap(i, j int) {
	heroslice[i], heroslice[j] = heroslice[j], heroslice[i]
}

func main() {
	var intSlice = []int{0, -1, 10, 7, 90}
	// 除了手撕冒泡排序等算法外，还可以用系统函数
	sort.Ints(intSlice) // 切片是引用类型
	fmt.Println(intSlice)

	// 结构体切片排序
	var heroSlices HeroSlice
	for i := 0; i < 10; i++ {
		hero := Hero{
			Name: fmt.Sprintf("hero-%d", rand.Intn(100)),
			Age:  rand.Intn(100),
		}
		heroSlices = append(heroSlices, hero)
	}

	// 排序前顺序
	for _, v := range heroSlices {
		fmt.Println(v)
	}
	// 排序
	sort.Sort(heroSlices)
	// 排序后顺序
	fmt.Println("排序后")
	for _, v := range heroSlices {
		fmt.Println(v)
	}
}
