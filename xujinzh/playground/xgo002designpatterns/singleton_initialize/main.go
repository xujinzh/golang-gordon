package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type singleton struct{}

var instance *singleton

var initialized uint32
var lock sync.Mutex

func GetInstance() *singleton {
	// 如果不是第一次被调用，那么说明已经创建了一个实例，则直接返回
	if atomic.LoadUint32(&initialized) == 1 {
		return instance
	}

	// 如果是第一次调用，那么为了避免同时创建实例抢夺资源，那么先对第一个创建实例的加锁
	lock.Lock()
	defer lock.Unlock()

	if instance == nil {
		instance = new(singleton)
		return instance
	}

	return instance

}

func (s *singleton) DoSomething() {
	fmt.Println("do something!")
}

func main() {
	s := GetInstance()
	s.DoSomething()

	s2 := GetInstance()
	s2.DoSomething()

	if s == s2 {
		fmt.Println("s == s2, 属于单例模式")
	}
}
