package main

import (
	"flag"
	"fmt"
	"net"
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
	fmt.Println("0. 退出")

	fmt.Scanln(&flagMenu)
	if flagMenu >= 0 && flagMenu <= 3 {
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

// 客户端主业务
func (this *Client) Run() {
	for this.flagMenu != 0 {
		for this.Menu() != true {

		}
		// 根据不同模式处理不同业务
		switch this.flagMenu {
		case 1: // 公聊模式
			fmt.Println("公聊模式...")
		case 2: // 私聊模式
			fmt.Println("私聊模式...")
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
	fmt.Println(">>>>>>连接服务器成功...")
	// 启动客户端业务
	client.Run()
}
