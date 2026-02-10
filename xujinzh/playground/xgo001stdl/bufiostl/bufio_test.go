package bufiostl_test

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

// bufio标准库

/*
bufio包实现了有缓冲的I/O
它包装了一个io.Reader或io.Writer接口对象，创建另一个也实现了该接口，且同时还提供了缓冲和一些文本I/O的帮助函数的对象。


*/

/*
1. bufio包原理

io操作本身的效率并不低，低的是频繁的访问本地磁盘的文件。
所以bufio就提供了缓冲区（分配一块内存），读和写都先在缓冲区中，最后再读写文件，来降低访问本地磁盘的次数，从而提高效率。

简单来说，把文件读取进缓冲区（内存）之后再读取的时候就可以避免文件系统的io从而提高效率。
同理，在进行写操作时，先把文件写入缓冲（内存），然后由缓冲写入文件系统。

缓冲区的设计是为了存储多次的写入，最后一口气把缓冲区内容写入文件。

1. 读取大于buf的内容：
程序读取大于buf的内容时直接从文件读取；
读取小于buf的内容时，从缓冲区读取，当buf为空时，一次性从文件中读取部分大小的内容到缓冲区

2. 写入大于buf的内容：
程序写入大于buf的内容时，直接写入文件
写入小于buf的内容时，先写入缓冲区，当buf没有足够空间的时候，把缓冲区的内容写入文件并清空缓冲区

bufio封装了io.Reader或io.Writer接口对象，并创建另一个也实现了该接口的对象
*/

/*
2. type Reader

2.1 基础介绍

bufio.Reader 是bufio中对io.Reader的封装

type Reader struct {
	buf				[]byte
	rd				io.Reader 	// reader provided by the client
	r, w			int			// buf read and write positions
	err				error
	lastByte		int 		// last byte read for UnreadByte; -1 means invalid
	lastRuneSize	int			// size of last rune read for UnreadRune; -1 means invalid
}

bufio.Read(p []byte)相当于读取大小len的内容，思路如下：
1. 当缓冲区有内容时，将缓冲区内容全部填入p并清空缓冲区
2. 当缓冲区没有内容的时候且len > len(buf)，即要读取的内容比缓冲区还要大，直接去文件读取即可
3. 当缓冲区没有内容的时候且len > len(buf)，即要读取的内容比缓冲区小，缓冲区从文件读取内容充满缓冲区，并将p填满（此时缓冲区有剩余内容）
4. 以后再次读取时缓冲区有内容，将缓冲区内容全部填入p并清空缓冲区（此时和情况1一样）
5. reader内部通过维护一个r，w即读入和写入的位置索引来判断是否缓冲区内容被全部读出
*/

/*
2.2 方法

func NewReaderSize

func NewReaderSize(rd io.Reader, size int) *Reader

NewReaderSize将rd封装成一个带缓冲的bufio.Reader对象，缓冲大小有size指定（如果小于16则会被设置为16）

如果rd的基类型就是有足够缓冲的bufio.Reader类型，则直接将rd转换为基类型返回


func NewReader

func NewReader(rd io.Reader) *Reader

NewReader 相当于 NewReaderSize（rd, 4096)


func (b *Reader) Reset(r io.Reader)

Reset丢弃缓冲中的数据，清除任何错误，将b重设为其下层从r读取数据
*/

func TestReader(t *testing.T) {
	// create a reader
	r1 := strings.NewReader("ABCDEFG")
	r2 := strings.NewReader("12345")

	// create bufio new reader
	br := bufio.NewReader(r1)

	// read a string
	b, _ := br.ReadString('\n')
	fmt.Printf("b: %v\n", b)

	// reset to new reader
	br.Reset(r2)

	// read a string
	b, _ = br.ReadString('\n')
	fmt.Printf("b: %v\n", b)

}

