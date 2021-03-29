package main

import "fmt"

type S struct {}

func f(x interface{}) {

}

func g(x *interface{}) {

}

/**
	永远不要使用一个指针指向一个接口类型，因为它已经是一个指针
 */
func test1() {
	s := S{}
	p := &s
	f(s) //A
	//g(s) //B
	f(p) //C
	//g(p) //D
}

type A struct {
	m string
}

func f1() *A {
	return &A{"1"}
}


func test2() {
	p := *f1()
	fmt.Println(p.m)

}
func main() {
	test1()
	test2()
}
