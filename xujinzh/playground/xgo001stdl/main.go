package main

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"math/big"
	"stl/bytesslice"
	"strconv"
	"strings"
)

func reverseSlice(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i] // 交换元素
	}
}

func byteString() {

	n, err := fmt.Println("hello")
	if err != nil {
		fmt.Println("error", err)

	} else {
		fmt.Println("len(hello) + len('\\n')=", n)
	}

	// 把十六进制字符串转成整数
	hexStr := "e4e"

	num, err := strconv.ParseInt(hexStr, 16, 64)
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println("0x0e4e=", num)

	// 把十六进制字节切片转成整数
	b := []byte{0x00, 0x00, 0xee, 0xff}
	// 先转成字符串在转成整数
	u32i, err := strconv.ParseInt(hex.EncodeToString(b), 16, 64)
	// fmt.Println(string(b))
	fmt.Println("hex encode to string:", hex.EncodeToString(b))
	fmt.Printf("% X\n", string(b))
	fmt.Println("先转成字符串再转成整数：", u32i)
	// 按照大端转换
	// 直接用binary转
	u64 := binary.BigEndian.Uint32(b)
	fmt.Printf("%X= %d\n", string(b), u64)
	fmt.Printf("%s= %d\n", strings.ToUpper(hex.EncodeToString(b)), u64)
	// 按照小端转换
	u32little := binary.LittleEndian.Uint32(b)
	reverseSlice(b)
	fmt.Printf("%X= %d\n", string(b), u32little)
	// 用math包，可处理大整数
	reverseSlice(b)
	bigIntValue := new(big.Int).SetBytes(b)
	fmt.Printf("Big Int value: %d, type: %T \n", bigIntValue, bigIntValue)

	// 去除字符串左右空格
	b = []byte("  hello, world  ")
	fmt.Println(string(b))
	bb := bytes.TrimSpace(b)
	fmt.Println(string(bb))

	// 判断r字符是否包含在 "!. "内
	f := func(r rune) bool {
		return bytes.ContainsRune([]byte("!. "), r)
	}

	// 去掉两边满足函数f的字符
	fmt.Printf("%q\n", bytes.TrimFunc([]byte(" hello. world!"), f))

}

func main() {
	byteString()
	bytesslice.ByteAppend()
	bytesslice.TestSlicesByteSplit()

}
