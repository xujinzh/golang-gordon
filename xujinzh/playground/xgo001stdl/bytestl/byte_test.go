package bytestl_test

import (
	"bytes"
	"fmt"
	"testing"
)

// bytes标准库

// bytes包提供了对字节切片进行读写操作的一系列函数
// 包括基本处理函数、比较函数、后缀检查函数、索引函数、分割函数、大小写处理函数和子切片处理函数等

// 常用函数
// 转换
// func ToUpper(s []byte) []byte 将s中的所有字符修改为大写格式返回
// func ToLower(s []byte) []byte 将s中的所有字符修改为小写格式返回
// func ToTitle(s []byte) []byte 将s中的所有字符修改为标题格式返回
// func ToUpperSpecial(_case unicode.SpecialCase, s []byte) []byte 使用指定的映射表将s中的所有字符修改为大写格式返回
// func ToLowerSpecial(_case unicode.SpecialCase, s []byte) []byte 使用指定的映射表将s中的所有字符修改为小写格式返回
// func ToTitleSpecial(_case unicode.SpecialCase, s []byte) []byte 使用指定的映射表将s中的所有字符修改为标题格式返回
// func Title(s []byte) []byte 将s中的所有单词的首字符修改为Title格式返回（不能很好的处理以Unicode标点符号分隔的单词）

// 字符串转换
func TestConvert(t *testing.T) {
	// 字符串小写转大写
	b := []byte("hainan")
	a := bytes.ToUpper(b)
	fmt.Printf("b: %s\n", b)
	fmt.Printf("a: %s\n", a)
	// 字符串大写转小写
	c := []byte("SHANGHAI")
	d := bytes.ToLower(c)
	fmt.Printf("c: %s\n", c)
	fmt.Printf("d: %s\n", d)
	// 字符串转为标题
	e := []byte("attention is all you need")
	f := bytes.ToTitle(e)
	fmt.Printf("e: %q\n", e)
	fmt.Printf("f: %q\n", f)
}

// 字符串比较
// func Compare(a, b []byte) int 比较两个[]byte, nil参数相当于空[]byte
// a < b 返回 -1；a == b 返回 0； a > b 返回1
// func Equal(a, b []byte) bool 判断 a, b 是否相等，nil 参数相当于空 []byte
// func EqualFold(s, t []byte) bool 判断 s, t 是否相似，忽略大写、小写、标题三种格式的区别

// 测试比较两个字符串
func TestCompare(t *testing.T) {
	a := []byte("beijing")
	b := []byte("BeiJing")
	// 比较字符串大小
	n := bytes.Compare(a, b)
	fmt.Printf("n: %v\n", n)
	// 比较字符串是否相等
	bt := bytes.Equal(a, b)
	fmt.Printf("bt: %v\n", bt)
	// 不区分大小写比较字符串是否相等
	btFold := bytes.EqualFold(a, b)
	fmt.Printf("btFold: %v\n", btFold)
}

// 清理

// func Trim(s []byte, cutset string) []byte 去掉s两边包含在cutset中的字符（返回s的切片）
// func TrimLeft(s []byte, cutset string) []byte 去掉s左边包含在cutset中的字符（返回s的切片）
// func TrimRight(s []byte, cutset string) []byte 去掉s右边包含在cutset中的字符（返回s的切片）
// func TrimFunc(s []byte, f func(r rune) bool) []byte 去掉s两边符合f函数===返回值是true还是false要求的字符（返回s的切片）
// func TrimLeftFunc(s []byte, f func(r rune) bool) []byte 去掉s左边符合f函数===返回值是true还是false要求的字符（返回s的切片）
// func TrimRightFunc(s []byte, f func(r rune) bool) []byte 去掉s右边符合f函数===返回值是true还是false要求的字符（返回s的切片）
// func TrimSpace(s []byte) []byte 去掉s两边的空白（unicode.IsSpace）（返回s的切片）
// func TrimPrefix(s, prefix []byte) []byte 去掉s的前缀prefix（返回s的切片）
// func TrimSuffix(s, suffix []byte) []byte 去掉s的后缀suffix（返回s的切片）

