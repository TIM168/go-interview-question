package main

import "fmt"

/**
数组或切片的截取操作

默认底层数组的长度，截取得到的切片长度和容量计算方法是 j-i、l-i
*/
func test1() {
	s := [3]int{1, 2, 3}
	a := s[:0]
	b := s[:2]
	c := s[1:2:cap(s)]
	fmt.Println("a:", a)
	fmt.Println("b:", b)
	fmt.Println("c:", c)
}

/**
	在 A 处只声明了map m ,并没有分配内存空间，不能直接赋值，需要使用 make()，都提倡使用 make() 或者字面量的方式直接初始化 map。
	B 处，v,ok := m["b"] 当 key 为 b 的元素不存在的时候，v 会返回值类型对应的零值，ok 返回 false
 */
func test2() {
	var m map[string]int //A m := make(map[string]int)
	m["a"] = 1
	if v := m["b"]; v != 0 { //B if v,ok := m["b"];ok {
		fmt.Println(v)
	}
}

type A interface {
	ShowA() int
}

type B interface {
	ShowB() int
}

type Work struct {
	i int
}

func (w Work) ShowA() int {
	return w.i + 10
}

func (w Work) ShowB() int {
	return w.i + 20
}

/**
	接口的静态类型。a、b 具有相同的动态类型和动态值，分别是结构体 work 和 {3}；
	a 的静态类型是 A，b 的静态类型是 B，接口 A 不包括方法 ShowB()，接口 B 也不包括方法 ShowA()，编译报错
 */
func test3() {
	c := Work{3}
	var a A = c
	var b B = c
	fmt.Println(a.ShowB())
	fmt.Println(b.ShowA())

}

func main() {
	//test1()
	test2()
	//test3()
}
