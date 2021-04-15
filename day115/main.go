package main

import "fmt"

var y int

func f(x int) int {
	return 7
}

/**

1.下面正确的是？

A.
switch y = f(2) {
case y == 7:
  return
}

B.
switch y = f(2); {
case y == 7:
  return
}

C.
switch y = f(2) {
case 7:
  return
}

D.
switch y = f(2); {
case 7:
  return
}


 */

func test1() {
	a := []int{1,2,3,4}
	b := variadic(a...)
	b[0],b[1] = b[1],b[0]
	fmt.Println(a)
}

func variadic(ints ...int) []int {
	return ints
}


func main() {
	test1()
}