/*
func (b *Reader) Read(p []byte) (n int, err error)

Read读取数据写入p
返回写入p的字节数
一次调用最多会调用下层Reader接口一次Read方法，因此返回值n可能小于len。
读取到达结尾时，返回值n将为0而err将为io.EOF
*/

func TestRead(t *testing.T) {
	// create reader
	r := strings.NewReader("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	// create bufio reader
	br := bufio.NewReader(r)
	// read
	p := make([]byte, 10)
	for {
		n, err := br.Read(p)
		if n == 0 || err == io.EOF {
			break
		} else {
			fmt.Printf("string(p[0:n]): %v\n", string(p[0:n]))
		}
	}
}

/*
func (b *Reader) Peek(n int) ([]byte, error)

Peek返回缓存的一个切片，该切片引用缓存中前n个字节的数据，该操作不会将数据读出，只是引用
引用的数据在下一次读取操作之前是有效的。
如果切片长度小于n，则返回一个错误信息说明原因
如果n大于缓存的总大小，则返回ErrBufferFull
*/

/*
func (b *Reader) ReadByte() (c byte, err error)

ReadByte读取并返回一个字节
如果没有可用的数据，会返回错误
*/

/*
func (b *Reader) UnreadByte() error

UnreadByte吐出最近一次读取操作读取的最后一个字节
只能吐出最后一个，多次调用会出问题
*/

func TestReadByte(t *testing.T) {
	// create reader
	r := strings.NewReader("ABCDEFG")
	// create bufio reader
	br := bufio.NewReader(r)
	// read byte
	c, _ := br.ReadByte()
	fmt.Printf("c: %q\n", c)

	// read byte
	c, _ = br.ReadByte()
	fmt.Printf("c: %q\n", c)

	// unreade byte
	// 吐出最后一次读取的字节
	br.UnreadByte()
	c, _ = br.ReadByte()
	fmt.Printf("c: %q\n", c)
}

/*
func (b *Reader) ReadRune() (r rune, size int, err error)

ReadRune读取一个UTF-8编码的unicode码值，返回该码值、其编码长度和可能得错误。
如果UTF-8编码非法，读取位置只移动1字节，返回U+FFFD，返回值size为1而err为nil
如果没有可用的数据，会返回错误
*/

/*
func (b *Reader) UnreadRune() error

UnreadRune吐出最后一次ReadRune调用读取的unicode码值
如果最后一次读取不是调用的ReadRune，会返回错误
UnreadRune比UnreadByte严格
*/

func TestReadRune(t *testing.T) {
	// create reader use chinese
	r := strings.NewReader("你好世界")
	// create bufio reader
	br := bufio.NewReader(r)
	// read rune
	c, size, _ := br.ReadRune()
	fmt.Printf("%c %v \n", c, size)

	// read rune
	c, size, _ = br.ReadRune()
	fmt.Printf("%c %v\n", c, size)

	// unread rune
	br.UnreadRune()
	c, size, _ = br.ReadRune()
	fmt.Printf("%c %v\n", c, size)
}

/*
func (b *Reader) ReadLine() (line []byte, isPrefix bool, err error)

ReadLine尝试返回一行数据，不包括行尾标志的字节。
如果行太长超过了缓冲，返回值isPrefix会被设为true，并返回行的前面一部分
该行剩下的部分将在之后的调用中返回，返回值isPrefix会在返回该行最后一个片段时才设为false
返回切片是缓冲的子切片，只在下一次读取操作之前有效
ReadLine要么返回一个非nil的line，要么返回一个非nil的err，两个返回值至少一个非nil

返回的文本不包含行尾的标志字节"\r\n" or "\n"
如果输入流结束时没有行尾标志字节，方法不会出错，也不会支出这一情况
在调用ReadLine之后调用UnreadByte会总是吐出最后一个读取的字节（很可能是该行的行尾标志字节），即使该字节不是ReadLine返回值的一部分。

*/

func TestReadLine(t *testing.T) {
	//create reader
	r := strings.NewReader("ABC\nDEF\r\nGHI\r\nJKL")
	// CREATE BUFIO READER
	br := bufio.NewReader(r)
	// read line
	w, isPrefix, _ := br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix)
	// read line
	w, isPrefix, _ = br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix)
	// read line
	w, isPrefix, _ = br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix)
	// read line
	w, isPrefix, _ = br.ReadLine()
	fmt.Printf("%q %v\n", w, isPrefix)
}

