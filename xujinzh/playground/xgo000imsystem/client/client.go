package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIP   string
	ServerPort int
	Name       string
	conn       net.Conn
	flagMenu   int // 当前client模式
}

func NewClient(serverIP string, serverPort int) *Client {
	// 创建客户端
	client := &Client{
		ServerIP:   serverIP,
		ServerPort: serverPort,
		flagMenu:   999,
	}
	// 连接server

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIP, serverPort))
	if err != nil {
		fmt.Println("net.Dial error:", err)
		return nil
	}

	client.conn = conn
	// 返回对象
	return client
}

func (this *Client) Menu() bool {
	var flagMenu int
	fmt.Println("1. 公聊模式")
	fmt.Println("2. 私聊模式")
	fmt.Println("3. 更新用户名")
	fmt.Println("4. 退出")
	fmt.Println(">>>请输入1-4之间的整数：<<<")
	fmt.Scanln(&flagMenu)

	// fmt.Println("-----------------------flagMenu:", flagMenu)
	if flagMenu >= 1 && flagMenu <= 4 {
		this.flagMenu = flagMenu
		return true
	} else {
		fmt.Println(">>>请输入范围内合法数字<<<")
		return false
	}
}

func (this *Client) SendMsg(msg string) (n int, err error) {
	n, err = this.conn.Write([]byte(msg))
	return
}

func (this *Client) UpdateName() bool {
	// 提示用户输入用户名
	fmt.Println(">>>请输入用户名：<<<")
	fmt.Scanln(&this.Name)

	// 发送消息
	msgForSend := "rename|" + this.Name + "\n"
	_, err := this.SendMsg(msgForSend)
	if err != nil {
		fmt.Println("client send message err:", err)
		return false
	}
	return true
}

// 处理server回应的消息，显示到标准输出
func (this *Client) DealResponse() {
	// 把 conn 中的消息拷贝到 os.stdout 中，永久阻塞
	// 一旦client.conn有数据，就直接copy到stdout标准输出上，永久阻塞监听（不是只拷贝一次，第二次不拷贝了）
	io.Copy(os.Stdout, this.conn)
}

// 客户端主业务
func (this *Client) Run() {
	fmt.Println("flagMenu:", this.flagMenu)
	for this.flagMenu != 4 {
		// 如果用户在客户端输入的不是1-4之间的数字，那么会一直请用户输入，直到输入1-4之间的值
		// 下面这个for循环保证用户只能输入1-4之间的整数
		for this.Menu() != true {
			// fmt.Println("this.menu is not true")
		}
		// 当用户输入1-4之前的某个值时，进入下面的代码，选择功能
		// 当用户输入4时，因为不符合下面的选项，会判断外层for循环的判断，不满足条件则退出
		// 根据不同模式处理不同业务
		switch this.flagMenu {
		case 1: // 公聊模式
			fmt.Println("----------------公聊模式...")
		case 2: // 私聊模式
			fmt.Println("----------------私聊模式...")
		case 3: // 更新用户名
			// fmt.Println("更新用户名...")
			this.UpdateName()
		}

	}
}

var serverIP string
var serverPort int

// 解析命令行方法：./client -ip 127.0.0.1 -port 8888
func init() {
	flag.StringVar(&serverIP, "ip", "127.0.0.1", "设置服务器IP地址")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器端口")
}
func main() {
	// 执行命令行解析
	flag.Parse()
	// 把解析结果带入执行
	client := NewClient(serverIP, serverPort)
	if client == nil {
		fmt.Println(">>>>>>连接服务器失败...")
		return
	}
	// 单独开启一个 goroutin 去处理 server 发送的消息
	go client.DealResponse()

	fmt.Println(">>>>>>连接服务器成功...")
	// 启动客户端业务
	client.Run()
}
