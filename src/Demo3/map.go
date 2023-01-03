package main

import "fmt"

func main() {
	//定义一个字典
	//学生id ===》 学生名字
	var idNames map[int]string
	//定义一个map，此时是不能直接使用的，它是空的

	//分配空间，使用make,可以不指定长度,但是建议直接指定长度
	idNames = make(map[int]string, 10)
	idScores := make(map[int]float64, 10)

	//定义时直接分配空间
	//idNames1 := make(map[int]string,10)

	idNames[0] = "aaa"
	idNames[1] = "bbb"

	for key, value := range idNames {
		fmt.Println(key, "===", value)
	}

	//如何确定一个key是否存在
	//在map中，不存在越界问题，他认为所有的key都是有效的
	//所以访问一个不存在的key不会崩溃，返回这个类型的零值
	// Boolean ==》 false 数字 ==》 0 字符串 ==》 空
	names3 := idNames[99]
	fmt.Println(names3)
	fmt.Println(idScores[100])

	//无法通过获取value值来判断一个key是否存在
	//因此我们需要一个能够校验key是否存在的方式
	value, ok := idNames[1]
	//如果id=1是存在的，那么value就是key=1的值，ok返回true
	//反之返回零值和false
	if ok {
		fmt.Println("id=1的key存在，值为：", value)
	}
	value, ok = idNames[1000]
	if ok {
		fmt.Println("id=1000的key存在，值为：", value)
	} else {
		fmt.Println("不存在")
	}

	//删除map元素，使用自由函数delete来删除指定的key
	delete(idNames, 0)
	fmt.Println("删除第一个元素后:", idNames)
	delete(idNames, 10000)
	fmt.Println("删除不存在的元素后:", idNames)

	//并发处理的时候，需要对map上锁

}
