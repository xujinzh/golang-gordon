package main

import (
	"fmt"
	"xgo14/utils"
)

func main() {
	fmt.Println("hello")
	p := utils.Person("xiaoming")
	fmt.Println(*p)
	p.SetAge(20)
	fmt.Printf("Name=%q, Age=%v\n", p.Name, p.GetAge())
}
