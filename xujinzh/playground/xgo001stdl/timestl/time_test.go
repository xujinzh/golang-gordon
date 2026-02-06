package timestl_test

import (
	"fmt"
	"sync"
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
func TestTimeStamp(t *testing.T) {
	// 获取当前时间
	now := time.Now().UTC()
	fmt.Printf("now: %v\n", now)
	// 当前时间的时间戳
	timestamp1 := now.Unix()
	// 当前时间的时间戳，单位纳秒
	timestamp2 := now.UnixNano()
	// 打印
	fmt.Printf("timestamp1: %v\n", timestamp1)
	fmt.Printf("timestamp2: %v\n", timestamp2)
	// 转成北京时
	nowBJ := now.Local()
	fmt.Printf("nowBJ: %v\n", nowBJ)
	fmt.Printf("nowBJ.Unix(): %v\n", nowBJ.Unix())
	fmt.Printf("now.Local().Unix(): %v\n", now.Local().Unix())
	// 可见：时间戳跟时区没有关系
	// 从毫秒时间戳转为时间
	fmt.Printf("time.UnixMilli(now.UnixMilli()): %v\n", time.UnixMilli(now.UnixMilli()))
}

// 把时间字符串转为时间
// Parse解析时间
// func Parse(layout, value string) (Time, error)
// 解析一个格式化的时间字符串并返回它代表的时间，如果缺少表示时区的信息，Parse会将时区设置为UTC
// func ParseInLocation(layout, value string, loc *Location) (Time, error)
// ParseInLocation类似Parse但有两个重要的不同之处：
// 第一，当缺少时区信息时，Parse将时间解释为UTC时间，而ParseInLocation将返回值的Location设置为loc
// 第二，当时间字符串提供了时区偏移量信息时，Parse会尝试去匹配本地时区，而ParseInLocation会去匹配loc
// layout的时间必须是"2006-01-02 15:04:05"这个时间。当然格式不一定是这个，但是时间值一定是，这是go诞生的时间
func TestParseTime(t *testing.T) {
	// Parse 解析
	tt, err := time.Parse("2006-01-02 15:04:05", "2026-02-06 09:51:13")
	if err != nil {
		t.Log(err)
	}
	fmt.Printf("tt: %v\n", tt)
	// 加载时区信息
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		t.Log(err)
	}
	// 使用时区解析
	tt, err = time.ParseInLocation("2006-01-02 15:04:05", "2026-02-06 09:51:13", loc)
	if err != nil {
		t.Log(err)
	}
	fmt.Printf("tt: %v\n", tt)
}

// 格式化时间显示. 把 time 转为字符串
func TestFormatTime(t *testing.T) {
	// // 获取当前时间
	// now := time.Now()
	// 加载时区信息
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		t.Log(err)
	}
	// 使用时区解析
	now, err := time.ParseInLocation("2006-01-02 15:04:05", "2026-02-06 19:51:13", loc)
	if err != nil {
		t.Log(err)
	}
	fmt.Printf("now: %v\n", now)
	// 格式化时间
	tt := now.Format("2006/01/02 15.04.05.000")
	// 打印
	fmt.Printf("tt: %v\n", tt)
	// 12小时制
	tt1 := now.Format("2006/01/02 03:04:05")
	fmt.Printf("tt1: %v\n", tt1)
	tt2 := now.Format("2006/01/02 03:04:05 pm") // 建议12小时制的时候带上pm，否则不知道是上午还是下午。pm是表示，自动会根据时间解析成上午am或下午pm
	fmt.Printf("tt2: %v\n", tt2)
}

