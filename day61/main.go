package main

import "fmt"

/**
计算等号左边的索引表达式和取址表达式，接着计算等号右边的表达式；
赋值；
所以本例，会先计算 s[k]，等号右边是两个表达式是常量，所以赋值运算等同于 k, s[1] = 0, 3
 */
func test1() {
	var k = 1
	var s = []int{1, 2}
	k, s[k] = 0, 3
	fmt.Println(s)
	fmt.Println(s[0] + s[1])
}

func test2() {
	var k = 9
	for k = range []int{} {
	}
	fmt.Println(k)

	for k = 0; k < 3; k++ {
	}

	fmt.Println(k)

	for k = range (*[4]int)(nil) {
	}
	fmt.Println(k)
}

func main() {
	//test1()
	test2()
}
