package main

import "fmt"

/**
1.关于 channel 下面描述正确的是？

A. 向已关闭的通道发送数据会引发 panic；
B. 从已关闭的缓冲通道接收数据，返回已缓冲数据或者零值；
C. 无论接收还是接收，nil 通道都会阻塞；

参考答案及解析：ABC
*/

type T struct {
	n int
}

func (t *T) Set(n int) {
	t.n = n
}

/**
1.直接返回的 T{} 不可寻址
2.不可寻址的结构体不能调用带结构体指针接收者的方法
*/
func getT() T {
	return T{}
}

func test1() {
	//getT().Set(1)
	t := getT()
	t.Set(2)
	fmt.Println(t.n)
}

func main() {
	test1()
}
