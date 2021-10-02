package main

import "fmt"

var studentName string

var (
	name string
	age int
	isOk bool
)

func main() {
	name = "dream"
	age = 16
	isOk = true
	fmt.Print(isOk)
	fmt.Println()
	fmt.Printf("name: %s\n", name)
	fmt.Println(age)

	var s1 string = "www"
	fmt.Println(s1)
	var s2 = "1111"
	fmt.Println(s2)
	s3:="aaaa"
	fmt.Println(s3)
	fmt.Println(studentName)

}
