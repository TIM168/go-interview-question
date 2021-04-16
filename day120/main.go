package main

import "fmt"

func test1() {
	fmt.Println(len("你好bj!"))
}

func GetValue(m map[int]string, id int) (string, bool) {
	if _, exist := m[id]; exist {
		return "存在数据", true
	}
	return "", false
}

func test2() {
	intmap := map[int]string{
		1: "a",
		2: "bb",
		3: "ccc",
	}

	v, err := GetValue(intmap, 4)
	fmt.Println(v, err)
}

func main() {
	//test1()
	test2()
}
