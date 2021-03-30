package main

import "fmt"

type Foo struct {
	bar string
}

func test1() {
	s1 := []Foo{
		{"A"},
		{"B"},
		{"C"},
	}
	s2 := make([]*Foo,len(s1))
	for i,value := range s1 {
		s2[i] = &value
	}
	fmt.Println(s1[0],s1[1],s1[2])
	fmt.Println(s2[0],s2[1],s2[2])
}

func test2() {
	var m = map[string]int {
		"A":21,
		"B":22,
		"C":23,
	}
	counter := 0
	for k,v := range m {
		if counter == 0 {
			delete(m,"A")
		}
		counter++
		fmt.Println(k,v)
	}
	fmt.Println("counter is ",counter)
}

func main() {
	test1()
	test2()
}
