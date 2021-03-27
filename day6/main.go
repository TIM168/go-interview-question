package main

import "fmt"

type MyInt1 int
type MyInt2 = int

/**
  类型别名与类型定义的区别
  第 5 行代码是基于类型 int 创建了新类型 MyInt1，第 6 行代码是创建了 int 的类型别名 MyInt2，注意类型别名的定义时 = 。
  所以，第 15 行代码相当于是将 int 类型的变量赋值给 MyInt1 类型的变量，Go 是强类型语言，编译当然不通过；而 MyInt2 只是 int 的别名，本质上还是 int，可以赋值
*/
func test1() {
	var i int = 0
	var i1 MyInt1 = i // var i1 MyInt1 = MyInt1(i)
	var i2 MyInt2 = i
	fmt.Println(i1, i2)
}

func test2() {
	a := []int{7, 8, 9}
	fmt.Printf("%v\n", a)
	ap(a)
	fmt.Printf("%v\n", a)
	app(a)
	fmt.Printf("%v\n", a)
}

func ap(a []int) {
	a = append(a, 10)
	fmt.Println(a)
}

func app(a []int) {
	a[0] = 1
}

func main() {
	test2()
}
