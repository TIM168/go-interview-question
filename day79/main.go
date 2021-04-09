package main

import "fmt"

/**
1.interface{} 是可以指向任意对象的 Any 类型，是否正确？

A. false
B. true

参考答案及解析：B
*/

type ConfigOne struct {
	Daemon string
}

func (c *ConfigOne) String() string {
	return fmt.Sprintf("print: %v", c)
}

/**
无限递归循环，栈溢出。知识点：类型的 String() 方法。
如果类型定义了 String() 方法，使用 Printf()、Print() 、 Println() 、 Sprintf() 等格式化输出时会自动使用 String() 方法
 */
func test1() {
	c := &ConfigOne{}
	c.String()
}

func main() {
	test1()
}
