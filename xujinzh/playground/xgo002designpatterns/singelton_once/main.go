/*
适合只需要创建一个实例的情况，如鼠标
*/
package main

import (
	"fmt"
	"sync"
)

type singleton struct{}

var instance *singleton

var once sync.Once

func GetInstance() *singleton {

	once.Do(func() {
		instance = new(singleton)

	})

	return instance
}

func (s *singleton) DoSomething() {
	fmt.Println("单例模式的某方法")
}

func main() {
	s := GetInstance()
	s.DoSomething()

	s2 := GetInstance()
	s2.DoSomething()

	if s == s2 {
		fmt.Println("s == s2，是单例模式")
	}
}
