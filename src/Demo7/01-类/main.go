package main

import "fmt"

//Person类 绑定方法，Eat，Run，Laugh
//public:private
/*
	class Person{
	public:
		string name
		int age

	public:
		Eat(){
			XXX
		}
	}
*/

type Person struct {
	//成员属性
	name   string
	age    int
	gender string
	score  float64
}

// 在类的外面绑定方法
func (p *Person) eat() {
	fmt.Println("Person is eating")
	//类的方法可以使用自己的成员
	fmt.Println(p.name + " is eating")
}

func (p Person) eat1() {
	fmt.Println("Person is eating")
	//类的方法可以使用自己的成员
	p.name = "Duke"

}

func main() {
	lily := Person{
		name:   "lily",
		age:    10,
		gender: "1",
		score:  20,
	}
	fmt.Println(lily)
	lily.eat()

	fmt.Println("eat2 不使用指针")
	lily.eat1()
	fmt.Println(lily)

}
