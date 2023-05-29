package main

import (
	"fmt"
	"net"
	"net/rpc/jsonrpc"
)

// World 定义类对象
type World struct {
}

// HelloWorld 绑定类方法
func (this *World) HelloWorld(name string, resp *string) error {
	*resp = name + " 你好！"
	return nil
}

func main() {
	////注册RPC服务，绑定对象方法
	//err := rpc.RegisterName("hello", new(World))
	//if err != nil {
	//	fmt.Println("注册RPC失败...", err)
	//	return
	//}
	RegisterService(new(World))

	//设置监听
	listener, err := net.Listen("tcp", "127.0.0.1:8800")
	if err != nil {
		println("net.Listen err :", err)
		return
	}
	defer listener.Close()

	fmt.Println("开始监听")

	//建立连接
	conn, err := listener.Accept()
	if err != nil {
		println("Accept() err :", err)
		return
	}
	defer conn.Close()
	println("连接成功......")

	//绑定服务
	//rpc.ServeConn(conn)
	jsonrpc.ServeConn(conn)
}
