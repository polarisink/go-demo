package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port int
	//在线用户map
	OnlineMap map[string]*User
	mapLock   sync.RWMutex
	//消息广播的channel
	Message chan string
}

// NewServer /*创建一个server接口*/
func NewServer(ip string, port int) *Server {
	return &Server{
		Ip:        ip,
		Port:      port,
		OnlineMap: make(map[string]*User),
		Message:   make(chan string),
	}
}

//监听广播channel的goroutine,一有消息就发送给全部在线用户
func (this *Server) ListenMessager() {
	for {
		msg := <-this.Message
		//msg发送给全部在线用户
		this.mapLock.Lock()
		for _, cli := range this.OnlineMap {
			cli.C <- msg
		}
		this.mapLock.Unlock()
	}
}

// BroadCast 广播消息
func (this *Server) BroadCast(user *User, msg string) {
	sendMsg := "[" + user.Addr + "]" + user.Name + ":" + msg
	this.Message <- sendMsg
}

func (this *Server) Handler(conn net.Conn) {

	user := NewUser(conn, this)
	//...当前连接的业务
	fmt.Println("连接建立成功")
	this.mapLock.Lock()
	this.OnlineMap[user.Name] = user
	this.mapLock.Unlock()

	user.Online()

	isLive := make(chan bool)

	//用户上线
	this.BroadCast(user, "已上线")

	//接受客户端发送的信息
	go func() {
		//4096先写死
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				user.Offline()
				return
			}
			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err:", err)
			}
			//提取消息，去除\n
			msg := string(buf[:n-1])
			//将得到的消息进行广播
			user.DoMessage(msg)
			//用户任意消息表示他在活跃状态
			isLive <- true
		}
	}()

	for {
		select {
		case <-isLive:
			//当前用户是活跃的，应该更新定时器
			//不做任何事情，为了激活select
		case <-time.After(time.Second * 100):
			//已经超时
			user.SendMsg("你被踢了")
			//销毁资源
			close(user.C)
			//关闭连接
			conn.Close()
		}
	}

	//block
	select {}
}

func (this *Server) Start() {
	//socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", this.Ip, this.Port))
	if err != nil {
		fmt.Println("net listen err: ", err)
		return
	}
	//close listen socket
	defer listener.Close()

	//启动监听message的goroutine
	go this.ListenMessager()

	//accept
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener accept err: ", err)
			continue
		}
		go this.Handler(conn)
	}

	//do handler

}
