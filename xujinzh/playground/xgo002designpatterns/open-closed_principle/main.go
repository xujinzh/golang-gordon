/*
开闭原则 Open-Closed Principle, OCP

类的改动是通过增加代码进行的，而不是修改源代码
*/
package main

import "fmt"

type AbstractBacker interface {
	DoBusi()
}

// 存款业务员
type SaveBanker struct {
}

func (sb *SaveBanker) DoBusi() {
	fmt.Println("存款业务")
}

// 转账业务员
type TransferBanker struct {
}

func (tb *TransferBanker) DoBusi() {
	fmt.Println("转账业务")
}

// +++++++++++++++++++++++++++
// 股票业务员
type SharesBanker struct {
}

func (sb *SharesBanker) DoBusi() {
	fmt.Println("股票业务")
}

// 实现一个架构层（基于抽象层进行业务封装-针对interface接口进行封装）
func BankBusiness(banker AbstractBacker) {
	// 通过接口向下来调用（多态现象）
	banker.DoBusi()
}

func main() {
	// 存款业务
	BankBusiness(&SaveBanker{})
	// 转账业务
	BankBusiness(&TransferBanker{})
	// 股票业务
	BankBusiness(&SharesBanker{})
}
