package main

import (
	"fmt"
	"sync"
)

/**
1.定义一个包内全局字符串变量，下面语法正确的是？

A. var str string
B. str := ""
C. str = ""
D. var str = ""

参考答案及解析：AD。全局变量要定义在函数之外，而在函数之外定义的变量只能用 var 定义。短变量声明 := 只能用于函数之内
*/



/**
存在两个问题：

在协程中使用 wg.Add()；
使用了 sync.WaitGroup 副本；
 */
func test1() {

	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		go func(wg sync.WaitGroup, i int) {
			wg.Add(1)
			fmt.Printf("i:%d\n", i)
			wg.Done()
		}(wg, i)
	}

	wg.Wait()

	fmt.Println("exit")
}

/**
修复代码
 */
func test1Fix(){
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Printf("i:%d\n", i)
			wg.Done()
		}(i)
	}

	wg.Wait()

	fmt.Println("exit")
}

func main() {
	//test1()
	test1Fix()
}
