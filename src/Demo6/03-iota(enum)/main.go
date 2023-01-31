package main

import "fmt"

// go里面没有枚举类型，但是我们可以使用 const + iota （常量计数器）来进行模拟
//模拟一个一周的枚举

const (
	Monday = iota //0

	Tuesday = iota //1

	Wednesday = iota //2

	Thursday = iota //...

	Friday // 4 没有赋值，默认与上一行相同

	Saturday

	Sunday

	M, N = iota, iota //const属于预编译期复制，不需要:=
)

const (
	JANU = iota
	FER
)

/*
1. iota是常量计数器
2. iota从0开始，每换行递增1
3. 常量组有个特点，如果默认不赋值，默认与上一行表达式相同
4. 如果同一行出现了两个iota，那么两个iota的值是相同的
5. 每个常量组的iota是独立的，如果遇到const iota会清零
*/
func main() {

	fmt.Println(Monday)
	fmt.Println(Tuesday)
	fmt.Println(Sunday)
	fmt.Println(M) //7
	fmt.Println(N) //7

	fmt.Println(JANU) //0
	fmt.Println(FER)  //1

	//var number int
	//var name string
	//var flag bool

	//可以使用变量组，统一定义变量
	//var(
	//	number int
	//	name string
	//	flag bool
	//)

}
