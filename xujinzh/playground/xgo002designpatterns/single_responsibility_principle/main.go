/*
单一职责原则 Single Responsibility Principle, SRP

类的职责单一，对外只提供一种功能，而引起类变化的原因都应该只有一个。
*/
package main

import "fmt"

type ClothesShop struct{}

func (cs *ClothesShop) Style() {
	fmt.Println("逛街的装扮")
}

type ClothesWork struct{}

func (cw *ClothesWork) Style() {
	fmt.Println("工作的装扮")
}

func main() {
	// 工作的业务
	cw := ClothesWork{}
	cw.Style()

	// 逛街的业务
	cs := ClothesShop{}
	cs.Style()
}
