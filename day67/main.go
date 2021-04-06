package main

import "fmt"

type T struct {
	n int
}

/**
9 [{0} {9}]。知识点：for-range 切片。for-range 切片时使用的是切片的副本，但不会复制底层数组，
换句话说，此副本切片与原数组共享底层数组
 */
func test1() {
	ts := [2]T{}
	for i := range ts[:] {
		switch i {
		case 0:
			ts[1].n = 9
		case 1:
			fmt.Print(ts[i].n, " ")
		}
	}
	fmt.Print(ts)

}

/**
9 [{3} {9}]。知识点：for-range 切片
 */
func test2() {
	ts := [2]T{}
	for i := range ts[:] {
		switch t := &ts[i]; i {
		case 0:
			t.n = 3
			ts[1].n = 9
		case 1:
			fmt.Print(t.n," ")
		}
	}
	fmt.Print(ts)
}

func main() {
	test1()
	test2()
}
