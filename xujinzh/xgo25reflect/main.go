package main

import (
	"fmt"
	"reflect"
)

func reflectTest(b interface{}) {
	// 通过反射获取传入的变量的 type，kind 和值
	// 1. 先获取 reflect.Type
	rType := reflect.TypeOf(b)
	fmt.Println("reflect type =", rType)

	// 2. 获取 reflect.Value
	rValue := reflect.ValueOf(b)
	fmt.Printf("reflect value = %v, reflect value type = %T\n", rValue, rValue)

	fmt.Println("2 + rValue =", 2+rValue.Int())

	// 3. 将 rValue 转成 interface{}
	iValue := rValue.Interface()
	// 将 interface{} 通过断言转成需要的类型
	num2 := iValue.(int)
	fmt.Println(num2)
}

func reflectTest02(b interface{}) {
	rType := reflect.TypeOf(b)
	fmt.Printf("iType = %v, iType Type = %T\n", rType, rType)

	rValue := reflect.ValueOf(b)

	iValue := rValue.Interface()
	fmt.Printf("iValue = %v, iValue Type = %T\n", iValue, iValue)

	// 这里必须使用类型断言，虽然从运行结果看 iValue 的类型是 Student 结构体类型，但是那是运行时，编译时无法判断，会报错
	// 将 interface() 通过类型断言转成需要的类型，如果类型不确定，可以使用 switch
	student, ok := iValue.(Student)
	if ok {
		fmt.Printf("student.Name = %v\n", student.Name)
	}

	// 获取变量对应的 Kind, 1. rValue.Kind(); 2. rType.Kind()
	kind1 := rValue.Kind()
	kind2 := rType.Kind()
	fmt.Printf("kind = %v, kind = %v\n", kind1, kind2)

}

type Student struct {
	Name string
	Age  int
}

func main() {
	// 基本数据类型、空接口interface{}、reflect.Value 进行反射基本操作
	// 先定义一个 int
	var num int = 100
	reflectTest(num)

	// 定义一个 Student 实例
	student := Student{
		Name: "tom",
		Age:  10,
	}
	reflectTest02(student)
}
