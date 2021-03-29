package main

import "fmt"

func test1() {
	m := map[int]string{0: "zero", 1: "one"}
	for k, v := range m {
		fmt.Println(k, v)
	}

}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}


/**
	程序执行到 main() 函数三行代码的时候，会先执行 calc() 函数的 b 参数，即：calc("10",a,b)，输出：10 1 2 3，得到值 3，因为
	defer 定义的函数是延迟函数，故 calc("1",1,3) 会被延迟执行；

	程序执行到第五行的时候，同样先执行 calc("20",a,b) 输出：20 0 2 2 得到值 2，同样将 calc("2",0,2) 延迟执行；

	程序执行到末尾的时候，按照栈先进后出的方式依次执行：calc("2",0,2)，calc("1",1,3)，则就依次输出：2 0 2 2，1 1 3 4
 */
func test2() {
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

func main() {
	test1()
	test2()
}
