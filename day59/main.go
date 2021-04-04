package main

import "fmt"

type N int

func (n *N) test() {
	fmt.Println(*n)
}

/**
13 13 13。知识点：方法值。当目标方法的接收者是指针类型时，那么被复制的就是指针
*/
func test1() {
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

/**
第44行代码会panic
因为左侧的 s[0] 中的 s 为 nil
 */
func test2() {
	var m map[int]bool //nil
	_ = m[123]
	var p *[5]string //nil
	for range p {
		_ = len(p)
	}
	var s []int //nil
	_ = s[:]
	s, s[0] = []int{1, 2}, 9
	fmt.Println(s)
}

func main() {
	test1()
	test2()
}