/*
func (b *Reader) ReadSlice(delim byte) (line []byte, err error)

ReadSlice读取直到第一次遇到delim字节，返回缓冲里的包含已读取的数据和delim字节的切片
该返回值只在下一次读取操作之前合法。如果ReadSlice放在读取到delim之前遇到了错误，它会返回在错误之前读取的数据在缓冲中的切片以及该错误（一般是io.EOF）
如果在读取delim之前缓冲就写满了，ReadSlice失败并返回ErrBufferFull。
因为ReadSlice的返回值会被下一次I/O操作重写，调用者应尽量使用ReadBytes或ReadString替代本方法。
当且仅当ReadBytes方法返回的切片不以delim结尾时，会返回一次非nil的错误。
*/

func TestReadSlice(t *testing.T) {
	// 创建一个Reader
	r := strings.NewReader("ABC,DEF,GHI,JKL")
	// 创建一个bufio Reader
	br := bufio.NewReader(r)
	// 使用ReadSlice读取
	w, _ := br.ReadSlice(',') // 当把,换成\n时就是ReadLine
	fmt.Printf("%q\n", w)
	// 使用ReadSlice读取
	w, _ = br.ReadSlice(',')
	fmt.Printf("%q\n", w)
	// 使用ReadSlice读取
	w, _ = br.ReadSlice(',')
	fmt.Printf("%q\n", w)

}

/*
func (b *Reader) ReadString(delim byte) (line string, err error)

ReadString读取直到第一次遇到delim字节，返回一个包含已读取的数据和delim字节的字符串。
如果ReadString方法在读取到delim之前遇到了错误，返回在错误之前读取的数据以及该错误（一般是io.EOF）。
当且仅当ReadString方法返回的切片不以delim结尾时，会返回一个非nil的错误。
*/

func TestReadString(t *testing.T) {
	// 创建reader
	r := strings.NewReader("ABC DEF GHI JKL")
	// 创建bufio reader
	br := bufio.NewReader(r)
	// 读取一个字符串
	w, _ := br.ReadString(' ')
	fmt.Printf("%q\n", w)
	// 读取一个字符串
	w, _ = br.ReadString(' ')
	fmt.Printf("%q\n", w)
	// 读取一个字符串
	w, _ = br.ReadString(' ')
	fmt.Printf("%q\n", w)
}

/*
func (b *Reader) WriteTo(w io.Writer) (n int64, err error)

WriteTo方法实现了io.WriterTo接口。
*/

func TestWriteTo(t *testing.T) {
	// 创建reader.
	r := strings.NewReader("ABCDEFGHIJKLMN")
	// 创建bufio reader.
	br := bufio.NewReader(r)
	// 创建一个buffer.
	b := bytes.NewBuffer(make([]byte, 0))
	// Write to buffer b.
	br.WriteTo(b)
	// 打印
	fmt.Printf("b: %s\n", b)
}

/*
3. type Writer

3.1 基础介绍

bufio.Writer 是 bufio 中对 io.Writer 的封装。

type Writer struct {
	err		error
	buf		[]byte
	n		int
	wr		io.Writer
}

bufio.Write(p []byte) 的思路如下：
1. 判断buf中可用容量是否可以放下p
2. 如果能放下，直接把p拼接到buf的后面，即把内容放到缓冲区
3. 如果缓冲区的可用容量不足以放下，且此时缓冲区是空的，直接把p写入文件即可
4. 如果缓冲区的可用容量不足以放下，且此时缓冲区有内容，则用p把缓冲区填满，把缓冲区所有内容写入文件，请清空缓冲区
5. 判断p的剩余内容大小能否放到缓冲区，如果能放下（此时和步骤1情况一样）则把内容放到缓冲区
6. 如果p的剩余内容依旧大于缓冲区，（注意此时缓冲区是空的，情况和步骤3一样）则把p的剩余内容直接写入文件
*/

