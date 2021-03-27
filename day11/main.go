package main

import "fmt"

func test1() {
	var i interface{}
	if i == nil {
		fmt.Println("nil")
		return
	}
	fmt.Println("not nil")
}

func test2() {
	s := make(map[string]int)
	delete(s,"h")
	fmt.Println(s["h"])
}

func main() {

}
