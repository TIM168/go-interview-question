package main

import "fmt"

/**
	append函数用于切片slice追加元素，第一个参数一个切片，第二个参数是切片存储元素类型的可变参数
	1、new([]int) 之后的 list 是一个 *[]int 类型的指针，不能对指针执行 append 操作
	2、map 和 channel 建议使用 make() 或字面量的方式初始化，不要用 new()
 */
func test1() {
	list := new([]int)
	list = append(list, 1) //
	fmt.Println(list)
}

/**
	append() 的第二个参数不能直接使用 slice，需使用 … 操作符，
	将一个切片追加到另一个切片上：append(s1,s2…)。或者直接跟上元素，形如：append(s1,1,2,3)
 */
func test2() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5}
	s1 = append(s1, s2)
	fmt.Println(s1)
}

//var (
//	size     := 1024
//	max_size = size * 2
//)

func main() {
	test2()
	//fmt.Println(size, max_size)
}
