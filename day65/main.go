package main

import "fmt"

/**
1.flag 是 bool 型变量，下面 if 表达式符合编码规范的是？

A. if flag == 1
B. if flag
C. if flag == false
D. if !flag

参考答案及解析：BCD
 */


func test1() {
	defer func() {
		fmt.Print(recover())
	}()
	defer func() {
		defer func() {
			fmt.Print(recover())
		}()
		panic(1)
	}()

	defer recover()
	panic(2)
}

func main() {
	test1()
}
