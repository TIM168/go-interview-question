package main

/**
1.下面两处打印的值是否相同？请简要说明

参考答案及解析：不同。知识点：栈增长、逃逸分析。每个 groutine 都会分配相应的栈内存，比如 Go 1.11 版本是 2Kb，随着程序运行，
栈内存会发生增长或缩小，协程会重新申请栈内存块。就像这个题目，循环调用 f()，发生深度递归，栈内存不断增大，当超过范围时，会重新申请栈内存，所以 val 的地址会变化。

这道题还有个特别注意的地方，如果将 println() 函数换成 fmt.Println() 会发现，打印结果相同。为什么？
因为函数 fmt.Println() 使变量 val 发生了逃逸，逃逸到堆内存，即使协程栈内存重新申请，val 变量在堆内存的地址也不会改变

 */
func test1() {
	var val int
	println(&val)
	f(10000)
	println(&val)
}

func f(i int) {
	if i--;i == 0 {
		return
	}
	f(i)
}

func f1(i int) {
	if i--;i == 0 {
		return
	}
	f(i)
}

/**
2.下面代码 A 处输出什么？请简要说明

A. true
B. false

参考答案及解析：A。这道题和上一道有一定联系，a 是指向变量 val 的指针，
我们知道 val 变量的地址发生了改变，a 指向 val 新的地址是由内存管理自动实现的
 */
func test2() {
	var val int

	a := &val
	println(a)

	f(10000)

	b := &val
	println(a)
	println(b)

	println(a == b) //A
}


func main() {
	//test1()
	test2()
}
