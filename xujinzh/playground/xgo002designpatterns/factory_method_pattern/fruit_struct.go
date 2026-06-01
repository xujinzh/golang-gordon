package main

import "fmt"

type Apple struct {
	// Fruit // 为了易于理解
}

func (apple *Apple) Show() {
	fmt.Println("I'm an apple")
}

type Banana struct {
}

func (banana *Banana) Show() {
	fmt.Println("I'm a banana")
}

// new add
type Peach struct {
}

func (peach *Peach) Show() {
	fmt.Println("I'm a peach")
}
