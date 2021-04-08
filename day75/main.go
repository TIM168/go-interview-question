package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

/**
defer 语句应该放在 if() 语句后面，先判断 err，再 defer 关闭文件句柄。
 */
func test1() {
	f, err := os.Open("file")
	defer f.Close()
	if err != nil {
		return
	}

	b, err := ioutil.ReadAll(f)
	println(string(b))
}

/**
修复代码
 */
func test1Fix(){
	f, err := os.Open("file")

	if err != nil {
		return
	}

	defer f.Close()

	b, err := ioutil.ReadAll(f)
	println(string(b))
}


/**
recover:1。
知识点：panic、recover()。当程序 panic 时就不会往下执行，可以使用 recover() 捕获 panic 的内容
 */
func test2(){
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("recover:%#v",r)
		}
	}()
	panic(1)
	panic(2)
}

func main() {
	test1()
	test2()
}