/*
方法

func NewWriter(w io.Writer) *Writer

NewWriter创建一个具有默认大小缓冲、写入w的*Writer。
NewWriter相当于NewWriterSize(wr, 4096).

func NewWriterSize(w io.Writer, size int) *Writer

NewWriterSize创建一个具有最少有size尺寸的缓冲、写入w的Writer。
如果参数w已经是一个具有足够大缓冲的Writer类型值，会返回w。

func (b *Writer) Reset(w io.Writer)

Reset丢弃缓冲中的数据，清楚任何错误，将b重置并将其输出写入w。
*/

func TestWrite(t *testing.T) {
	// 创建一个缓冲区和Writer，并写入字符串
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)
	bw.WriteString("123456789")

	// 新创建一个缓冲区，将上面的Writer重设为新缓冲区，并写入字符串
	c := bytes.NewBuffer(make([]byte, 0))
	bw.Reset(c)
	bw.WriteString("456")

	// 将 writes any buffered data to the underlying io.Writer.
	bw.Flush()

	// 打印缓冲区
	fmt.Printf("b: %v\n", b)
	fmt.Printf("c: %v\n", c)
}

/*
func (b *Writer) Buffered() int

Buffered返回缓冲中已使用的字节数


func (b *Writer) Available() int

Available返回缓冲中还有多少字节未使用

func (b *Writer) Write(p []byte) (n int, err error)

Write将p的内容写入缓冲。返回写入的字节数。如果返回值 n < len，还会返回一个错误说明原因。

func (b *Writer) WriteString(s string) (int, error)

WriteString写入一个字符串。返回写入的字节数。如果返回值 n < len，还会返回一个错误说明原因。

func (b *Writer) WriteByte(c byte) error

WriteByte写入单个字节

func (b *Writer) WriteRune(r rune) (size int, err error)

WriteRune写入一个Unicode码值，返回写入的字节数和可能的错误

func (b *Writer) Flush() error

Flush方法将缓冲中的数据写入下层的io.Writer接口

func (b *Writer) ReadFrom(r io.Reader) (n int64, err error)

ReadFrom实现了io.ReadFrom接口
*/

func TestAvailableBuffered(t *testing.T) {
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)

	fmt.Printf("bw.Available(): %v\n", bw.Available()) // 4096
	fmt.Printf("bw.Buffered(): %v\n", bw.Buffered())   // 0

	bw.WriteString("ABCDEFGHIJKLMN")

	fmt.Printf("bw.Available(): %v\n", bw.Available()) // 4082
	fmt.Printf("bw.Buffered(): %v\n", bw.Buffered())   // 14
	fmt.Printf("b: %q\n", b)                           // ""，当前的数据还未flush

	bw.Flush()

	fmt.Printf("bw.Available(): %v\n", bw.Available()) // 4096
	fmt.Printf("bw.Buffered(): %v\n", bw.Buffered())   // 0
	fmt.Printf("b: %q\n", b)                           // 显示写入的内容

	// 创建另一个缓冲区
	b1 := bytes.NewBuffer(make([]byte, 0))
	r1 := strings.NewReader("Hello 世界！")
	bw1 := bufio.NewWriter(b1)
	bw1.ReadFrom(r1)
	// bw.Flush() // ReadFrom无需使用Flush，其自己已经写入
	fmt.Printf("b1: %v\n", b1)
}

