package main

import "fmt"

func main() {

	names := [7]string{"a", "b", "c", "d", "e"}

	//基于names创建一个新的数组
	names1 := [3]string{}
	names1[0] = names[0]

	//切片可以基于一个数组灵活的创建新的数组
	names2 := names[0:3] //左闭右开
	fmt.Println(names2)

	names2[2] = "Hello"
	fmt.Println("修改names2之后：names2：", names2)
	fmt.Println("修改names2之后：names：", names)

	//如果从0元素开始截取，那么冒号左边的数字可以省略
	names3 := names[:5]
	fmt.Println(names3)

	//如果截取到数组最后一个，冒号右边可以省略
	names4 := names[4:]
	fmt.Println(names4)

	//如果全部都要，冒号左边右边可以省略
	names5 := names[:]
	fmt.Println(names5)

	//也可以基于一个字符串进行截取，截取字符串
	sub1 := "HelloWorld"[5:7]
	fmt.Println(sub1)

	//可以创建空切片的时候，指定切片的容量,这样可以减少内存消耗
	//创建一个容量是20，当前长度是0的string类型切片
	str := make([]string, 0, 20)
	fmt.Println(len(str))
	fmt.Println(cap(str))

	//如果想让切片完全独立于原始数组，可以使用copy函数来完成
	namesCopy := make([]string, len(names))
	//names是一个数组，copy在接受参数类型是切片，所以需要使用[:]
	copy(namesCopy, names[:])
	namesCopy[0] = "香港"
	fmt.Println("namesCopy:", namesCopy)
	fmt.Println("names:", names)

}
