package main

import (
	"fmt"
	"os"
)

var Num int = 7

// 定义emp
type Emp struct {
	Id   int
	Name string
	Next *Emp
}

// 显示雇员信息
func (emp *Emp) Show(hashtable HashTable) {
	fmt.Printf("链表%d 找到雇员 %d \n", hashtable.HashFun(emp.Id), emp.Id)
}

// 定义 emplink
// 不带表头，第一个节点就存放雇员
type EmpLink struct {
	Head *Emp
}

// 添加员工的方法
// 单链表，从小到大
func (emplink *EmpLink) Insert(emp *Emp) {
	// 辅助指针
	cur := emplink.Head
	var pre *Emp = nil // 始终在 cur 前面
	// 如果当前的 EmpLink 是空链表
	if cur == nil {
		emplink.Head = emp
		return
	}
	// 如果不是空链表，给 emp 找到对应的位置并插入
	// 思路：让 cur 和 emp 比较，并让 pre 保持在 cur 的前面
	for {
		if cur != nil {
			// 比较
			if cur.Id >= emp.Id { // 找到位置
				break
			}
			// 保证 pre 在 cur 前面
			pre = cur
			cur = cur.Next
		} else { // 链表末尾 cur == nil
			break
		}
	}
	// 退出时，我们看下是否将 emp 添加到链表最后
	if pre == nil {
		emplink.Head = emp
	} else {
		pre.Next = emp
	}
	emp.Next = cur

}

// 显示当前链表的信息
func (emplink *EmpLink) ShowLink(no int) {
	if emplink.Head == nil {
		fmt.Printf("当前链表 %d 为空\n", no)
		return
	}
	// 遍历当前链表显示数据
	cur := emplink.Head // 辅助指针
	for {
		if cur != nil {
			fmt.Printf("链表 %d，雇员id=%d，名字=%s ->", no, cur.Id, cur.Name)
			cur = cur.Next
		} else {
			break
		}
	}
	fmt.Println()
}

// 根据ID查找雇员
func (emplink *EmpLink) FindById(id int) *Emp {
	cur := emplink.Head // 辅助指针
	for {
		if cur != nil && cur.Id == id {
			return cur
		} else if cur == nil {
			break
		}
		cur = cur.Next
	}
	return nil
}

// 定义hashtable
// 含有一个链表数组
type HashTable struct {
	LinkArr [7]EmpLink
}

// 给 hashtable 编写 insert 雇员的方法
func (hashtable *HashTable) Insert(emp *Emp) {
	// 使用散列函数，确定将该雇员添加到哪个链表
	linkNo := hashtable.HashFun(emp.Id)
	// 使用对应的链表添加
	hashtable.LinkArr[linkNo].Insert(emp)
}

// 显示 hashtable 所有雇员
func (hashtable *HashTable) ShowAll() {
	for i := 0; i < len(hashtable.LinkArr); i++ {
		hashtable.LinkArr[i].ShowLink(i)
	}
}

// 查找
func (hashtable *HashTable) FindById(id int) *Emp {
	// 先确定在哪个链表
	linkNo := hashtable.HashFun(id)
	return hashtable.LinkArr[linkNo].FindById(id)
}

// 编写一个散列方法
func (hashtable *HashTable) HashFun(id int) int {
	return id % 7
}

func main() {
	key := ""
	id := 0
	name := ""
	var hashtable HashTable
	for {
		fmt.Println("===========雇员系统菜单===========")
		fmt.Println("input 表示添加雇员")
		fmt.Println("show 表示添加雇员")
		fmt.Println("find 表示添加雇员")
		fmt.Println("exit 表示添加雇员")
		fmt.Println("请输入你的选择")
		fmt.Scanln(&key)
		switch key {
		case "input":
			fmt.Println("输入雇员id")
			fmt.Scanln(&id)
			fmt.Println("输入雇员name")
			fmt.Scanln(&name)
			emp := &Emp{
				Id:   id,
				Name: name,
			}
			hashtable.Insert(emp)
		case "show":
			hashtable.ShowAll()
		case "find":
			fmt.Println("请输入id")
			fmt.Scanln(&id)
			emp := hashtable.FindById(id)
			if emp == nil {
				fmt.Printf("id=%d 的雇员不存在\n", id)
			} else {
				emp.Show(hashtable)
			}
		case "exit":
			os.Exit(0)
		default:
			fmt.Println("输入错误")
		}
	}
}
