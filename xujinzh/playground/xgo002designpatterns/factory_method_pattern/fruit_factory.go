package main

type AppleFactory struct {
}

func (factory *AppleFactory) CreateFruit() Fruit {
	var fruit Fruit
	fruit = new(Apple)
	return fruit
}

type BananaFactory struct {
}

func (factory *BananaFactory) CreateFruit() Fruit {
	fruit := new(Banana)
	return fruit
}

// new add
type PeachFactory struct {
}

func (factory *PeachFactory) CreateFruit() Fruit {
	fruit := new(Peach)
	return fruit
}
