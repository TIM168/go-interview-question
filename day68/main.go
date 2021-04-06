package main

import "fmt"

/**
goto 不能跳转到其他函数或者内层代码。编译报错：
 */
func test1() {
	for i := 0; i < 10; i++ {
	loop:
		println(i)
	}
	goto loop
}

/**
知识点：defer()、for-range。for-range
虽然使用的是 :=，但是 v 不会重新声明，可以打印 v 的地址验证下
 */
func test2() {
	x := []int{0, 1, 2}
	y := [3]*int{}

	for i, v := range x {
		defer func() {
			print(v)
		}()
		fmt.Println(v)
		y[i] = &v
	}
	fmt.Println(y)
	print(*y[0], *y[1], *y[2])

}

func main() {
	test2()
}
