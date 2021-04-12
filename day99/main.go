package main

import "fmt"

/**
参考答案即解析：编译通过。
true 是预定义标识符可以用作变量名，但是不建议这么做
 */
func test1() {
	true := false
	fmt.Println(true)
}

/**
参考答案即解析：100。
知识点：变量作用域和defer 返回值
 */
func watShadowDefer(i int) (ret int) {
	ret = i * 2
	if ret > 10 {
		ret := 10
		defer func() {
			ret = ret + 1
		}()
	}
	return
}

func test2() {
	result := watShadowDefer(50)
	fmt.Println(result)
}

func main() {
	test1()
	test2()
}
