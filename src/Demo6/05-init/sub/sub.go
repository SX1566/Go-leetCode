package sub

import "fmt"

//同一层级目录下不允许出现多个包名

// init没有参数，没有返回值，原型固定如下
// 一个包中包含多个init时，调用的顺序是不确定的
// init函数不允许用户显示调用的
// 有的时候引用一个包，可能只想使用这个包里面的init函数
// 不想使用这个包里面的其他函数，为了防止编译器报错，可以使用下划线处理

func init() {
	fmt.Println("this is first init() package")
}

func init() {
	fmt.Println("this is second init() package")
}

func Sub(a, b int) int {
	//init() ===> 不允许显示调用
	return a - b
}
