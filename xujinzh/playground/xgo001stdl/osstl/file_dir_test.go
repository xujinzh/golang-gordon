package osstl

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"testing"
	"time"
)

// panic 介绍
// 主动抛出不可恢复的严重错误，导致程序中断执行。仅用于不可预测的重大错误（如系统运行状态不一致、严重配置错误）。普通业务错误应优先使用error
// 配合defer, recover可以在程序崩溃前进行必要的清理工作。
// func main() {
// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("恢复了:", r) // 捕获并处理panic
// 		}
// 	}()

// 	panic("发生严重错误") // 主动触发
// 	fmt.Println("这行代码不会执行")
// }

// 测试创建一个文件
func TestCreate(t *testing.T) {
	// 创建一个文件。内部就是调用：OpenFile(name, O_RDWR|O_CREATE|O_TRUNC, 0666)
	// 1. 文件不存在会创建一个新文件，O_CREATE
	// 2. 文件存在会先清空文件再创建新文件，O_TRUNC
	// 特别注意，当待创建的文件路径名的上一层目录不存在，那么创建失败，报 no such file or directory
	f, err := os.Create("./assets/test.txt")
	if err != nil { // 如果创建失败，报异常
		// 触发panic
		panic(err)
	}
	// 记得关闭文件
	defer f.Close()
	// 最后一个defer，在函数运行完后先执行
	defer fmt.Printf("程序运行结束。\n")
}

// 测试创建单级目录
func TestMkdir(t *testing.T) {
	// 创建一个文件夹，而不是多级目录。注意非递归创建，即创建的文件夹上一层目录必须存在
	err := os.Mkdir("./assets", os.ModePerm)
	// 如果创建失败，那么报异常信息
	if err != nil {
		panic(err)
	}
}

// 测试创建多级目录
func TestMkdirAll(t *testing.T) {
	// 递归创建多级目录
	err := os.MkdirAll("./assets/a/b", os.ModePerm)
	// 如果创建失败，那么报异常信息
	if err != nil {
		panic(err)
	}
}

// 测试删除文件或空目录
func TestRemove(t *testing.T) {
	// 删除空目录
	err := os.Remove("./assets/test")
	// 如果删除空目录失败
	if err != nil {
		t.Log("删除空目录失败：", err)
	}
	// 删除文件
	err = os.Remove("./assets/test.txt")
	// 如果删除文件失败
	if err != nil {
		t.Log("删除文件失败：", err)
	}

}

// 测试删除多级目录和目录中的文件
func TestRemoveAll(t *testing.T) {
	// 删除多级目录和目录中的文件
	err := os.RemoveAll("./assets/a")
	// 如果删除多级目录和目录中的文件失败
	if err != nil {
		t.Log("删除目录失败：", err)
	}
}

// 测试获取当前路径
func TestGetwd(t *testing.T) {
	// 获取当前路径
	wd, err := os.Getwd()
	// 如果失败，打印错误信息
	if err != nil {
		t.Log("获取当前路径失败：", err)
	}
	// 打印当前路径
	fmt.Printf("当前工作目录: %v\n", wd)

}

// 测试切换目录
func TestChdir(t *testing.T) {
	// 尝试切换到指定目录
	err := os.Chdir("/opt")
	// 如果切换失败
	if err != nil {
		t.Log("切换目录失败：", err)
	} else { // 切换成功
		t.Log("切换目录成功。")
	}
	// 获取切换后的工作目录
	wd, err := os.Getwd()
	// 如果获取工作目录失败
	if err != nil {
		t.Log("获取当前目录失败：", err)
	}
	// 打印当前目录
	fmt.Printf("wd: %v\n", wd)
	// fmt.Println(os.Getwd())
}

// 测试获取临时目录
func TestTempDir(t *testing.T) {
	// 获取存放临时文件的目录，不同操作系统结果可能不同
	tmp := os.TempDir()
	fmt.Printf("tmp: %v\n", tmp)
}

// 测试重命名文件
func TestRename(t *testing.T) {
	// // 重命名文件
	// err := os.Rename("./assets/test.txt", "./assets/text.txt")
	// // 捕获重命名失败的异常
	// if err != nil {
	// 	t.Log("重命名文件失败：", err)
	// }
	// 重命名文件夹
	err := os.Rename("./assets", "./resources")
	if err != nil { // 如果失败
		t.Log("重命名文件夹失败。")
	} else { // 成功
		t.Log("重命名文件夹成功。")
	}
}

