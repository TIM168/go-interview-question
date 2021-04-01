package main

import "fmt"

/**
	自增、自减不在是运算符，只能作为独立语句，而不是表达式；
	不像其他语言，Go 语言中不支持 ++i 和 --i 操作；
 */
func test1() {
	data := []int{1,2,3}
	i := 0
	//++i
	//i++
	fmt.Println(data[i])
	//fmt.Println(data[i++])
}

/**
使用变量简短声明符号 := 时，如果符号左边有多个变量，
只需要保证至少有一个变量是新声明的，并对已定义的变量尽进行赋值操作。
但如果出现作用域之后，就会导致变量隐藏的问题
 */
func test2() {
	x := 1
	fmt.Println(x)
	{
		fmt.Println(x)
		i,x := 2,2
		fmt.Println(i,x)
	}
	fmt.Println(x)
}


func main() {
	test1()
	test2()
}
