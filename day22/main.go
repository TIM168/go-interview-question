package main

import "fmt"

/**
	golang 的字符串类型是不能赋值 nil 的，也不能跟 nil 比较
*/
func test1() {
	var x string = nil
	if x == nil {
		x = "default"
	}
	fmt.Println(x)
}

var a bool = true

/**
	defer 关键字后面的函数或者方法想要执行必须先注册，return 之后的 defer 是不能注册的，
	也就不能执行后面的函数或方法
 */
func test2() {
	defer func() {
		fmt.Println("1")
	}()
	if a == true {
		fmt.Println("2")
		return
	}
	defer func() {
		fmt.Println("3")
	}()
}

func main() {
	test1()
	test2()
}
