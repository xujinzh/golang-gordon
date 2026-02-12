package sortstl_test

import (
	"fmt"
	"slices"
	"sort"
	"testing"
)

/*
sort标准库


sort包提供了排序切片和用户自定义数据集以及相关功能函数。

sort包主要针对[]int, []float64, []string, 以及其他自定义切片排序。

主要包括：
- 对基本数据类型切片的排序支持
- 基本数据元素查找
- 判断基本数据类型切片是否已经排好序
- 对排好序的数据集逆序
*/

/*
1. 排序接口

type Inferface interface {
	Len()			int		// 获取数据集元素个数
	Less(i, j int) 	bool	// i > j，即i索引表示后一个数、j索引表示前一个数。如果i索引的数小于j索引的数据是返回true，调用Swap方法，表示升序；反之，表示降序。
	Swap(i, j int)			// 交换i和j索引的两个元素的位置
}
*/

// 声明一个新的数据类型
type NewInts []uint

// 实现接口的Len方法
func (intSlice NewInts) Len() int {
	return len(intSlice) // 返回切片长度
}

// 实现接口的Less方法
func (intSlice NewInts) Less(i, j int) bool {
	// fmt.Printf("i=%d, j=%d, intSlice[i]=%d, intSlice[j]=%d, intSlice[i] < intSlice[j] : %t, intSlice: %v\n",
	// 	i, j, intSlice[i], intSlice[j], intSlice[i] < intSlice[j], intSlice) // 辅助打印
	return intSlice[i] < intSlice[j] // 如果后一个值（索引i）小于前一个值（索引j），那么返回true，调用下面的交换Swap方法。即升序
}

// 实现接口的Swap方法
func (intSlice NewInts) Swap(i, j int) {
	intSlice[i], intSlice[j] = intSlice[j], intSlice[i] // 如果Less(i, j int)返回true，则执行交换
}

func TestSort(t *testing.T) {
	// 初始化一个NewInts切片
	intSlice := NewInts{1, 3, 2, 4}
	fmt.Printf("intSlice是否已经排好序：%t\n", sort.IsSorted(intSlice))
	fmt.Printf("排序前intSlice: %v\n", intSlice)
	// time.Sleep(time.Second)
	sort.Sort(intSlice) // 调用其排序方法
	// time.Sleep(time.Second)
	fmt.Printf("排序后intSlice: %v\n", intSlice)

	fmt.Printf("intSlice是否已经排好序：%t\n", sort.IsSorted(intSlice))
}

/*
2. 相关函数汇总

func Ints(a []int)
func IntsAreSorted(a []int) bool
func SearchInts(a []int, x int) int
func Float64s(a []float64)
func Float64sAreSorted(a []float64) bool
func SearchFlaot64s(a []float64, x float64) int
func SearchFloat64s(a []float64, x float64) bool
func Strings(a []string)
func StringsAreSorted(a []string) bool
func SearchStrings(a []string, x string) int
func Sort(data Interface)
func Stable(data Interface)
func Reverse(data Interface) Interface
func IsSorted(data Interface) bool
func Search(n int, f func(int) bool) int
*/

/*
3. 数据集合排序

3.1 Sort排序方法

对数据集合（包括自定义数据类型的集合）排序，需要实现sort.Interface接口的三个方法，即：

type Inferface interface {
	Len()			int		// 获取数据集元素个数
	Less(i, j int) 	bool	// i > j，即i索引表示后一个数、j索引表示前一个数。如果i索引的数小于j索引的数据是返回true，调用Swap方法，表示升序；反之，表示降序。
	Swap(i, j int)			// 交换i和j索引的两个元素的位置
}

实现这三个方法后，即可调用该包的Sort()方法进行排序。Sort()方法定义如下：

func Sort(data Interface)

Sort()方法唯一的参数就是待排序的数据集合。
*/

/*
3.2 IsSorted方法

判断数据集是否已经排好序。IsSorted方法的内部实现依赖于我们自己实现的Len()和Less()方法：

func IsSorted(data Interface) bool {
	n := data.Len()
	for i := n - 1; i > 0; i-- {
		if data.Less(i, i-1) {
			return false
		}
	}
	return true
}
*/

/*
3.3 Reverse逆序排序方法

将数据按Less()定义的排序方式逆序排序，而不必修改Less()代码。

func Reverse(data Interface) Interface

Reverse()的内部实现如下，其返回一个sort.Interface接口类型的值：

type reverse struct {
	Interface
}

// reverse结构类型的Less()方法拥有嵌入的Less()方法相反的行为。
// Len()和Swap()方法则会保持嵌入类型的方法行为。

func (r reverse) Less(i, j int) bool {
	return r.Interface.Less(j, i)
}

// 返回新的实现Interface接口的数据类型
func Reverse(data Interface) Interface {
	return &reverse{data}
}
*/

