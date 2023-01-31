package main

import "fmt"

func main() {
	//标签 label1
	//goto label1 ===> 下次进入循环时，i的状态不会被保存，重新开始计算
	//break label1 ===> continue会跳到指定的位置，但是会记录之前的状态
	//continue label1 ===> 直接跳出指定位置的循环

	//标签的名字是自定义的，后面要加冒号
LABEL1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if j == 3 {
				//goto LABEL1
				//continue LABEL1
				break LABEL1
			}
			fmt.Println("i: ", i, "j: ", j)
		}
	}
	fmt.Println("over!")

}
