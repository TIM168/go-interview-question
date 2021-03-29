package main

import "fmt"

/**
	当且仅当接口的动态值和动态类型都为 nil 时，接口类型值才为 nil
 */
func test1() {
	var i interface{}
	if i == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println("not nil")
}

/**
	删除 map 不存在的键值对时，不会报错，相当于没有任何作用；获取不存在的减值对时，
	返回值类型对应的零值，所以返回 0
 */
func test2() {
	s := make(map[string]int)
	delete(s,"h")
	fmt.Println(s["h"])
}

func main() {
	test1()
	test2()
}
