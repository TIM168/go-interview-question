package main

import "fmt"

/**
	a 的类型是 int，b 的类型是 float，两个不同类型的数值不能相加，编译报错
 */
func test1() {
	a := 5
	b := 8.1
	fmt.Println(a + b)
}

/**
	Go 中的数组是值类型，可比较，另外一方面，数组的长度也是数组类型的组成部分，所以 a 和 b 是不同的类型，是不能比较的
 */
func test2() {
	a := [5]int{1, 2, 3, 4, 5}
	t := a[3:4:4]
	fmt.Println(t[0])
}

func test3() {
	a := [2]int{5, 6}
	b := [3]int{5, 6}
	if a == b {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}

func main() {
	test2()
}
