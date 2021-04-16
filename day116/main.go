package main

import "fmt"


const (
	one1 = 1 << iota
	two1
)

/**
1 2
 */
func test1(){
	fmt.Println(one1,two1)
}

const (
    greeting = "Hello, Go"
    one = 1 << iota
    two
)

/**
2 4
 */
func test2(){
	fmt.Println(one,two)
}

func main() {
	//test1()
	test2()
}