/*
4. type ReadWriter

4.1 基本介绍

ReadWriter类型保管了指向Reader和Writer类型的指针，实现了io.ReadWriter接口

type ReadWriter struct {
	*Reader
	*Writer
}

4.2 方法

func NewReadWriter(r *Reader, w *Writer) *ReadWriter

NewReadWriter申请创建一个新的、将读写操作分派给r和w的ReadWriter
*/

func TestReadWriter(t *testing.T) {

	// 创建 writer
	b := bytes.NewBuffer(make([]byte, 0))
	bw := bufio.NewWriter(b)

	// 创建 reader
	r := strings.NewReader("123")
	br := bufio.NewReader(r)

	// 创建 readerWriter
	rw := bufio.NewReadWriter(br, bw)

	// 根据换行符读
	p, _ := rw.ReadString('\n')

	// 打印读出的数据
	fmt.Printf("string(p): %v\n", string(p))

	// 写入数据，然后刷盘
	rw.WriteString("asdf")
	rw.Flush()

	// 打印
	fmt.Printf("b: %v\n", b)
}

/*
5. type SplitFunc

type SplitFunc func(data []byte, atEOF bool) (advance int, token []byte, err error)

SplitFunc 类型代表用于对输出做词法分析的分割函数

参数data是尚未处理的数据的一个开始部分的切片，参数atEOF表示是否Reader接口不能提供更多的数据。
返回值是解析位置前进的字节数，将要返回给调用者的token切片，以及可能遇到的错误。
如果数据不足以（保证）生成一个完成的token，例如需要一整行数据但data里没有换行符，SplitFunc可以返回(0, nil, nil)来告诉Scanner读取更多的数据，
写入切片，然后用从同一位置起始、长度更长的切片再试一次（调用SplitFunc类型函数）。

如果返回值err非nil，扫描将终止并将该错误返回给Scanner的调用者。

除非atEOF为真，永远不会使用空切片data调用SplitFunc类型函数。然而，如果atEOF为真，data却可能是非空的、且包含着未处理的文本。

SplitFunc的作用很简单，从data中找出你感兴趣的数据，然后返回并告诉调用者，data中有多少数据你已经处理过了。

*/

func TestSplitFunc(t *testing.T) {
	// 打开文件
	file, err := os.Open("./resources/a.txt")
	if err != nil {
		log.Fatal(err)

	}
	// 关闭文件
	defer file.Close()

	// 创建扫描器
	fileScanner := bufio.NewScanner(file)
	// 按照单词切分扫描
	fileScanner.Split(bufio.ScanWords)
	// 进行扫描，并把扫描的文本内容打印
	for fileScanner.Scan() {
		fmt.Printf("fileScanner.Text(): %v\n", fileScanner.Text())
	}
}

/*
6. type Scanner

6.1 基本介绍

type Scanner struct {
	r				io.Reader		// The reader provided by the client
	split			SplitFunc		// The function to split the token
	maxTokenSize	int				// Maximum size of a token; modified by tests
	token			[]byte			// Last token returned by split
	buf				[]byte			// Buffer used an argument to split
	start			int				// Fist non-processed byte in buf
	end				int				// End of data in buf
	err				error			// Sticky error
}

Scanner类型提供了方便的读取数据的接口，如从换行符分隔的文本里读取每一行。成功调用的Scann方法会逐步提供文件的token，跳过token之间的字节。
token有SplitFunc类型的分割函数指定；默认的分割函数会将输入分割为多个行，并去掉行尾的换行标志。

本包预定义的分割函数可以将文件分割为行、字节、unicode码值、空白分隔的word。
调用者可以定制自己的分割函数。
扫描会在抵达输入流结尾、遇到的第一个I/O错误、token过大不能保存进缓冲时，不可恢复的停止。
当扫描停止后，当前读取位置可能会在最后一个获得的token后面。
需要更多对错误管理的控制或token很大，或必须从reader连续扫描的程序，应使用buffio.reader代替。
*/

