package main

import "fmt"

func test1(){
	s := make([]int,3,9)
	fmt.Println(len(s))
	fmt.Println(cap(s))
	s1 := s[4:8:9]
	fmt.Println(len(s1))
	fmt.Println(cap(s1))
}

type N int

func (n N) test(){
	fmt.Println(n)
}

/**
知识点：方法值。当指针值赋值给变量或者作为函数参数传递时，会立即计算并复制该方法执行所需的接收者对象，
与其绑定，以便在稍后执行时，能隐式第传入接收者参数
 */
func test2(){
	var n N = 10
	p := &n

	n++
	f1 := n.test

	n++
	f2 := p.test

	n++
	fmt.Println(n)

	f1()
	f2()

}
func main() {
	test1()
	test2()
}
