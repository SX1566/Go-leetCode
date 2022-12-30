package main

import "fmt"

func main() {
	//定义数组
	nums := [10]int{1, 2, 3}
	//遍历，方式1
	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}
	//方式二 for range ===> python支持
	//key是数组下标，value是数组的值
	for key, value := range nums {
		value += 1
		fmt.Println("key:", key, "value:", value, "num", nums[key])
	}

	//在go中，如果想忽略某个值，可以使用 _ 代替
	for _, value := range nums {
		fmt.Println("忽略key,value:", value)
	}

	for key, _ := range nums {
		fmt.Println("忽略value,key:", key)
	}

	//如果两个都忽略，直接用 =
	for _, _ = range nums {
		fmt.Println("忽略value和key:")
	}

	//不定长数组定义
	//使用make创建数组

}
