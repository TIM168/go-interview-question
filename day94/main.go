package main

/**
参考答案及解析：a。Go 语言里面 ^ 表示按位异或，而不是求幂。

1、0010 ^ 1111 == 1101   (2^15 == 13)
2、0100 ^ 1111 == 1011   (4^15 == 11)
 */

func test1() {
	a := 2 ^ 15
	b := 4 ^ 15
	if a > b {
		println("a")
	} else {
		println("b")
	}
}

/**
参考答案及解析：C() 函数不能通过编译。C() 函数的 default 属于关键字。
string 和 len 是预定义标识符，可以在局部使用。
nil 也可以当做变量使用，不过不建议写这样的代码，可读性不好
 */

func A(string string) string {
	return string + string
}

func B(len int) int {
	return len + len
}

//func C(val,default1 string) string {
//	if val == "" {
//		return default
//	}
//	return val
//}

func main() {
	test1()
}
