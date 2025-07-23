package utils

import (
	"fmt"
)

type person struct {
	Name string
	age  int
}

func (p *person) SetAge(age int) {
	if age < 0 || age > 150 {
		fmt.Println("年龄不合法")
		return
	}
	p.age = age
}

func (p *person) GetAge() int {
	return p.age
}

func Person(name string) *person {
	return &person{
		Name: name,
	}
}
