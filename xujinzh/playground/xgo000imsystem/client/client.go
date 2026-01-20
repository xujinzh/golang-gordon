package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net"
	"os"
	"strings"
	"time"
)

/*
fmt.Scanln() 接收用空格分隔的值，直到遇到换行符，但如果输入包含空格，它只会读取第一个空格前的部分，更适合接收单个词语或数字
*/

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

// 发送消息
func (this *Client) SendMsg(msg string) (n int, err error) {
	n, err = this.conn.Write([]byte(msg))
	return
}

// 查询当前在线用户
func (this *Client) QueryOnlineUser() (err error) {
	// 1. 通过发送who消息，查询当前在线的用户有哪些
	_, err = this.SendMsg("who\n")
	return
}

// 私聊模式
func (this *Client) PrivateChat() {
	// 0. 先提示用户当前在线用户
	this.QueryOnlineUser()
	// 1. 提示用户选择要私聊的用户
	fmt.Println("请选择您要私聊的对象(输入exit退出私聊)")
	reader := bufio.NewReader(os.Stdin)
	// 1.1 判断读取用户输入私聊对象是否成功，如果不成功，那么退出，成功则赋值给privateObj
	privateObj, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取用户输入私聊对象失败：", err)
		return
	}
	// 1.2 私聊对象后面的换行符去掉，只保留私聊对象用户名
	privateObj = strings.Split(privateObj, "\n")[0]

	// 2. 如果用户输入的不是退出（exit）和空（直接回车），而是用户名，那么进入私聊空间
	for privateObj != "exit" {
		// 1.3 如果用户输入的不是空
		if len(privateObj) != 0 {

			// 判断私聊对象是否正确
			// todo

			// 2.1 首先提示用户私聊对象和输入私聊信息，想要退出，则输入exit
			fmt.Printf("您的私聊对象是[%s]，请输入聊天信息(输入exit退出与该对象私聊)：", privateObj)
			// 2.2 读取用户输入的聊天信息
			privateChatMsg, err := reader.ReadString('\n') // 该私聊信息字符串最后包含有回车符"\n"
			if err != nil {
				fmt.Println("读取用户输入私聊信息失败：", err)
				return
			}
			// 2.2.1 将用户输入的内容后的回车符去掉
			privateChatMsg = strings.Split(privateChatMsg, "\n")[0]

			// 2.3 如果用户输入的不是exit，那么用户可以持续发送私聊信息
			for privateChatMsg != "exit" {
				fmt.Printf("--------------===[%s]===\n", privateChatMsg)
				// 2.4 如果用户输入的不是空内容，那么发送用户消息
				// 2.2.2  如果用户输入的不是空（回车）
				if len(privateChatMsg) != 0 {
					// 2.4.1 先格式化消息
					privateChatMsgFormatted := "to|" + privateObj + "|" + privateChatMsg
					// 2.4.2 发送消息
					_, err := this.SendMsg(privateChatMsgFormatted + "\n")
					// 2.4.3  如果发送失败，那么退出
					if err != nil {
						fmt.Printf("发送消息给[%s]失败：%s", privateObj, err)
						return
					}
				}
				// privateChatMsg = ""
				// 2.1 再次提示用户私聊对象和输入私聊信息，想要退出，则输入exit
				fmt.Printf("您的私聊对象是[%s]，请输入聊天信息(输入exit退出与该对象私聊)：", privateObj)
				// 2.2 读取用户输入的聊天信息
				privateChatMsg, err = reader.ReadString('\n') // 该私聊信息字符串最后包含有回车符"\n"
				fmt.Printf("--------------------[%s]-\n", privateChatMsg)
				if err != nil {
					fmt.Println("读取用户输入私聊信息失败：", err)
					return
				}
				// 2.2.1 将用户输入的内容后的回车符去掉
				privateChatMsg = strings.Split(privateChatMsg, "\n")[0]
				fmt.Printf("-------------------=[%s]=\n", privateChatMsg)
			}
			privateObj = ""
			// 1. 提示用户选择其他要私聊的用户
			fmt.Println("请选择您要私聊的对象(输入exit退出私聊)")
			// 1.1 判断读取用户输入私聊对象是否成功，如果不成功，那么退出，成功则赋值给privateObj
			privateObj, err := reader.ReadString('\n')
			if err != nil {
				fmt.Println("读取用户输入私聊对象失败：", err)
				return
			}
			// 1.2 私聊对象后面的换行符去掉，只保留私聊对象用户名
			privateObj = strings.Split(privateObj, "\n")[0]

		}
	}
}

// 公聊模式
func (this *Client) PublicChat() {
	// 1. 提示用户输入消息
	fmt.Println(">>>请输入要公聊的消息(输入exit退出公聊)：")
	// 2. 定义变量接收用户输入的消息
	var publicChatMsg string
	// 2.1 替换为bufio
	// fmt.Scanln(&publicChatMsg)
	reader := bufio.NewReader(os.Stdin)
	// 2.2 这种方式读取的字符串已经包含了回车符号
	publicChatMsg, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("读取用户输入错误：", err)
		return
	}
	// 3. 当用户输入的内容不是退出公聊模式的exit时，持续要用户进行输入
	for publicChatMsg != "exit\n" {
		// 4. 当用户输入的内容不为空时，则广播消息
		if len(publicChatMsg) != 0 {
			_, err := this.SendMsg(publicChatMsg)
			// 4.1 如果发送失败，那么可能是网络等问题，退出当前公聊模式
			if err != nil {
				fmt.Println("广播消息出错：", err)
				break
			}
			// 4.1 等待发送公聊信息完成后，再进行提示公聊
			time.Sleep(10 * time.Millisecond)
		}
		// 5. 将publicChatMsg置为空，再次提示用户输入公聊消息
		publicChatMsg = ""
		fmt.Println(">>>请输入要公聊的消息(输入exit退出公聊)：")
		// 6. 再次接收用户输入的消息到变量publicChatMsg
		// 6.1 替换为bufio
		// fmt.Scanln(&publicChatMsg)
		publicChatMsg, err = reader.ReadString('\n')
		if err != nil {
			fmt.Println("读取用户输入错误：", err)
			return
		}

	}

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
			// fmt.Println("----------------公聊模式...")
			this.PublicChat()
		case 2: // 私聊模式
			// fmt.Println("----------------私聊模式...")
			this.PrivateChat()
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
