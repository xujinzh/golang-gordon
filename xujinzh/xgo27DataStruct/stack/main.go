package main

import (
	"errors"
	"fmt"
)

// 使用数组来模拟一个栈
type Stack struct {
	MaxTop int    // 表示栈中最大存放的元素个数
	Top    int    // 栈顶
	arr    [5]int // 用数组模拟栈
}

// 向栈中增加元素
func (stack *Stack) Push(val int) (err error) {
	// 先判断栈是否满了
	if stack.Top == stack.MaxTop-1 {
		fmt.Println("栈已满")
		return errors.New("栈满")
	}
	stack.Top++
	// 放入数据
	stack.arr[stack.Top] = val
	return
}

// 出栈
func (stack *Stack) Pop() (val int, err error) {
	// 判断栈是否为空
	if stack.Top == -1 {
		fmt.Println("栈为空")
		return 0, errors.New("栈为空")
	}

	// 先取值，再 Top--
	val = stack.arr[stack.Top]
	stack.Top--
	return val, nil
}

// 遍历栈，需要从栈顶开始遍历
func (stack *Stack) List() {
	// 先判断栈是否为空
	if stack.Top == -1 {
		fmt.Println("栈为空")
		return
	}
	fmt.Println("栈的情况如下")
	for i := stack.Top; i >= 0; i-- {
		fmt.Printf("arr[%d]=%d\n", i, stack.arr[i])
	}
}

func main() {
	stack := &Stack{
		MaxTop: 5,
		Top:    -1, // 栈顶为-1时，表示栈为空
	}
	fmt.Println(*stack)
	// 入栈
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)
	stack.Push(6)
	// 打印
	stack.List()
	// 出栈
	value, err := stack.Pop()
	if err != nil {
		fmt.Println("出栈错误")
	} else {
		fmt.Println("出栈值：", value)
	}
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	stack.Pop()
	// 打印
	stack.List()
}
