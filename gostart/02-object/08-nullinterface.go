package main

import "fmt"

func myfunc(iface interface{}) {
	fmt.Println(iface)
}

func myfunc2(ifaces ...interface{}) {
	for _, iface := range ifaces {
		fmt.Println(iface)
	}
}

func main() {
	var i interface{}
	fmt.Printf("type: %T, value: %v", i, i)

	//1、保存值
	i = 1
	fmt.Println(i)
	i = "hello"
	fmt.Println(i)
	i = false
	fmt.Println(i)

	//2、函数接收任意类型的值
	a := 10
	b := "hello"
	c := true
	myfunc(a)
	myfunc(b)
	myfunc(c)
	myfunc2(a, b)

	//空接口可以保存任何值，但空接口赋值给固定类型会报错
	/*var i3 interface{} = a
	var b2 int = i3*/

	//当空接口承载数组和切片后，该对象无法再进行切片
	sli := []int{2, 3, 5, 7, 11, 13}
	var i4 interface{}
	i4 = sli
	/*cannot slice i4 (type interface {})
	g := i4[1:3]
	fmt.Println(g)*/
	fmt.Println(i4)
}
