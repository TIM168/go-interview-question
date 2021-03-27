package main

import "fmt"


/**
	1、相同类型的结构体才能够进行比较，结构体是否相同不但与属性类型有关，还与属性顺序相关
	2、可比较的有bool、数值型、字符、指针、数组等
	3、切片、map、函数等是不能比较的
 */
func test1() {
	sn1 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	sn2 := struct {
		age  int
		name string
	}{age: 11, name: "qq"}

	if sn1 == sn2 {
		fmt.Println("sn1 == sn2")
	}

	sm1 := struct {
		age int
		m map[string] string
	}{age:11,m: map[string]string{"a":"1"}}

	sm2 := struct {
		age int
		m map[string] string
	}{age:11,m: map[string]string{"a":"1","b":"2"}}

	if sm1 == sm2 {
		fmt.Println("sm1 == sm2")
	}

}

func main() {
	test1()
}
