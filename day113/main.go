package main

import (
	"fmt"
)

/**
1.关于 const 常量定义，下面正确的使用方式是？

A.
const Pi float64 = 3.14159265358979323846
const zero= 0.0

B.
const (
    size int64= 1024
    eof = -1
)

C.
const (
    ERR_ELEM_EXISTerror = errors.New("element already exists")
    ERR_ELEM_NT_EXISTerror = errors.New("element not exists")
)

D.
const u, vfloat32 = 0, 3
const a,b, c = 3, 4, "foo"

参考答案及解析：ABD
*/

//const Pi float64 = 3.14159265358979323846
//const zero= 0.0
//
//const (
//	size int64= 1024
//	eof = -1
//)
//
//const (
//	ERR_ELEM_EXISTerror = errors.New("element already exists")
//	ERR_ELEM_NT_EXISTerror = errors.New("element not exists")
//)
//
//const u, vfloat32 = 0, 3
//const a,b, c = 3, 4, "foo"

func link(p ...interface{}) {
	fmt.Println(p)
}

/**
2.修改下面的代码，使得第二个输出 [seek 1 2 3 4]
*/
func test1() {
	link("seek", 1, 2, 3, 4) //输出 [seek 1 2 3 4]
	a := []int{1, 2, 3, 4}
	s := append(a)
	link("seek", s) // 输出 [seek [1 2 3 4]]

	//以下为实现代码
	tmplink := make([]interface{}, 0, len(a)+1)
	tmplink = append(tmplink, "seek")
	for _, ii := range a {
		tmplink = append(tmplink, ii)
	}
	link(tmplink...)
}

func main() {
	test1()
}
