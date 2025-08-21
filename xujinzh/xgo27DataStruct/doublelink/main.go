package main

import "fmt"

type HeroNode struct {
	no       int
	name     string
	nickname string
	pre      *HeroNode // 指向前一个节点
	next     *HeroNode // 指向后一个节点
}

// 插入节点
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	tmp := head
	for {
		if tmp.next == nil {
			break
		}
		tmp = tmp.next
	}

	tmp.next = newHeroNode
	newHeroNode.pre = tmp
}

// 根据编号给链表添加结点，从小到大
func InsertHeroNodeByNo(head *HeroNode, newHeroNode *HeroNode) {
	// 先找到链表最后结点
	// 用辅助结点（跑龙套）
	tmp := head
	flag := true
	for {
		if tmp.next == nil {
			// 链表最后
			break
		} else if tmp.next.no > newHeroNode.no {
			// newHeroNode 插入tmp后面
			break
		} else if tmp.next.no == newHeroNode.no {

			flag = false
			break
		}
		tmp = tmp.next
	}

	// 将 newHeroNode 添加到 tmp 后面
	if !flag {
		fmt.Println("链表中已经有改ID", newHeroNode.no)
	} else {
		newHeroNode.next = tmp.next
		newHeroNode.pre = tmp
		if tmp.next != nil {
			tmp.next.pre = newHeroNode
		}
		tmp.next = newHeroNode
	}
}

// 显示链表信息
func ListHeroNode(head *HeroNode) {
	// 用辅助结点
	tmp := head
	// 先判断该链表是不是空链表
	if tmp.next == nil {
		fmt.Println("link is null")
		return
	}
	// 遍历链表
	for {
		fmt.Printf("[%d, %s, %s]->", tmp.next.no, tmp.next.name, tmp.next.nickname)
		tmp = tmp.next
		if tmp.next == nil {
			break

		}
	}
}

// 删除一个节点
func DelHeroNode(head *HeroNode, id int) {
	tmp := head
	flag := false

	for {
		if tmp.next == nil {
			break
		} else if tmp.next.no == id {
			// found it
			flag = true
			break
		}
		tmp = tmp.next
	}
	// delete
	if flag {
		tmp.next = tmp.next.next
		if tmp.next != nil {
			tmp.next.pre = tmp
		}
	} else {
		fmt.Printf("id = %d 不存在", id)
	}
}

// 显示链表信息
func ListHeroNodeConverse(head *HeroNode) {
	// 用辅助结点
	tmp := head
	// 先判断该链表是不是空链表
	if tmp.next == nil {
		fmt.Println("link is null")
		return
	}

	// 先定位到最后一个节点
	for {
		if tmp.next == nil {
			break
		}
		tmp = tmp.next
	}
	// 遍历链表
	for {
		fmt.Printf("[%d, %s, %s]->", tmp.no, tmp.name, tmp.nickname)
		tmp = tmp.pre
		if tmp.pre == nil {
			break

		}
	}
}

func main() {
	// 1. 先创建一个头结点
	head := &HeroNode{}
	// 2. 创建一个新的 HeroNode
	hero1 := &HeroNode{
		no:       1,
		name:     "宋江",
		nickname: "及时雨",
	}
	hero2 := &HeroNode{
		no:       2,
		name:     "卢俊义",
		nickname: "玉麒麟",
	}
	hero3 := &HeroNode{
		no:       3,
		name:     "林冲",
		nickname: "豹子头",
	}
	// add
	// InsertHeroNode(head, hero1)
	// InsertHeroNode(head, hero2)
	// InsertHeroNode(head, hero3)
	InsertHeroNodeByNo(head, hero3)
	InsertHeroNodeByNo(head, hero2)
	InsertHeroNodeByNo(head, hero1)
	// show
	ListHeroNode(head)
	fmt.Println()
	ListHeroNodeConverse(head)
	// delete
	DelHeroNode(head, 2)
	// show
	fmt.Println()
	ListHeroNodeConverse(head)
}
