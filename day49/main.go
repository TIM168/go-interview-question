package main

import (
	"encoding/json"
	"fmt"
)

/**

 */
func test1() {
	//ch := make(chan int, 1)
	var ch chan int
	select {
	case v, ok := <-ch:
		println(v, ok)
	default:
		println("default")
	}
}

/**
结构体访问控制，因为 name 首字母是小写，导致其他包不能访问，所以输出为空结构体
 */
type People struct {
	//name string `json:"name"`
	Name string `json:"name"`
}

func test2() {
	js := `{
		"name":"seekload"
	}`
	var p People
	err := json.Unmarshal([]byte(js), &p)
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println(p)
}

func main() {
	test1()
	test2()
}
