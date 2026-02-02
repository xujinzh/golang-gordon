package timestl_test

import (
	"fmt"
	"testing"
	"time"
)

// time 标准库中日历的计算采用的是公历

// time 代表一个纳秒精度的时间点
// 程序中应该使用 Time 类型值来保存和传递时间，而不能用指针。即，表示时间的变量或字段应该为 time.Time 类型，而不是 *time.Time 类型
// 一个 Time 类型值可以被多个 goroutine 同时使用
// 时间点可以使用 Before, After, Equal 方法进行比较
// Sub 方法让两个时间点相减，生成一个 Duration 类型值(代表时间段)
// Add 方法给一个时间点加上一个时间段，生成一个新的 Time 类型的时间点
// Time 零值代表时间点 January 1, year 1, 00:00:00.000000000 UTC, 因为本时间点一般不会在使用中出现， IsZero 方法提供了检验时间是否显示初始化的一个简单路径
// 每一个时间都具有一个地点信息（即对应地点的时区信息），当计算时间的表示格式时，如 Format, Hour, Year 等方法，都会考虑该信息
// Local, UTC, In 方法返回一个指定时区（但指向同一个时间点）的 Time。修改地点或时区信息只是会改变其表示，不会修改被表示的时间点，因此不会影响其计算。

func TestTime(t *testing.T) {
	now := time.Now()
	fmt.Printf("now: %v\n", now)
	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	// 02d 输出的整数不足两位 用0补足，如把 1-1 补为 01-01
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
}

// 时间戳
// 时间戳是自1970年1月1日0时0分0秒至当前日期的总毫秒数。它也被称为 Unix 时间戳(UnixTimestamp)
// 这里指的UTC时间，比北京时间晚8个小时
