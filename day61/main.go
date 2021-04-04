package main

import "fmt"

func test1() {
	var k = 1
	var s = []int{1, 2}
	k, s[k] = 0, 3
	fmt.Println(s[0] + s[1])
}

func test2() {
	var k = 9
	for k = range []int{} {
	}
	fmt.Println(k)

	for k = 0; k < 3; k++ {
	}

	fmt.Println(k)

	for k = range (*[3]int)(nil) {
	}
	fmt.Println(k)
}

func main() {
	//test1()
	test2()
}
