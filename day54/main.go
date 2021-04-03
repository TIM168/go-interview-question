package main

import "fmt"

//func (n N) value() {
//	n++
//	fmt.Printf("v:%p,%v\n", &n, n)
//}
//
//func (n *N) pointer() {
//	n++
//	fmt.Printf("v:%p,%v\n", n, *n)
//}
//
//func test1() {
//	var a N = 25
//
//	p := &a
//	p1 := &p
//
//	p1.value()
//	p1.pointer()
//}

type N int

func (n N) test() {
	fmt.Println(n)
}

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
