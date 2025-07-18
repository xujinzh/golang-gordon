package main

import (
	"fmt"
	"xgo2/utils"
)

var MainAge = testVarFirst()

func testVarFirst() int {
	fmt.Println("testVarFirst function called")
	return 20
}

func init() {
	fmt.Println("main init function called")
}

func main() {
	fmt.Println("main()", utils.Age)
	fmt.Println("main()", utils.Name)
	fmt.Println("main()", MainAge)
}
