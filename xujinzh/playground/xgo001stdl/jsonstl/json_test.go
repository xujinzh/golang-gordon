package jsonstl_test

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
)

/*
json标准库

json包可以实现json的编码和解码，就是将json字符串转换为struct，或者将struct转化为json。

1. 核心函数

1.1 Marshal

func Marshal(v any) ([]byte, error)

将struct编码成json，可以接受任意类型

- 布尔型转换为json后仍是布尔型，如true->true
- 浮点型和整数型转化后为json里面的常规数字，如1.23->1.23
- 字符串将以UTF-8编码转化输出位Unicode字符集的字符串，特殊字符比如<将会被转义为\u003c
- 数组和切片被转换为json里面的数组，[]byte类会被转换为base64编码后的字符串，slice的零值被转换为 null
- 结构体会转化为json对象，并且只有结构体里边以大写字母开头的可被导出的字段才会被转化输出，而这些可导出的字段会作为json对象的字符串索引
- 转化一个map类型的数据结构时，该数据的类型必须是map[string]T（T可以使encoding/json包支持的任意数据类型）


1.2 Unmarshal

func Unmarshal(data []byte, v any) error

将json转码为struct结构体

这个函数会把传入的data作为一个json来进行解析，解析后的数据存储在参数v中。
这个参数v也是任意类型的参数（但一定是一个类型的指针），因为以此参数进行json解析的时候，这个函数不知道这个传入参数的具体类型，所以它需要接受所有的类型。
*/

// 定义一个结构体
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
	Married bool   `json:"married"`
}

func TestMarshal(t *testing.T) {
	// 创建一个实例
	p := Person{
		Name:    "zhangsan",
		Age:     20,
		Email:   "zhangsan@gmail.com",
		Married: true,
	}
	// 编码为json
	b, _ := json.Marshal(p)
	// 打印
	fmt.Printf("string(b): %v\n", string(b))

	// 反编码会person类
	jsonStr := string(b)
	p1 := &Person{}
	json.Unmarshal([]byte(jsonStr), p1)
	fmt.Printf("p1: %#v\n", p1)

}

/*
2.1 Decoder

	type Decoder struct {
		// contains filtered or unexported fields
	}

从输入流读取并解析json，应用于io流Reader Writer可以扩展到http websocket 等场景。
*/
func TestDecoder(t *testing.T) {
	// 打开文件
	f, _ := os.Open("./resources/test.json")
	// 关闭文件
	defer f.Close()

	// 创建解码器
	d := json.NewDecoder(f)

	// 声明一个map存放解码后的内容
	// var v map[string]interface{}
	var v map[string]any
	// 解码
	d.Decode(&v)
	// 打印
	fmt.Printf("v: %v\n", v)

}

/*
2.2 Encoder

type Encoder struct {
	// contains filtered or unexported fields
}

写json到输出流，应用于io流Reader Writer可以扩展到http websocket等场景。
*/

func TestEncoder(t *testing.T) {
	// 创建一个实例
	p := Person{
		Name:    "zhangsna",
		Age:     20,
		Email:   "zs@163.com",
		Married: false,
	}
	// 打开一个文件，用于将编码后的内容写入文件中
	f, _ := os.OpenFile("./resources/a.json", os.O_WRONLY|os.O_CREATE, 0755)
	defer f.Close()
	// 创建一个编码器
	d := json.NewEncoder(f)
	// 进行编码
	d.Encode(p)

}
