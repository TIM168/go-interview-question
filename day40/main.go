package main

import "fmt"

/**
1.关于select机制，下面说法正确的是?

	A. select机制用来处理异步IO问题；
	B. select机制最大的一条限制就是每个case语句里必须是一个IO操作；
	C. golang在语言级别支持select关键字；
	D. select关键字的用法与switch语句非常类似，后面要带判断条件；

参考答案及解析：ABC
 */

/**
	有方向的 channel 不可以被关闭
 */
//func Stop(stop <- chan bool) {
//	close(stop)
//}


type Param map[string]interface{}

type Show struct {
	*Param
}

/**
1.map 需要初始化才能使用；
2.指针不支持索引
 */
//func test1() {
//	s := new(Show)
//	s.Param["day"] = 2
//}

/**
test1修复
 */
func test2() {
	s := new(Show)
	p := make(Param)
	p["day"] = 2
	s.Param = &p
	tmp := *s.Param
	fmt.Println(tmp["day"])
}

func main() {
	test2()
}
