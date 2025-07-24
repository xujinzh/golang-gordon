package main

import (
	"fmt"
)

type Goods struct {
	Name  string
	Price float64
}

type Brand struct {
	Name    string
	Address string
}

type TV struct {
	Goods
	*Brand
}

func main() {
	fmt.Println()
	tv := TV{
		Goods{"电视机SN30239", 5000.9},
		&Brand{"海尔", "shandong"},
	}
	fmt.Println(tv.Goods)
	fmt.Println(*tv.Brand)
}
