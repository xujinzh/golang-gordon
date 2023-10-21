package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var i int8 = 12
	x := i + 10
	fmt.Printf("%T, %d\n", i, unsafe.Sizeof(i))
	fmt.Printf("%T, %d", x, unsafe.Sizeof(x))
}