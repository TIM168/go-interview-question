package main

import (
	"fmt"
	"sync"
)

type S1 struct {
}

func (s1 S1) f() {
	fmt.Println("S1.f()")
}

func (s1 S1) g() {
	fmt.Println("S1.g()")
}

type S2 struct {
	S1
}

func (s2 S2) f() {
	fmt.Println("S2.f()")
}

type I interface {
	f()
}

/**
类型断言，结构体嵌套。
结构体 S2 嵌套了结构体 S1，S2 自己没有实现 g() ，调用的是 S1 的 g()
 */
func printType(i I) {
	fmt.Printf("%T\n", i)
	if s1, ok := i.(S1); ok {
		s1.f()
		s1.g()
	}
	if s2, ok := i.(S2); ok {
		s2.f()
		s2.g()
	}
}

func test1() {
	printType(S1{})
	printType(S2{})
}

/**
协程里面，使用 wg.Add(1) 但是没有 wg.Done()，导致 panic()
 */
func test2() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("1")
		wg.Done()
		wg.Add(1)  // 1
	}()
	wg.Wait()
}

func main() {
	test1()
	test2()
}
