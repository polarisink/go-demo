package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Clint struct {
	ServerIp   string
	ServerPort int
	Name       string
	coon       net.Conn
	flag       int //客户端模式
}

func NewClient(serverIp string, serverPort int) *Clint {
	//创建客户端对象
	clint := &Clint{
		ServerIp:   serverIp,
		ServerPort: serverPort,
		flag:       999,
	}
	//链接server
	coon, err := net.Dial("tcp", fmt.Sprintf("%s:%d", serverIp, serverPort))
	if err != nil {
		fmt.Println("net.Dial err:", err)
		return nil
	}
	clint.coon = coon
	//返回
	return clint
}

func (client *Clint) DealResponse(){
	//一旦client.coon有数据,拷贝到stdout中,永久阻塞监听
	io.Copy(os.Stdout,client.coon)

}

func (client *Clint) menu() bool {
	var flag int
	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")
	//等待用户输入
	fmt.Scanln(&flag)
	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println("请输入合法数字")
		return false
	}
}

var serverIp string
var serverPort int

//从命令行获取:./client -ip 127.0.0.1 -p 8889
//最好在init函数中做
func init() {
	flag.StringVar(&serverIp, "ip", "127.0.0.1", "设置服务器IP地址(默认127.0.0.1)")
	flag.IntVar(&serverPort, "port", 8888, "设置服务器IP地址(默认8888)")
}

func (client *Clint) UpdateName() bool{

	fmt.Println("请输入用户名")
	fmt.Scanln(&client.Name)
	sendMsg:="rename|"+client.Name+"\n"
	_,err:=client.coon.Write([]byte(sendMsg))
	if err!=nil{
		fmt.Println("oon.Write err:",err)
		return false
	}
	return true
}

func (client *Clint) Run() {
	for client.flag != 0 {
		for client.menu() != true {}
		switch client.flag{
		case 1:
			fmt.Println("公聊模式")
			break
		case 2:
			fmt.Println("私聊模式")
			break
		case 3:
			fmt.Println("更新用户名模式")
			client.UpdateName()
			break
		case 0:
		}
	}
}

func main() {
	client := NewClient(serverIp, serverPort)
	if client == nil {
		fmt.Println("链接服务器失败")
	}
	fmt.Println("链接服务器成功")

	//开启goroutine处理server回执
	go client.DealResponse()

	fmt.Println("链接服务器成功")

}
