package main

import (
	"fmt"
	"io"
	"net"
	"runtime"
	"sync"
	"time"
)

type Server struct {
	Ip        string
	Port      int
	OnlineMap map[string]*User // 在线用户列表
	mapLock   sync.RWMutex
	Message   chan string // 消息广播的channel
}

// 广播消息方法
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := fmt.Sprintf("[addr-> %s] name-> %s: msg-> %s", user.Addr, user.Name, msg)
	this.Message <- sendMsg
}

// 监听message广播消息channel的goroutine，一旦有消息就发送给全部的在线user
func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message
		// 将msg发送给全部的在线用户
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

func (this *Server) Handler(conn net.Conn) {
	user := NewUser(conn, this)
	fmt.Printf("%s - %s-%s已上线\n", time.Now().Format("2006-01-02 15:04:05.000"), user.Addr, user.Name)
	// 用户上线
	// // 将用户加入onlineMap中
	// this.mapLock.Lock()
	// this.OnlineMap[user.Name] = user
	// this.mapLock.Unlock()
	// // 广播当前用户上线消息
	// this.BroadCast(user, fmt.Sprintf("%s已上线", user.Name))
	user.Online()

	// 监听用户是否活跃
	isLive := make(chan bool)

	// 接收客户端发送的消息
	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			// 客户端合法关闭client
			if n == 0 {
				// this.BroadCast(user, fmt.Sprintf("%s下线", user.Name))
				user.Offline()
				fmt.Printf("%s - %s-%s下线\n", time.Now().Format("2006-01-02 15:04:05.000"), user.Addr, user.Name)
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("conn read err:", err)
				return
			}
			// 提取用户的消息(去除'\n')
			msg := string(buf[:n-1])
			// 将得到的消息进行广播
			// this.BroadCast(user, msg)
			user.HandleMessage(msg)
			// 用户的任意消息，代表当前用户是活跃的
			isLive <- true
		}
	}()
	// 当前handler阻塞
	for {

		select {
		case <-isLive:
			// 当用户活动状态下，该管道会触发。同时会运行下一条CASE，重置计时器为10秒；如果该条没有运行，那么
			// 不做任何事情，为了激活select，更新下面的定时器
		case <-time.After(time.Second * 600):
			// 已经超时，将当前的用户强制关闭
			// 如果该条触发，那么表示用户10秒未活动，剔除用户下线
			// user.conn.Write([]byte("长期未活跃，被迫下线\n"))
			user.SendMsg("长期未活跃，被迫下线\n")
			// 销毁使用的资源
			close(user.C)
			// 关闭连接
			conn.Close()
			// 退出当前的handler
			// return // or runtime.Goexit()
			runtime.Goexit()
		}
	}
}

func (this *Server) Start() {
	fmt.Println("Listening...")

	// 创建 socket
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net listen err:", err)
	}

	// 关闭连接
	defer listener.Close()

	// 启动监听 message 的 goroutine
	go this.ListenMessager()

	for {
		conn, err := listener.Accept()

		if err != nil {
			fmt.Println("listener accept err:", err)
			continue
		}
		go this.Handler(conn)

	}

}

func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
	return server
}
