package main

import "fmt"

/**
1.函数执行时，如果由于 panic 导致了异常，则延迟函数不会执行。这一说法是否正确？

A. true
B. false

参考答案及解析：B。
由 panic 引发异常以后，程序停止执行，然后调用延迟函数（defer），就像程序正常退出一样
 */


/**

 */
func test1() {
	a := [3]int{0,1,2}
	s := a[1:2]
	fmt.Println(s)
	fmt.Println(a)

	s[0] = 11
	fmt.Println(s)

	s = append(s,12)
	fmt.Println(s)

	s = append(s,13)
	fmt.Println(s)
	s[0] = 21

	fmt.Println(a)
	fmt.Println(s)
}

func main() {
	test1()
}
