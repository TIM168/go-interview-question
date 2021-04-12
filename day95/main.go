package main

import (
	"fmt"
	"time"
)

type foo struct{ Val int }

type bar struct{ Val int }

/**
参考答案及解析：false true true。
这道题唯一有疑问的地方就在第一个比较，Go 语言里没有引用变量，每个变量都占用一个惟一的内存位置，所以第一个比较输出 false
 */
func test1() {
	a := &foo{Val: 5}
	b := &foo{Val: 5}
	c := foo{Val: 5}
	d := bar{Val: 5}
	e := bar{Val: 5}
	f := bar{Val: 5}
	fmt.Print(a == b, c == foo(d), e == f)
}

func A() int {
	time.Sleep(100 * time.Millisecond)
	return 1
}

func B() int {
	time.Sleep(1000 * time.Millisecond)
	return 2
}

/**
参考答案及解析：1、2随机输出
 */
func test2() {
	ch := make(chan int, 1)
	go func() {
		select {
		case ch <- A():
		case ch <- B():
		default:
			ch <- 3
		}
	}()
	fmt.Println(<-ch)
}

func main() {
	//test1()
	test2()
}