func TestReverse(t *testing.T) {
	// 初始化一个NewInts切片
	intSlice := NewInts{1, 3, 2, 4}
	fmt.Printf("intSlice是否已经排好序：%t\n", sort.IsSorted(intSlice))
	fmt.Printf("排序前intSlice: %v\n", intSlice)
	// 默认是升序，现在改为降序
	intSliceForReverse := sort.Reverse(intSlice)
	sort.Sort(intSliceForReverse) // 调用其排序方法
	fmt.Printf("排序后intSliceForReverse: %v\n", intSliceForReverse)

	fmt.Printf("intSliceForReverse是否已经排好序：%t\n", sort.IsSorted(intSliceForReverse))
}

/*
3.4 Search查询位置方法

func Search(n int, f func(int) bool) int

Search()方法会使用“二分查找”算法，来搜索某指定切片[0:n]，并返回能够使 f(i) = true的最小i(0 <= i < n)值。
并且会假定：如果 f(i)=true，那么f(i+1)=true。
即对于切片[0:n]，i之前的切片元素会是f()函数返回false，i及i之后的元素会使f()函数返回true。
但是，当在切片中无法找到时，f(i)=true的i时（此时切片元素都不能使f()函数返回true），Search()方法会返回n。

注意，这里假定数据集都是排好序的。
*/

func TestSearch(t *testing.T) {
	// 初始化一个NewInts切片
	intSlice := NewInts{1, 3, 2, 4}
	fmt.Printf("intSlice是否已经排好序：%t\n", sort.IsSorted(intSlice))
	fmt.Printf("排序前intSlice: %v\n", intSlice)

	sort.Sort(intSlice) // 调用其排序方法
	fmt.Printf("排序后intSlice: %v\n", intSlice)

	fmt.Printf("intSlice是否已经排好序：%t\n", sort.IsSorted(intSlice))

	// 在排序后的数据集中查找索引
	index := sort.Search(intSlice.Len(), func(i int) bool { return intSlice[i] >= 3 })
	fmt.Printf("index: %v\n", index)
}

/*
4. sort包支持的内部数据类型

4.1 []int排序

sort包定义了一个IntSlice类型，并且实现了sort.Interface接口：

type IntSlice []int
func (p IntSlice) Len() int {return len(p)}
func (p IntSlice) Less(i, j int) bool {return p[i] < p[j]}
func (p IntSlice) Swap(i, j int) {p[i], p[j] = p[j], p[i]}
// IntSlice类型定义了Sort()方法，包装了sort.Sort()函数
func (p IntSlice) Sort() {Sort(p)}
// IntSlice类型定义了Search()方法，包装了SearchInts()函数
func (p IntSlice) Search(x int) int {return SearchInts(p, x)}

并且，提供的sort.Ints()方法使用了该IntSlice类型：
func Ints(a []int) {Sort(IntSlice(a))}

所以，对[]int切片升序排序，经常使用sort.Ints()，而不是直接使用IntSlice类型。

如果要使用降序排序，要用前面提到的Reverse()方法。

如果要查找整数x在切片中的位置，sort包提供了SearchInts()
func SearchInts(a []int, x int) int
注意，SearchInts()的使用条件是切片a已经升序排序。
*/
func TestInts(t *testing.T) {
	// 升序
	iSlice := []int{3, 5, 1, 2, 4}
	fmt.Printf("排序前 iSlice: %v\n", iSlice)
	sort.Ints(iSlice)
	fmt.Printf("排序后 iSlice: %v\n", iSlice)

	// 降序
	iSlice1 := []int{333, 555, 111, 222, 444}
	sort.Sort(sort.Reverse(sort.IntSlice(iSlice1)))
	fmt.Printf("iSlice1: %v\n", iSlice1)

	// 查找
	iSlice2 := []int{5, 7, 9, 3, 6}
	sort.Ints(iSlice2)                   // 先排序
	index := sort.SearchInts(iSlice2, 3) // 再查找
	fmt.Printf("index: %v\n", index)

	// 使用slices提供的方法
	jSlice := []int{33, 55, 11, 22, 44}
	slices.Sort(jSlice)

	fmt.Printf("jSlice: %v\n", jSlice)
	fmt.Printf("slices.IsSorted(jSlice): %v\n", slices.IsSorted(jSlice))
	fmt.Printf("slices.Index([]int{1, 3, 5, 4, 2}, 5): %v\n", slices.Index([]int{1, 3, 5, 4, 2}, 5))

	jSlice1 := []int{3, 5, 7, 2, 4, 6}
	slices.Sort(jSlice1)
	slices.Reverse(jSlice1)
	fmt.Printf("jSlice1: %v\n", jSlice1)

}

