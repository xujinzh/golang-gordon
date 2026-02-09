package iostl_test

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"testing"
)

// io标准库

// io包中提供I/O原始操作的一系列接口。它主要包装了一些已有的实现，如os包中的那些，并将这些抽象成为实用性的功能和一些其他相关的接口。

// 1. 错误变量

// var EOF = errors.New("EOF")
// 正常输入结束Read返回EOF，如果在一个结构化数据流中EOF在不期望的位置出现了，则应返回错误ErrUnexpectedEOF或者其他给出更多细节的错误

// var ErrClosedPipe = errors.New("io: read/write on closed pipe")
// 当从一个已关闭的Pipe读取或写入时，会返回ErrClosedPipe

// var ErrNoProgress = errors.New("multiple Read calls return no data or error")
// 某些使用io.Reader接口的客户端如果多次调用Read都不返回数据也不返回错误时，就会返回本错误，一般来说是io.Reader实现有问题的标志

// var ErrShortBuffer = errors.New("short buffer")
// ErrShortBuffer表示读取操作需要大缓冲，但提供的缓冲不够大

// var ErrShortWrite = errors.New("short write")
// ErrShortWrite表示写入操作写入的数据比提供的少，却没有显示的返回错误

// var ErrUnexpectedEOF = errors.New("unexpected EOF")
// ErrUnexpectedEOF表示在读取一个固定尺寸的块或者数据结构时，在读取为完全时遇到了EOF

// 2. 基础接口

// Reader接口
/*
type Reader interface {
	Read(p []byte) (n int, err error)
}

Read将len个字节读取到p中
它返回读取的字节数n (0 <= n <= len) 以及任何遇到的错误。
即使Read返回的n < len，它也会在调用过程中使用p的全部作为暂存空间。
若一些数据可用但不到len个字节，Read会照例返回可用的东西，而不是等待更多。

当Read在成功读取n > 0个字节后遇到一个错误或EOF情况，它就会返回读取的字节数，它会从相同的调用中返回（非nil的）错误，或从随后的调用中返回错误（和n==0）
这种一般情况的一个例子就是Reader在输入流结束时会返回一个非零的字节数，可能的返回不是 err == EOF 就是 err == nil。
无论如何，下一个Read都应当返回0，EOF

调用者应当总在考虑到错误err前处理 n > 0 的字节。这样做可以在读取一些字节，以及允许的EOF行为
*/

func TestReader(t *testing.T) {
	// 创建一个文件
	f, err := os.Open("./resources/a.txt")
	if err != nil {
		t.Log(err)
	}
	// 关闭文件
	defer f.Close()

	// 实例化一个长度为4的[]byte
	// 缓冲
	buf := make([]byte, 4)
	// 读取的内容汇总
	var body []byte
	// 读取
	for {
		n, err2 := f.Read(buf)
		if n == 0 || err2 == io.EOF {
			fmt.Println("文件读取完毕")
			break
		}
		body = append(body, buf[:n]...)
	}
	fmt.Println(string(body))
	fmt.Println("\033[0;31m=============\033[0m")
}

// Writer接口

/*
type Writer interface {
	Writer(p []byte) (n int, err error)
}

Write将len个字节从p中写入到基本数据流中。
它返回从p中被写入的字节数n(0 <= n <= len) 以及任何遇到的引起写入提前停止的错误
若Write返回的 n < len，它就必须返回一个非nil的错误
Write不能修改此切片的数据，即便它是临时的
*/

func TestWriter(t *testing.T) {
	// 以读写模式打开文件，并且在写操作室将数据附件到文件末尾
	f, err := os.OpenFile("./resources/b.txt", os.O_CREATE|os.O_RDWR|os.O_APPEND, 0775)
	if err != nil {
		t.Log(err)
	}
	// 关闭文件
	defer f.Close()
	// 写入数据
	f.Write([]byte(" hello, golang"))
}

// Seeker 接口

