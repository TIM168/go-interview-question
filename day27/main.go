package main

import "fmt"

type Direction int

const (
	North Direction = iota
	East
	South
	West
)

func (d Direction) String() string {
	return [...]string{"North", "East", "South", "West"}[d]
}

/**
	根据 iota 的用法推断出 South 的值是 3；另外，如果类型定义了 String() 方法，
	当使用 fmt.Printf()、fmt.Print() 和 fmt.Println() 会自动使用 String() 方法，实现字符串的打印
 */
func test1() {
	fmt.Println(South)
}

type Math struct {
	x, y int
}

var m = map[string]*Math{
	"foo": {2, 3},
}

/**
	对于类似 X = Y的赋值操作，必须知道 X 的地址，才能够将 Y 的值赋给 X，但 go 中的 map 的 value 本身是不可寻址的
 */
func test2() {
	m["foo"].x = 4
	fmt.Println(m["foo"].x)
}

func main() {
	test1()
	test2()
}
