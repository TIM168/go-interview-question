package main

import "fmt"

func hello(i int) {
	fmt.Println(i)
}

/**
	hello() 函数的参数在执行 defer 语句的时候会保存一份副本，
	在实际调用 hello() 函数时用，所以是 5
 */
func test1() {
	i := 5
	defer hello(i)
	i = i + 10
}

type People struct{}

func (p *People) ShowA() {
	fmt.Println("showA")
	p.ShowB()
}

func (p *People) ShowB() {
	fmt.Println("showB")
}

type Teacher struct {
	People
}

func (t *Teacher) ShowB() {
	fmt.Println("teacher showB")
}

/**
	Teacher 没有自己 ShowA()，所以调用内部类型 People 的同名方法
 */
func test2() {
	t := Teacher{}
	t.ShowA()
}

func main() {
	test1()
	test2()
}
