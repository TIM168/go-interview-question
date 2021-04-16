package main

import (
	"errors"
	"fmt"
)

/**
1.关于slice或map操作，下面正确的是

A.
var s []int
s = append(s,1)

B.
var m map[string]int
m["one"] = 1

C.
var s []int
s = make([]int, 0)
s = append(s,1)

D.
var m map[string]int
m = make(map[string]int)
m["one"] = 1

参考答案及解析：ACD
*/

var ErrDidNotWork = errors.New("did not work")

func DoTheThing(reallyDoIt bool) (err error) {
	if reallyDoIt {
		result, err := tryTheThing()
		fmt.Println(err)
		if err != nil || result != "it worked" {
			err = ErrDidNotWork
		}
	}
	return err
}

func tryTheThing() (string, error) {
	return "", ErrDidNotWork
}

func test1() {
	fmt.Println(DoTheThing(true))
	fmt.Println(DoTheThing(false))
}

func main() {
	test1()
}
