package utils

import (
	"fmt"
)

type FamilyAccount struct {
	// 声明一个字段，保存用户输入的选项
	key string

	// 声明一个字段，用于控制退出主循环
	loop bool

	// 定义账户余额
	balance float64
	// 每次收支的余额
	money float64
	// 每次收支的说明
	note string
	// 定义变量，记录是否有收支行为
	flag bool
	// 收支的详情
	details string
}

// 工厂模式
func FactoryFamilyAccount() *FamilyAccount {
	return &FamilyAccount{
		key:     "",
		loop:    true,
		balance: 10000,
		money:   0.0,
		note:    "",
		flag:    false,
		details: "收支\t账户余额\t收支余额\t说明",
	}
}

// 给该结构体绑定方法
// 显示明细
func (FamilyAccount *FamilyAccount) ShowDetails() {
	fmt.Println("-----------------当前收支明细记录-----------------")
	if FamilyAccount.flag {
		fmt.Println(FamilyAccount.details)
	} else {
		fmt.Println("当前没有收支明细...来一笔吧！")
	}
}

// 收入
func (FamilyAccount *FamilyAccount) Income() {
	fmt.Print("本次收入金额: ")
	fmt.Scanln(&FamilyAccount.money)
	fmt.Print("本次收入说明: ")
	fmt.Scanln(&FamilyAccount.note)
	FamilyAccount.balance += FamilyAccount.money // 修改账户余额
	FamilyAccount.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", FamilyAccount.balance, FamilyAccount.money, FamilyAccount.note)
	FamilyAccount.flag = true
}

// 支出
func (FamilyAccount *FamilyAccount) Pay() {
	fmt.Print("本次支出金额: ")
	fmt.Scanln(&FamilyAccount.money)
	fmt.Print("本次支出说明: ")
	fmt.Scanln(&FamilyAccount.note)
	if FamilyAccount.money > FamilyAccount.balance {
		fmt.Println("余额不足")
	}
	FamilyAccount.balance -= FamilyAccount.money // 修改账户余额
	FamilyAccount.details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", FamilyAccount.balance, FamilyAccount.money, FamilyAccount.note)
	FamilyAccount.flag = true
}

// 退出系统
func (FamilyAccount *FamilyAccount) Exit() {
	fmt.Println("你确定要退出吗？y/n")
	choice := ""
	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Println("你的输入有误，请重新输入 y/n")
	}
	if choice == "y" {
		FamilyAccount.loop = false
	}
}

// 显示主菜单
func (FamilyAccount *FamilyAccount) MainMenu() {
	// 显示主菜单
	for {
		fmt.Println()
		fmt.Println("-----------------家庭收支记账软件-----------------")
		fmt.Println("                  1 收支明细")
		fmt.Println("                  2 登记收入")
		fmt.Println("                  3 登记支出")
		fmt.Println("                  4 退出软件")
		fmt.Print("请选择(1-4): ")

		// fmt.Scanln(&key)
		fmt.Scanf("%s", &FamilyAccount.key)

		switch FamilyAccount.key {
		case "1":
			FamilyAccount.ShowDetails()
		case "2":
			FamilyAccount.Income()
		case "3":
			FamilyAccount.Pay()
		case "4":
			FamilyAccount.Exit()
		default:
			fmt.Println("请输入正确的选项...")
		}
		if !FamilyAccount.loop {
			break
		}
	}
	fmt.Println("你退出了家庭记账软件")
}
