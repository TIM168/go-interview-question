package main

import "fmt"

/**
运行时错误。知识点：多重赋值。

多重赋值分为两个步骤，有先后顺序：

计算等号左边的索引表达式和取址表达式，接着计算等号右边的表达式；
赋值；
 */
func test1() {
	var a []int = nil
	a, a[0] = []int{1, 2}, 9
	fmt.Println(a)
}


/**
2.下面代码中的指针 p 为野指针，因为返回的栈内存在函数结束时会被释放？

A. false
B. true

参考答案及解析：A。Go语言的内存回收机制规定，只要有一个指针指向引用一个变量，
那么这个变量就不会被释放（内存逃逸），因此在 Go 语言中返回函数参数或临时变量是安全的

 */
type TimesMatcher struct {
	base int
}

func NewTimesMatcher(base int) *TimesMatcher {
	return &TimesMatcher{base: base}
}

func test2() {
	p := NewTimesMatcher(3)
	fmt.Println(p)
}

func main() {
	//test1()
	test2()
}
