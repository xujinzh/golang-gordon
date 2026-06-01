package main

type Fruit interface {
	Show()
}

type AbstractFactory interface {
	CreateFruit() Fruit // 生产水果类的生产器方法
}
