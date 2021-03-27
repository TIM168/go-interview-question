package main

import "fmt"


/**
	1、init() 函数是用于程序执行前做包的初始化的函数，比如初始化包里的变量等;
	2、一个包可以出线多个 init() 函数,一个源文件也可以包含多个 init() 函数；
	3、同一个包中多个 init() 函数的执行顺序没有明确定义，但是不同包的init函数是根据包导入的依赖关系决定的（看下图）;
	4、init() 函数在代码中不能被显示调用、不能被引用（赋值给函数变量），否则出现编译错误;
	5、一个包被引用多次，如 A import B,C import B,A import C，B 被引用多次，但 B 包只会初始化一次；
	6、引入包，不可出现死循坏。即 A import B,B import A，这种情况编译失败；
 */
func init() {
	fmt.Println(1)
}

func init() {
	fmt.Println(2)
}

func hello() []string {
	return nil
}

func test1() {
	h := hello
	if h == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}

func GetValue() interface{} {
	return 1
}

/**
	i.(type)，其中 i 是接口，type 是固定关键字，需要注意的是，只有接口类型才可以使用类型选择
 */
func test2() {
	i := GetValue()
	switch i.(type) {
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")

	}
}

func main() {
	test1()
}
