package main

import (
	"fmt"
	"sync"
	"time"
)

func doubleScore(source float32) (score float32) {
	defer func() {
		if score < 1 || score >= 100 {
			score = source
		}
	}()
	return source * 2
}

/**
参考答案及解析：输出 0 40 50。知识点：defer 语句与返回值。函数的 return value 不是原子操作，
而是在编译器中分解为两部分：返回值赋值 和 return
 */
func test1() {
	fmt.Println(doubleScore(0))
	fmt.Println(doubleScore(20.0))
	fmt.Println(doubleScore(50.0))
}

var mu sync.RWMutex
var count int

/**
A. 不能编译；
B. 输出 1；
C. 程序 hang 住；
D. fatal error；

参考答案及解析：D。
当写锁阻塞时，新的读锁是无法申请的（有效防止写锁饥饿），导致死锁
 */
func test2() {
	go A()
	time.Sleep(2 * time.Second)
	mu.Lock()
	defer mu.Unlock()
	count++
	fmt.Println(count)
}

func A() {
	mu.RLock()
	defer mu.RUnlock()
	B()
}

func B() {
	time.Sleep(5 * time.Second)
	C()
}

func C() {
	mu.RLock()
	defer mu.RUnlock()
}

func main() {
	test1()
	test2()
}
