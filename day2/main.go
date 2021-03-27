package main

import "fmt"

/**
	test1函数里 for range 循环的时候会创建每个元素的副本，而不是元素的引用，
	所以 m[key] = &val 取的都是变量 val 的地址，
	所以最后 map 中的所有元素的值都是变量 val 的地址，因为最后 val 被赋值为3，所有输出都是3
 */
func test1() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		m[key] = &val
	}

	for k, v := range m {
		fmt.Println(k, "->", v)  //k -> 地址
		fmt.Println(k, "->", *v) //k -> 值
	}
}

/**
	正确写法:
 */
func test2() {
	slice := []int{0,1,2,3}
	m := make(map[int]*int)

	for key,val := range slice {
		value := val           //这里用value作为中间量，val副本赋值给value
		m[key] = &value
	}

	for k, v := range m {
		fmt.Println(k, "->", *v) //k -> 值
	}

}


func main() {
	test2()
}

