package main

import "fmt"

type CatNode struct {
	no   int
	name string
	next *CatNode
}

func InsertCatNode(head *CatNode, newCatNode *CatNode) {
	// 先判断是不是第一只猫
	if head.next == nil {
		head.no = newCatNode.no
		head.name = newCatNode.name
		head.next = head
		fmt.Println(newCatNode, "加入到单节点环形链表中")
		return

	}
	// 定义辅助变量，帮助获取到环形链表的最后节点
	tmp := head
	for {
		if tmp.next == head {
			break
		}
		tmp = tmp.next
	}
	// 加入链表
	tmp.next = newCatNode
	newCatNode.next = head
}

// 删除一个节点
func DelCatNode(head *CatNode, id int) *CatNode {
	cur := head
	bef := head
	// 空环形链表
	if cur.next == nil {
		fmt.Println("circle link is null, can not delete")
		return head
	}
	// 单节点环形链表
	if cur.next == cur {
		cur.next = nil
		return head
	}
	// 多个节点的环形链表
	// 先把 bef 移动到最后一个节点，保持在 cur（目前是指向 head）前面
	for {
		if bef.next == head {
			break
		}
		bef = bef.next
	}

	flag := true
	for {
		// 如果到最后一个节点，但是还没与其比较
		if cur.next == head {
			break
		}
		if cur.no == id {
			if cur == head { // 如果删除的事头结点
				head = head.next
			}
			// 找到要删除的节点
			bef.next = cur.next
			fmt.Printf("%d被删除\n", id)
			flag = false // 已经删除成功
			break
		}
		// 继续移动
		cur = cur.next
		bef = bef.next
	}
	if flag { // 说明还没有删除，还要比较一次 最后的节点
		if cur.no == id {
			bef.next = cur.next
			fmt.Printf("%d 被删除\n", id)
		} else {
			fmt.Println("没有该节点")
		}
	}
	return head

}

// 打印链表
func ListCircleLink(head *CatNode) {
	fmt.Println("circle link info is")
	tmp := head
	if tmp.next == nil {
		fmt.Println("circle link is null")
		return
	}
	for {
		// fmt.Println("cat info =", *tmp, "->")
		fmt.Printf("cat info = [id=%d name=%s] ->\n", tmp.no, tmp.name)
		if tmp.next == head { //如果tmp是链表最后结点
			break
		}
		tmp = tmp.next
	}
}

func main() {
	head := &CatNode{}

	cat1 := &CatNode{
		no:   1,
		name: "tom",
	}

	cat2 := &CatNode{
		no:   2,
		name: "tom",
	}
	cat3 := &CatNode{
		no:   3,
		name: "tom",
	}
	InsertCatNode(head, cat1)
	InsertCatNode(head, cat2)
	InsertCatNode(head, cat3)
	fmt.Println()
	head = DelCatNode(head, 30)
	fmt.Println()
	ListCircleLink(head)
}
