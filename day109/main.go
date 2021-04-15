package main

import (
	"fmt"
	"sort"
	"time"
)

/**
参考答案及解析：panic。
协程开启还未来得及执行，chan 就已经 close() ，往已经关闭的 chan 写数据会 panic
 */
func test1() {
	ch := make(chan int, 1000)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()
	go func() {
		for {
			a, ok := <-ch
			if !ok {
				fmt.Println("close")
				return
			}
			fmt.Println("a: ", a)
		}
	}()
	close(ch)
	fmt.Println("ok")
	time.Sleep(time.Second * 1)
}

type S struct {
	v int
}


func test2() {
	s := []S{{1}, {3}, {5}, {2}}
	//A
	sort.Slice(s, func(i, j int) bool {
		return s[i].v < s[j].v
	})
	fmt.Printf("%#v", s)

}

func main() {
	//test1()
	test2()
}