/*
type Seeker interface {
	Seek(offset int64, whence int) (int64, error)
}

Seeker用来移动数据的读写指针
Seek设置下一次读写操作的指针位置，每次的读写操作都是从指针位置开始

whence的含义
- 如果whence为0：表示从数据的开头开始移动指针
- 如果whence为1：表示从数据的当前指针位置开始移动指针
- 如果whence为2：表示从数据的尾部开始移动指针

offset是指针移动的偏移量

返回移动后的指针位置和移动过程中遇到的任何错误。
*/

func TestSeeker(t *testing.T) {
	// 打开文件，指针默认在文件开头
	f, err := os.Open("./resources/a.txt")
	if err != nil {
		t.Log()
	}
	// 关闭文件
	defer f.Close()
	// 设置指针光标位置
	f.Seek(7, 0)
	// 设置缓冲区
	buf := make([]byte, 10)
	// 读取内容到缓冲区
	n, err := f.Read(buf)
	if err != nil {
		t.Log(err)
	}
	// 打印输出
	fmt.Printf("n: %v\n", n)
	fmt.Printf("string(buf): %q\n", string(buf))
}

// Closer接口

/*
type Closer interface {
	Close() error
}

Closer关闭的接口，带有Close()方法，但是行为没有定义，所以可以特定行为来实现
在整个标准库内都没有对Closer的引用，只有实现，用法都是开启某连接、流，在用完、报错后进行Close操作
*/

// 组合接口

// 组合接口是对多个接口进行组合，当同时实现多个接口时，可以使用组合接口进行传递

// 3.1 ReadWriter接口
/*
type ReadWriter interface {
	Reader
	Writer
}

ReadWriter接口聚合了基本的读写操作
*/

// 3.2 ReadCloser接口
/*
type ReadCloser interface {
	Reader
	Closer
}

ReadCloser就是Reader+Closer，例如在ioutil中的NopCloser方法返回的就是一个ReadCloser，但是里面的Closer就是个空函数，毫无作用
*/

// 3.3 WriteCloser接口
/*
type WriteCloser interface {
	Writer
	Closer
}

WriteCloser接口聚合了基本的写入和关闭操作
*/

// 3.4 ReadWriteCloser接口
/*
type ReadWritCloser interface {
	Reader
	Writer
	Closer
}

ReadWriteCloser接口聚合了基本的读写和关闭操作
*/

// 3.5 ReadSeeker接口
/*
type ReadSeeker interface {
	Reader
	Seeker
}

ReadSeeker接口聚合了基本的读取和移位操作
*/

// 3.6 WriteSeeker接口
/*
type WriteSeeker interface {
	Writer
	Seeker
}

WriteSeeker接口聚合了基本的写入和移位操作
*/

// 3.7 ReadWriteSeeker接口
/*
type ReadWriteSeeker interface {
	Reader
	Writer
	Seeker
}

ReadWriteSeeker接口聚合了基本的读写和移位操作
*/

// 4. 指定读写器读写接口

// 4.1 ReaderFrom接口
/*
type ReaderFrom interface {
	ReadFrom(r Reader) (n int64, err error)
}

ReadFrom从r中读取数据到对象的数据流中
直到r返回EOF或r出现读取错误为止
返回值n是读取的字节数
返回值err就是r的返回值err
*/

// 4.2 WriterTo接口
/*
type WriterTo interface {
	WriterTo(w Writer) (n int64, err error)
}

WriterTo将对象的数据流写入到w中
直到对象的数据流全部写入完毕或遇到写入错误为止
返回值n是写入的字节数
返回值err就是w的返回值err
*/

// 5. 指定偏移量读写接口

// 5.1 ReaderAt接口
/*
type ReaderAt interface {
	ReadAt(p []byte, off int64) (n int, err error)
}

ReadAt 从对象数据流的off处读出数据到p中
- 忽略数据的读写指针，从数据的起始位置偏移off处开始读取
- 如果对象的数据流只有部分可用，不足以填满p则ReadAt将等待所有数据可用之后，继续向p中写入直到将p填满后返回，在这点上ReadAt要比ReadAt更严格
- 返回读取的字节数n和读取是遇到的错误
- 如果n<len，则需要返回一个err值来说明为什么没有将p填满（比如EOF）
- 如果n>len，而且对象的数据没有全部读完，则err将返回nil
- 如果n=len，而且对象的数据刚好全部读完，则err将返回EOF或者nil（不确定）
*/

