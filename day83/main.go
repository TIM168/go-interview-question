package main

import "fmt"

/**
1.同级文件的包名不允许有多个，是否正确？

A. true
B. false


参考答案及解析：A。一个文件夹下只能有一个包，可以多个.go文件，但这些文件必须属于同一个包
 */

type data struct {
	name string
}

func (p *data) print() {
	fmt.Println("name:",p.name)
}

type printer interface {
	print()
}

/**
参考答案及解析：编译报错。

1cannot use data literal (type data) as type printer in assignment:
2data does not implement printer (print method has pointer receiver)

结构体类型 data 没有实现接口 printer
 */
func test1(){
	d1 := data{"one"}
	d1.print()

	var in printer = data{"two"}
	in.print()
}

func main() {
	
}
