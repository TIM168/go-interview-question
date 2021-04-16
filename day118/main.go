package main

import "fmt"

/**
1、下面说法正确的是

A. Go 语言中，声明的常量未使用会报错；
B. cap() 函数适用于 array、slice、map 和 channel;
C. 空指针解析会触发异常；
D. 从一个已经关闭的 channel 接收数据，如果缓冲区中为空，则返回一个零值；


参考答案及解析：CD。A.声明的常量未使用不会报错；B.cap() 函数不适用 map
*/

const (
	_      = iota
	c1 int = (10 * iota)
	c2
	d = iota
)


/**
A. compile error
B. 1 - 2 - 3
C. 10 - 20 - 30
D. 10 - 20 - 3

参考答案及解析：D。iota 的使用

 */
func test1(){
	fmt.Printf("%d - %d - %d",c1,c2, d)
}

func main() {
	test1()
}
