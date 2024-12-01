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





package main

func main() {

	strings := []string{"candy", "toy"}
	for i, s := range strings {
		fmt.Println(i, s)
	}
	输出：  0 candy
	       1 toy
	
*****************************************************************************
	numbers := [6]int{1, 2, 3, 5}
	for i, x := range numbers {
	fmt.Printf("第 %d 位 x 的值= %d\n", i, x)
	}
	输出：
	
		    第 0 位 x 的值 = 1
		    第 1 位 x 的值 = 2
		    第 2 位 x 的值 = 3
		    第 3 位 x 的值 = 5
		    第 4 位 x 的值 = 0
		    第 5 位 x 的值 = 0
******************************************************************************
	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0

	读取key和value
		for key, value := range map1 {
			fmt.Printf("key is:%d--value is: %f\n", key, value)
		}

		输出：
		key is:1--value is: 1.000000
	        key is:2--value is: 2.000000
	        key is:3--value is: 3.000000
	
        ******************************************
	for key := range map1 {
		fmt.Printf("key is: %d\n", key)
	}

	 读取 key
	
		key is: 1
		key is: 2
		key is: 3
		key is: 4
        *******************************************
	for _, value := range map1 {
		fmt.Printf("value is: %f\n", value)
	}
	 读取 value
	
		value is: 4.000000
		value is: 1.000000
		value is: 2.000000
		value is: 3.000000
	
}

****************************************************
package main

import "fmt"

func main() {
	a := 100
	b := 200
	res := max(a, b)
	fmt.Print(res)
}

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

输出：200
*********************************************************

package main

import (
	"fmt"
	"main/caculate"
)

func main() {
	fmt.Println(caculate.Add(1, 2))
}

package caculate

func Add(a, b int) int {
	return a + b
}******在另一个子文件夹中
