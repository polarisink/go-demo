package main

import "fmt"

type Phone interface {
	call()
}
type Nokia struct {
	name string
}

//实现
func (phone Nokia) call() {
	fmt.Printf("我是:%v,是一台电话\n", phone.name)
}

func (phone Nokia) send_wechat() {
	fmt.Println("Hello, Wechat.")
}

//隐式转换
func printType(i interface{})  {
	switch i.(type) {
	case int:
		fmt.Println("参数的类型是 int")
	case string:
		fmt.Println("参数的类型是 string")
	}
}

func main() {
	var nokia1 Phone
	nokia1 = Nokia{"lubia"}
	nokia1.call()
	//不能直接使用，没这方法
	//nokia1.send_wechat()

	//解决方法是不显示的声明为Nokia类型
	nokia2:=Nokia{"8010"}
	nokia2.call()
	nokia2.send_wechat()

	//传参的时候进行隐式转换
	a:=1000
	b:="Stream"
	printType(a)
	printType(b)

	/*会报错
	a11 := 10
	switch a11.(type) {
	case int:
		fmt.Println("参数的类型是 int")
	case string:
		fmt.Println("参数的类型是 string")
	}*/

	//手动转换
	a22 := 10
	switch interface{}(a22).(type) {
	case int:
		fmt.Println("参数的类型是 int")
	case string:
		fmt.Println("参数的类型是 string")
	}
}