// 5.2 WriterAt接口
/*
type WriterAt interface {
	WriterAt(p []byte, off int64) (n int, err error)
}

WriterAt将p中的数据写入到对象数据流的off处
- 忽略数据的读写指针，从数据的起始位置偏移off处开始写入
- 返回写入的字节数和写入时遇到的错误
- 如果n<len，则必须返回一个err值来说明为什么没有将p完全写入
*/

// 6 单个字节读写接口

// 6.1 ByteReader接口
/*
type ByteReader interface {
	ReadByte() (byte, error)
}

ByteReader是基本的ReadByte方法的包装
ReadByte读取输入中的单个字节并返回
如果没有字节可读取，会返回错误
*/

// 6.2 ByteScanner接口

/*
type ByteScanner interface {
	ByteReader
	UnreadByte() error
}

ByteScanner接口在基本的ReadByte方法之外还添加了UnreadByte方法
UnreadByte方法让下一次调用ReadByte时返回之前调用ReadByte时返回的同一个字节
连续调用两次UnreadByte方法而中间没有调用ReadByte时，可能会导致错误
*/

// 6.3 ByteWriter接口

/*
type ByteWriter interface {
	WriteByte(c byte) error
}

包装WriterByte单个字节写入方法的接口
*/

// 6.4 RuneReader接口

/*
type RuneReader interface {
	ReadRune() (r rune, size int, err error)
}

ReadRune方法的包装，读取单个UTF-8编码的Unicode字符，并返回rune及其字节大小。如果没有可用字符，将设置err
*/

// 6.5 RuneScanner接口

/*
type RuneScanner interface {
	RuneReader
	UnreadRune() error
}

RuneScanner接口在基本的ReadRune方法之外还添加了UnreadRune方法
UnreadRune方法让下一次调用ReadRune时返回之前调用ReadRune时返回的同一个UTF-8字符
连续调用两次UnreadRune方法而中间没有调用ReadRune时，可能会导致错误
*/

// 6.6 StringWriter接口
/*
type StringWriter interface {
	WriteString(s string) (n int, err error)
}

字符串写入方法WriteString的包装
*/

// 7 结构体

// 7.1 LimitedReader
/*
type LimitedReader struct {
	R	Reader	// underlying reader
	N 	int64	// max bytes remaining
}

LimitedReader从R读取，但将返回的数据量限制为N个字节。
每次读取更新N以标记剩余可以读取的字节数。
Read在N<=0时或基础R返回EOF时返回EOF

具体实现方法为：func LimitReader(r Reader, n int64) Reader
*/

// 7.2 PipeReader

/*
type PipeReader struct {
	// 内含隐藏或非导出字段

}

PipeReader是一个管道的读取端

具体实现方法有：
- func (r *PipeReader) Read(data []byte) (n int, err error) Read实现了标准的读取接口：它从管道中读取数据，阻塞直到写入端到达或写入端被关闭。如果用错误关闭写入端，则返回错位为ERR；否则ERR为EOF
- func (r *PipeReader) Close() error Close关闭读取器：关闭后如果对管道的写入端进行写入操作，就会返回（0， ErrClosedPipe）
*/

// 8. 供外部调用的函数

// 8.1 Copy
/*
func Copy(dst writer, src Reader) (written int64, err error)
将副本从src复制到dst，直到在src上达到EOF或发生错误
它返回复制的字节数和复制时遇到的第一个错误（如果有）
成功的复制将返回err==nil而不是err==EOF，
因为复制被定义为从src读取直到EOF，所以它不会讲读取的EOF视为要报告的错误
如果src实现WriterTo接口，则通过调用src.WriterTo(dst)实现该副本
否则， 如果dst实现了ReaderFrom接口，则通过调用dst.ReadFrom(src)实现该副本
*/

