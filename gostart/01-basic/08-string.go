package main

import "fmt"

func main() {
	path := "'D:\\Go\\src\\code'"
	path5 := `"lqs"`
	fmt.Println(path)
	fmt.Printf("path5"+path5)
	s := "I'm OK"
	fmt.Println(s)
	// 多行的字符串
	s2 := `
	世情薄
				人情恶
		雨送黄昏花易落
	`
	fmt.Println(s2)
	s3 := `D:\Go\src\code.oldboyedu.com\studygo\day01`
	fmt.Println(s3)

	//字符串相关操作
	//len
	fmt.Println(len(s2))

	//拼接
	s4 := s2 + s3
	fmt.Println(s4)

	//格式化
	name:=`name`
	world:="world"
	ss1:=fmt.Sprintf("%s%s",name,world)
	fmt.Println(ss1)
}
