package main

import "fmt"

/**
1.下面列举的是 recover() 的几种调用方式，哪些是正确的？

A.
func main() {
   recover()
   panic(1)
}

B.
func main() {
    defer recover()
    panic(1)
}

C.
func main() {
   defer func() {
        recover()
    }()
    panic(1)
}

D.
func main() {
    defer func() {
        defer func() {
            recover()
       }()
    }()
    panic(1)
}

参考答案及解析：C。recover() 必须在 defer() 函数中直接调用才有效。
上面其他几种情况调用都是无效的：直接调用 recover()、在 defer() 中直接调用 recover() 和 defer() 调用时多层嵌套。
*/

/**
recover() 必须在 defer() 函数中调用才有效，所以第 56 行代码捕获是无效的。
在调用 defer() 时，便会计算函数的参数并压入栈中，所以执行第 52 行代码时，此时便会捕获 panic(2)；此后的 panic(1)，会被上一层的 recover() 捕获。所以输出 21
 */
func test1() {
	defer func() {
		fmt.Print(recover())
	}()

	defer func() {
		defer fmt.Print(recover())
		panic(1)
	}()

	defer recover()
	panic(2)
}

func main() {
	test1()
}
