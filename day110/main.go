package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

type T struct {
	V int
}

func (t *T) Incr(wg *sync.WaitGroup) {
	t.V++
	wg.Done()
}

func (t *T) Print() {
	time.Sleep(1)
	fmt.Print(t.V)
}

/**

A. 输出 12345678910；
B. 输出 0123456789；
C. 输出 9999999999；
D. panic；

参考答案及解析：C。
这道题需要注意的一点是 for range 循环里的变量 t 是临时变量
 */
func test1() {
	var wg sync.WaitGroup
	wg.Add(10)
	var ts = make([]T,10)
	for i := 0;i < 10;i++ {
		ts[i] = T{i}
	}
	for _,t := range ts {
		go t.Incr(&wg)
	}
	wg.Wait()
	for _,t := range ts {
		go t.Print()
	}
	time.Sleep(5 * time.Second)
}

const N = 26


/**
2.下面的代码可以随机输出大小写字母，
尝试在 A 处添加一行代码使得字母先按大写再按小写的顺序输出
 */
func test2() {
	const GOMAXPROCS = 1
	runtime.GOMAXPROCS(GOMAXPROCS)

	var wg sync.WaitGroup
	wg.Add(2 * N)
	for i := 0;i <N;i++ {
		go func(i int) {
			defer wg.Done()
			//A
			//runtime.Gosched()
			fmt.Printf("%c",'a' + i)
		}(i)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("%c",'A'+i)
		}(i)
	}
	wg.Wait()
}

func main() {
	test1()
	test2()
}
