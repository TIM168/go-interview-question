package main

import "fmt"

func test1() {
	var x interface{}
	var y interface{} = []int{3, 5}
	_ = x == x
	_ = y == y
	_ = y == y

}

var o = fmt.Print

func test2() {
	c := make(chan int, 1)
	for range [3]struct{}{} {
		select {
		default:
			o(1)
		case <-c:
			o(2)
			c = nil
		case c <- -1:
			o(3)
		}
	}
}
func main() {
	test2()
}
