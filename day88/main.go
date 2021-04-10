package main

import (
	"fmt"
	"os"
)

/**
参考答案及解析：能通过编译。

上面的代码可以理解成：

func main() {
    m := make(map[string]int)
    v := m["foo"]
    v++
    m["foo"] = v
    fmt.Println(m["foo"])
}
 */
func test1() {
	m := make(map[string]int)
	m["foo"]++
	fmt.Println(m["foo"])
}


func Foo() error {
	var err *os.PathError = nil

	return err
}

/**
参考答案及解析：nil false。
知识点：接口值与 nil 值。
只有在值和动态类型都为 nil 的情况下，接口值才为 nil。Foo() 函数返回的 err 变量，值为 nil、动态类型为 *os.PathError，与 nil（值为 nil，动态类型为 nil）显然是不相等
 */
func test2() {
	err := Foo()
	fmt.Printf("%#v\n",err)
	fmt.Println(err)
	fmt.Println(err == nil)

}


func main() {
	test1()
	test2()
}
