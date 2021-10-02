package main

import (
	"fmt"
	"reflect"
)

type Man struct {
	Name   string `label:"Name is: "`
	Age    int    `label:"Age is: "`
	Gender string `label:"Gender is: " default:"unknown"`
}

func Print(obj interface{}) error {
	v := reflect.ValueOf(obj)
	for i := 0; i < v.NumField(); i++ {
		//取tag
		field := v.Type().Field(i)
		tag := field.Tag
		label := tag.Get("label")
		defaultValue := tag.Get("default")
		value := fmt.Sprintf("%v", v.Field(i))
		if value == "" {
			value = defaultValue
		}
		fmt.Println(label + value)
	}
	return nil
}

func main() {
	man := Man{
		Name: "lqs",
		Age: 22,
	}
	Print(man)
}
