package main

import "fmt"

/**
	常量，Go 语言中的字符串是只读的
 */
func test1() {
	str := "hello"
	//str[0] = 'x'
	fmt.Println(str)
}

func incr(p *int) int {
	*p++
	return *p
}

/**
	指针，incr() 函数里的 p 是 *int 类型的指针，指向的是 main() 函数的变量 p 的地址。
	第 15 行代码是将该地址的值执行一个自增操作，incr() 返回自增后的结果
 */
func test2() {
	p := 1
	incr(&p)
	fmt.Println(p)
}

/**
	可变函数
 */
func add(args ...int) int {

	sum := 0
	for _, arg := range args {
		sum += arg
	}
	return sum
}

func main() {
	test2()
	fmt.Println(add([]int{1, 2, 7}...))
}
