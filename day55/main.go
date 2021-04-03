package main

import "fmt"

/**
1.关于 channel 下面描述正确的是？

A. close() 可以用于只接收通道；
B. 单向通道可以转换为双向通道；
C. 不能在单向通道上做逆向操作（例如：只发送通道用于接收）；

参考答案及解析：C
 */

type T struct {
	n int
}

func getT() T {
	return T{}
}

/**
直接返回的 T{} 无法寻址，不可直接赋值
 */
func test1() {
	//getT().n = 1

}

func test2(){
	t := getT()
	p := &t.n
	*p = 1
	fmt.Println(t.n)
}
func main() {
	test2()
}
