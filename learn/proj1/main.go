package main

import "fmt"
import "unsafe"

func main() {
	fmt.Println("hello, \rgo!")

	var x uint = '中'
	fmt.Printf("%d", unsafe.Sizeof(x))
}