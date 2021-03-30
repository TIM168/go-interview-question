package main

import "fmt"

/**
1.关于异常的触发，下面说法正确的是？
	A. 空指针解析；
	B. 下标越界；
	C. 除数为0；
	D. 调用panic函数；

参考答案：ABD

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
