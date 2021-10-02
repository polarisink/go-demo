package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

// Human 属性，自定义json返回名称
//空tag和不设置tag效果一样
type Human struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	//omitempty：没有就不打印
	Addr string `json:"addr,omitempty"`
}

func main() {
	p1 := Human{
		Name: "Jack",
		Age:  22,
	}
	p2 := Human{
		Name: "Jack",
		Age:  22,
		Addr: "China",
	}
	data1, err := json.Marshal(p1)
	if err != nil {
		panic(err)
	}
	data2, err := json.Marshal(p2)
	if err != nil {
		panic(err)
	}
	// p1 没有 Addr，就不会打印了
	fmt.Printf("%s\n", data1)
	// p2 则会打印所有
	fmt.Printf("%s\n", data2)

	fmt.Println("reflect...")
	//反射
	field, _ := reflect.TypeOf(p2).FieldByName("Name")
	field2 := reflect.ValueOf(p2).Type().Field(1)
	field3 := reflect.ValueOf(&p2).Elem().Type().Field(1)
	fmt.Println(field)
	fmt.Println(field2)
	fmt.Println(field3)

	tag := field2.Tag
	//get是lookup的简单封装
	labelValue := tag.Get("json")
	labelValue2, _ := tag.Lookup("json")
	fmt.Println(tag)
	fmt.Println(labelValue)
	fmt.Println(labelValue2)
}
