package main

import "fmt"

/**
1.关于channel的特性，下面说法正确的是？

	A. 给一个 nil channel 发送数据，造成永远阻塞
	B. 从一个 nil channel 接收数据，造成永远阻塞
	C. 给一个已经关闭的 channel 发送数据，引起 panic
	D. 从一个已经关闭的 channel 接收数据，如果缓冲区中为空，则返回一个零值

参考答案及解析：ABCD
 */


const i = 100

var j = 123

/**
	常量不同于变量的在运行期分配内存，常量通常会被编译器在预处理阶段直接展开，
	作为指令数据使用，所以常量无法寻址
 */
func test1() {
	fmt.Println(&j, i)
	fmt.Println(&i, j)
}

/**
	nil 可以用作 interface、function、pointer、map、slice 和 channel 的“空值”。
	但是如果不特别指定的话，Go 语言不能识别类型，所以会报错
 */
func GetValue(m map[int]string, id int) (string, bool) {

	if _, exist := m[id]; exist {
		return "exist", true
	}
	return nil, false
}


func test2() {
	intmap := map[int]string{
		1: "a",
		2: "b",
		3: "c",
	}

	v, err := GetValue(intmap, 3)
	fmt.Println(v, err)
}

func main() {
	//test1()
	test2()
}
