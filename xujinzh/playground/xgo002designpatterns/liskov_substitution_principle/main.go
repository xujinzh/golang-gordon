/*
里氏代换原则 Liskov Substitution Principle, LSP
任何抽象类（interface接口）出现的地方都可以用他的实现类进行替换，实际就是虚拟机制，语言级别实现面向对象功能。

依赖倒转原则 Dependence Inversion Principle, DIP
依赖于抽象（接口），不要依赖具体的实现（类），也就是针对接口编程。
*/
package main

import "fmt"

// 依赖倒转原则

// ---> 抽象层 <---
type Car interface {
	Run()
}

type Driver interface {
	Drive(car Car)
}

// ---> 实现层 <---

type BenZ struct {
}

func (b *BenZ) Run() {
	fmt.Println("BenZ is running...")
}

type Bmw struct{}

func (b *Bmw) Run() {
	fmt.Println("Bmw is running...")
}

type ZhangSan struct{}

func (zs *ZhangSan) Drive(car Car) {
	fmt.Println("ZhangSan drive car")
	car.Run()
}

type LiSi struct{}

func (ls *LiSi) Drive(car Car) {
	fmt.Println("LiSi drive car")
	car.Run()
}

// +++
type WangWu struct{}

func (ww *WangWu) Drive(car Car) {
	fmt.Println("WangWu drive car")
	car.Run()
}

// ---> 业务逻辑层 <---

func main() {
	// ZhangSan drive BenZ
	var benz Car
	benz = new(BenZ)
	var zhangsan Driver
	zhangsan = new(ZhangSan)
	zhangsan.Drive(benz)

	// LiSi drive Bmw
	var bmw Car
	bmw = new(Bmw)
	var lisi Driver
	lisi = new(LiSi)
	lisi.Drive(bmw)

	// zhangsan drive bmw
	zhangsan.Drive(bmw)

	// +++
	var wangwu Driver = new(WangWu)
	wangwu.Drive(benz)
	wangwu.Drive(bmw)
}
