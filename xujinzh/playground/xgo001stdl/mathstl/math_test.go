package mathstl_test

import (
	"fmt"
	"math"
	"math/rand"
	"testing"
	"time"
)

/*
math标准库

math包包含一些常量和一些有用的数学计算函数，例如三角函数、随机数、绝对值、平方等。
*/

/*
1. 常量
*/

func TestConstant(t *testing.T) {
	fmt.Printf("Float64的最大值: %.f\n", math.MaxFloat64)
	fmt.Printf("Float64最小值: %.f\n", math.SmallestNonzeroFloat64)
	fmt.Printf("Float32最大值: %.f\n", math.MaxFloat32)
	fmt.Printf("Float32最小值: %.f\n", math.SmallestNonzeroFloat32)
	fmt.Printf("Int8最大值: %d\n", math.MaxInt8)
	fmt.Printf("Int8最小值: %d\n", math.MinInt8)
	fmt.Printf("Uint8最大值: %d\n", math.MaxUint8)
	fmt.Printf("Int16最大值: %d\n", math.MaxInt16)
	fmt.Printf("Int16最小值: %d\n", math.MinInt16)
	fmt.Printf("Uint16最大值: %d\n", math.MaxUint16)
	fmt.Printf("Int32最大值: %d\n", math.MaxInt32)
	fmt.Printf("Int32最小值: %d\n", math.MinInt32)
	fmt.Printf("Uint32最大值: %d\n", math.MaxUint32)
	fmt.Printf("Int64最大值: %d\n", math.MaxInt64)
	fmt.Printf("Int64最小值: %d\n", math.MinInt64)
	fmt.Printf("圆周率默认值: %v\n", math.Pi)

}

/*
2. 常用函数

2.1 IsNaN函数

func IsNaN(f float64) (is bool)

报告f是否表示一个NaN(Not A Number)值，如果是数字，那么返回false；如果不是数值，那么返回true。
*/

func TestIsNaN(t *testing.T) {
	fmt.Println(math.IsNaN(32.1))
	// fmt.Println(math.IsNaN(float64("a")))
	fmt.Printf("math.IsNaN(math.Mod(1, 0)): %v\n", math.IsNaN(math.Mod(1, 0)))
	fmt.Println("\033[0;31m-------------------\033[0m")
}

/*
2.2 Ceil函数

func Ceil(x float64) float64
返回一个不小于x的最小整数，即向上取整。

2.3 Floor函数

func Floor(x float64) float64
返回一个不大于x的最小整数，即向下取整。

2.4 Trunc函数

func Trunc(x float64) float64
返回x整数部分，与Floor函数效果一样。
*/

func TestRoundNum(t *testing.T) {
	// ceil
	fmt.Printf("math.Ceil(3.1415926): %v\n", math.Ceil(3.1415926))
	// floor
	fmt.Printf("math.Floor(3.1415926): %v\n", math.Floor(3.1415926))
	// trunc
	fmt.Printf("math.Trunc(3.1415926): %v\n", math.Trunc(3.1415926))
}

/*
2.5 Abs 函数

func Abs(x float64) float64
返回x的绝对值。

2.6 Max函数

func Max(x, y float64) float64
返回x，y中的最大值。

2.7 Min函数

func Min(x, y float64) float64
返回想x，y中的最小值。

2.8 Dim函数

func Dim(x, y float64) float64
返回x-y和0中的最大值，max(x-y, 0)
*/

func TestUpperBound(t *testing.T) {
	// abs
	fmt.Printf("math.Abs(-3.1415926): %v\n", math.Abs(-3.1415926))
	// max
	fmt.Printf("math.Max(3.14, 3.1415926): %v\n", math.Max(3.14, 3.1415926))
	// min
	fmt.Printf("math.Min(3.14, 3.1415926): %v\n", math.Min(3.14, 3.1415926))
	// dim
	fmt.Printf("math.Dim(3.14, 3.15): %v\n", math.Dim(3.14, 3.15))
	fmt.Printf("math.Dim(3.15, 3.14): %v\n", math.Dim(3.15, 3.14))

}

/*
2.9 Mod函数

func Mod(x, y float64) float64
取余运算，可以理解为x-Trunc(x/y)*y，结果的正负号和x相同。
*/

func TestMod(t *testing.T) {
	// mod
	fmt.Printf("math.Mod(3, 0): %v\n", math.Mod(3, 0))
	fmt.Printf("math.Mod(-3, 2): %v\n", math.Mod(-3, 2))
	fmt.Printf("math.Mod(10, -3): %v\n", math.Mod(10, -3))

}

/*
2.10 Sqrt函数

func Sqrt(x float64) float64
返回x的二次方根，即平方根。

2.11 Cbrt函数

func Cbrt(x float64) float64
返回x的三次方根，即立方根。

2.12 Hypot函数

func Hypot(p, q float64) float64
返回Sqrt(p * p + q * q)，注意要避免不必要的溢出或下溢。

2.13 Pow函数

func Pow(x, y float64) float64
求幂，x的y次方。
*/

