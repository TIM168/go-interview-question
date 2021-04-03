package main

import "fmt"

type X struct {
}

func (x *X) test() {
	println(x)
}

/**
X{} 是不可寻址的，不能直接调用方法。
知识点：在方法中，指针类型的接收者必须是合法指针（包括 nil）,或能获取实例地址
*/
func test1() {
	var a *X
	a.test()

	//X{}.test()
	var x = X{}
	x.test()
}

/**
检查 map 是否含有某一元素，直接判断元素的值并不是一种合适的方式。
最可靠的操作是使用访问 map 时返回的第二个值
*/
func test2() {
	x := map[string]string{"one": "a", "two": "", "three": "c"}

	//if v := x["two"]; v == "" {
	if _, ok := x["two1"]; !ok {
		fmt.Println("no entry")
	}
}

func main() {
	test1()
	test2()
}
