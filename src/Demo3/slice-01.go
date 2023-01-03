package main

import "fmt"

func main() {
	//切片: slice，底层也是数组，可以动态改变长度
	//定义一个切片
	//names := [10]string{"aa","bb","dd"} 定长
	names := []string{"aa", "bb", "dd"}
	for i, v := range names {
		fmt.Println(i, v)
	}

	//1.追加数据 append()
	names1 := append(names, "ee")
	fmt.Println(names1)

	//2.对于一个切片,不仅有长度的概念len(),还有个容量的概念cap()
	fmt.Println("追加前")
	fmt.Println(len(names), cap(names))
	fmt.Println("追加后")
	fmt.Println(len(names1), cap(names1))
	names2 := append(names1, "ff", "gg", "hh")
	fmt.Println("再次追加")
	fmt.Println(len(names2), cap(names2))
}
