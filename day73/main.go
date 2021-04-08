package main

/**
闭包延迟求值。for 循环局部变量 i，
匿名函数每一次使用的都是同一个变量。（说明：i 的地址，输出可能与上面的不一样）
 */
func test() []func() {
	var funs []func()
	for i := 0; i < 2; i++ {
		funs = append(funs, func() {
			println(&i, i)
		})
	}
	return funs
}

func test1() {
	funs := test()
	for _, f := range funs {
		f()
	}
}

var f = func(i int) {
	print("x")
}

/**
10x。这道题一眼看上去会输出 109876543210，
其实这是错误的答案，这里不是递归。
假设 main() 函数里为 f2()，外面的为 f1()，当声明 f2() 时，调用的是已经完成声明的 f1()
 */
func test2() {
	f := func(i int) {
		print(i)
		if i > 0 {
			f(i - 1)
		}
	}
	f(10)
}

func main() {
	test1()
	test2()
}
