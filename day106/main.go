package main

import (
	"fmt"
	"runtime"
	"time"
)

/**

A. 不能编译；
B. 输出 1；
C. 输出 0；
D. panic；

参考答案及解析：D。ch 未被初始化，关闭时会报错
 */
func test1() {
	var ch chan int
	var count int
	go func() {
		ch <- 1
	}()
	go func() {
		count++
		close(ch)
	}()
	<-ch
	fmt.Println(count)
}

/**

A. 不能编译；
B. 一段时间后总是输出 #goroutines: 1；
C. 一段时间后总是输出 #goroutines: 2；
D. panic；

参考答案即解析：C。程序执行到第二个 groutine 时，ch 还未初始化，导致第二个 goroutine 阻塞。
需要注意的是第一个 goroutine 不会阻塞

 */
func test2() {
	var ch chan int
	go func() {
		ch = make(chan int, 1)
		ch <- 1
	}()
	go func(ch chan int) {
		time.Sleep(time.Second)
	}(ch)

	c := time.Tick(1 * time.Second)
	for range c {
		fmt.Printf("#goroutines:%d\n", runtime.NumGoroutine())
	}
}

func main() {
	//test1()
	test2()
}
