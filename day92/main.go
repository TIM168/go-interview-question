package main

import "fmt"

var x int

func init() {
	x++
}

/**
参考答案及解析：编译失败。
init() 函数不能被其他函数调用，包括 main() 函数
*/
func test1() {
	init()
	fmt.Println(x)
}

/**
min() 函数是求两个数之间的较小值，能否在 该函数中添加一行代码将其功能补全
*/
func min(a int, b uint) {
	var min = 0
	min = copy(make([]struct{}, a), make([]struct{}, b))
	fmt.Printf("The min of %d and %d is %d\n", a, b, min)
}

func test2() {
	min(1225, 256)
}

func main() {
	test1()
	test2()

}
