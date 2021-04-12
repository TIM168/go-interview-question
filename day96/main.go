package main

import "fmt"

type Point struct {
	x, y int
}

/**
参考答案及解析：输出 [{1 2} {3 4}]。
知识点：for range 循环。range 循环的时候，获取到的元素值是副本，就比如这里的 p
*/
func test1() {
	s := []Point{
		{1, 2},
		{3, 4},
	}
	for _, p := range s {
		p.x, p.y = p.y, p.x
	}
	fmt.Println(s)
}

/**
修复代码
*/
func test1Fix() {
	s := []*Point{
		&Point{1, 2},
		&Point{3, 4},
	}
	for _, p := range s {
		p.x, p.y = p.y, p.x
	}
	fmt.Println(*s[0])
	fmt.Println(*s[1])
}

func get() []byte {
	raw := make([]byte, 10000)
	fmt.Println(len(raw), cap(raw), &raw[0])
	return raw[:3]
}

/**
参考答案及解析：get() 函数返回的切片与原切片公用底层数组，
如果在调用函数里面（这里是 main() 函数）修改返回的切片，将会影响到原切片
*/
func test2() {
	data := get()
	fmt.Println(len(data), cap(data), &data[0])
}

/**
修复代码
*/
func getFix() []byte {
	raw := make([]byte, 10000)
	fmt.Println(len(raw), cap(raw), &raw[0])
	res := make([]byte, 3)
	copy(res, raw[:3])
	return res
}

func test2Fix() {
	data := getFix()
	fmt.Println(len(data), cap(data), &data[0])
}

func main() {
	test1()
	test2()
	test2Fix()
}
