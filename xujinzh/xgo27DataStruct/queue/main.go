package main

import (
	"errors"
	"fmt"
	"os"
)

type Queue struct {
	maxSize int
	array   [5]int // 数组
	front   int    // 队列首，不含
	rear    int    // 队列尾，含
}

// 添加数据到队列
func (queue *Queue) AddQueue(val int) (err error) {
	// 先判断队列是否已满
	if queue.rear == queue.maxSize-1 {
		return errors.New("queue full")
	}

	queue.rear++
	queue.array[queue.rear] = val
	return
}

// 显示队列
func (queue *Queue) ShowQueue() {
	fmt.Println("当前队列情况：")
	// 找到队首，然后遍历到队尾
	for i := queue.front + 1; i <= queue.rear; i++ { // queue.front 不含队首的元素
		fmt.Printf("array[%d]=%d\n", i, queue.array[i])
	}
}

// 从队列中取数据
func (queue *Queue) GetQueue() (val int, err error) {
	// 先判断队列是否为空
	if queue.rear == queue.front {
		err = errors.New("queue empty")
		return 0, err
	}
	queue.front++
	val = queue.array[queue.front]
	err = nil
	return val, err
}

func main() {
	// 先创建一个队列
	queue := &Queue{
		maxSize: 5,
		front:   -1,
		rear:    -1,
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
			err := queue.AddQueue(val)
			if err != nil {
				fmt.Println("添加数据到队列失败")
			} else {
				fmt.Println("添加数据到队列成功")
			}
		case "get":
			fmt.Println("get")
			v, err := queue.GetQueue()
			if err != nil {
				fmt.Println("get val error")
			} else {
				fmt.Println("get value =", v)
			}
		case "show":
			queue.ShowQueue()
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入有误")
		}
	}
}
