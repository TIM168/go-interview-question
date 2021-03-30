package main

import "fmt"

//type MyInt int
//var i int = 1
//var j MyInt = i

//type MyInt int
//var i int = 1
//var j MyInt = (MyInt)i
//
//
type MyInt int
var i int = 1
var j MyInt = MyInt(i)
//
//type MyInt int
//var i int = 1
//var j MyInt = i.(MyInt)

type Integer int
func (a Integer) Add(b Integer) Integer {
	return a + b
}

//type Integer int
//func (a Integer) Add(b *Integer) Integer {
//	return a + *b
//}

//type Integer int
//func (a *Integer) Add(b Integer) Integer {
//	return *a + b
//}

//type Integer int
//func (a *Integer) Add(b *Integer) Integer {
//	return *a + *b
//}

func test1() {
	var a Integer = 1
	var b Integer = 2
	var i interface{} = &a
	sum := i.(*Integer).Add(b)
	fmt.Println(sum)
}

func main() {
	test1()
}
