package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/test", doRequest)      //   设置访问路由
	err := http.ListenAndServe(":2333", nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}

}

func doRequest(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		return
	} //解析url传递的参数，对于POST则解析响应包的主体（request body）
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("value:", strings.Join(v, ""))
	}
	_, err = fmt.Fprintln(w, "My first Golang Service")
	if err != nil {
		return
	}
	fmt.Fprintln(w, "service start...") //这个写入到w的是输出到客户端的 也可以用下面的 io.WriteString对象

	//注意:如果没有调用ParseForm方法，下面无法获取表单的数据
	//query := r.URL.Query()
	var uid string // 初始化定义变量
	if r.Method == "GET" {
		uid = r.FormValue("uid")
	} else if r.Method == "POST" {
		uid = r.PostFormValue("uid")
	}
	fmt.Fprintln(w, "method is :"+r.Method)
	io.WriteString(w, "uid = "+uid)

}
