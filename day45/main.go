package main

import "fmt"

/**
变量重复声明。不能在单独的声明中重复声明一个变量，
但在多变量声明的时候是可以的，但必须保证至少有一个变量是新声明的
 */
func test1() {
	one := 0
	one := 1
}

/**
编译错误，第21行代码没有逗号。用字面量初始化数组、slice 和 map 时，
最好是在每个元素后面加上逗号，即使是声明在一行或者多行都不会出错
 */
func test2() {
	x := []int{
		1,
		2
	}
	_ = x
}

func test(x byte) {
	fmt.Println(x)
}

func test3() {
	var a byte = 0x11
	var b uint8 = 0x12
	fmt.Println(a)
	var c uint8 = a+b
	fmt.Println(c)

	//test(c)
}

func main() {
	test3()
}
