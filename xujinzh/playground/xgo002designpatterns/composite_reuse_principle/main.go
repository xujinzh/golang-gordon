/*
1. 接口隔离原则（Interface Segregation Principle, ISP）不应该强迫用户的程序依赖他们不需要的接口方法。一个接口应该只提供一种对外功能，不应该把所有操作都封装到一个接口中去。

2. 合成复用原则（Composite Reuse Principle, CRP）如果使用继承，会导致父类的任何变换都可能影响到子类的行为。如果使用对象组合，就降低了这种依赖关系。对于继承和组合，优先使用组合。

*/

package main

import "fmt"

/*
有一个结构体animal具有eat的方法，我想创建一个cat结构体，其不仅具有eat方法还有sleep方法。
*/
type Animal struct {
}

func (a *Animal) Eat() {
	fmt.Println("eating")
}

// Cat have eat and sleep

type Cat struct {
	// method 1
	// a *Animal
}

// method 1
// func (c *Cat) Eat(){
// 	c.Eat()
// }

// method 2
func (c *Cat) Eat(a *Animal) {
	a.Eat()
}

func (c *Cat) Sleep() {
	fmt.Println("sleeping")
}

func main() {
	a := &Animal{}
	a.Eat()

	fmt.Println("---------")
	c := &Cat{}
	// method 1
	// c.Eat()

	// method 2
	c.Eat(a)
	c.Sleep()
}