/*
6.2 方法

func ScanBytes(data []byte, atEOF bool) (advance int, token []byte, err error)

ScanBytes是用于Scanner类型的分割函数（符合SplitFunc），本函数会将每个字节作为一个token返回。


func ScanRunes(data []byte, atEOF bool) (advance int, token []byte, err error)

ScanRunes是用于Scanner类型的分割函数（符合SplitFunc），本函数将每个UTF-8编码的unicode码值作为一个token返回。
本函数返回的rune序列和range一个字符串的输出rune序列相同。
错误的UTF-8编码会翻译为U+FFFD="\xef\xbf\xbd"，但只会消耗一个字节。
调用者无法区分正确编码的rune和错误编码的rune


func ScanWords(data []byte, atEOF bool) (advnace int, token []byte, err error)

ScanRunes是用于Scanner类型的分割函数（符合SplitFunc），本函数会将空白分隔的片段（去掉前后空白后）作为一个token返回。
本函数永远不会返回空字符串。
用来找出data中的单行数据并返回（包括空行）。


func ScanLines(data []byte, atEOF bool) (advance int, token []byte, err error)

ScanLines是用于Scnanner类型的分割函数（符合SplitFunc），本函数会将每一行文本去掉末尾的换行标记作为一个token返回。
返回的行可以是空字符串。
换行标记为一个可选的回车后跟一个必选的换行符。
最后一行即使没有换行符也会作为一个token返回。

func NewScanner(r io.Reader) *Scanner

NewScanner创建并返回一个从r读取数据的Scanner，默认的分割函数是ScanLines。


func (s *Scanner) Split(split SplitFunc)

Split设置该Scanner的分割函数。本方法必须在Scan之前调用。


*/

func TestScanSplit(t *testing.T) {
	// 创建reader
	r := strings.NewReader("ABC DEF GHI JKL")

	// 创建scanner
	bs := bufio.NewScanner(r)

	// 按照单词（以空格分隔）划分
	bs.Split(bufio.ScanWords)
	// 扫描每一个单词并打印
	for bs.Scan() {
		fmt.Printf("bs.Text(): %v\n", bs.Text())
	}
}

/*
func (s *Scanner) Scan() bool

Scan方法获取当前位置的token（该token可以通过Bytes或Text方法获得），并让Scanner的扫描位置移动到下一个token。
当扫描因为抵达输入流结尾或者遇到错误而停止时，本方法会返回false。
在Scan方法返回false后，Err方法将返回扫描时遇到的任何错误；除非是io.EOF，此时Err会返回nil。
*/
func TestScan(t *testing.T) {
	// 创建一个reader
	r := strings.NewReader("Hello 中国！")
	// 创建一个scanner
	bscanner := bufio.NewScanner(r)
	// 扫描时划分字符的方法选择按照字节
	bscanner.Split(bufio.ScanBytes)
	// 开始扫描，并打印扫描的结果
	for bscanner.Scan() {
		fmt.Printf("bscanner.Text(): %q\n", bscanner.Text())
	}

}

/*
func (s *Scanner) Bytes() []byte

Bytes方法返回最近一次Scan调用生成的token。
底层数组指向的数据可能会被下一次Scan的调用重写。

*/

func TestScanRune(t *testing.T) {
	// 创建一个reader
	r := strings.NewReader("hello，西双版纳")
	// 创建一个Scanner
	bscanner := bufio.NewScanner(r)
	// 定义扫描的分割方式，按照中文字符
	bscanner.Split(bufio.ScanRunes)
	// 开始扫描并打印字符
	for bscanner.Scan() {
		fmt.Printf("bscanner.Text(): %v\n", bscanner.Text())
	}

}

/*
func (s *Scanner) Text() string

Bytes方法返回最近一次Scan调用生成的token，会申请创建一个字符串保存token并返回该字符串。

func (s *Scanner) Err() error

Err返回Scanner遇到的第一个非EOF的错误。
*/
