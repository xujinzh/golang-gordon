package main

import (
	"bytes"
	"fmt"
	"strings"
)

// 拼接[]byte的方法
func ByteAppend() {

	// 1. 使用 append 函数，最常用
	b1 := []byte{0x00, 0x00}
	b2 := []byte{0xee, 0xff}
	b3 := []byte{0x00, 0x00, 0xee, 0xff}
	// method1: use append
	result1 := append(b1, b2...)
	fmt.Printf("append1: % x\n", result1)

	// method2: append multi slice
	result2 := append(b1, append(b2, b3...)...)
	fmt.Printf("append2: % x\n", result2)

	// method3: use init empty slice
	var result3 []byte
	result3 = append(result3, b1...)
	result3 = append(result3, b2...)
	result3 = append(result3, b3...)
	fmt.Printf("append3: % x\n", result3)

	// 2. 使用 bytes.Buffer，适合大量拼接
	var buf bytes.Buffer
	// 写入字节切片
	buf.Write(b1)
	buf.Write(b2)
	buf.Write(b3)
	// 也可以写入单个字节
	buf.WriteByte(0xaa)
	// 写入字符串，会自动转换为字节
	buf.WriteString("hello")
	// 转为[]byte
	result4 := buf.Bytes()
	fmt.Printf("result4: % x \n", result4)
	fmt.Println(result4)

	// 3. 使用 copy 函数，性能最好
	// 先预先分配足够空间
	totalLen := len(b1) + len(b2) + len(b3)
	// 根据已知空间大小创建切片，切片的cap容量是已知空间大小
	result5 := make([]byte, 0, totalLen)
	// 方法1：先扩容再copy
	result5 = result5[:totalLen]
	offset := 0
	offset += copy(result5[offset:], b1)
	offset += copy(result5[offset:], b2)
	offset += copy(result5[offset:], b3)
	fmt.Printf("result5: % x \n", result5)

	// 方法2：
	result6 := make([]byte, totalLen)
	n := 0
	n += copy(result6[n:], b1)
	n += copy(result6[n:], b2)
	n += copy(result6[n:], b3)
	fmt.Printf("result6: % x \n", result6)

	// 4. 使用 bytes.Join，类似于字符串的 Join
	slices := [][]byte{
		{0x00, 0x00},
		{0xee, 0xff},
		{0x00, 0x00, 0xee, 0xff},
	}
	// 使用空字节的切片作为分隔符
	result7 := bytes.Join(slices, []byte{})
	fmt.Printf("result7: % x \n", result7)

	// 使用指定的分隔符
	result8 := bytes.Join(slices, []byte{0xaa, 0xaa})
	fmt.Printf("result8: % x \n", result8)

	// 5. 使用 strings.Builder，适合频繁拼接
	var builder strings.Builder
	// 接受字节片
	builder.Write(b1)
	builder.Write(b2)
	builder.Write(b3)
	// 转为[]byte
	result9 := []byte(builder.String())
	fmt.Printf("result9: % x \n", result9)
	// 多种方式组合使用
	// 使用 concatBytes
	result10 := concatBytes(b1, b2, b3, []byte{0xaa}, []byte{0xbb})
	fmt.Printf("result10: % x \n", result10)
	// 带分隔符的拼接
	slices = [][]byte{b1, b2, b3}
	separator := []byte{0xaa, 0xbb}
	result11 := bytes.Join(slices, separator)
	fmt.Printf("result11: % X \n", result11)

	var withSeparator []byte
	for i, s := range slices {
		withSeparator = append(withSeparator, s...)
		if i < len(slices)-1 {
			withSeparator = append(withSeparator, separator...)
		}
	}
	fmt.Printf("withSeparator: % X \n", withSeparator)

	result12 := efficientConcatBytes(slices)
	fmt.Printf("result12: % X \n", result12)
}

func concatBytes(slices ...[]byte) []byte {
	// 计算总长度
	total := 0
	for _, s := range slices {
		total += len(s)
	}
	// 分配空间并拼接
	result := make([]byte, 0, total)
	for _, s := range slices {
		result = append(result, s...)
	}
	return result
}

func efficientConcatBytes(slices [][]byte) []byte {
	total := 0
	for _, s := range slices {
		total += len(s)
	}
	result := make([]byte, total)
	offset := 0
	for _, s := range slices {
		offset += copy(result[offset:], s)
	}
	return result
}
