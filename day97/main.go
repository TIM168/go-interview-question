package main

import "fmt"

/**
1.关于map，下面说法正确的是？

A. map 反序列化时 json.unmarshal() 的入参必须为 map 的地址；
B. 在函数调用中传递 map，则子函数中对 map 元素的增加不会导致父函数中 map 的修改；
C. 在函数调用中传递 map，则子函数中对 map 元素的修改不会导致父函数中 map 的修改；
D. 不能使用内置函数 delete() 删除 map 的元素；

参考答案及解析：A。知识点：map 的使用
 */

type Foo struct {
	val int
}

func (f Foo) Inc(inc int) {
	f.val += inc
}

/**
参考答案及解析：输出 0。使用值类型接收者定义的方法，调用的时候，使用的是值的副本，对副本操作不会影响的原来的值。
如果想要在调用函数中修改原值，可以使用指针接收者定义的方法
 */
func test1() {
	var f Foo
	f.Inc(100)
	fmt.Println(f.val)
}

func main() {
	test1()
}
