package main

import (
	"fmt"
)

type Cat struct {
	Name string
	Age  int
}

func main() {
	allChan := make(chan interface{}, 3)
	allChan <- 10
	allChan <- "tom"
	cat := Cat{"xiaomao", 4}
	allChan <- cat

	// 希望获得管道中第三个元素，先将前两个退出
	<-allChan
	<-allChan

	newCat := <-allChan

	fmt.Printf("newCat type = %T, newCat value = %v\n", newCat, newCat)

	newCat2, ok := newCat.(Cat) // 类型推断
	if ok {
		fmt.Println(newCat2.Name)
	}
}
