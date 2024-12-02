#GOlang 

**Day 1**

组成
- 包声明
- 引入包
- 函数
- 变量
- 语句&表达式
- 注释

注：
**‘’ { ‘’ 不能单独放在一行**
**Go 语言中变量的声明必须使用空格隔开**

### 语言变量

- 用 **var** 定义：var identifier type

  例如：var a string = "sun"

  ​           var b,c int = 1, 2

- **1.没有初始化，变量默认为零**

  **2.根据值自行判断变量类型**
  **3.直接赋值 例：intval := 1相当于var intval int  /  intval = 1**

  ​    **注:出现在 := 左侧的变量不应该是已经声明过的**

- **4.全局变量的声明：**

  ```
  var (
      vname1 v_type1
      vname2 v_type2
  )
  ```

- **5.如果在相同的代码块中，我们不可以再次对于相同名称的变量使用初始化声明，
例如：a := 20 就是不被允许的，编译器会提示错误 no new variables on left side of :=，
但是 a = 20 是可以的，因为这是给相同的变量赋予一个新的值**

### 语言常量

- 定义格式： const identifier [type] = value

  显式类型：const b string = "abc"

  隐式类型：const b = "abc"

- **itoa** :const 中每新增一行常量声明将使 iota 计数一次



### 语言运算符

- 基本的算术运算符、关系运算符、逻辑运算符与C语言大同小异，不再赘述

- **位运算符**（虽然和C语言相同，但比较重要）

  A = 0011 1100

  B = 0000 1101

  -----------------

  A&B = 0000 1100 （同1为1，其余为0）

  A|B = 0011 1101   （有1为1，全0为0）

  A^B = 0011 0001   （相同为0，不同为1）

  左移运算符"<<"是双目运算符。左移n位就是乘以2的n次方。

  A << 2 结果为 240 ，二进制为 1111 0000

  右移n位就是除以2的n次方。

  A >> 2 结果为 15 ，二进制为 0000 1111

- **运算符优先级**

  | 优先级 | 运算符           |
  | :----- | :--------------- |
  | 5      | * / % << >> & &^ |
  | 4      | + - \| ^         |
  | 3      | == != < <= > >=  |
  | 2      | &&               |
  | 1      | \|\|             |


# Day2 

## 语言条件语句

- if condition1 {
      // 当 condition1 为 true 时执行的代码
  } else if condition2 {
      // 当 condition1 为 false 且 condition2 为 true 时执行的代码
  } else {
      // 当所有条件都不满足时执行的代码
  }

  ****补充****

  if condition1 && expensiveFunction() {
      // 如果 condition1 为 false，则不会调用 expensiveFunction()
  }

  if condition1 || cheapFunction() {
      // 如果 condition1 为 true，则不会调用 cheapFunction()
  }

- switch expression {
  case value1:
      // 当 expression 等于 value1 时执行的代码
  case value2:
      // 当 expression 等于 value2 时执行的代码
  default:
      // 当没有 case 匹配时执行的代码
  }

  

### 循环语句

- init： 一般为赋值表达式，给控制变量赋初值；

- condition： 关系表达式或逻辑表达式，循环控制条件；

- post： 一般为赋值表达式，给控制变量增量或减量。

  

- 类C语言的 for循环

  for init; condition; post { }

- 类C语言的while

  for condition { }

- 类C的for( ; ; )

  for { }

  

- 例子：

  package main

  import "fmt"

  // 计算 1 到 10 的数字之和
  func main() {
  	sum := 0
  	for i := 0; i <= 10; i++ {
  		sum += i
  	}
  	// 55
  	fmt.Println(sum)
  }

- range 循环，对字符串、数组、切片进行迭代输出

  for 循环的 range 格式可以省略 key 和 value

  

  

**今天补一下输出的一些区别，敲代码的时候注意到的：Print系列函数会将内容输出到系统的标准输出，
区别在于Print函数直接输出内容，Printf函数支持格式化输出字符串，Println函数会在输出内容的结尾添加一个换行符。**

func Print(a ...interface{}) (n int, err error)

func Printf(format string, a ...interface{}) (n int, err error)

func Println(a ...interface{}) (n int, err error)

举例如下：
func main() {
    fmt.Print("在终端打印该信息。")
    name := "tree"
    fmt.Printf("我是：%s\n", name)
    fmt.Println("在终端打印单独一行显示")
}
输出：
    在终端打印该信息。我是：tree
    在终端打印单独一行显示

# Day 3

### 函数

- 定义：

  func function_name( [parameter list] ) [return_types] {
     函数体
  }

  - function_name：函数名称，参数列表和返回值类型构成了函数签名。
  - parameter list：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。
  - return_types：返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的

  **无需多言，和C语言差不多，就是注意参数列表的同种类型声明在最后，且只用一次**

- 函数调用

  - 同一个包直接调用

  - 不同的包在import处声明，例如：

    package caculate

    func ....

    ******************************************************************

    package main

    import(

    "fmt"

    "caculate"

    )

**注意：在调用其他包定义的函数时，只有这个函数名首字母大写的才可以被调用，例如函数名为add就会出现如下情况**



- 不同包的调用：在一个主文件夹中创建子文件夹进行调用：

  例如：在/scr中创建两个文件夹，caculate文件夹以及main文件夹，在caculate文件夹定义函数如Add，在main中定义并进行调用



# Day 4
### 语言变量

- 局部变量：函数体内定义
- 全局变量：函数体外定义

**全局变量与局部变量名称可以相同，但是函数内的局部变量会被优先考虑**



### 数组

- index从0开始，第一个元素索引为 0，第二个索引为 1，以此类推

- 数组声明：var arrayName [size]dataType

  其中，**arrayName** 是数组的名称，**size** 是数组的大小，**dataType** 是数组中元素的数据类型

**数组的大小是类型的一部分，因此不同大小的数组是不兼容的，也就是说 [5]int和 [10]int 是不同的类型。**

**如果数组长度不确定，可以使用 ... 代替数组的长度，编译器会根据元素个数自行推断数组的长度**

- 通过指定下标来初始化元素：

   将索引为 1 和 3 的元素初始化
  balance := [5]float32{1:2.0,3:7.0}







