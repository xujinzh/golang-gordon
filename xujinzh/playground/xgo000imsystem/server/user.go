package main

import (
	"fmt"
	"net"
	"strings"
)

type User struct {
	Name   string
	Addr   string
	C      chan string
	conn   net.Conn
	server *Server
}

// 创建一个用户
func NewUser(conn net.Conn, server *Server) *User {
	userAddr := conn.RemoteAddr().String()
	user := &User{
		Name:   fmt.Sprintf("%s.user", userAddr),
		Addr:   userAddr,
		C:      make(chan string),
		conn:   conn,
		server: server,
	}
	// 启动监听消息管道channel的goroutine
	go user.ListenMessage()
	return user
}

// 监听当前用户channel的方法，一旦有消息，就直接发送给对端客户端
func (this *User) ListenMessage() {
	for {
		// 从管道读取消息
		msg := <-this.C
		// 写入消息到客户端
		this.conn.Write([]byte(msg + "\n"))
	}
}

// 用户上线功能
func (this *User) Online() {
	// 将用户加入onlineMap中
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()
	// 广播当前用户上线消息
	this.server.BroadCast(this, fmt.Sprintf("%s已上线", this.Name))
}

// 用户下线功能
func (this *User) Offline() {
	// 将用户从 onlineMap 中删除
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()
	// 广播当前用户下线消息
	this.server.BroadCast(this, fmt.Sprintf("%s下线", this.Name))
}

func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

// 处理消息
func (this *User) HandleMessage(msg string) {

	if msg == "who" { // 用户查询谁在线
		for _, user := range this.server.OnlineMap {
			onlineUserMsg := fmt.Sprintf("%s-%s 在线...", user.Addr, user.Name)
			// this.conn.Write([]byte(onlineUserMsg + "\n"))
			this.SendMsg(onlineUserMsg + "\n")
		}
	} else if len(msg) > 8 && msg[:7] == "rename|" { // 修改自己的用户名
		// 消息格式符合改名字
		newName := strings.Split(msg, "|")[1]
		// 改名字之前，先判断这个名字是否已经被占用了
		_, ok := this.server.OnlineMap[newName]
		if ok { //如果用户名在onlineMap中已经存在，那么就不能再使用该名字了
			// this.conn.Write([]byte(fmt.Sprintf("%s已经存在\n", newName)))
			this.SendMsg(fmt.Sprintf("%s已经存在\n", newName))
		} else {
			// 从server的在线MAP中先删除旧名字，再增加新名字
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()
			// 更新当前用户名
			oldName := this.Name
			this.Name = newName
			// 通知客户端更新名字成功
			// this.conn.Write([]byte(fmt.Sprintf("更新名字%s到%s成功\n", oldName, newName)))
			this.SendMsg(fmt.Sprintf("更新名字%s到%s成功\n", oldName, newName))
		}

	} else if len(msg) > 4 && msg[:3] == "to|" { // 对指定用户发送消息
		// 消息格式：to|zhangsan|message

		// 1. 获取对方用户名
		remoteName := strings.Split(msg, "|")[1]
		if remoteName == "" {
			this.SendMsg("消息格式不正确，请使用\"to|name|message\"格式。\n")
			return
		}
		// 2. 根据用户名得到对方user对象
		remoteUser, ok := this.server.OnlineMap[remoteName]
		if !ok {
			this.SendMsg("该用户名不存在\n")
			return
		}
		// 3. 获取消息内容，通过对方的user对象将消息内容发过去
		content := strings.Split(msg, "|")[2]
		if content == "" {
			this.SendMsg("无消息内容，请重发\n")
			return
		}
		remoteUser.SendMsg(this.Name + "说：" + content + "\n")
	} else {

		this.server.BroadCast(this, msg)
	}
}
