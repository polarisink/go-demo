package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
)

func sayHelloName(w http.ResponseWriter,r *http.Request){
	r.ParseForm()//解析参数
	fmt.Println(r.Form)
	fmt.Println("path",r.URL.Path)
	fmt.Println("scheme",r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k,v :=range r.Form {
		fmt.Println("key:",k)
		fmt.Println("value:",strings.Join(v,""))
	}
	fmt.Fprintf(w,"Hello World")//写入w的输出到客户端
}

func main() {
	http.HandleFunc("/",sayHelloName)
	err:=http.ListenAndServe(":9000",nil)
	if err != nil {
		log.Fatal("ListenAndServe",err)
	}
}
