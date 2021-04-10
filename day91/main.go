package main

import "fmt"

/**
除 init() 函数之外，一个包内不允许有其他同名函数
*/
//
//func f() {}
//func f() {}

func init() {}
func init() {}

type User map[string]string

func (m map[string]string) Set(Key string, value string) {
	m[key] = value
}

/**
修复代码
*/
func (m User) SetFix(Key string, value string) {
	m[Key] = value
	fmt.Println(m)
}

func test1() {
	m := make(map[string]string)
	m.Set("A", "One")
}

func test2() {
	m := make(User)
	m.SetFix("A", "One")
}

func main() {
	test2()
}
