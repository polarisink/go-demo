package main

import (
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

// NewUser 创建用户
func NewUser(conn net.Conn, server *Server) *User {
	user := &User{
		conn.RemoteAddr().String(),
		conn.RemoteAddr().String(),
		make(chan string),
		conn,
		server,
	}

	//启动监听的goroutine
	go user.ListenMessage()
	return user
}

// ListenMessage 监听当前用户的信道,一旦有消息就发送到客户端
func (this *User) ListenMessage() {
	for {
		msg := <-this.C
		this.conn.Write([]byte(msg + "\n"))
	}
}

//用户上线
func (this *User) Online() {
	this.server.mapLock.Lock()
	this.server.OnlineMap[this.Name] = this
	this.server.mapLock.Unlock()
	this.server.BroadCast(this, "已上线")
}

//用户下线
func (this *User) Offline() {
	this.server.mapLock.Lock()
	delete(this.server.OnlineMap, this.Name)
	this.server.mapLock.Unlock()
	this.server.BroadCast(this, "已下线")
}

func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

//用户处理消息
func (this *User) DoMessage(msg string) {
	//todo 业务功能可进行重构
	if msg == "who" {
		//查询当前用户有哪些人
		this.server.mapLock.Lock()
		for _, user := range this.server.OnlineMap {
			onlineMsg := "[" + user.Addr + "]" + user.Name + ":" + "在线...\n"
			this.SendMsg(onlineMsg)
		}
		this.server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		//改名字格式:rename|newName
		newName := strings.Split(msg, "|")[1]
		//判断newname是否已存在
		_, ok := this.server.OnlineMap[newName]
		if ok {
			//已存在
			this.SendMsg("当前用户名已被使用\n ")
		} else {
			this.server.mapLock.Lock()
			delete(this.server.OnlineMap, this.Name)
			this.server.OnlineMap[newName] = this
			this.server.mapLock.Unlock()
			this.Name = newName
			this.SendMsg("你已经更新用户名:" + this.Name + "\n ")
		}
	} else if len(msg) > 4 && msg[:3] == "to|" {
		//私聊格式: to|name|msg
		name := strings.Split(msg, "|")[1]
		if name == "" {
			this.SendMsg("消息格式不正确,请使用to|name|msg格式")
			return
		}
		remoteUser,ok:=this.server.OnlineMap[name]
		if !ok{
			this.SendMsg("该用户不存在")
			return
		}
		//根据用户名得到user对象
		msg := strings.Split(msg, "|")[2]
		if msg=="" {
			this.SendMsg("消息为空,请重发")
			return
		}
		remoteUser.SendMsg(this.Name+"对您说:"+msg)
	} else {
		this.server.BroadCast(this, msg)
	}
}
