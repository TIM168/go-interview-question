package main

import (
	"fmt"
	"os"
	"runtime"
)

/**
for {} 独占 CPU 资源导致其他 Goroutine 饿死
*/
func test1() {
	runtime.GOMAXPROCS(1)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
	}()

	for {
		var a = 1
		fmt.Println(a)
	}
}

/**
可以通过阻塞的方式避免 CPU 占用

修复代码
*/

func test1Fix() {
	runtime.GOMAXPROCS(1)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(i)
		}
		os.Exit(0)
	}()

	select {}
}

/**
2.假设 x 已声明，y 未声明，下面 4 行代码哪些是正确的。错误的请说明原因？

1、 x, _ := f()  // 1
2、 x, _ = f()  // 2
3、 x, y := f()  // 3
4、 x, y = f()  // 4

参考答案及解析：2、3正确。知识点：简短变量声明。使用简短变量声明有几个需要注意的地方：

只能用于函数内部；
短变量声明语句中至少要声明一个新的变量；
*/

func main() {
	test1()
}
