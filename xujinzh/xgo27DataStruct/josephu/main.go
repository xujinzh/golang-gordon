package main

import "fmt"

type Boy struct {
	No   int
	Next *Boy
}

// 编写一个函数，构成单向的环形链表
// num 表示环形链表中小孩的个数；*Boy 表示环形链表的头指针
func AddBoy(num int) *Boy {
	first := &Boy{} // 空节点
	cur := &Boy{}   // 辅助指针
	if num < 1 {
		fmt.Println("num 值小于1")
		return first
	}
	// 循环构建这个链表
	for i := 1; i <= num; i++ {
		boy := &Boy{
			No: i,
		}
		if i == 1 {
			first = boy
			cur = boy
			cur.Next = first
		} else {
			cur.Next = boy
			cur = boy
			boy.Next = first
		}
	}
	return first
}

func PlayGame(firstBoy *Boy, startNo int, countNum int) {
	// 空链表
	if firstBoy.Next == nil {
		fmt.Println("null")
		return
	}
	// startNo 必须小于小孩的总数
	countBoy := 1
	curBoy := firstBoy
	for {
		if curBoy.Next == firstBoy {
			break
		}
		countBoy++
		curBoy = curBoy.Next
	}
	if startNo > countBoy {
		fmt.Println("起始数大于小孩的总数")
		return
	}
	// 创建辅助结点，一定到最后一个位置
	tailBoy := firstBoy
	for {
		if tailBoy.Next == firstBoy {
			break
		}
		tailBoy = tailBoy.Next
	}
	// 先入 firstBoy 移动到 startNo 节点
	for i := 1; i <= startNo-1; i++ {
		firstBoy = firstBoy.Next
		tailBoy = tailBoy.Next
	}
	// 开始数 countNum 下删除 firstBoy 指向的小孩
	for {
		for i := 1; i <= countNum-1; i++ {
			firstBoy = firstBoy.Next
			tailBoy = tailBoy.Next
		}
		fmt.Printf("boy %d out ->\n", firstBoy.No)
		// 删除 firstBoy 指向的小孩
		firstBoy = firstBoy.Next
		tailBoy.Next = firstBoy
		if firstBoy == tailBoy {
			break
		}
	}
	fmt.Printf("boy %d out\n", firstBoy.No)

}

// 显示单向环形链表
func Show(firstBoy *Boy) {
	// 链表空
	if firstBoy.Next == nil {
		fmt.Println("null")
		return
	}

	// 说明至少有一个小孩
	// 创建一个指针帮助遍历
	curBoy := firstBoy
	for {
		fmt.Printf("boy no = %d -> ", curBoy.No)
		if curBoy.Next == firstBoy {
			break
		}
		curBoy = curBoy.Next
	}

}
func main() {
	firstBoy := AddBoy(16)
	fmt.Println()
	Show(firstBoy)
	fmt.Println("\n\nBegin Game:")
	PlayGame(firstBoy, 3, 2)
}