func TestTrim(t *testing.T) {
	// 定义字节切片二维数组
	bs := [][]byte{
		[]byte("Hello World!"),
		[]byte("Hello 世界！"),
		[]byte("hello golang."),
	}
	// 单个字符判断函数
	// 判断r字符是否包含在 "!！."内
	f := func(r rune) bool {
		return bytes.ContainsRune([]byte("!！."), r)
	}
	// 对二维数组每一个元素（[]byte切片）进行判断
	for _, b := range bs {
		fmt.Printf("去掉两边：%q\n", bytes.TrimFunc(b, f))

	}
	// 去掉H
	fmt.Printf("去掉两边的H：%q\n", bytes.Trim(bs[0], "H"))
	// 去掉前缀He
	fmt.Printf("去掉前缀He：%q\n", bytes.TrimPrefix(bs[0], []byte("He")))
}

// 拆合

// func Split(s, sep []byte) [][]byte -> Split以sep为分隔符将s切分成多个子串，结果不包含分隔符。如果sep为空，则将s切分成Unicode字符列表
// func SplitN(s, sep []byte, n int) [][]byte -> SplitN可以指定切分次数n，超出n的部分将不进行切分
// func SplitAfter(s, sep []byte) [][]byte -> 功能同Split，只不过结果包含分隔符（在各个子串尾部）
// func SplitAfterN(s, sep []byte, n int) [][]byte -> 功能同SplitN，只不过结果包含分隔符（在各个子串尾部）
// func Fields(s []byte) [][]byte -> 以连续空白为分隔符将s切分成多个子串，结果不包含分隔符
// func FieldsFunc(s []byte, f func(rune) bool) [][]byte -> 以符合f的字符为分隔符将s切分成多个子串，结果不包含分隔符
// func Join(s [][]byte, sep []byte) []byte -> 以sep为连接符，将子串列表s连接成一个字符串
// func Repeat(b []byte, count int) []byte -> 将子串b重复count次后返回

func TestSplit(t *testing.T) {
	// 初始化一个[]byte切片
	b := []byte("  Hello   World ! ")
	fmt.Printf("b:%q\n", b)
	// 按照单个空格切分
	fmt.Printf("%q\n", bytes.Split(b, []byte{' '}))
	fmt.Printf("%q\n", bytes.SplitN(b, []byte{' '}, 2))  // 至少2个连续分隔符，作为最终的分隔符，这里是两个空格"  "
	fmt.Printf("%q\n", bytes.SplitAfter(b, []byte{' '})) // 切分后的子串后面带着分隔符
	// 按照连续空白切分
	fmt.Printf("%q\n", bytes.Fields(b))
	// 按照函数定义的方式切分，切分符号包含" "和"!"
	f := func(r rune) bool {
		return bytes.ContainsRune([]byte(" !"), r)
	}
	// 符合函数切分类型的都进行切分，这里是" "和"!"
	fmt.Printf("%q\n", bytes.FieldsFunc(b, f))

	// 拼接
	// 定义字节切片二维数组
	bs := [][]byte{
		[]byte("Hello World!"),
		[]byte("Hello 世界！"),
		[]byte("hello golang."),
	}
	fmt.Printf("%q\n", bytes.Join(bs, []byte{'-'}))
	fmt.Printf("%q\n", bytes.Repeat([]byte("hello"), 3))
}

// 字串

// func HasPrefix(s, prefix []byte) bool 判断s是否有前缀prefix
// func HasSuffix(s, suffix []byte) bool 判断s是否有后缀suffix
// func Contains(b, subslice []byte) bool 判断b中是否包含子串subslice
// func ContainRune(b []byte, r rune) bool 判断b中是否包含子串字符r
// func ContainsAny(b []byte, chars string) bool 判断b中是否包含chars中的任何一个字符
// func Index(s, sep []byte) int 查找子串sep在s中第一次出现的位置，找不到则返回-1
// func IndexByte(s []byte, c byte) int 查找子串字节c在s中第一次出现的位置，找不到则返回-1
// func IndexRune(s []byte, r rune) int 查找子串字符r在s中第一次出现的位置，找不到则返回-1
// func IndexAny(s []byte, chars string) int 查找chars中的任何一个字符在s中第一次出现的位置，找不到则返回-1
// func IndexFunc(s []byte, f func(r rune) bool) int 查找符合f的字符在s中第一次出现的位置，找不到则返回-1
// func LastIndex(s, sep []byte) int 功能同上，只不过查找最后一次出现的位置
// func LastIndexByte(s []byte, c byte) int 功能同上，只不过查找最后一次出现的位置
// func LastIndexAny(s []byte, chars string) int 功能同上，只不过查找最后一次出现的位置
// func LastIndexFunc(s []byte, f func(r rune) bool) int 功能同上，只不过查找最后一次出现的位置
// func Count(s, sep []byte) int 获取sep在s中出现的次数（sep不能重叠）