func TestCopy(t *testing.T) {
	// 读取字符串
	r := strings.NewReader("some io.Reader stream to be read\n")
	// 拷贝
	if _, err := io.Copy(os.Stdout, r); err != nil {
		// os.Stdout将内容输出到控制台
		log.Fatal(err)

	}

	// 拷贝文件内容，赋值到另一个文件中
	fmt.Printf("\033[0;32m=======================\033[0m\n")
	// 原文件
	src, err := os.Open("./resources/b.txt")
	if err != nil {
		t.Log(err)
	}
	// 关闭原文件
	defer src.Close()
	// 目标文件
	dst, err := os.OpenFile("./resources/c.txt", os.O_CREATE|os.O_RDWR, 0644)
	if err != nil {
		t.Log(err)
	}
	// 关闭目标文件
	defer dst.Close()
	// 拷贝原文件内容到目标文件
	if _, err := io.Copy(dst, src); err != nil {
		t.Fatal(err)
	}

}

// 8.2 CopyBuffer

/*
func CopyBuffer(dst Writer, src Reader, buf []byte) (written int64, err error)

CopyBuffer和Copy相同，区别在于CopyBuffer逐步遍历提供的缓冲区（如果需要），而不是分配临时缓冲区
如果buf为nil，则分配一个；如果长度为0，则CopyBuffer会panic报错
如果src实现WriterTo或dst实现ReaderFrom，则buf将不用于执行复制
*/

func TestCopyBuffer(t *testing.T) {
	// 创建reader
	r1 := strings.NewReader("first reader\n")
	r2 := strings.NewReader("second reader\n")
	buf := make([]byte, 4)
	// 使用buf拷贝
	if _, err := io.CopyBuffer(os.Stdout, r1, buf); err != nil {
		log.Fatal(err)
	}

	if _, err := io.CopyBuffer(os.Stdout, r2, buf); err != nil {
		log.Fatal(err)
	}
}

// 8.3 CopyN

/*
func CopyN(dst Writer, src Reader, n int64) (written int64, err error)

CopyN将n个字节（或直接出错）从src复制到dst
返回复制的字节数和复制时遇到的最早错误
返回时，只有err==nil时，written==n
如果dst实现了ReaderFrom接口，则使用该接口实现副本
*/

func TestCopyN(t *testing.T) {
	// 创建reader
	r3 := strings.NewReader("some io.Reader stream to be read")
	// 拷贝4个字节
	if _, err := io.CopyN(os.Stdout, r3, 4); err != nil {
		log.Fatal(err)
	}
}

// 8.4 LimitReader

/*
func LimitReader(r Reader, n int64) Reader

LimitReader从r读取，但将返回的数据量限制为n个字节
每次读取更新n以标记剩余可以读取的字节数
Read在n<=0时或基础r返回EOF时返回EOF
*/

func TestLimitReader(t *testing.T) {
	// 创建reader
	r := strings.NewReader("some io.Reader stream to be read\n")

	lr := io.LimitReader(r, 4)

	// 拷贝
	if _, err := io.Copy(os.Stdout, lr); err != nil {
		log.Fatal(err)
	}
}

// 8.5 MultiReader

/*
func MultiReader(readers ...Reader) Reader

MultiReader返回一个Reader，它是所提供的输入阅读器的逻辑串联
它们被顺序读取，一旦所有输入均返回EOF，读取将返回EOF
如果任何读取器返回非零，非EOF错误，则Read将返回该错误
*/

func TestMultiReader(t *testing.T) {
	// 创建多个reader
	r1 := strings.NewReader("first reader ")
	r2 := strings.NewReader("second reader ")
	r3 := strings.NewReader("third reader\n")

	// 创建multireader
	r := io.MultiReader(r1, r2, r3)

	// 拷贝
	if _, err := io.Copy(os.Stdout, r); err != nil {
		t.Fatal(err)
	}
}

// 8.6 MultiWriter

