package main

import "fmt"

/**
1.关于 main() 函数，下面说法正确的是？

A、不能带参数；
B、不能定义返回值；
C、所在的包必须为 main 包；
D、可以使用 flag 包来获取和解析命令行参数；

参考答案及解析：ABCD
*/

type User struct {
	Name string
}

func (u *User) SetName(name string) {
	u.Name = name
	fmt.Println(u.Name)
}

type Employee User

/**
参考答案及解析：编译不通过。
当使用 type 声明一个新类型，它不会继承原有类型的方法集
 */
func test1() {
	employee := new(Employee)
	employee.SetName("Jack")
}

func main() {

}
