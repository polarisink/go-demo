package main

import "fmt"

func main() {
	f1 := 1.234567
	//default float64
	fmt.Printf("%T\n", f1)
	f2:=float32(1.234567)
	fmt.Printf("%T\n", f2)
}
