package main

import "fmt"

/**
https://www.cnblogs.com/bigdataZJ/p/go-iota.html
ioda 每行才会加一，同一行相同值
 */
const (
	a1 = iota
	a2
	a3
)

func main() {
	fmt.Println(a1)
	fmt.Println(a2)
	fmt.Println(a3)
}
