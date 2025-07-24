package main

import (
	"fmt"
)

type USB interface {
	// 声明或定义了两个没有实现的方法
	Start()
	Stop()
}

type Phone struct {
	Name string
}

// 让 Phone 实现 USB 接口的（所有）方法
func (phone Phone) Start() {
	fmt.Printf("%v started\n", phone.Name)
}

func (phone Phone) Stop() {
	fmt.Printf("%v stoped\n", phone.Name)
}

type Camera struct {
	Weight float64
}

// 让 Camera 实现 USB 接口的（所有）方法
func (camera Camera) Start() {
	fmt.Printf("Camera weighing %v is started\n", camera.Weight)
}

func (camera Camera) Stop() {
	fmt.Printf("Camera weighing %v is stoped\n", camera.Weight)
}

type Computer struct {
}

// 编写一个方法 working 方法，接收一个 USB 接口类型的变量
// 只要是实现了 USB 接口（指实现了USB接口声明的所有方法）
func (computer Computer) Working(usb USB) {
	usb.Start()
	usb.Stop()
}

func main() {
	fmt.Println()
	// 创建接口体变量
	computer := Computer{}
	phone := Phone{"iPhone"}
	camera := Camera{2.3}

	computer.Working(phone)

	computer.Working(camera)
}
