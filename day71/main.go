package main

import "fmt"

/**
1.判断题：对变量x的取反操作是 ~x？

参考答案及解析：错。Go 语言的取反操作是 ^，它返回一个每个 bit 位都取反的数。作用类似在 C、C#、Java 语言中中符号 ~，
对于有符号的整数来说，是按照补码进行取反操作的（快速计算方法：对数 a 取反，结果为 -(a+1) ），对于无符号整数来说就是按位取反
*/

type Slice []int

func NewSlice() Slice {
	return make(Slice, 0)
}

func (s *Slice) Add(elem int) *Slice {
	*s = append(*s, elem)
	fmt.Print(elem)
	return s
}

/**
参考答案及解析：132。这一题有两点需要注意：
1.Add() 方法的返回值依然是指针类型 *Slice，所以可以循环调用方法 Add()；
2.defer 函数的参数（包括接收者）是在 defer 语句出现的位置做计算的，而不是在函数执行的时候计算的，所以 s.Add(1) 会先于 s.Add(3) 执行
*/
func test1() {
	s := NewSlice()
	defer s.Add(1).Add(2)
	s.Add(3)
}

func main() {
	test1()
}
