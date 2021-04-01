package main

import "fmt"

type foo struct {
	bar int
}

/**
:= 操作符不能用于结构体字段赋值
*/
func test1() {
	var f foo
	var tmp int
	//f.bar, tmp := 1, 2
	f.bar, tmp = 1, 2
	fmt.Println(tmp)
}

func test2() {
	//fmt.Println(~2)
}

func main() {
	test1()
}