// 修改文件或文件夹权限
func TestChmod(t *testing.T) {
	// 尝试改变文件的权限
	err := os.Chmod("./resources/text.txt", 0o750)
	if err != nil { // 如果失败
		t.Log("改变文件权限失败:", err)
	} else { // 成功
		t.Log("改变文件权限成功。")
	}
}

// 测试改变文件的所有者和所有组
func TestChown(t *testing.T) {
	err := os.Chown("./resources/text.txt", 0, 0)
	if err != nil {
		t.Log("改变文件所有者和所有组失败：", err)
	} else {
		t.Log("成功")
	}

}

// 测试文件统计
func TestFileStat(t *testing.T) {
	// 先打开文件
	f, err := os.OpenFile("./resources/text.txt", os.O_RDWR|os.O_CREATE, 0755)
	// 如果打开文件失败
	if err != nil {
		t.Log(err)
	}
	// 函数退出后关闭文件
	defer f.Close()
	// 文件信息统计
	fileInfo, err := f.Stat()
	if err != nil {
		t.Log(err)
	}
	// 打印文件信息
	fmt.Printf("fileInfo.Size(): %v\n", fileInfo.Size())
	fmt.Printf("fileInfo: %v\n", fileInfo)

}

// 测试读取文件
func TestRead(t *testing.T) {
	// 打开文件
	f, err := os.OpenFile("./resources/text.txt", os.O_RDONLY, 0755)
	// 如果读取失败
	if err != nil {
		t.Log(err)
	}
	// 开始读取文件内容
	// 创建存放文件内容的切片
	// var body = make([]byte, 0) // 注意，这里面不要初始化大于0的切片，否则前面默认填充0。或者只声明不分配内存。
	var body []byte
	// 循环读取，一次读取一小段
	for {
		// 存放每次读取的内容
		var buf = make([]byte, 4)
		// 读取一小段内容
		n, err := f.Read(buf)
		// 如果读取失败
		if err != nil {
			// 如果读取到文件末尾
			if err == io.EOF {
				body = append(body, buf[:n]...)
				break
			} else { // 如果读取出错

				t.Log(err)
				break
			}
		}
		body = append(body, buf[:n]...)
	}
	// 打印读取的内容
	fmt.Printf("body: %v\n", string(body))

}

// 测试从指定的位置开始读
func TestReadAt(t *testing.T) {
	// 先打开文件
	f, err := os.OpenFile("./resources/text.txt", os.O_RDONLY, 0755)
	// 如果读取失败
	if err != nil {
		t.Log(err)
	}
	// 别忘记关闭文件
	defer f.Close()
	// 设置读取的偏移量
	var offset = 7
	// 创建切片，存放读取的内容
	var buf = make([]byte, 6)
	// 从指定的位置开始读
	n, err := f.ReadAt(buf, int64(offset))
	// 如果读取失败
	if err != nil {
		t.Log("读取失败")
	}
	// 打印读取的内容
	fmt.Printf("buf: %s\n", buf)
	fmt.Printf("n: %v\n", n)
}

// 测试读取目录
func TestReadDir(t *testing.T) {
	// 打开文件夹
	f, err := os.Open(".")
	// 如果读取失败
	if err != nil {
		t.Log(err)
	}
	// 关闭文件
	defer f.Close()
	// 读取文件夹下的文件或子文件夹
	dirs, err := f.ReadDir(-1) // 读取所有
	// 如果读取失败
	if err != nil {
		t.Log(err)
	}
	// 打印读取的文件或子文件夹
	for _, v := range dirs {
		fmt.Println(v.Name(), "is dir or not:", v.IsDir())
	}
}

// 测试文件搜索读取
func TestSeek(t *testing.T) {
	// 打开文件
	f, err := os.OpenFile("./resources/text.txt", os.O_RDONLY, 0755)
	// 如果读取失败
	if err != nil {
		t.Log(err)
	}
	// 记得关闭文件
	defer f.Close()
	// 设置偏移量
	// whence: 0 means relative to the origin of the file, 1 means relative to the current offset, and 2 means relative to the end.
	nOffset, err := f.Seek(7, 0)
	// 如果设置偏移量失败
	if err != nil {
		t.Log(err)
	}
	// 设置后的偏移量
	fmt.Printf("nOffset: %v\n", nOffset)
	// 创建读取的内容存放的位置
	var buf = make([]byte, 10)
	// 读取所有内容
	n, err := f.Read(buf)
	// 如果读取失败
	if err != nil {
		t.Log(err)
	}
	// 打印读取的内容
	fmt.Printf("buf: %s\n", buf)
	fmt.Printf("n: %v\n", n)

}

