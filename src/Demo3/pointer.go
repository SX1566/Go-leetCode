package main

import "fmt"

func main() {

	name := "name"
	ptr := &name
	fmt.Println("name ptr", ptr, " ==>指针")

	name2Ptr := new(string)
	*name2Ptr = "name2"
	ptr2 := &name2Ptr
	fmt.Println(ptr2)

	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2**i
	}
	for _, value := range pow {
		fmt.Printf("%d\n", value)
	}

	//可以返回栈上的指针，编译器在编译程序时会自动判断这段代码，将city分配在堆上，内存逃逸
	res := testPtr()
	fmt.Println(*res)
	fmt.Println(res)

	//空指针 go: nil C: null C++: nullptr
	if res == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("非空")
	}

}

// 定义一个函数，返回一个string类型的指针，go里面返回值写在参数列表后面
func testPtr() *string {
	city := "武汉"
	ptr := &city
	return ptr

}
