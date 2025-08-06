// progress_simple.go
package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

type SimpleProgress struct {
	total    int64 // 总任务量（如字节数、条目数）
	progress int64 // 已完成的量
	width    int   // 进度条宽度（字符数），常用 30~50
}

// NewSimpleProgress 创建一个进度条实例
func NewSimpleProgress(total int64, width int) *SimpleProgress {
	if width <= 0 {
		width = 30
	}
	return &SimpleProgress{
		total: total,
		width: width,
	}
}

// Add 增加已完成的数量（线程安全）
func (p *SimpleProgress) Add(delta int64) {
	atomic.AddInt64(&p.progress, delta)
}

// Render 将当前进度渲染到终端（不换行）
func (p *SimpleProgress) Render() {
	cur := atomic.LoadInt64(&p.progress)
	if cur > p.total {
		cur = p.total
	}
	// 计算已完成的比例
	ratio := float64(cur) / float64(p.total)
	filled := int(ratio * float64(p.width))

	// 构造条形图
	bar := fmt.Sprintf("\r[%s%s] %6.2f%% (%d/%d)",
		// 已完成的部分使用 =，未完成使用空格
		stringRepeat("=", filled),
		stringRepeat(" ", p.width-filled),
		ratio*100,
		cur, p.total,
	)
	fmt.Print(bar)

	if cur == p.total {
		fmt.Print("\n") // 完成后换行
	}
}

// 辅助函数：重复字符串 n 次
func stringRepeat(s string, n int) string {
	if n <= 0 {
		return ""
	}
	b := make([]byte, len(s)*n)
	pos := 0
	for i := 0; i < n; i++ {
		copy(b[pos:], s)
		pos += len(s)
	}
	return string(b)
}

/* ------------------- Demo ------------------- */
func main() {
	const total = 1000 // 假设我们要处理 1000 条记录
	p := NewSimpleProgress(total, 40)

	// 模拟并发生产者/消费者
	go func() {
		for i := 0; i < total; i++ {
			time.Sleep(5 * time.Millisecond) // 业务耗时
			p.Add(1)                          // 完成 1 条
			p.Render()
		}
	}()

	// 主协程阻塞，等进度条走完
	for {
		if atomic.LoadInt64(&p.progress) >= total {
			break
		}
		time.Sleep(50 * time.Millisecond)
	}
}