// 测试写数据到文件
func TestWrite(t *testing.T) {
	// 打开一个文件
	f, err := os.OpenFile("./resources/text.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0755)
	// 如果打开失败
	if err != nil {
		t.Log(err)
	}
	// 记得关闭文件
	defer f.Close()
	// 写入字节数据
	f.Write([]byte("hello, golang []byte\n"))
	// 写入字符串
	f.WriteString("hello, golang string\n")
}

// 测试在文件的指定位置写入内容
func TestWriteAt(t *testing.T) {
	// 打开一个文件
	f, err := os.OpenFile("./resources/text.txt", os.O_RDWR|os.O_CREATE, 0755)
	// if open file error
	if err != nil {
		t.Log(err)
	}
	// close the f
	defer f.Close()
	// write info at the specified position, this write will conver write the info in file
	_, err = f.WriteAt([]byte("insert content\n"), 5)
	// if error, then print
	if err != nil {
		t.Log(err)
	}
}

// 打印进程相关信息
func TestOsInfo(t *testing.T) {
	fmt.Printf("\033[0;31m%s\n\033[0m", strings.Repeat("-", 33))
	// 获取当前正在运行的进程id
	fmt.Printf("os.Getpid(): %v\n", os.Getpid())
	// 获取当前正在运行的进程的父id
	fmt.Printf("os.Getppid(): %v\n", os.Getppid())
	fmt.Println(os.FindProcess(os.Getppid()))
	// 设置新进程的属性
	attr := &os.ProcAttr{
		// Files指定新进程继承的活动文件对象
		// 前三个分别为，标准输入、标准输出、标准错误输出
		Files: []*os.File{os.Stdin, os.Stdout, os.Stderr},
		// 新进程的环境变量
		Env: os.Environ(),
	}
	// 开始一个新进程
	p, err := os.StartProcess("/usr/bin/ping", []string{"/usr/bin/ping", "www.baidu.com"}, attr)
	// 如果开启失败
	if err != nil {
		t.Log(err)
	}
	// 打印新进程
	fmt.Printf("p: %v\n", p)
	// 打印新进程id
	fmt.Printf("p.Pid: %v\n", p.Pid)
	// 通过进程id查找进程
	p2, _ := os.FindProcess(p.Pid)
	// 打印进程
	fmt.Printf("p2: %v\n", p2)
	// 等待2秒，执行函数
	time.AfterFunc(time.Second*2, func() {
		// 向p进程发送退出信号
		p.Signal(os.Kill)
	})
	// 等待进程 p 退出，返回进程状态
	ps, _ := p.Wait()
	fmt.Printf("ps.String(): %v\n", ps.String())
}

// 测试信号
func TestSignal(t *testing.T) {
	// 打印当前进程，根据该进程ID，在终端中执行 "kill 该id" 查看程序运行效果
	fmt.Printf("os.Getpid(): %v\n", os.Getpid())
	// 创建协程
	go func() {
		// // 注册路由到默认多路复用器
		// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "hello golang") })
		// // 传入 nil，使用 defaultServeMux
		// http.ListenAndServe(":8088", nil)
		fmt.Println("hello golang")
	}()
	// 创建协程
	go func() {
		// // 注册路由到默认多路复用器
		// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprintln(w, "hello golang") })
		// // 传入 nil，使用 defaultServeMux
		// http.ListenAndServe(":8089", nil)
		fmt.Println("hello golang")
	}()
	// 创建signal channel, 必须要有一个1
	ch := make(chan os.Signal, 1)
	// 接收通知消息。一直等待程序运行
	// signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)
	// 读取接收到的消息
	c := <-ch
	// 打印接收到的消息
	log.Println(c)

}

// 测试获取或更改环境变量
func TestEnv(t *testing.T) {
	// 获取所有环境变量
	env := os.Environ()
	// 打印环境变量
	fmt.Printf("env: %v\n", env)
	// 获取某个环境变量
	gorootEnv := os.Getenv("GOROOT")
	fmt.Printf("gorootEnv: %v\n", gorootEnv)
	// 设置环境变量
	err := os.Setenv("env1", "env1Value")
	if err != nil {
		t.Log(err)
	}
	// 获取环境变量，如果没有该环境变量，打印空串
	s2 := os.Getenv("aaa")
	fmt.Printf("不存在的环境变量s2: %v\n", s2)
	// 查找
	s3, b := os.LookupEnv("env1")
	fmt.Printf("b: %v\n", b)
	fmt.Printf("s3: %v\n", s3)
	// 清空环境变量
	// 删除当前程序已有的所有环境变量。不会影响当前电脑系统的环境变量，这些环境变量都是对当前go程序而言的
	os.Clearenv()
}
