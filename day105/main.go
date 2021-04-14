package main

import (
	"fmt"
	"sync"
)

var c = make(chan int)
var a int

func f() {
	a = 1
	<-c
}

/**

A. 不能编译；
B. 输出 1；
C. 输出 0；
D. panic；


参考答案及解析：B。能正确输出，不过主协程会阻塞 f() 函数的执行
*/
func test1() {
	go f()
	c <- 0
	print(a)
}

type MyMutex struct {
	count int
	sync.Mutex
}

/**

A. 不能编译；
B. 输出 1, 1；
C. 输出 1, 2；
D. fatal error；

参考答案及解析：D。加锁后复制变量，会将锁的状态也复制，所以 mu1 其实是已经加锁状态，再加锁会死锁
 */
func test2() {
	var mu MyMutex
	mu.Lock()
	var mu1 = mu  //此处复制锁
	mu.count++
	mu.Unlock()
	mu1.Lock()
	mu1.count++
	mu1.Unlock()
	fmt.Println(mu.count, mu1.count)
}

func main() {
	test1()
	test2()
}
