package main

func main() {
	// create apple and show
	var appleFactory AbstractFactory
	appleFactory = new(AppleFactory)

	var apple Fruit
	apple = appleFactory.CreateFruit()
	apple.Show()

	// create banana and show
	bananaFactory := new(BananaFactory)
	banana := bananaFactory.CreateFruit()
	banana.Show()

	// new add
	peachFactory := new(PeachFactory)
	peach := peachFactory.CreateFruit()
	peach.Show()
}
