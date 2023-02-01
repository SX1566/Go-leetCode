package main

import (
	"fmt"
	"os"
)

func main() {
	//延迟，关键字，可以用于修饰语句、函数，用于确保这条语句可以在当前栈退出的时候执行
	//一般用于做资源清理工作
	//解锁，关闭文件
	//在一个函数多次调用defer，执行时类似于栈的机制：先入后出

	readFile("main.go")
}

func readFile(filename string) {
	//func Open(name string) (*File, error)
	//Go一般会将错误码作为最后一个参数返回
	f1, err := os.Open(filename)

	//匿名函数，没有名字，一次性的逻辑 ===》 lambda表达式
	defer func(a int) {
		err := f1.Close()
		if err != nil {
			return
		}
		fmt.Println("准备关闭文件流,code:", a)
	}(100) //创建一个匿名函数同时调用

	//err一般nil代表没有错误，非nil代表执行失败
	if err != nil {
		fmt.Println("os.Open(\"main.go\") ===> 打开文件失败, err:", err)
		return
	}

	/*
		222
		111
		000
		准备关闭文件流
	*/
	defer fmt.Println("000")
	defer fmt.Println("111")
	defer fmt.Println("222")

	buf := make([]byte, 1024)
	n, err := f1.Read(buf)
	fmt.Println("读取文件的实际长度", n)
	fmt.Println("读取文件内容", string(buf))
}
