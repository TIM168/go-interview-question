package main

type T int

/**
我们将这两道题目放到一块做一个解析，
第一题不能通过编译，第二题可以通过编译。
我们知道不同类型的值是不能相互赋值的，即使底层类型一样，所以第一题编译不通过；
对于底层类型相同的变量可以相互赋值还有一个重要的条件，即至少有一个不是有名类型（named type）
 */
func F(t T) {}

func test1() {
	var q int
	F(q)
}

type M []int

func Z(m M) {}

func test2() {
	var q []int
	Z(q)
}

func main() {
	test2()
}
