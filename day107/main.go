package main

import (
	"fmt"
	"sync"
)

/**

A. 不能编译；
B. 输出 1；
C. 输出 0；
D. panic；

参考答案及解析：D。sync.Map 没有 Len() 方法
 */
func test1() {
	var m sync.Map
	m.LoadOrStore("a", 1)
	m.Delete("a")
	//fmt.Println(m.Len())
}

/**

A. 不能编译；
B. 输出 2000；
C. 输出可能不是 2000；
D. panic；

参考答案及解析：C。append() 并不是并发安全的，下面可以用锁来实现
 */
func test2() {
	var wg sync.WaitGroup
	var mu sync.Mutex //加入锁
	wg.Add(2)
	var ints = make([]int, 0, 1000)
	go func() {
		for i := 0; i < 1000; i++ {
			mu.Lock()   //写锁
			ints = append(ints, i)
			mu.Unlock() //释放写锁
		}
		wg.Done()
	}()
	go func() {
		for i := 0; i < 1000; i++ {
			mu.Lock() //写锁
			ints = append(ints, 1)
			mu.Unlock() //释放写锁
		}
		wg.Done()
	}()
	wg.Wait()
	fmt.Println(len(ints))
}

func main() {
	test1()
	test2()
}
