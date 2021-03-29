package main

import "fmt"

/**
	defer 是闭包引用，返回值被修改
 */
func test1() (r int) {
	r = 0

	defer func() {
		r++
	}()
	return
}

/**
	没涉及返回值 r 的操作，所以返回 5
 */
func test2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

/**
	r 是作为函数参数使用，是一份复制，defer 语句里面的 r 和 外面的 r 其实是两个变量，
	里面变量的改变不会影响外层变量 r，所以不是返回 6 ，而是返回 1
 */
func test3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 1
}

func main() {
	fmt.Println(test1())
	fmt.Println(test2())
	fmt.Println(test3())
}
