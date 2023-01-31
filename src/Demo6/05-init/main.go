package main

import (
	_ "05-init/sub" //此时只会调用sub的init函数，编译不会报错
)

// 同一个包的多个文件中都可以有init（）
func main() {
	//add.Add(1, 2)
}
