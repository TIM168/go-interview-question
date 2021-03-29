package main

import "fmt"



func (i int) PrintInt() {
	fmt.Println(i)
}

/**
	基于类型创建的方法必须定义在同一个包内，上面的代码基于 int 类型创建了 PrintInt() 方法，
	由于 int 类型和方法 PrintInt() 定义在不同的包内，所以编译出错
 */
func test1() {
	var i int = 1
	i.PrintInt()
}

type People interface {
	Speak(string) string
}

type Student struct {

}

func (stu *Student) Speak(think string) (talk string) {
	if think == "speak" {
		talk = "speak"
	} else {
		talk = "hi"
	}
	return
}

/**
	值类型 Student 没有实现接口的 Speak() 方法，而是指针类型 *Student 实现该方法
 */
func test2() {
	var peo People = Student{}
	think := "speak"
	fmt.Println(peo.Speak(think))
}

func main() {
	test1()
}
