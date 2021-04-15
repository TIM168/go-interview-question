package main

import "fmt"

/**
参考答案及解析：200。因为原切片的容量已经满了，执行 append 操作之后会创建一个新的底层数组，
并将原切片底层数组的值拷贝到新的数组，原数组保持不变
*/
func test1() {
	x := []int{100, 200, 300, 400, 500, 600, 700}
	twohundred := &x[1]
	x = append(x, 800)
	for i := range x {
		x[i]++
	}
	fmt.Println(*twohundred)

}

/**
test1()拓展
 */
func test1_1() {
	x := make([]int, 0, 7)
	x = append(x, 100, 200, 300, 400, 500, 600, 700)
	twohundred := &x[1]
	x = append(x, 800)
	for i := range x {
		x[i]++
	}
	fmt.Println(*twohundred) // 输出 200

	x = make([]int, 0, 8) // 指向另一个切片
	x = append(x, 100, 200, 300, 400, 500, 600, 700)
	twohundred = &x[1]
	x = append(x, 800) // 执行 append 操作，容量足够，不会重新申请内存
	for i := range x {
		x[i]++
	}
	fmt.Println(*twohundred) // 输出 201
}


/**
参考答案及解析：输出 []。
对一个切片执行 [i,j] 的时候，i 和 j 都不能超过切片的长度值
 */
func test2() {
	a := []int{0, 1}
	fmt.Printf("%v", a[len(a):])
}

func main() {
	test1()
	test2()
	test1_1()
}