func TestSubChar(t *testing.T) {
	// 初始化字节切片
	b := []byte("hello golang")
	subslice1 := []byte("hello")
	subslice2 := []byte("Hello")
	// 判断是否包含
	fmt.Printf("%v\n", bytes.Contains(b, subslice1))
	fmt.Printf("%v\n", bytes.Contains(b, subslice2))

	// 初始化新的字节切片
	s := []byte("helloooooooo")
	sep1 := []byte("h")
	sep2 := []byte("l")
	sep3 := []byte("o")
	// 统计包含多少个字串
	fmt.Printf("%v\n", bytes.Count(s, sep1))
	fmt.Printf("%v\n", bytes.Count(s, sep2))
	fmt.Printf("%v\n", bytes.Count(s, sep3))
	// 判断是否包含前缀
	fmt.Printf("%v\n", bytes.HasPrefix(s, sep1))
	// 查找索引位置
	fmt.Printf("%v\n", bytes.Index(s, sep3))
}

// 替换

// func Replace(s, old, new []byte, n int) []byte 将s中前n个old替换为new，n<0则替换全部
// func ReplaceAll(s, old, new []byte) []byte 将s中所有old替换为new
// func Map(mapping func(r rune) rune, s []byte) []byte 将s中的字符替换为mapping的返回值，如果mapping返回负值，则丢弃该字符
// func Runes(s []byte) []rune 将s转换为[]rune类型返回

func TestReplace(t *testing.T) {
	// 初始化字节切片
	s := []byte("hello, world, open door")
	old := []byte("o")
	new := []byte("ee")
	fmt.Printf("%q\n", bytes.Replace(s, old, new, 0))  // 替换0次，即不替换
	fmt.Printf("%q\n", bytes.Replace(s, old, new, 1))  // 替换1次
	fmt.Printf("%q\n", bytes.Replace(s, old, new, 2))  // 替换2次
	fmt.Printf("%q\n", bytes.Replace(s, old, new, -1)) // 替换所有
	fmt.Printf("%q\n", bytes.ReplaceAll(s, old, new))  // 替换所有

	// 初始化字节切片
	s1 := []byte("你好世界")
	r1 := bytes.Runes(s1)
	fmt.Printf("%q转换前字符串(%T)的长度：%v\n", s1, s1, len(s1))
	fmt.Printf("%#q转换后字符串(%T)的长度：%v\n", r1, r1, len(r1))
}

// Buffer 类型

// 缓冲区是具有读取和写入方法的可变大小的字节缓冲区
// Buffer的零值是准备使用的空缓冲区
/*
type Buffer struct {
	buf			[]byte	// contents are the bytes buf[off: len(buf)]
	off			int		// read at &buf[off], write at &buf[len(buf)]
	lastRead	readOp	// last read operation, so that Unread* can work correctly
}

*/

// 声明buffer
// var b bytes.Buffer 直接定义一个Buffer变量，不用初始化，可以直接使用
// b := new(bytes.Buffer) 使用New返回Buffer变量
// b := bytes.NewBuffer(s []byte) 从一个[]byte切片构造一个Buffer
// b := bytes.NewBufferString(s string) 从一个string变量构造一个Buffer

// 向Buffer中写入数据
// b.Write(d []byte) 将切片d写入Buffer数据
// b.WriteString(s string) 将字符串s写入Buffer尾部
// b.WriteByte(c byte) 将字符c写入Buffer尾部
// b.WriteRune(r rune) 将一个rune类型的数据放到缓冲区的尾部
// b.WriteTo(w io.Writer) 将Buffer中的内容输出到实现了io.Writer接口的可写入对象中

// 从Buffer中读取数据
// b.Read(c) 一次读取8个byte到c容器中，每次读取新的8个byte覆盖c中原来的内容
// b.ReadByte() 读取第一个byte，b的第一个byte被拿掉，赋值给a，这里a, _ := b.ReadByte()
// b.ReadRune() 读取第一个rune，b的第一个rune被拿掉，赋值给r, 这里r, _ := b.ReadRune()
// b.ReadBytes(delimiter byte) 需要一个byte作为分隔符，读的时候从缓冲区里找第一个出现的分隔符（delim）
// 找到后，把从缓冲区头部开始到分隔符之间的所有byte进行返回，作为byte类型的slice
// 返回后，缓冲区也会空掉一部分
// b.ReadString(delimiter byte) 需要一个byte作为分隔符，读的时候从缓冲区里找第一个出现的分隔符（delim）
// 找到后，把从缓冲区头部开始到分隔符之间的所有byte进行返回，作为字符串返回
// 返回后，缓冲区也会空掉一部分
// b.ReadFrom(i io.Reader) 从一个实现io.Reader接口的i，把i里的内容读到缓冲区里，n返回读的数量

