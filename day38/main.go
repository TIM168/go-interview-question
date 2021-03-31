package main

import "fmt"

/**
1.关于异常的触发，下面说法正确的是？
	A. 空指针解析；
	B. 下标越界；
	C. 除数为0；
	D. 调用panic函数；

参考答案：A B C D

 */

func test1() {
	x := []string{"a", "b", "c"}
	for v := range x {
		fmt.Print(v)
	}
}

type User struct {
}

type User1 User
type User2 = User

func (i User1) m1() {
	fmt.Println("m1")
}

func (i User) m2() {
	fmt.Println("m2")
}

/**
	第 26 行代码基于类型 User 创建了新类型 User1，第 27 行代码是创建了 User 的类型别名 User2，注意使用 = 定义类型别名。
	因为 User2 是别名，完全等价于 User，所以 User2 具有 User 所有的方法。但是 i1.m1() 是不能执行的，因为 User1 没有定义该方法
 */
func test2() {
	var i1 User1
	var i2 User2
	i1.m1()
	i2.m2()
}
func main() {
	test1()
	test2()
}
