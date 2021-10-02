package main

import (
	"fmt"
	"time"
)

func mytest() {
	fmt.Println("hello, go")
}

func main() {
	go mytest()
	fmt.Println("hello, world")
	time.Sleep(time.Second)
}
