package main

import "fmt"

func findType(i interface{}) {
	switch x := i.(type) {
	case int:
		fmt.Println(x, "is int")
	case string:
		fmt.Println(x, "is string")
	case nil:
		fmt.Println(x, "is nil")
	default:
		fmt.Println(x, "not type matched")
	}
}

func main() {
	var i interface{} = 10
	t1 := i.(int)
	fmt.Println(t1)
	/*t2 := i.(string)
	fmt.Println(t2)

	var i2 interface{} // nil
	var _ = i2.(interface{})*/

	findType(18)
	findType("Hello")
	findType(10.23)
	var  key interface{}
	findType(key)
}
