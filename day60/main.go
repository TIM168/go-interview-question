package main

import (
	"fmt"
	"sync"
	"time"
)

type T struct {
}

func (*T) foo() {

}

func (T) bar() {

}

type S struct {
	*T
}

/**
第31行
因为 s.bar 将被展开为 (*s.T).bar，而 s.T 是个空指针，解引用会 panic
 */
func test1() {
	s := S{}
	fmt.Printf("%#v",s)
	_ = s.foo
	s.foo()
	//_ = s.bar
}

type data struct {
	sync.Mutex
}

/**
锁失效。
将 Mutex 作为匿名字段时，相关的方法必须使用指针接收者，否则会导致锁机制失效
 */
//func (d data) test(s string) {
func (d *data) test(s string) {    // 指针接收者
	d.Lock()
	defer d.Unlock()

	for i := 0; i < 5; i++ {
		fmt.Println(s, i)
		time.Sleep(time.Second)
	}
}



func test2() {
	var wg sync.WaitGroup
	wg.Add(2)
	var d data

	go func() {
		defer wg.Done()
		d.test("read")
	}()

	go func() {
		defer wg.Done()
		d.test("write")
	}()

	wg.Wait()
}

func main() {
	test1()
	//test2()
}