func TestRoot(t *testing.T) {
	// sqrt
	fmt.Printf("math.Sqrt(141): %v\n", math.Sqrt(141))
	fmt.Printf("math.Sqrt2: %v\n", math.Sqrt2)
	fmt.Printf("math.SqrtE: %v\n", math.SqrtE)
	fmt.Printf("math.SqrtPhi: %v\n", math.SqrtPhi)
	fmt.Printf("math.SqrtPi: %v\n", math.SqrtPi)
	// cbrt
	fmt.Printf("math.Cbrt(8): %v\n", math.Cbrt(8))
	// hypot
	fmt.Printf("math.Hypot(12, 5): %v\n", math.Hypot(12, 5))
	fmt.Printf("math.Hypot(3, 4): %v\n", math.Hypot(3, 4))
	// pow
	fmt.Printf("math.Pow(2, 3): %v\n", math.Pow(2, 3))
	fmt.Printf("math.Pow10(3): %v\n", math.Pow10(3))
}

/*
2.14 Sin函数

func Sin(x float64) float64
求正弦。

2.15 Cos函数

func Cos(x float64) float64
求余弦。

2.16 Tan函数

func Tan(x float64) float64
求正切。
*/

func TestTrigonometirc(t *testing.T) {
	// sin
	fmt.Printf("math.Sin(3): %v\n", math.Sin(3))
	// cos
	fmt.Printf("math.Cos(3): %v\n", math.Cos(3))
	// tan
	fmt.Printf("math.Tan(3): %v\n", math.Tan(3))

}

/*
2.17 Log函数

func Log(x float64) float64
求自然对数。

2.18 Log2函数

func Log2(x float64) float64
求以2为底的对数。

2.19 Log10函数

func Log10(x float64) float64
求以10为底的对数。
*/

func TestLog(t *testing.T) {
	// log ln
	fmt.Printf("math.Log(2.71828): %v\n", math.Log(2.71828))
	// log2
	fmt.Printf("math.Log2(2.1718): %v\n", math.Log2(2.1718))
	// log10
	fmt.Printf("math.Log10(9.98): %v\n", math.Log10(9.98))
	// other
	fmt.Printf("math.Log1p(0.1111): %v\n", math.Log1p(0.1111))
}

/*
2.20 Signbit函数

func Signbit(x float64) bool
如果x是一个负数，返回true。
*/

func TestSign(t *testing.T) {
	// signbit
	fmt.Printf("math.Signbit(-3.14): %v\n", math.Signbit(-3.14))
	// fmt.Printf("math.Signbit(-0): %v\n", math.Signbit(-0))

}

/*
3. 随机数 math/rand

math/rand包是go提供用来产生各种各样随机数的包，注意：rand生成的数值虽然说是随机数，但它其实是伪随机数。

rand实现的几个方法：

- func (r *Rand) Int() int 返回一个非负的伪随机int值
- func (r *Rand) Int32() int32 返回一个int32类型的31位伪随机数
- func (r *Rand) Intn(n int) int 返回一个取值范围在[0, n)的伪随机 int值，如果n<0会panic
- func (r *Rand) Int63() int64 返回一个int64类型的63位伪随机数
- func (r *Rand) Uint32() uint32 返回一个uint32类型的非负32位伪随机数
- func (r *Rand) Uint64() uint64 返回一个uint64类型的非负的64位伪随机数
- func (r *Rand) Int31n(n int32) int32 返回一个取值范围在[0, n)的伪随机int32值，如果n<0会panic
- func (r *Rand) Int63n(n int64) int64 返回一个取值范围在[0, n)的伪随机int64值，如果n<0会panic
- func (r *Rand) Float32() float32 返回一个取值范围在[0.0, 1.0)的伪随机float32值
- func (r *Rand) Float64() float64 返回一个取值范围在[0.0, 1.0)的伪随机float64值
- func (r *Rand) Perm(n int) []int 返回一个有n个元素的，[0,n)范围内整数的伪随机切片
- func (r *Rand) Read(p []byte) (n int, err error) 生成len个伪随机数，伪随机数的范围是0-255，并将伪随机数存入p，返回len和可能发生的错误
- func (r *Rand) NewSource(seed int64) Source 使用给定的种子创建一个伪随机源
- func (r *Rand) New(src Source) *Rand 返回一个使用src随机源生成的一个Rand

*/

func TestRand(t *testing.T) {
	// 不设置种子时，直接使用rand生成随机数，每次运行都是一样的
	// 高版本的go好像不设随机种子也能生成随机的数
	fmt.Printf("rand.Int(): %v\n", rand.Int())
	fmt.Printf("rand.Int31(): %v\n", rand.Int31())
	fmt.Printf("rand.Intn(5): %v\n", rand.Intn(5))
	// // 设置随机种子后
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("rand.Int(): %v\n", rand.Int())
	fmt.Printf("rand.Int31(): %v\n", rand.Int31())
	fmt.Printf("rand.Intn(5): %v\n", rand.Intn(5))

	// 使用当前的纳秒生成一个随机源，即随机种子
	// NewSource()方法等同于前面的rand.Seed()方法，都是用来设置随机种子
	// 这两种方法本质上没有区别
	source := rand.NewSource(time.Now().UnixNano())
	// 生成一个rand
	rander := rand.New(source)
	// 生成随机数
	fmt.Printf("rander.Int(): %v\n", rander.Int())
}
