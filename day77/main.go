package main

import (
	"fmt"
	"testing"
)

/**
1.关于 cap 函数适用下面哪些类型？

A. 数组；
B. channel;
C. map；
D. slice；

参考答案即解析：ABD。cap() 函数的作用：

array 返回数组的元素个数；
slice 返回 slice 的最大容量；
channel 返回 channel 的容量；
*/

func hello(num ...int) {
	num[0] = 18
}
/**
A. 18
B. 5
C. Compilation error

参考答案及解析：A。可变函数是指针传递
 */
func Test13(t *testing.T) {
	i := []int{5, 6, 7}
	hello(i...)
	fmt.Println(i[0])
}

func test1() {
	t := &testing.T{}
	Test13(t)
}

func main() {
	test1()
}
