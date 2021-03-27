package main

import "fmt"

func main() {
	defer_call()
}

/**
	1、defer执行顺序后进先出，当出现panic时，会先按defer执行顺序后进先出，最后执行panic
	2、利用recover可以捕获panic异常，防止报错，程序正常执行下去
 */
func defer_call() {
	defer func() {
		fmt.Println("打印前")

	}()

	defer func() {
		fmt.Println("打印中")
	}()

	defer func() {
		fmt.Println("打印后")
		err := recover()
		fmt.Println(err)
	}()

	panic("触发异常")
}