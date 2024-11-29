package main

import "fmt"

func main() {
	d := "runboo"
	fmt.Println(d)
}

输出： runboo


常量的应用
package main

import "fmt"

func main() {
	const length int = 10
	const width int = 5
	var area int
	const a, b, c = 1, false, "str"

	area = length * width
	fmt.Println(" 面积为 :", area)
	println()
	println(a, b, c)
}

输出：

面积为 : 50

1 false str


itoa的用法
package main

import "fmt"

func main() {
	const (
		a = iota
		b
		c
		d = "what"
		e
		f = 100
		g
		h = iota
		i
	)
	fmt.Println(a, b, c, d, e, f, g, h, i)
}

输出：0 1 2 what what 100 100 7 8
******************************************
第二个例子
package main

import "fmt"

const (
	a = 1 << iota
	b = 3 << iota
	c
	d
)

func main() {
	fmt.Println(a, b, c, d)
}

输出：
1 6 12 24

由以上两个例子发现可能存在的规律是：iota的计算规律决定之后所有的计算准则

