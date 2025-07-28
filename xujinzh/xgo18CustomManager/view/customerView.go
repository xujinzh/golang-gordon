package main

import (
	"fmt"
	"xgo18/model"
	"xgo18/service"
)

type customerView struct {
	key             string // 用户输入
	loop            bool   // 表示是否循环显示菜单
	customerService *service.CustomerService
}

// 显示所有客户信息
func (cv *customerView) list() {
	customers := cv.customerService.List()
	// 显示
	fmt.Println("-----------------------客户列表-----------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")

	for i := 0; i < len(customers); i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("-----------------------客户列表完成-------------------")
	fmt.Println()
}

func (cv *customerView) add() {
	fmt.Println("-----------------------添加客户-----------------------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("邮箱：")
	email := ""
	fmt.Scanln(&email)

	// id 号没有让用户输入，是唯一的，需要系统分配
	customer := model.FactoryCustomer2(name, gender, age, phone, email)
	if cv.customerService.Add(customer) {
		fmt.Println("-----------------------添加完成-----------------------")
	} else {
		fmt.Println("-----------------------添加失败-----------------------")
	}

	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t邮箱")

}

// 得到用户的输入，删除ID对应的客户
func (cv *customerView) delete() {
	fmt.Println("---------------------删除客户---------------------")
	fmt.Println("请选择待删除客户编号(-1退出): ")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return // 放弃删除
	}
	fmt.Println("确认是否删除(y/n): ")
	choice := ""
	fmt.Scanln(&choice)
	if choice == "y" {
		// 调用 customerService 的 Delete 方法
		if cv.customerService.Delete(id) {
			fmt.Println("---------------------删除成功---------------------")
		} else {
			fmt.Println("-----------------删除失败，ID不存在----------------")
		}
	}

}

// 退出软件
func (cv *customerView) exit() {
	fmt.Println("确认是否退出(y/n): ")
	for {
		fmt.Scanln(&cv.key)
		if cv.key == "y" || cv.key == "n" {
			break
		}
		fmt.Println("你输入有误，确认是否退出(y/n): ")
	}
	if cv.key == "y" {
		cv.loop = false
	}

}

func (cv *customerView) mainMenu() {
	for {
		fmt.Println("---------------客户信息管理软件---------------")
		fmt.Println("               1 添加客户")
		fmt.Println("               2 修改客户")
		fmt.Println("               3 删除客户")
		fmt.Println("               4 客户列表")
		fmt.Println("               5 退    出")
		fmt.Print("请选择(1-5): ")

		fmt.Scanln(&cv.key)
		switch cv.key {
		case "1":
			// fmt.Println("添加客户")
			cv.add()
		case "2":
			fmt.Println("修改客户")
		case "3":
			// fmt.Println("删除客户")
			cv.delete()
		case "4":
			// fmt.Println("客户列表")
			cv.list()
		case "5":
			cv.exit()
		default:
			fmt.Println("输入有误，请重新输入")
		}

		if !cv.loop {
			break
		}
	}
	fmt.Println("你退出了系统")
}

func main() {
	fmt.Println("ok")
	cv := customerView{
		key:             "",
		loop:            true,
		customerService: service.FactoryCustomerService(),
	}
	cv.mainMenu()
}
