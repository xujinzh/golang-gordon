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

func main() {
	fmt.Println()
	// 创建接口数组
	var usb [3]USB
	usb[0] = Phone{"苹果"}
	usb[1] = Phone{"小米"}
	usb[2] = Camera{3.14}
	fmt.Println(usb)
}