/*
func MultiWriter(writers ...Writer) Writer

MultiWriter创建一个Writers，将其写入复制到所有提供的写入器中，类似于Unix tee命令
每个写入一次写入每个列出的写入器
如果列出的写程序返回错误，则整个写操作将停止并返回错误
它不会在列表中继续下去
*/
func TestMultiWriter(t *testing.T) {
	// 创建reader
	r := strings.NewReader("some io.Reader stream to be read \n")

	// 创建多个writer
	var buf1, buf2 bytes.Buffer

	w := io.MultiWriter(&buf1, &buf2)

	// 拷贝
	if _, err := io.Copy(w, r); err != nil {
		t.Fatal(err)
	}

	// 打印写入的结果
	fmt.Printf("buf1.String(): %v\n", buf1.String())
	fmt.Printf("buf2.String(): %v\n", buf2.String())
}

// 8.7 Pipe

/*
func Pipe() (*PipeReader, *PipeWriter)

Pipe创建一个同步的内存管道
可用于连接期望io.Reader的代码和期望io.Writer的代码

管道上的读和写是一对一匹配的，除非需要多次读取才能使用单词写入。也就是说，每次对PipeWriter的写入都将阻塞，直到它满足从PipeReader读取的一个或多个读取，这些读取会完全消耗已写入的数据。

数据直接从Write复制到相应的Read（或Reads）；没有内部缓冲

对读的并行调用和对写的并行调用也是安全的：单个调用将按顺序执行
*/

func TestPip(t *testing.T) {
	// 创建管道
	r, w := io.Pipe()
	// 写入数据
	go func() {
		fmt.Fprint(w, "some io.Reader stream to be read\n")
		w.Close()
	}()

	// 读取数据
	if _, err := io.Copy(os.Stdout, r); err != nil {
		t.Fatal(err)
	}

}

// 8.8 ReadAll

/*
func ReadAll(r Reader) ([]byte, error)

ReadAll从r读取，直到出现错误或EOF，并返回其读取的数据。成功的调用返回err==nil，而不是err==EOF。
由于ReadAll定义为从src读取直到EOF，因此它不会将读取的EOF视为要报告的错误
*/

func TestReadAll(t *testing.T) {
	// 创建reader
	r := strings.NewReader("go is a general-purpose language designed with systems programming in mind")
	// 读取所有数据
	b, err := io.ReadAll(r)
	if err != nil {
		t.Log(err)

	}
	fmt.Printf("b: %q\n", b)
}

// 8.9 ReadAtLeat

/*
func ReadAtLeast(r Reader, buf []byte, min int) (n int, err error)

ReadAtLeast从r读取到buf，直到它至少读取了min字节
它返回复制的字节数n，如果读取的字节数少则返回错误
仅当未读取任何字节时，错误才是EOF
如果在读取少于最小字节后发生EOF，则ReadAtLeast返回ErrUnexpectedEOF
如果min大于buf的长度，则ReadAtLeast返回ErrShortBuffer
返回时，当且仅当err==nil时，n >= min
*/

func TestReadAtLeast(t *testing.T) {

	r := strings.NewReader("some io.Reader stream to be read \n")

	buf := make([]byte, 4)

	if _, err := io.ReadAtLeast(r, buf, 4); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("buf: %q\n", buf)

	// buffer smaller than minimal read size
	shortBuf := make([]byte, 3)

	if _, err := io.ReadAtLeast(r, shortBuf, 4); err != nil {
		fmt.Printf("err: %v\n", err)

	}

}


// 8.10 ReadFull

/* 
func ReadFull(r Reader, buf []byte) (n int, err error)

ReadFull将r中的len(buf)个字节准确地读取到buf中
它返回复制的字节数，如果读取的字节数少则返回错误。
仅当未读取任何字节时，错误才是EOF
如果在读取了一些但不是全部字节后发生EOF，则ReadFull返回ErrUnexpectedEOF
返回时，当且仅当err==nil时，n==len(buf)
*/

func TestReadFull(t *testing.T) {

	r := strings.NewReader("some io.Reader stream to be read\n")

	buf := make([]byte, 4)

	if _,err  := io.ReadFull(r, buf); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("buf: %q\n", buf)

	// minimal read size bigger than io.Reader stream
	longBuf := make([]byte, 64)

	if _, err := io.ReadFull(r, longBuf); err != nil{
		log.Fatal(err)
	}
}