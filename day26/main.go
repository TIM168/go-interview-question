package main

import "fmt"

const (
	a = iota
	b = iota
)

const (
	name = "name"
	c = iota
	d = iota
)

/**
	iota 在 const 关键字出现时将被重置为0，const中每新增一行常量声明将使 iota 计数一次
 */
func test1() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
}

type People interface {
	Show()
}

type Student struct {

}

func (stu *Student) Show(){

}

/**
	当且仅当动态值和动态类型都为 nil 时，接口类型值才为 nil。上面的代码，给变量 p 赋值之后，p 的动态值是 nil，
	但是动态类型却是 *Student，是一个 nil 指针，所以相等条件不成立
 */
func test2() {
	var s *Student
	if s == nil {
		fmt.Println("s is nil")
	} else {
		fmt.Println("s is not nil")
	}
	var p People = s
	if p == nil {
		fmt.Println("p is nil")
	} else {
		fmt.Println("p is not nil")
	}

}



func main() {
	test1()
	test2()
}
