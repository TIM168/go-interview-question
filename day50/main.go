package main

import "fmt"

type T struct {
	ls []int
}

func foo(t T) {
	t.ls[0] = 100
}

/**
调用 foo() 函数时虽然是传值，但 foo() 函数中，
字段 ls 依旧可以看成是指向底层数组的指针
 */
func test1() {
	var t = T{
		ls: []int{1, 2, 3},
	}
	foo(t)
	fmt.Println(t.ls)
	fmt.Println(t.ls[0])
}

/**
Go 语言的 switch 语句虽然没有"break"，但如果 case 完成程序会默认 break，可以在 case 语句后面加上关键字 fallthrough，
这样就会接着走下一个 case 语句（不用匹配后续条件表达式）。
或者，利用 case 可以匹配多个值的特性。
 */
func test2() {
	isMatch := func(i int) bool {
		switch (i) {
		case 1:
		case 2:
			return true
		}
		return false
	}
	fmt.Println(isMatch(1))
	fmt.Println(isMatch(2))
}

func test3() {
	isMatch := func(i int) bool {
		switch (i) {
		case 1:
			fallthrough
		case 2:
			return true
		}
		return false
	}
	fmt.Println(isMatch(1))
	fmt.Println(isMatch(2))

	match := func(i int) bool{
		switch (i) {
		case 1,2:
			return true
		}
		return false
	}
	fmt.Println(match(1))
	fmt.Println(match(2))
}

func main() {
	test1()
	//test2()
	test3()
}
