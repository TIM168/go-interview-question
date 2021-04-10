package main

import (
	"fmt"
	"strings"
)

/**
参考答案及解析：输出空字符。这是一个大多数人遇到的坑，
TrimRight() 会将第二个参数字符串里面所有的字符拿出来处理，只要与其中任何一个字符相等，便会将其删除。
想正确地截取字符串，可以参考 TrimSuffix() 函数
 */
func test1() {
	fmt.Println(strings.TrimRight("ABBA", "BA"))
}

/**
参考答案及解析：输出 []。知识点：拷贝切片。copy(dst, src) 函数返回 len(dst)、len(src) 之间的最小值。
如果想要将 src 完全拷贝至 dst，必须给 dst 分配足够的内存空间
 */
func test2() {
	var src, dst []int
	src = []int{1, 2, 3}
	copy(dst, src)
	fmt.Println(dst)
}

func main() {
	//test1()
	test2()
}
