package main

import (
	"fmt"
	"sync"
)

/**
1.Go 语言中中大多数数据类型都可以转化为有效的JSON文本，下面几种类型除外

A. 指针
B. channel
C. complex
D. 函数

参考答案及解析：BCD
*/


/**
2.下面代码输出什么？如果想要代码输出 10，应该如何修改？

参考答案及解析：输出 1。知识点：并发、引用
 */
const N = 10

func test1() {
	m := make(map[int]int)

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			fmt.Println(i)
			m[i] = i
			mu.Unlock()
		}()
	}
	wg.Wait()
	println(len(m))
	fmt.Println(m)
}

/**
修复代码
 */
func test1Fix(){
	m := make(map[int]int)

	wg := &sync.WaitGroup{}
	mu := &sync.Mutex{}
	wg.Add(N)
	for i := 0; i < N; i++ {
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			fmt.Println(i)
			m[i] = i
			mu.Unlock()
		}(i)
	}
	wg.Wait()
	println(len(m))
	fmt.Println(m)
}

func main() {
	test1()
	test1Fix()
}
