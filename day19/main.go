package main

import "fmt"

type Person struct {
	age int
}


/**
	1.person.age 此时是将 28 当做 defer 函数的参数，会把 28 缓存在栈中，等到最后执行该 defer 语句的时候取出，即输出 28；
	2.defer 缓存的是结构体 Person{28} 的地址，最终 Person{28} 的 age 被重新赋值为 29，
	所以 defer 语句最后执行的时候，依靠缓存的地址取出的 age 便是 29，即输出 29；
	3.闭包引用，输出 29
 */
func test1() {
	person := &Person{28}

	//1
	defer fmt.Println(person.age)

	//2
	defer func(p *Person) {
		fmt.Println(p.age)
	}(person)

	//3
	defer func() {
		fmt.Println(person.age)
	}()

	person.age = 29
}

func main() {
	test1()
}
