package main

import (
	"errors"
	"fmt"
	"os"
)

type CircleQueue struct {
	maxSize int
	array   [5]int
	head    int
	tail    int
}

// 添加元素到队列
func (circlequeue *CircleQueue) Push(val int) (err error) {
	// 先判断是否已满
	if circlequeue.IsFull() {
		return errors.New("circle queue full")
	}
	circlequeue.array[circlequeue.tail] = val
	circlequeue.tail = (circlequeue.tail + 1) % circlequeue.maxSize
	return
}

// 出队列
func (circlequeue *CircleQueue) Pop() (val int, err error) {
	if circlequeue.IsEmpty() {
		return 0, errors.New("circle queue empty")
	}
	// 非空，取值
	// head 是队首，含元素
	val = circlequeue.array[circlequeue.head]
	circlequeue.head = (circlequeue.head + 1) % circlequeue.maxSize
	return val, nil
}

// 显示队列
func (circlequeue *CircleQueue) Show() {
	size := circlequeue.Size()
	if size == 0 {
		fmt.Println("circle queue is empty")
	}

	// 遍历队列
	tmp := circlequeue.head
	for i := 0; i < size; i++ {
		fmt.Printf("array[%d] = %d \n", tmp, circlequeue.array[tmp])
		tmp = (tmp + 1) % circlequeue.maxSize
	}
}

// 判断队列为空
func (circlequeue *CircleQueue) IsEmpty() bool {
	return circlequeue.tail == circlequeue.head
}

// 判断队列满
func (circlequeue *CircleQueue) IsFull() bool {
	return (circlequeue.tail+1)%circlequeue.maxSize == circlequeue.head
}

// 获取队列元素个数
func (circlequeue *CircleQueue) Size() int {
	return (circlequeue.tail + circlequeue.maxSize - circlequeue.head) % circlequeue.maxSize
}

func main() {
	// 先创建一个队列
	queue := &CircleQueue{
		maxSize: 5,
		head:    0,
		tail:    0,
	}

	// 添加数据
	var key string
	var val int
	for {
		fmt.Println("1. 输入 add 添加数据到队列")
		fmt.Println("2. 输入 get 获取队列中数据")
		fmt.Println("3. 输入 show 显示队列的数据")
		fmt.Println("4. 输入 exit 退出程序")

		fmt.Scanln(&key)
		switch key {
		case "add":
			fmt.Println("输入要添加进队列的数")
			fmt.Scanln(&val)
			err := queue.Push(val)
			if err != nil {
				fmt.Println("添加数据到队列失败")
			} else {
				fmt.Println("添加数据到队列成功")
			}
		case "get":
			fmt.Println("get")
			v, err := queue.Pop()
			if err != nil {
				fmt.Println("get val error")
			} else {
				fmt.Println("get value =", v)
			}
		case "show":
			queue.Show()
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入有误")
		}
	}
}
