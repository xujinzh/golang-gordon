package main

import "fmt"

type HeroNode struct {
	no       int
	name     string
	nickname string
	next     *HeroNode // 指向下一个结点
}

// 单链表的增删改查
// 给链表添加结点
func InsertHeroNode(head *HeroNode, newHeroNode *HeroNode) {
	// 先找到链表最后结点
	// 用辅助结点（跑龙套）
	tmp := head
	for {
		if tmp.next == nil {
			break
		}
		tmp = tmp.next
	}
	// 将 newHeroNode 加入到链表末尾
	tmp.next = newHeroNode
}

// 根据编号给链表添加结点，从小到大
func InsertHeroNodeByNo(head *HeroNode, newHeroNode *HeroNode) {
	// 先找到链表最后结点
	// 用辅助结点（跑龙套）
	tmp := head
	flag := true
	for {
		if tmp.next == nil {
			break
		} else if tmp.next.no > newHeroNode.no {
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
		tmp.next = newHeroNode
	}
}

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
	} else {
		fmt.Printf("id = %d 不存在", id)
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
	// delete
	DelHeroNode(head, 2)
	// show
	fmt.Println()
	ListHeroNode(head)

}
