package main

import (
	"fmt"
	"sync"
)

/**
1.ch := make(chan interface{}) 和 ch := make(chan interface{},1) 有什么区别？

参考答案及解析：第一个是声明无缓存通道，第二个是声明缓存为 1 的通道。
无缓存通道需要一直有接收者接收数据，写操作才会继续，不然会一直阻塞；而缓冲为 1 则即使没有接收者也不会阻塞，因为缓冲大小是 1 ，只有当放第二个值的时候，第一个还没被人拿走，这时候才会阻塞。注意这两者还是有区别的
 */

var mu sync.Mutex
var chain string

func main() {
	chain = "main"
	A()
	fmt.Println(chain)
}


/**
A. 不能编译；
B. 输出 main --> A --> B --> C；
C. 输出 main；
D. fatal error；

参考答案即解析：D。使用 Lock() 加锁后，
不能再继续对其加锁，直到利用 Unlock() 解锁后才能再加锁
 */
func A() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + "--> A"
	B()
}

func B() {
	chain = chain + "--> B"
	C()
}

func C() {
	mu.Lock()
	defer mu.Unlock()
	chain = chain + " --> C"
}



