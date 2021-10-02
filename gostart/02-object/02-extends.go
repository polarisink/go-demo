package main

import "fmt"

type company struct {
	companyName string
	companyAddr string
}

/*组合实现继承
在 Go 语言中，函数名的首字母大小写非常重要，它被来实现控制对方法的访问权限。
当方法的首字母为大写时，这个方法对于所有包都是Public，其他包可以随意调用
当方法的首字母为小写时，这个方法是Private，其他包是无法访问的。
*/

type staff struct {
	name     string
	age      int
	gender   string
	position string
	company
}

func main() {
	myCom := company{
		companyName: "Tencent",
		companyAddr: "深圳市南山区",
	}
	staffInfo := staff{
		name:     "小明",
		age:      28,
		gender:   "男",
		position: "云计算开发工程师",
		company:  myCom,
	}
	fmt.Printf("%s 在 %s 工作\n", staffInfo.name, staffInfo.companyName)
	fmt.Printf("%s 在 %s 工作\n", staffInfo.name, staffInfo.company.companyName)

	//实例化方法
	//1、正常实例化
	com2 := company{
		companyName: "lggk",
		companyAddr: "wuhan",
	}
	fmt.Println(com2)
	//2、new
	com3 := new(company)
	com3.companyAddr = "shanghai"
	com3.companyName = "mhy"
	fmt.Println(com3)
	//3、&
	var com4 = &company{}
	var com5 *company = &company{}//这两个等价
	com4.companyAddr = "shenzhen"
	com4.companyName = "tecent"
	fmt.Println(com4)
	fmt.Println(com5)

	//选择器 .
	//指针直接解引用
	fmt.Println(com4.companyName)
}
