# Go语言

##  一、简介

### 1. 语言特色

- 简介、快速、安全
- 并行、有趣、开源
- 内存管理、数组安全、编译迅速

### 2. 语言用途

用于搭载web服务器，存储集群或类似用途的巨型中央服务器的系统编程语言。

 对于高性能分布式系统领域而言，Go 语言无疑比大多数其它语言有着更高的开发效率。它提供了海量并行的支持，这对于游戏服务端的开发而言是再好不过了。 

### 3. 主要特性

- 自动垃圾回收
- 更丰富的内置类型
- 函数多返回值
- 错误处理
- 匿名函数和闭包
- 类型和接口
- 并发编程
- 反射
- 语言交互性

## 二、环境搭建

### 1.  下载

安装包下载地址

1.  https://golang.org/dl/ 
2.  https://golang.google.cn/dl/ 

### 2.  Windows系统下安装

Windows 下可以使用 .msi 后缀(在下载列表中可以找到该文件，如go1.4.2.windows-amd64.msi)的安装包来安装。

默认情况下 **.msi** 文件会安装在 **c:\Go** 目录下。你可以将 **c:\Go\bin** 目录添加到 **Path** 环境变量中。添加后你需要重启命令窗口才能生效。

### 3. Unix/Linux/Mac OS X或FreeBSD安装

1. 下载二进制包： go1.4.linux-amd64.tar.gz 

2. 解压至/usr/local目录

   ```shell
   tar -C /usr/local -xzf  go1.4.linux-amd64.tar.gz
   ```

3. 将/usr/local/go/bin添加到PATH环境变量

   ```shell
   export PATH=$PATH:/usr/local/go/bin
   ```

## 三、语言结构

### 1. 基础组成

- 包声明
- 引入包
- 函数
- 变量
- 语句&表达式
- 注释

### 2.  实例

```go
package main  //定义包名，指明这个文件属于哪个包

import "fmt"   // 告诉编译器需要fmt包，fmt包实现了格式化IO（输入/输出）

func main() {
   /* 这是我的第一个简单的程序 */
   fmt.Println("Hello, World!")
}
```

 当标识符（包括常量、变量、类型、函数名、结构字段等等）以一个大写字母开头，如：Group1，那么使用这种形式的标识符的对象就可以被外部包的代码所使用（客户端程序需要先导入这个包），这被称为导出（像面向对象语言中的 public）；标识符如果以小写字母开头，则对包外是不可见的，但是他们在整个包的内部是可见并且可用的（像面向对象语言中的 protected  ）。 

*注意*： 

