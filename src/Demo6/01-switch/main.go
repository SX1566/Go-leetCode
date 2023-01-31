package main

import (
	"fmt"
	"os"
)

//从命令输入参数，在switch处理

func main() {
	//GO : os.Args ==> 直接获取前命令输入，是一个字符串切片

	cmds := os.Args

	//os.Args[0] ==> 程序名字
	//os.Args[1] ==> 第一个参数，以此类推
	for key, cmd := range cmds {
		fmt.Println("key:", key, "cmd:", cmd, "len:", len(cmds))

	}
	if len(cmds) < 2 {
		fmt.Println("清正确输入参数")
		return
	}

	fmt.Println("switch:")
	switch cmds[1] {
	case "hello":
		fmt.Println("hello")
		//GO里面 switch默认加break了，不需要手动处理
		//如果想向下穿透的话，需要加上关键字: fallthrough
		fallthrough

	//	$ ./main.exe  hello 111 22233 33444
	//key: 0 cmd: D:\GolandWorkSpace\Goland\src\Demo6\01-switch\main.exe len: 5
	//key: 1 cmd: hello len: 5
	//key: 2 cmd: 111 len: 5
	//key: 3 cmd: 22233 len: 5
	//key: 4 cmd: 33444 len: 5
	//	switch:
	//	hello
	//	default

	default:
		fmt.Println("default")
	}

}
