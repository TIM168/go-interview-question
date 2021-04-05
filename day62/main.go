package main

import "fmt"

/**
第 11 行，当前作用域中，预定义的 nil 被覆盖，此时 nil 是 int 类型值，不能赋值给 map 类型
 */
func test1() {
	nil := 123
	fmt.Println(nil)
	var _ map[string]int = nil
}

/**
输出-128。因为溢出
 */
func test2() {
	var x int8 = -128
	var y  = x/-1
	fmt.Println(y)
}

func main() {
	test2()
}
