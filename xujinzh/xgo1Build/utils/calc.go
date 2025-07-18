package utils

import (
	_ "fmt"
	_ "math"
)

func Add(a int, b int) (int, int) {
	sum := a + b
	diff := a - b
	return sum, diff
}