package main

import "fmt"

type person struct {
	name string
}

/**
	打印一个 map 中不存在的值时，返回元素类型的零值。
	这个例子中，m 的类型是 map[person]int，因为 m 中不存在 p，所以打印 int 类型的零值，即 0
 */
func test1() {
	var m map[person]int
	p := person{"mike"}
	fmt.Println(m[p])
}

func hello(num ...int) {
	num[0] = 18
}


func test2() {
	i := []int{5,6,7}
	hello(i...)
	fmt.Println(i[0])
}

func main() {
	test2()
}
