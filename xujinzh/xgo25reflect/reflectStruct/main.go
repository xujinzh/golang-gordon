package main

import (
	"fmt"
	"reflect"
)

// define a Monster struct
type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"monster-age"`
	Score float32
	Sex   string
}

// method, show struct info
func (monster Monster) Print() {
	fmt.Println(monster)
}

// method, get sum
func (monster Monster) GetSum(n1, n2 int) int {
	return n1 + n2
}

// method, set monster field value
func (monster Monster) Set(name string, age int, score float32, sex string) {
	monster.Name = name
	monster.Age = age
	monster.Score = score
	monster.Sex = sex
}

func TestStruct(any interface{}) {
	// get reflect.Type
	rType := reflect.TypeOf(any)
	// get reflect.Value
	rValue := reflect.ValueOf(any)
	// judge is struct
	anyKind := rValue.Kind()
	if anyKind != reflect.Struct {
		fmt.Println("expect struct")
		return
	}

	// get num of field of struct
	numField := rValue.NumField()
	fmt.Printf("struct has %d fields\n", numField)

	// traversal all field
	for i := 0; i < numField; i++ {
		fmt.Printf("Field %d: value = %v\n", i, rValue.Field(i))
		// get tag of struct
		tagValue := rType.Field(i).Tag.Get("json")
		if tagValue != "" {
			fmt.Printf("Field %d: tag = %v\n", i, tagValue)
		}
	}

	// get num of method
	numMethod := rValue.NumMethod()
	fmt.Printf("struct has %d methods\n", numMethod)

	// call struct method, input arg is [] reflect.Value, here is null
	rValue.Method(1).Call(nil)

	// call struct method with arguments, input arg is []reflect.Value, and output arg is also []reflect.Value
	var params []reflect.Value
	params = append(params, reflect.ValueOf(10))
	params = append(params, reflect.ValueOf(40))
	res := rValue.Method(0).Call(params)
	fmt.Println("res =", res[0].Int())
}

func main() {
	// define a Monster instance
	monster := Monster{
		Name:  "monkey",
		Age:   500,
		Score: 88,
	}
	TestStruct(monster)
}
