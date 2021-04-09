package main

/**
1.关于 switch 语句，下面说法正确的是？

A. 单个 case 中，可以出现多个结果选项；
B. 需要使用 break 来明确退出一个 case;
C. 只有在 case 中明确添加 fallthrought 关键字，才会继续执行紧跟的下一个 case;
D. 条件表达式必须为常量或者整数；

参考答案及解析：AC
*/

func alwaysFalse() bool {
	return false
}

/**
可以编译通过，输出：true。知识点：Go 代码断行规则
 */
func test1() {
	switch alwaysFalse()
	{
	case true:
		println(true)
	case false:
		println(false)
	}
}

func main() {
	test1()
}
