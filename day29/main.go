package main

import (
	"fmt"
	"time"
)

func test1() {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
		fmt.Println(v)
	}
}

func test2() {
	var m = [...]int{1, 2, 3}

	for i, v := range m {
		go func(i,v int) {
			fmt.Println(i, v)
		}(i,v)
	}

	time.Sleep(time.Second * 3)
}

func main() {
	test1()
	test2()
}
