package main

import "fmt"

/**
类型断言。类型断言语法：i.(Type)，其中 i 是接口，Type 是类型或接口。编译时会自动检测 i 的动态类型与 Type 是否一致。
但是，如果动态类型不存在，则断言总是失败
 */
func test1() {
	x := interface{}(nil)
	y := (*int)(nil)
	a := y == x
	b := y == nil
	_, c := x.(interface{})
	println(a, b, c)
}


func test2() {
	var s []int
	s = append(s, 1)
	fmt.Println(s)
	var m map[string]int
	m = make(map[string]int)
	m["one"] = 1
	fmt.Println(m)
}


func main() {
	test1()
	test2()
}
