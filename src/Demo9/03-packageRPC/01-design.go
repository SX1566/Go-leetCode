package main

import (
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 要求服务端在注册rpc对象时能让编译器监测出注册对象是否合法

// MyInterface 创建接口 在接口中定义原型
type MyInterface interface {
	HelloWorld(string, *string) error
}

// RegisterService 调用该方法时需要给I传参，传递的参数应该是实现了HelloWorld 方法的类对象
func RegisterService(i MyInterface) {
	err := rpc.RegisterName("hello", i)
	if err != nil {
		return
	}
}

// 像调用本地函数一样去调用远程函数
// 定义一个类
type Myclient struct {
	c *rpc.Client
}

// 由于使用的 c 去调用 Call 因此需要初始化
func (c *Myclient) InitClient(addr string) {
	conn, _ := jsonrpc.Dial("tcp", addr)
	c.c = conn
}

// 实现函数 参照上面的interface实现
func (this *Myclient) HelloWorld(a string, b *string) error {
	// 参数1 参照上面interface RegisterName 而来
	return this.c.Call("hello.HelloWorld", a, b)
}
