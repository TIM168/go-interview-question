package main

import "fmt"

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
312。对比昨天的第二题，本题的 s.Add(1).Add(2) 作为一个整体包在一个匿名函数中，会延迟执行
 */
func test1() {
	s := NewSlice()
	defer func() {
		s.Add(1).Add(2)
	}()
	s.Add(3)
}

type Orange struct {
	Quantity int
}

func (o *Orange) Increase(n int) {
	o.Quantity += n
}

func (o *Orange) Decrease(n int) {
	o.Quantity -= n
}

func (o *Orange) String() string {
	return fmt.Sprintf("%#v", o.Quantity)
}

/**
{5}。这道题容易忽视的点是，
String() 是指针方法，而不是值方法，所以使用 Println() 输出时不会调用到 String() 方法
 */
func test2() {
	var orange Orange
	orange.Increase(10)
	orange.Decrease(5)
	fmt.Println(orange)
}

func main() {
	test1()
	test2()
}
