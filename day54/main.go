package main

import "fmt"

func (n N) value() {
	n++
	fmt.Printf("v:%p,%v\n", &n, n)
}

func (n *N) pointer() {
	n++
	fmt.Printf("v:%p,%v\n", n, *n)
}

/**
不能使用多级指针调用方法
*/
func test1() {
	var a N = 25

	p := &a
	//p1 := &p

	p.value()
	//p1.pointer()
}

type N int

func (n N) test() {
	fmt.Println(n)
}

/**
方法表达式。通过类型引用的方法表达式会被还原成普通函数样式，接收者是第一个参数，调用时显示传参。
类型可以是 T 或 *T，只要目标方法存在于该类型的方法集中就可以
*/
func test2() {
	var n N = 10
	fmt.Println(n)

	n++
	f1 := N.test
	f1(n)

	n++
	f2 := (*N).test
	f2(&n)
}

func main() {
	test2()
}
