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

	switch cmds[1] {
	case "hello":
		fmt.Println("hello")
	//GO里面 switch默认加break了，不需要手动处理
	default:
		fmt.Println("default")
	}

}
