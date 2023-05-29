package main

import (
	"net/rpc/jsonrpc"
)

func main() {
	// 用rpc连接服务器
	//conn, err := rpc.Dial("tcp", "127.0.0.1:8800")
	conn, err := jsonrpc.Dial("tcp", "127.0.0.1:8800")
	if err != nil {
		println("Dial err :", err)
		return
	}
	defer conn.Close()

	// 调用远程函数

	//接受返回值
	var reply string
	err = conn.Call("hello.HelloWorld", "Luke", &reply)
	if err != nil {
		println("conn.Call err :", err)
		return
	}
	println(reply)

}
