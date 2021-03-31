package main

import (
	"fmt"
	"time"
)

/**
1.关于无缓冲和有冲突的channel，下面说法正确的是？

	A. 无缓冲的channel是默认的缓冲为1的channel；
	B. 无缓冲的channel和有缓冲的channel都是同步的；
	C. 无缓冲的channel和有缓冲的channel都是非同步的；
	D. 无缓冲的channel是同步的，而有缓冲的channel是非同步的；

参考答案及解析：D
*/

func Foo(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}

/**
interface 的内部结构，我们知道接口除了有静态类型，还有动态类型和动态值，
当且仅当动态值和动态类型都为 nil 时，接口类型值才为 nil。这里的 x 的动态类型是 *int，所以 x 不为 nil
*/
func test1() {
	var x *int = nil
	fmt.Println(x)
	Foo(x)
}

func test2() {
	ch := make(chan int, 100)
	//A
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	//B
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	//close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 10)
}

func main() {
	test1()
	test2()
}
