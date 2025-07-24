package main

import (
	"fmt"
)

func main() {
	fmt.Println()

	// 声明一个变量，保存用户输入的选项
	key := ""

	// 声明一个变量，用于控制退出主循环
	loop := true

	// 定义账户余额
	balance := 10000.0
	// 每次收支的余额
	money := 0.0
	// 每次收支的说明
	note := ""
	// 定义变量，记录是否有收支行为
	flag := false
	// 收支的详情
	details := "收支\t账户余额\t收支余额\t说明"
	details += fmt.Sprintf("\n账户\t%v\t\t%v\t\t%v", balance, 0.0, "家庭资产")

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
		fmt.Scanf("%s", &key)

		switch key {
		case "1":
			fmt.Println("-----------------当前收支明细记录-----------------")
			if flag {
				fmt.Println(details)
			} else {
				fmt.Println("当前没有收支明细...来一笔吧！")
			}
		case "2":
			fmt.Print("本次收入金额: ")
			fmt.Scanln(&money)
			fmt.Print("本次收入说明: ")
			fmt.Scanln(&note)
			balance += money // 修改账户余额
			details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%v", balance, money, note)
			flag = true
		case "3":
			fmt.Print("本次支出金额: ")
			fmt.Scanln(&money)
			fmt.Print("本次支出说明: ")
			fmt.Scanln(&note)
			if money > balance {
				fmt.Println("余额不足")
				break
			}
			balance -= money // 修改账户余额
			details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%v", balance, money, note)
			flag = true
		case "4":
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
				loop = false
			}
		default:
			fmt.Println("请输入正确的选项...")
		}
		if !loop {
			break
		}
	}
	fmt.Println("你退出了家庭记账软件")
}
