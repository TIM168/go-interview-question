package main

import "fmt"

/**
函数只能与 nil 比较
 */
func test1() {
	var fn1 = func() {}
	//var fn2 = func() {}

	//if fn1 != fn2 {
	if fn1 != nil {
		println("fn1 not equal fn2")
	}
}

type T struct {
	n int
}

/**
map[key]struct 中 struct 是不可寻址的，所以无法直接赋值
 */
func test2() {
	m := make(map[int]T)

	t := T{1}
	m[0] = t
	//m[0].n = 1
	fmt.Println(m[0].n)
}

func main() {
	test2()
}
