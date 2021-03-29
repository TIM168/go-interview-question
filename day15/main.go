package main

import "fmt"

/**
	nil 切片和空切片。nil 切片和 nil 相等，一般用来表示一个不存在的切片；
	空切片和 nil 不相等，表示一个空的集合
 */
func test1() {
	var s1 []int
	var s2 = []int{1}
	fmt.Println(s2)
	if s1 == nil {
		fmt.Println("yes nil")
	} else {
		fmt.Println("no nil")
	}
}

/**
	UTF-8 编码中，十进制数字 65 对应的符号是 A
 */
func test2() {
	i := 65
	fmt.Println(string(i))
}

type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}

func (w Work) ShowB() int {
	return w.i + 20
}

/**
	一种类型实现多个接口，结构体 Work 分别实现了接口 A、B，
	所以接口变量 a、b 调用各自的方法 ShowA() 和 ShowB()，输出 13、23
 */
func test3() {
	c := Work{3}
	var a A = c
	var b B = c
	fmt.Println(a.ShowA())
	fmt.Println(b.ShowB())
}

func main() {
	test1()
	test2()
	test3()
}