/*
4.2 []float64排序

实现和Ints类似。内部实现：

type Float64Slice []float64
func (p Float64Slice) Len() int {return len(p)}
func (p Float64Slice) Less(i, j int) bool {return p[i] < p[j] || isNaN(p[i]) && !isNaN(p[j])}
func (p Float64Slice) Swap(i, j int) {p[i], p[j] = p[j], p[i]}
func (p Float64Slice) Sort() {Sort(p)}
func (p Float64Slice) Search(x float64) int {return SearchFloat64s(p, x)}

与 Sort(), IsSorted(), Search()相对应的三个方法：
func Float64s(a []float64)
func Float64sAreSorted(a []float64) bool
func SearchFloat64s(a []float64, x float64) int

其他如Search()方法与Ints类似。

需要注意：在上面Float64Slice类型定义的Less方法中，有一个内部函数isNaN()，其与math包中IsNaN()实现完全相同，sort包之所以不适用math.IsNaN()，完全是基于包依赖性的考虑。
应当看到，sort包的实现不依赖其他任何包。
*/

func TestFloat64(t *testing.T) {
	fSlice := []float64{1.1, 2.2, 5.5, 4.4, 0.0}
	sort.Float64s(fSlice)
	fmt.Printf("fSlice: %v\n", fSlice)

}

/*
4.3 []string排序

两个string对象之间的大小比较是基于“字典序”的。

实现与Ints类似

内部实现：

type StringSlice []string
func (p StringSlice) Len() int {return len(p)}
func (p StringSlice) Less(i, j int) bool {return p[i] < p[j]}
func (p StringSlice) Swap(i, j int) {p[i], p[j] = p[j], p[i]}
func (p StringSlice) Sort() {Sort(p)}
func (p StringSlice) Search(x sting) int {return SearchStrings(p, x)}

与 Sort(), IsSorted(), Search()相对应的三个方法：
func Strings(a []string)
func StringsAreSorted(a []string) bool
func SearchStrings(a []string, x string) int
*/
func TestString(t *testing.T) {
	// 字符串数字排序
	ls := sort.StringSlice{
		"100",
		"42",
		"41",
		"3",
		"2",
	}

	fmt.Printf("ls: %v\n", ls)
	sort.Strings(ls)
	fmt.Printf("ls: %v\n", ls)

	// 字符串字母排序
	ws := sort.StringSlice{
		"d",
		"ac",
		"c",
		"ab",
		"e",
	}
	fmt.Printf("ws: %v\n", ws)
	sort.Strings(ws)
	fmt.Printf("ws: %v\n", ws)

	// 字符串中文排序
	cs := sort.StringSlice{
		"啊",
		"博",
		"次",
		"得",
		"饿",
		"周",
	}
	fmt.Printf("cs: %v\n", cs)
	sort.Strings(cs)
	fmt.Printf("cs: %v\n", cs)

	for _, v := range cs {
		fmt.Println(v, []byte(v))
	}
}

/*
4.4 复杂结构：[][]int 排序


*/

type testSlice [][]int

func (ts testSlice) Len() int           { return len(ts) }
func (ts testSlice) Less(i, j int) bool { return ts[i][1] < ts[j][1] } // 根据第2个值升序
func (ts testSlice) Swap(i, j int)      { ts[i], ts[j] = ts[j], ts[i] }

func TestComplex(t *testing.T) {
	ts := testSlice{
		{1, 4},
		{9, 3},
		{7, 5},
	}
	fmt.Printf("ts: %v\n", ts)
	sort.Sort(ts)
	fmt.Printf("ts: %v\n", ts)

}

/*
4.5 复杂结构体排序

*/

type testMapSlice []map[string]float64

func (tms testMapSlice) Len() int           { return len(tms) }
func (tms testMapSlice) Less(i, j int) bool { return tms[i]["a"] < tms[j]["a"] } // 按照"a"对应的值排序
func (tms testMapSlice) Swap(i, j int)      { tms[i], tms[j] = tms[j], tms[i] }

func TestComples2(t *testing.T) {
	tms := testMapSlice{
		{"a": 4, "b": 12},
		{"a": 3, "b": 11},
		{"a": 5, "b": 10},
	}
	fmt.Printf("tms: %v\n", tms)
	sort.Sort(tms)
	fmt.Printf("tms: %v\n", tms)

}
