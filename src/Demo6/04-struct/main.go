package main

import (
	"fmt"
)

// Student go语言使用type + struct来处理
type Student struct {
	name   string
	age    int
	gender string
	score  float64
}

// MyInt 自定义类型
type MyInt int //相当于 typedef

func main() {

	var i, j MyInt
	i, j = 10, 20
	fmt.Println(i, j)

	lily := Student{
		"lily",
		15,
		"1",
		90.5, //最后一个元素必须加逗号，如果不加逗号，则必须与}同行
	}

	fmt.Println(lily)

	//指针
	s1 := &lily
	fmt.Println("使用指针打印")
	fmt.Println((*s1).age, (*s1).gender, s1.name, s1.score)

	//在定义期间对结构体赋值时，那么字段的名字可以不写
	//如果只对局部变量赋值，必须明确指定变量的名字
	Duke := Student{
		name: "Duke",
		age:  0,
		//gender: "",
		//score:  0,
	}
	Duke.gender = "2"
	Duke.score = 120
	fmt.Println(Duke)

}
