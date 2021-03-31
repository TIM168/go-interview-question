package main

import "fmt"

var gvar int

/**
未使用变量。如果有未使用的变量代码将编译失败。但也有例外，
函数中声明的变量必须要使用，
但可以有未使用的全局变量。函数的参数未使用也是可以的
*/
func test1() {
	var one int
	two := 2
	var three int
	three = 3

	func(unused string) {
		fmt.Println("Unused arg. No compile error")
	}("what?")
}

type ConfigOne struct {
	Daemon string
}

func (c *ConfigOne) String() string {
	return fmt.Sprintf("print: %v", c)
}

func test2() {
	c := &ConfigOne{}
	c.String()
}

/**
	a 在 for range 过程中增加了两个元素，
len 由 5 增加到 7，但 for range 时会使用 a 的副本 a' 参与循环，副本的 len 依旧是 5，
因此 for range 只会循环 5 次，也就只获取 a 对应的底层数组的前 5 个元素
*/
func test3() {
	var a = []int{1, 2, 3, 4, 5}
	var r = make([]int, 0)

	for i, v := range a {
		if i == 0 {
			a = append(a, 6, 7)
		}

		r = append(r, v)
	}

	fmt.Println(r)
}

func main() {
	//test2()
	test3()
}
