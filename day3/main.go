package main

import "fmt"

/**
	返回 [0 0 0 0 0 1 2 3 4 5]
*/
func test1() {
	s := make([]int, 5)
	s = append(s, 1, 2, 3, 4, 5)
	fmt.Println(s)
}

/**
	返回 [1 2 3 4 5]
 */
func test2() {
	s := make([]int, 0)
	s = append(s, 1, 2, 3, 4, 5)
	fmt.Println(s)
}

/**
    当函数有多个返回值时，当有一个函数返回值命名了，其它返回值也必须命名
 */
func funcMui(x, y int) (sum int, err error) {
	return x + y, nil
}

func main() {
	test1()
	test2()
	fmt.Println(funcMui(1,2))
}

/**
	new和make区别：
	1、new(T)返回的是一个指针，指针指向新分配、类型为T的零值，适用于数组、结构体等
	2、make(T,args)返回的是一个T类型的值，适用于slice切片，map映射，channel通道
 */
