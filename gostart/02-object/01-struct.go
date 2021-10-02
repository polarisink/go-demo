package main

import (
	"fmt"
	"unsafe"
)

type Person struct {
	id int
	name string
	address string
}
//绑定方法到结构体Person上
func (person Person) fmtPerson(){
	fmt.Printf("名字：%s\n", person.name)
	fmt.Printf("id：%d\n", person.id)
	fmt.Printf("地址：%s\n", person.address)
}

type P struct {}

func main() {
	xiaoming:=Person{
		id:1,
		name:"lqs",
		address: "湖北武汉",
	}
	//	属性名要么全写，要么全不写
	xiaoming2:=Person{
		2,
		"sk",
		"安徽芜湖",
	}
	//只指定部分字段，未指定的使用默认值
	xiaoming3:=Person{
		name:"lxl",
	}
	fmt.Println(xiaoming)
	fmt.Println(xiaoming2)
	fmt.Println(xiaoming3)
	xiaoming2.fmtPerson()

	//空结构体，也有作用
	p:=P{}
	fmt.Println(unsafe.Sizeof(p))
}
