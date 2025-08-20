package main

import "fmt"

type Hero struct {
	No    int
	Name  string
	Left  *Hero
	Right *Hero
}

// 前序遍历（先root结点，再左子树，最后右子树）
func PreOrder(node *Hero) {
	if node != nil {
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
		PreOrder(node.Left)
		PreOrder(node.Right)
	}
}

// 中序遍历（先左子树，再root结点，最后右子树）
func InfixOrder(node *Hero) {
	if node != nil {
		InfixOrder(node.Left)
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
		InfixOrder(node.Right)
	}
}

// 后序遍历（先左子树，再右子树， 最后root结点）
func PostOrder(node *Hero) {
	if node != nil {
		PostOrder(node.Left)
		PostOrder(node.Right)
		fmt.Printf("no=%d name=%s\n", node.No, node.Name)
	}
}

func main() {
	root := &Hero{
		No:   1,
		Name: "lvluo",
	}

	left := &Hero{
		No:   2,
		Name: "duorou",
	}

	right := &Hero{
		No:   3,
		Name: "mudan",
	}
	root.Left = left
	root.Right = right

	left1 := &Hero{
		No:   4,
		Name: "cimei",
	}
	left.Right = left1

	right1 := &Hero{
		No:   5,
		Name: "qiandai",
	}
	right.Right = right1

	// 前序遍历
	fmt.Println("前序遍历：")
	PreOrder(root)

	// 中序遍历
	fmt.Println("中序遍历：")
	InfixOrder(root)

	// 后序遍历
	fmt.Println("后序遍历：")
	PostOrder(root)
}
