package main
import "fmt"

//DAY1
func main() {

	d := "runboo"
	fmt.Println(d)
        //输出： runboo

	
        //常量的应用
	const length int = 10
	const width int = 5
	var area int
	const a, b, c = 1, false, "str"

	area = length * width
	fmt.Println(" 面积为 :", area)
	println(a, b, c)
	/*
        输出：
        面积为 : 50
        1 false str
        */

        //itoa的用法
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
	/*
        输出：0 1 2 what what 100 100 7 8
        */

        //第二个例子
        const (
	      a = 1 << iota
	      b = 3 << iota
	      c
	      d
        )
	
	fmt.Println(a, b, c, d)
	/*
        输出：
        1 6 12 24
	*/
        //由以上两个例子发现可能存在的规律是：iota的计算规律决定之后所有的计算准则
}

//DAY2
func main(){
        ***
	strings := []string{"candy", "toy"}
	for i, s := range strings {
		fmt.Println(i, s)
	}
	/*
	输出：  0 candy
	       1 toy
        */

	
        ***
	numbers := [6]int{1, 2, 3, 5}
	for i, x := range numbers {
	fmt.Printf("第 %d 位 x 的值= %d\n", i, x)
	}
	/*
	输出：
	
		    第 0 位 x 的值 = 1
		    第 1 位 x 的值 = 2
		    第 2 位 x 的值 = 3
		    第 3 位 x 的值 = 5
		    第 4 位 x 的值 = 0
		    第 5 位 x 的值 = 0
         */

	
	***
	map1 := make(map[int]float32)
	map1[1] = 1.0
	map1[2] = 2.0
	map1[3] = 3.0

	//读取key和value
	for key, value := range map1 {
		fmt.Printf("key is:%d--value is: %f\n", key, value)
	}
	/*
	输出：
	key is:1--value is: 1.000000
	key is:2--value is: 2.000000
	key is:3--value is: 3.000000
        */

	
        ***
	/*
        读取 key
	输出：
	key is: 1
	key is: 2
	key is: 3
	key is: 4
        */
	for key := range map1 {
		fmt.Printf("key is: %d\n", key)
	}
}

//DAY3
func main(){
        ***
	for _, value := range map1 {
		fmt.Printf("value is: %f\n", value)
	}
	读取 value
	value is: 4.000000
	value is: 1.000000
	value is: 2.000000
	value is: 3.000000
	

        ***
	a := 100
        b := 200
	res := max(a, b)
	fmt.Print(res)

	
        func max(num1, num2 int) int {
	   if num1 > num2 {
		   return num1
	   } else {
		   return num2
	   }
        }
        输出：200



	***
        import (
	   "fmt"
	   "main/caculate"
        )
	fmt.Println(caculate.Add(1, 2))

        package caculate
        func Add(a, b int) int {
	return a + b
        }//在另一个文件夹中
}

//DAY 4
const MAX int = 3
func main() {

   ***
   a := []int{10,100,200}
   var i int
   var ptr [MAX]*int;

   for  i = 0; i < MAX; i++ {
      ptr[i] = &a[i] /* 整数地址赋值给指针数组 */
   }

   for  i = 0; i < MAX; i++ {
      fmt.Printf("a[%d] = %d\n", i,*ptr[i] )
   }
   /*
   输出：
   a[0] = 10
   a[1] = 100
   a[2] = 200
   */

   var a int
   var ptr *int
   var pptr **int
   a = 3000
   /* 指针 ptr 地址 */
   ptr = &a
   /* 指向指针 ptr 地址 */
   pptr = &ptr
   /* 获取 pptr 的值 */
   fmt.Printf("变量 a = %d\n", a )
   fmt.Printf("指针变量 *ptr = %d\n", *ptr )
   fmt.Printf("指向指针的指针变量 **pptr = %d\n", **pptr)
   /*
   输出：
   变量 a = 3000
   指针变量 *ptr = 3000
   指向指针的指针变量 **pptr = 3000
   */

}

//DAY 5

//斐波那契数列
package main

import "fmt"

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-2) + fibonacci(n-1)
}

func main() {
	var i int
	for i = 0; i < 10; i++ {
		fmt.Printf("%d\t", fibonacci(i))
	}
}