func TestBuffer(t *testing.T) {
	//
	rd := bytes.NewBufferString("Hello World!")
	fmt.Printf("rd: %q, type: %T\n", rd, rd)
	// 获取数据切片
	b := rd.Bytes()
	fmt.Printf("获取数据切片: %s\n", b)
	// 读出一部分数据，看看切片有没有变化
	buf := make([]byte, 6)
	rd.Read(buf)
	fmt.Printf("读取后rd: %q\n", rd.String()) // 读取后，前6个字节被读取了，缓冲区还剩下后面的内容
	fmt.Printf("获取数据切片: %s\n", b)          // 但是，获取的数据切片内容不变

	// 写入一部分数据，看看切片有没有变化
	rd.Write([]byte(" golang"))
	fmt.Printf("写入内容后rd: %q\n", rd.String())
	fmt.Printf("获取数据切片: %s\n", b) // 但是，获取的数据切片内容不变

	// 再读取一部分数据，看看切片有没有变化
	rd.Read(buf)
	fmt.Printf("读取后rd: %q\n", rd.String())
	fmt.Printf("获取数据切片%s\n", b)

	// 重置
	rd.Reset()
	fmt.Printf("reset后rd: %q\n", rd.String())
}

// 其他方法

// func (b *Buffer) Len() int 未读取部分的数据长度
// func (b *Buffer) Cap() int 获取缓存的容量
// func (b *Buffer) Next(n int) []byte 读取前n字节的数据并以切片形式返回，如果数据长度小于n，则全部读取。切片只在下一次读写操作前合法
// func (b *Buffer) Bytes() []byte 引用未读取部分的数据切片（不移动读取位置）
// func (b *Buffer) String() string 返回未读取部分的数据字符串（不移动读取位置）
// func (b *Buffer) Grow(n int) 自动增加缓存容量，以保证有n字节的剩余空间，如果n小于0或无法增加容量则会panic
// func (b *Buffer) Truncate(n int) 将数据长度截短到n字节，如果n小于0或大于Cap则panic
// func (b *Buffer) Reset() 重设缓冲区，清空所有数据（包括初始内容）

// Reader类型

/*
	type Reader struct {
		s			[]byte
		i			int64 		// current reading index
		prevRune	int			// index of previous rune; or < 0
	}

# Reader 实现了 io.Reader, io.ReaderAt, io.WriterTo, io.Seeker, io.ByteScanner, io.RuneScanner 接口

func NewReader(b []byte) *Reader 将b包装成bytes.Reader对象
func (r *Reader) Len() int 返回未读取部分的数据长度
func (r *Reader) Size() int64 返回底层数据总长度，方便ReadAt使用，返回值永远不变
func (r *Reader) Reset(b []byte) 将底层数据切换为b，同时复位所有标记（读取位置等信息）
*/
func TestReader(t *testing.T) {
	// 初始化字符串
	data := "123456789"
	// 创建Reader
	reader := bytes.NewReader([]byte(data))
	// 返回未读取部分的长度
	fmt.Printf("未读取数据长度reader.Len(): %v\n", reader.Len())
	// 返回底层数据总长度
	fmt.Printf("底层数据总长度reader.Size(): %v\n", reader.Size())

	// 读取数据
	buf := make([]byte, 2)
	for {
		fmt.Printf("剩余数据长度reader.Len(): %v\n", reader.Len())
		// 每次读取长度2
		n, err := reader.Read(buf)
		if err != nil {
			break
		}
		// 打印读取的内容
		fmt.Println(string(buf[:n]))
	}

	// 初始化字符串
	data = "123456789"
	// 创建Reader
	reader = bytes.NewReader([]byte(data))
	// 读取数据
	// 设置偏移量
	reader.Seek(1, 0)
	for {
		// 一个字节一个字节的读取
		b, err := reader.ReadByte()
		if err != nil {
			break
		}
		fmt.Printf("b: %q\n", b)
	}

	// 初始化字符串
	data = "123456789"
	// 创建Reader
	reader = bytes.NewReader([]byte(data))
	// 读取数据
	buf = make([]byte, 2)
	off := int64(1)
	for {
		n, err := reader.ReadAt(buf, off)
		if err != nil {
			break
		}
		off += int64(n)
		fmt.Println("偏移量：", off, "，读取的值", string(buf[:n]))
	}
}
