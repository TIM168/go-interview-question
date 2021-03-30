package main

import "fmt"

func change(s ...int) {
	s = append(s,3)
}

/**
	Go 提供的语法糖...，可以将 slice 传进可变函数，不会创建新的切片。第一次调用 change() 时，append() 操作使切片底层数组发生了扩容，原 slice 的底层数组不会改变；
	第二次调用change() 函数时，使用了操作符[i,j]获得一个新的切片，假定为 slice1，它的底层数组和原切片底层数组是重合的，
	不过 slice1 的长度、容量分别是 2、5，所以在 change() 函数中对 slice1 底层数组的修改会影响到原切片
 */
func test1() {
	slice := make([]int,5,5)
	slice[0] = 1
	slice[1] = 2
	change(slice...)
	fmt.Println(slice)
	change(slice[0:2]...)
	fmt.Println(slice)
}

/**
	切片在 go 的内部结构有一个指向底层数组的指针，当 range 表达式发生复制时，副本的指针依旧指向原底层数组，
	所以对切片的修改都会反应到底层数组上，所以通过 v 可以获得修改后的数组元素
 */
func test2() {
	var a = []int{1,2,3,4,5}
	var r [5]int

	for i,v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
 	}
 	fmt.Println("r = ",r)
 	fmt.Println("a = ",a)
}

func main() {
	test1()
	test2()
}
