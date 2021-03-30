package main

import "fmt"

func f(n int) (r int) {
	defer func() {
		fmt.Println(r)
		r += n
		recover()
		fmt.Println("2",r)
	}()

	var f func()

	defer f()
	f = func() {
		r += 2
		fmt.Println("1",r)
	}
	return n + 1
}

/**
	第一步执行r = n +1，接着执行第二个 defer，由于此时 f() 未定义，引发异常，随即执行第一个 defer，异常被 recover()，程序正常执行，最后 return
 */
func test1() {
	fmt.Println(f(3))
}

func test2() {
	var a = [5]int{1,2,3,4,5}
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
