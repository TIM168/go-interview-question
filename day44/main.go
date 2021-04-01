package main

import "fmt"

/**
1.使用 make 创建 map 变量时可以指定第二个参数，不过会被忽略。
2.cap() 函数适用于数组、数组指针、slice 和 channel，不适用于 map，可以使用 len() 返回 map 的元素个数
 */
func test1() {
	m := make(map[string]int, 2)
	cap(m)
}

/**
nil 用于表示 interface、函数、maps、slices 和 channels 的“零值”。
如果不指定变量的类型，编译器猜不出变量的具体类型，导致编译错误
 */
func test2() {
	var x = nil
	_ = x
}


type info struct {
	result int
}

func work() (int, error) {
	return 13, nil
}

func test3() {
	var data info
	var err error
	data.result, err = work()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("info %v\n", data)
}

func main() {
	test3()
}
