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

func (phone Phone) Call() {
	fmt.Printf("%v calling \n", phone.Name)
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
	// 如果USB指向Phone结构体变量，则需要调用Call方法
	if phone, ok := usb.(Phone); ok {
		phone.Call()
	}
	usb.Stop()
}

func main() {
	fmt.Println()
	// 创建接口体数组
	var usbArr [3]USB
	usbArr[0] = Phone{"小米"}
	usbArr[1] = Phone{"苹果"}
	usbArr[2] = Camera{3.14}

	computer := Computer{}
	for _, device := range usbArr {
		computer.Working(device)
	}
}