// 时间间隔
// time.Duration是time包定义的一个类型，它代表两个时间点之间经过的时间，以纳秒为单位
// time.Duration表示一段时间间隔，可表示的最长时间间隔大约是290年
// time包中定义的时间间隔类型常量如下：
/*
const (
	Nanosecond Duration = 1
	Microsecond         = 1000 * Nanosecond
	Millisecond         = 1000 * Microsecond
	Second              = 1000 * Millisecond
	Minute              = 60 * Second
	Hour                = 60 * Minute
)
*/
// time.Duration表示1纳秒，time.Second表示1秒
func TestTimeDuration(t *testing.T) {
	// 表示1纳秒
	nano := time.Nanosecond
	fmt.Printf("nano: %v\n", nano)
	// 表示5秒
	sec := 5 * time.Second
	fmt.Printf("sec: %v\n", sec)
	// 在当前时间上增加5分钟
	// 先获取当前时间
	now := time.Now()
	fmt.Printf("now: %v\n", now)
	// 增加5分钟后的时间
	nowAddFiveMinute := now.Add(5 * time.Minute)
	fmt.Printf("nowAddFiveMinute: %v\n", nowAddFiveMinute)
	// 在当前时间上减少5分钟
	nowMinusFiveMinute := now.Add(-5 * time.Minute)
	fmt.Printf("nowMinusFiveMinute: %v\n", nowMinusFiveMinute)
	// 两个时间相减
	diff := nowAddFiveMinute.Sub(nowMinusFiveMinute)
	fmt.Printf("diff: %v\n", diff)
}

// 时间比较
// Equal: func (t Time) Equal(u Time) bool
// 判断两个时间是否相同，会考虑时区的影响，因此不同时区标准的时间也可以正确比较
// 与用 t == u 不同，这种方法还会比较地点和时区信息
// Befor: func (t Time) Before(u Time) bool
// 如果 t 代表的时间点在 u 之前，返回真，否则返回假
// After: func (t Time) After(u Time) bool
// 如果 t 代表的时间点在 u 之后，返回真，否则返回假
func TestCompare(t *testing.T) {
	// 获取当前时间
	now := time.Now()
	now1 := time.Now()
	fmt.Printf("now: %v\n", now)
	fmt.Printf("now1: %v\n", now1)
	// 比较是否相等
	fmt.Printf("now.Equal(now1): %v\n", now.Equal(now1))
	// 比较是否在前
	fmt.Printf("now.After(now1): %v\n", now.After(now1))
	// 比较是否在后
	fmt.Printf("now.Before(now1): %v\n", now.Before(now1))
}

// 定时器
// 使用time.Tick(time.Duration)来设置定时器，定时器的本质是一个通道（channel）
func TestTick(t *testing.T) {
	// 创建一个定时器，每隔一定时间发送一个数据
	ticker := time.Tick(time.Second)
	// 打印发送的数据
	for i := range ticker {
		fmt.Printf("i: %v\n", i) // 每秒执行一次该任务
		break
	}
	// 5秒钟后执行一个方法
	time.AfterFunc(5*time.Second, func() {
		fmt.Printf("\"5秒之后执行\": %v\n", "5秒之后执行")
	})
	fmt.Printf("hello, 我先执行\n")
	time.Sleep(6 * time.Second)
}

// ticker 只要定义完成，从此刻开始计时，不需要任何其他的操作，每隔固定时间都会出发
// timer 定时器，是到固定时间后会执行一次
// 如果timer定时器要每隔间隔的时间执行，实现ticker的效果，使用 func (t *Timer) Reset(d Duration) bool

func TestTimer(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(2)
	// NewTimer 创建一个 Timer，它会在最少过去时间段 d 后到期，向其自身的 C 字段发送当时的时间
	timer1 := time.NewTimer(2 * time.Second)
	go func(t *time.Timer) {
		defer wg.Done()
		for {
			<-t.C
			fmt.Println("get timer", time.Now().Format("2006-01-02 15:04:05"))
			// reset 使 t 重新开始计时，本方法返回后再等待时间段 d 过去后到期，如果调用时 t 还在等待中会返回真
			// 如果 t 已经到期或者被停止了会返回假
			t.Reset(2 * time.Second)
		}
	}(timer1)
	// NewTicker 返回一个新的 Ticker，该 Ticker 包含一个通道字段，并会每隔时间段 d 向该通道发送当时的时间
	// 它会调整时间间隔或者丢弃tick信息以适应反应慢的接收者，如果d <= 0会触发panic，关闭该 Ticker 可以释放相关资源
	ticker1 := time.NewTicker(2 * time.Second)

	go func(t *time.Ticker) {
		defer wg.Done()
		for {
			<-t.C
			fmt.Println("get ticker1", time.Now().Format("2006-01-02 15:04:05"))
		}
	}(ticker1)

	wg.Wait()
}
