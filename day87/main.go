package main

import "fmt"

/**
1.关于协程，下面说法正确是()

A.协程和线程都可以实现程序的并发执行；
B.线程比协程更轻量级；
C.协程不存在死锁问题；
D.通过 channel 来进行协程间的通信；

参考答案及解析：AD
 */

/**
参考答案及解析：false。在 Go 语言里面，^ 作为二元运算符时表示按位异或：对应位，相同为 0，相异为 1。所以第一段代码输出 true 是因为：

10011 ^ 0010 == 0001   (3^2 == 1)
20100 ^ 0010 == 0110   (4^2 == 6)
30101 ^ 0010 == 0111   (5^2 == 7)
1+6=7，这当然是相等的。你来试试分解下第二段代码的数学表达式
*/
func test2() {
	fmt.Println(3^2 + 4^2 == 5^2)
	fmt.Println(6^2 + 8^2 == 10^2)
}


func main() {
	test2()
}