需要注意的是 { 不能单独放在一行，所以以下代码在运行时会产生错误： 

```go
func main()  
{  // 错误，{ 不能在单独的行上
    fmt.Println("Hello, World!")
}
```

### 3. 执行Go程序

输入命令

```shell
go run xxx.go
```

还可以使用go build命令来生成二进制文件

```shell
go build xxx.go
```

### 4. 一般结构

```go
// 当前程序的包名
package main

// 导入其他包
import . "fmt"

// 常量定义
const PI = 3.14

// 全局变量的声明和赋值
var name = "gopher"

// 一般类型声明
type newType int

// 结构的声明
type gopher struct{}

// 接口的声明
type golang interface{}

// 由main函数作为程序入口点启动
func main() {
    Println("Hello World!")
}
```



## 四、基础语法

### 1. 标记

Go 程序可以由多个标记组成，可以是关键字，标识符，常量，字符串，符号。如以下 GO 语句由 6 个标记组成：

```go
fmt.Println("Hello, World!")
```

### 2. 行分隔符

在 Go 程序中，一行代表一个语句结束。每个语句不需要像 C 家族中的其它语言一样以分号 ; 结尾，因为这些工作都将由 Go 编译器自动完成。

如果你打算将多个语句写在同一行，它们则必须使用 ; 人为区分，但在实际开发中我们并不鼓励这种做法。

### 3. 注释

```go
// 单行注释
/*
多行注释
多行注释
*/
```

### 4. 标识符

第一个字符必须是字母或下划线而不能是数字。

以下是有效的标识符：

```
mahesh   kumar   abc   move_name   a_123
myname50   _temp   j   a23b9   retVal
```

### 5. 字符串连接

通过‘+’进行连接：“Google” + "Runoob"  ——> “GoogleRunoob”

### 6. 关键字

|  break  |   default   |  func  | interface | select |
| :-----: | :---------: | :----: | :-------: | :----: |
|  case   |    defer    |   go   |    map    | struct |
|  chan   |    else     |  goto  |  package  | switch |
|  const  | fallthrough |   if   |   range   |  type  |
| ontinue |     for     | import |  return   |  var   |

### 7. Go语言的空格

 Go 语言中变量的声明必须使用空格隔开，如： 

```go
var age int
```

## 五、数据类型

### 1. 布尔型

 布尔型的值只可以是常量 true 或者 false。一个简单的例子：var b bool = true。 

### 2. 数字类型

 整型 int 和浮点型 float32、float64，Go 语言支持整型和浮点型数字，并且支持复数，其中位的运算采用补码。 

| 序号 | 类型和描述                                                   |
| ---- | ------------------------------------------------------------ |
| 1    | **uint8**：无符号 8 位整型 (0 到 255)                        |
| 2    | **uint16**：无符号 16 位整型 (0 到 65535)                    |
| 3    | **uint32**：无符号 32 位整型 (0 到 4294967295)               |
| 4    | **uint64**：无符号 64 位整型 (0 到 18446744073709551615)     |
| 5    | **int8**：有符号 8 位整型 (-128 到 127)                      |
| 6    | **int16**：有符号 16 位整型 (-32768 到 32767)                |
| 7    | **int32**：有符号 32 位整型 (-2147483648 到 2147483647)      |
| 8    | **int64**：有符号 64 位整型 (-9223372036854775808 到 9223372036854775807) |

### 3. 字符串类型

 字符串就是一串固定长度的字符连接起来的字符序列。Go 的字符串是由单个字节连接起来的。Go 语言的字符串的字节使用 UTF-8 编码标识 Unicode 文本。 

### 4. 其他数字类型

| 序号 | 类型和描述                             |
| ---- | -------------------------------------- |
| 1    | byte：类似unit8                        |
| 2    | rune：int32                            |
| 3    | uint:：32或64位                        |
| 4    | int：与uint一样大小                    |
| 5    | unintptr：无符号整型，用于存放一个指针 |

## 六、语言变量

### 1. 变量声明

- 第一种：指定变量类型，如果没有初始化，则变量默认为零值

   零值就是变量没有做初始化时系统默认设置的值。 

  ```go
  package main
  import "fmt"
  func main() {
  
      // 声明一个变量并初始化
      var a = "RUNOOB"
      fmt.Println(a)
  
      // 没有初始化就为零值
      var b int
      fmt.Println(b)
  
      // bool 零值为 false
      var c bool
      fmt.Println(c)
  }
  
  // 结果
  // RUNOOB
  // 0
  // false
  ```

  1. 数值类型默认为0

  2. 布尔类型默认为false

  3. 字符串默认为""(空字符串)

  4. 以下几种类型为nil:

     ```go
     var a *int
     var a []int
     var a map[string] int
     var a chan int
     var a func(string) int
     var a error // error 是接口
     ```

- 第二种：根据值自行判定变量类型

  ```go
  package main
  import "fmt"
  func main() {
      var d = true
      fmt.Println(d)
  }
  
  // 结果
  // true
  ```

- 第三种：省略var，注意**:=**左侧如果没有声明新的变量，就会产生编译错误，格式：

  ```go
  var intVal int 
  
  intVal :=1 // 这时候会产生编译错误
  
  intVal,intVal1 := 1,2 // 此时不会产生编译错误，因为有声明新的变量，因为 := 是一个声明语句
  ```

   可以将 var f string = "Runoob" 简写为 f := "Runoob"： 

### 2. 多变量声明

​	```

```go
//类型相同多个变量, 非全局变量
var vname1, vname2, vname3 type
vname1, vname2, vname3 = v1, v2, v3

var vname1, vname2, vname3 = v1, v2, v3 // 和 python 很像,不需要显示声明类型，自动推断

vname1, vname2, vname3 := v1, v2, v3 // 出现在 := 左侧的变量不应该是已经被声明过的，否则会导致编译错误


// 这种因式分解关键字的写法一般用于声明全局变量
var (
    vname1 v_type1
    vname2 v_type2
)
```

````go
package main

var x, y int
var (  // 这种因式分解关键字的写法一般用于声明全局变量
    a int
    b bool
)

var c, d int = 1, 2
var e, f = 123, "hello"

//这种不带声明格式的只能在函数体中出现
//g, h := 123, "hello"

func main(){
    g, h := 123, "hello"
    println(x, y, a, b, c, d, e, f, g, h)
}

// 结果
// 0 0 0 false 1 2 123 hello 123 hello
````

### 3. 值类型和引用类型

 所有像 int、float、bool 和 string 这些基本类型都属于值类型，使用这些类型的变量直接指向存在内存中的值

 当使用等号 `=` 将一个变量的值赋值给另一个变量时，如：`j = i`，实际上是在内存中将 i 的值进行了拷贝，此时 i 和 j 的内存地址是不一样的。

 你可以通过 &i 来获取变量 i 的内存地址，例如：0xf840000040（每次的地址都可能不一样）。值类型的变量的值存储在栈中。 

 内存地址会根据机器的不同而有所不同，甚至相同的程序在不同的机器上执行后也会有不同的内存地址。因为每台机器可能有不同的存储器布局，并且位置分配也可能不同。 

 更复杂的数据通常会需要使用多个字，这些数据一般使用引用类型保存。 

 一个引用类型的变量 r1 存储的是 r1 的值所在的内存地址（数字），或内存地址中第一个字所在的位置。 



### 4. 简短形式，使用 := 赋值操作符

 我们知道可以在变量的初始化时省略变量的类型而由系统自动推断，声明语句写上 var 关键字其实是显得有些多余了，因此我们可以将它们简写为 a := 50 或 b := false。 

 a 和 b 的类型（int 和 bool）将由编译器自动推断。 

 这是使用变量的首选形式，但是它只能被用在函数体内，而不可以用于全局变量的声明与赋值。使用操作符 := 可以高效地创建一个新的变量，称之为初始化声明。



## 七、语言常量

### 1. 常量定义格式

```go
const identifier [type] = value
```

​	 你可以省略类型说明符 [type]，因为编译器可以根据变量的值来推断其类型。 

- 显式类型定义：const b string = "abc"
- 隐式类型定义：const b = "abc"

多个相同类型的声明可以简写为：

```go
const c_name1, c_name2 = value1, value2
```

常量还可以用作枚举：

```go
const (
    Unknown = 0
    Female = 1
    Male = 2
)

// 数字 0、1 和 2 分别代表未知性别、女性和男性。
```

 常量可以用len(), cap(), unsafe.Sizeof()函数计算表达式的值。常量表达式中，函数必须是内置函数，否则编译不过： 

```go
package main

import "unsafe"
const (
    a = "abc"
    b = len(a)
    c = unsafe.Sizeof(a)
)

func main(){
    println(a, b, c)
}

// abc 3 16
```

###  2. iota

iota，特殊常量，可以认为是一个可以被编译器修改的常量。

iota 在 const关键字出现时将被重置为 0(const 内部的第一行之前)，const 中每新增一行常量声明将使 iota 计数一次(iota 可理解为 const 语句块中的行索引)。

iota 可以被用作枚举值：

```go
package main

import "fmt"

func main() {
    const (
            a = iota   //0
            b          //1
            c          //2
            d = "ha"   //独立值，iota += 1
            e          //"ha"   iota += 1
            f = 100    //iota +=1
            g          //100  iota +=1
            h = iota   //7,恢复计数
            i          //8
    )
    fmt.Println(a,b,c,d,e,f,g,h,i)
}
```

## 八、运算符

### 1. 算术运算符

| 运算符 | 描述 |
| :----: | :--: |
|   +    | 相加 |
|   -    | 相减 |
|   *    | 相乘 |
|   /    | 相除 |
|   %    | 求余 |
|   ++   | 自增 |
|        | 自减 |



### 2.关系运算符

假定A = 10 ，B  =  20

| 运算符 |        实例        |
| :----: | :----------------: |
|   ==   | (A == B)  为 False |
|   !=   |  (A != B) 为 True  |
|   >    |  (A > B) 为 False  |
|   <    |  (A < B) 为 True   |
|   >=   | (A >= B) 为 False  |
|   <=   |  (A <= B) 为 True  |



### 3. 逻辑运算符

假定A = True ，B = False

| 运算符 | 实例              |
| ------ | ----------------- |
| &&     | (A && B) 为 False |
| \|\|   | (A \|\| B) 为True |
| ！     | !(A && B) 为 True |

### 4. 其他运算符

位运算符，赋值运算符等不再详细介绍，有基础的同学跳过此节介绍，无基础的同学百度学习运算符，各语言的运算符都是通用的。



## 九、条件语句

### 1. if

```go
if 布尔表达式 {
   /* 在布尔表达式为 true 时执行 */
}
```

### 2. if ... else...

```go
if 布尔表达式 {
   /* 在布尔表达式为 true 时执行 */
} else {
  /* 在布尔表达式为 false 时执行 */
}
```

### 3. if嵌套

```go
if 布尔表达式 1 {
   /* 在布尔表达式 1 为 true 时执行 */
   if 布尔表达式 2 {
      /* 在布尔表达式 2 为 true 时执行 */
   }
}
```

### 4. if...else if...else

```go
if 布尔表达式 {
   /* 在布尔表达式为 true 时执行 */
} else if{
  /* 在布尔表达式为 false 时执行 */
} else {
  /* 在布尔表达式为 false 时执行 */
}
```

### 5. switch

```go
switch var1 {
    case val1:
        ...
    case val2:
        ...
    default:
        ...
}
```

## 十、循环语句

### 1. 语法结构

```go
for init; condition; post { }  // == for(init:condition:post){ }
for condition { }              // == for(:condition:){ }
for { }						 // == for(::){ }

/*
init: 一般为赋值表达式，给控制变量赋初值；
condition:关系表达式或逻辑表达式，循环控制条件；
post:一般为赋值表达式，给控制变量增量或减量。
*/
```

### 2. 实例1——计算1到10的数字之和

```go
package main

import "fmt"

func main() {
        sum := 0
        for i := 0; i <= 10; i++ {
                sum += i
        }
        fmt.Println(sum)
}
```

### 3. 实例2——在 sum 小于 10 的时候计算 sum 自相加后的值 

```go
package main

import "fmt"

func main() {
        sum := 1
        for ; sum <= 10; {
                sum += sum
        }
        fmt.Println(sum)

        // 这样写也可以，更像 While 语句形式
        for sum <= 10{
                sum += sum
        }
        fmt.Println(sum)
}
```

### 4. 无限循环

```go
package main

import "fmt"

func main() {
        sum := 0
        for {
            sum++ // 无限循环下去
        }
        fmt.Println(sum) // 无法输出
}
//要停止无限循环，可以在命令窗口按下ctrl-c 。
```

### 5. For-each range循环

```go
package main
import "fmt"

func main() {
        strings := []string{"google", "runoob"}
        for i, s := range strings {
                fmt.Println(i, s)
        }


        numbers := [6]int{1, 2, 3, 5}
        for i,x:= range numbers {
                fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
        }  
}
/*结果
0 google
1 runoob
第 0 位 x 的值 = 1
第 1 位 x 的值 = 2
第 2 位 x 的值 = 3
第 3 位 x 的值 = 5
第 4 位 x 的值 = 0
第 5 位 x 的值 = 0
*/
```



## 十一、语言函数

### 1. 函数定义

```go
func function_name( [parameter list] ) [return_types] {
   函数体
}
```

- func：函数由 func 开始声明

- function_name：函数名称，函数名和参数列表一起构成了函数签名。

- parameter list：参数列表，参数就像一个占位符，当函数被调用时，你可以将值传递给参数，这个值被称为实际参数。参数列表指定的是参数类型、顺序、及参数个数。参数是可选的，也就是说函数也可以不包含参数。

- return_types：返回类型，函数返回一列值。return_types 是该列值的数据类型。有些功能不需要返回值，这种情况下 return_types 不是必须的。

- 函数体：函数定义的代码集合。

  ```go
  /* 函数返回两个数的最大值 */
  func max(num1, num2 int) int {
     /* 声明局部变量 */
     var result int
  
     if (num1 > num2) {
        result = num1
     } else {
        result = num2
     }
     return result
  }
  ```

### 2. 函数调用

```go
package main

import "fmt"

func main() {
   /* 定义局部变量 */
   var a int = 100
   var b int = 200
   var ret int

   /* 调用函数并返回最大值 */
   ret = max(a, b)

   fmt.Printf( "最大值是 : %d\n", ret )
}

/* 函数返回两个数的最大值 */
func max(num1, num2 int) int {
   /* 定义局部变量 */
   var result int

   if (num1 > num2) {
      result = num1
   } else {
      result = num2
   }
   return result
}
```

### 3.返回多个值

```go
package main

import "fmt"

func swap(x, y string) (string, string) {
   return y, x
}

func main() {
   a, b := swap("Google", "Runoob")
   fmt.Println(a, b)
}
```

## 十二、变量作用域

作用域为已声明标识符所表示的常量、类型、变量、函数或包在源代码中的作用范围。

Go 语言中变量可以在三个地方声明：

- 函数内定义的变量称为局部变量
- 函数外定义的变量称为全局变量
- 函数定义中的变量称为形式参数

## 十三、语言数组

### 1.声明数组

```go
var variable_name [SIZE] variable_type
var balance [10] float32 //定义了数组 balance 长度为 10 类型为 float32
```

### 2. 初始化数组

```go
var balance = [5]float32{1000.0, 2.0, 3.4, 7.0, 50.0}  // {}中元素个数不大于[]
var balance = [...]float32{1000.0, 2.0, 3.4, 7.0, 50.0}  // 根据{}元素个数而定
balance[4] = 50.0 
```

### 3. 访问数组元素

```gO
package main

import "fmt"

func main() {
   var n [10]int /* n 是一个长度为 10 的数组 */
   var i,j int

   /* 为数组 n 初始化元素 */        
   for i = 0; i < 10; i++ {
      n[i] = i + 100 /* 设置元素为 i + 100 */
   }

   /* 输出每个数组元素的值 */
   for j = 0; j < 10; j++ {
      fmt.Printf("Element[%d] = %d\n", j, n[j] )
   }
}
```

### 4. 多维数组

```go
var variable_name [SIZE1][SIZE2]...[SIZEN] variable_type   //多维数组

var arrayName [ x ][ y ] variable_type  //二维数组
```

 1. 初始化二维数组

    ```go
    a = [3][4]int{  
     {0, 1, 2, 3} ,   /*  第一行索引为 0 */
     {4, 5, 6, 7} ,   /*  第二行索引为 1 */
     {8, 9, 10, 11},   /* 第三行索引为 2 */
    }
    //注意：以上代码中倒数第二行的 } 必须要有逗号，因为最后一行的 } 不能单独一行，也可以写成这样： 
    a = [3][4]int{  
     {0, 1, 2, 3} ,   /*  第一行索引为 0 */
     {4, 5, 6, 7} ,   /*  第二行索引为 1 */
     {8, 9, 10, 11}}   /* 第三行索引为 2 */
    ```

    

 2. 访问二维数组

    ```go
    val := a[2][3]
    或
    var value int = a[2][3]
    ```

    

## 十四、 语言指针

### 1.定义

一个指针变量指向了一个值的内存地址。

类似于变量和常量，在使用指针前你需要声明指针。指针声明格式如下：

```go
var var_name *var-type
var ip *int        /* 指向整型*/
var fp *float32    /* 指向浮点型 */
```

### 2.使用指针

```go
package main

import "fmt"

func main() {
   var a int= 20   /* 声明实际变量 */
   var ip *int        /* 声明指针变量 */

   ip = &a  /* 指针变量的存储地址 */

   fmt.Printf("a 变量的地址是: %x\n", &a  )

   /* 指针变量的存储地址 */
   fmt.Printf("ip 变量储存的指针地址: %x\n", ip )

   /* 使用指针访问值 */
   fmt.Printf("*ip 变量的值: %d\n", *ip )
}
```

### 3. 空指针

