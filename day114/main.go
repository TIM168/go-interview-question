package main

import "fmt"

func test1() {
	ns := []int{010: 200, 005: 100}
	print(len(ns))
}

/**
参考答案即解析：2。知识点：select 的使用
 */
func test2() {
	i := 0
	f := func() int {
		i++
		return i
	}
	fmt.Println(i)
	c := make(chan int, 1)
	c <- f()
	fmt.Println(i)
	select {
	case c <- f():
	default:
		fmt.Println(i)
	}
}

func main() {
	//test1()
	test2()
}
