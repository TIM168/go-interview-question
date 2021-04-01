package main

import "fmt"

/**
常量是一个简单值的标识符，在程序运行时，
不会被修改的量。不像变量，常量未使用是能编译通过的
 */
func test1() {
	const x = 123
	const y = 1.23
	fmt.Println(x)
}


const (
	x uint16 = 120
	y
	s = "abc"
	z
)

/**
常量组中如不指定类型和初始化值，则与上一行非空常量右值相同
*/
func test2() {
	fmt.Printf("%T %v\n",y,y)
	fmt.Printf("%T %v\n",z,z)
}


func test3() {
	//var x string = nil
	var x string

	//if x == nil {
	if x == "" {
		x = "default"
	}
	fmt.Println(x)
}

func main() {
	test1()
	test2()
	test3()
}
