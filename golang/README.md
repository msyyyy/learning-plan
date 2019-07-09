安装 : [liunx下安装](<http://c.biancheng.net/view/3993.html>)



## go语言的特点

1. 引入包的概念，go语言的每一个文件都要属于一个包，不能单独存在

2. 引入垃圾回收机制，内存自动回收

3. 天然并发(支持大并发)

   - 从语言层面支持并发，实现简单
   - 从**goroutine**，轻量级线程，可实现大并发处理，高效利用多核
   - 基于CPS并发模型实现

4. 吸收管道通信机制，**channel**，可以实现不同goroutine的通信

5. go函数可以支持返回多个值

   ```go
   func getSumAndSub(n1 int,n2 int) (int,int){
       sum := n1 + n2 // go语句后面不需要分号
       sub := n1 - n2
       return sum,sub
   }
   ```

6. 增加切片 slice ,延迟执行 defer



编译:`go build hello.go` `./hello`

​         或者 `go run hello.go`



```
go build  -o myhello.exe hello.go
```



编译时，编译器将程序依赖运行的库文件包含在可执行文件中，所以可执行文件可以在没go环境的机器上运行



```
package main

import (
        "fmt"
        "math/rand"
        "time"
)

// 数据生产者
func producer(header string, channel chan<- string) {
     // 无限循环, 不停地生产数据
     for {
            // 将随机数和字符串格式化为字符串发送给通道
            channel <- fmt.Sprintf("%s: %v", header, rand.Int31())
            // 等待1秒
            time.Sleep(time.Second)
        }
}

// 数据消费者
func customer(channel <-chan string) {
     // 不停地获取数据
     for {
            // 从通道中取出数据, 此处会阻塞直到信道中返回数据
            message := <-channel
            // 打印数据
            fmt.Println(message)
        }
}

func main() {
    // 创建一个字符串类型的通道
    channel := make(chan string)
    // 创建producer()函数的并发goroutine
    go producer("cat", channel)
    go producer("dog", channel)
    // 数据消费函数
    customer(channel)
}
/*
对代码的分析：
第03行，导入格式化（fmt）、随机数（math/rand）、时间（time）包参与编译。
第10行，生产数据的函数，传入一个标记类型的字符串及一个只能写入的通道。
第13行，for{}构成一个无限循环。
第15行，使用rand.Int31()生成一个随机数，使用fmt.Sprintf()函数将header和随机数格式化为字符串。
第18行，使用time.Sleep()函数暂停1秒再执行这个函数。如果在goroutine中执行时，暂停不会影响其他goroutine的执行。
第23行，消费数据的函数，传入一个只能写入的通道。
第26行，构造一个不断消费消息的循环。
第28行，从通道中取出数据。
第31行，将取出的数据进行打印。
第35行，程序的入口函数，总是在程序开始时执行。
第37行，实例化一个字符串类型的通道。
第39行和第40行，并发执行一个生产者函数，两行分别创建了这个函数搭配不同参数的两个goroutine。
第42行，执行消费者函数通过通道进行数据消费。

整段代码中，没有线程创建，没有线程池也没有加锁，仅仅通过关键字 go 实现 goroutine，和通道实现数据交换。
*/
```

【实例】HTTP 文件服务器是常见的 Web 服务之一。开发阶段为了测试，需要自行安装 Apache 或 Nginx 服务器，下载安装配置需要大量的时间。使用 Go语言实现一个简单的 HTTP 服务器只需要几行代码，如下所示。

```
package main

import (
    "net/http"
)

func main() {
    http.Handle("/", http.FileServer(http.Dir(".")))
    http.ListenAndServe(":8080", nil)
}

/*
下面是代码说明：
第 1 行，标记当前文件为 main 包，main 包也是 Go 程序的入口包。
第 3~5 行，导入 net/http 包，这个包的作用是 HTTP 的基础封装和访问。
第 7 行，程序执行的入口函数 main()。
第 8 行，使用 http.FileServer 文件服务器将当前目录作为根目录（/目录）的处理器，访问根目录，就会进入当前目录。
第 9 行，默认的 HTTP 服务侦听在本机 8080 端口。
*/
```

![](./tupian/http.png)

## Go语言工程结构简单

Go语言的源码无须头文件，编译的文件都来自于后缀名为`.go`的源码文件。

Go语言无须解决方案、工程文件和 Make File，只要将工程文件按照 `GOPATH` 的规则进行填充，即可使用 go build/go install 进行编译，编译安装的二进制可执行文件统一放在 bin 文件夹下。

后面的章节会介绍 GOPATH 及 go build/go install 的详细使用方法。

`for循环`

```
for a := 0;a<10;a++{
    // 循环代码
}
```

`if判断`

```
if 表达式{
    // 表达式成立
}
```

`自增i++`

在 Go语言中，自增操作符不再是一个操作符，而是一个语句,只有`i++`

| 目录名 | 说明                                                         |
| ------ | ------------------------------------------------------------ |
| api    | 每个版本的 api 变更差异                                      |
| bin    | go 源码包编译出的编译器（go）、文档工具（godoc）、格式化工具（gofmt） |
| blog   | Go 博客的模板，使用 Go 的网页模板，有一定的学习意义          |
| doc    | 英文版的 Go 文档                                             |
| lib    | 引用的一些库文件                                             |
| misc   | 杂项用途的文件，例如 [Android](http://c.biancheng.net/android/) 平台的编译、git 的提交钩子等 |
| pkg    | Windows 平台编译好的中间文件                                 |
| src    | 标准库的源码                                                 |
| test   | 测试用例                                                     |

# 1. GO 语言的基本语法

## 1.1 go语言变量的声明（使用var关键字）

声明变量的一般形式是使用 var 关键字：`var 变量名 变量类型 `

例如 `var a, b *int`

Go语言的基本类型有：

- bool
- string
- int、int8、int16、int32、int64
- uint、uint8、uint16、uint32、uint64、uintptr
- byte // uint8 的别名
- rune // int32 的别名 代表一个 Unicode 码
- float32、float64
- complex64、complex128

当一个变量被声明之后，系统自动赋予它该类型的零值：int 为 0，float 为 0.0，bool 为 false，string 为空字符串，指针为 nil 等。

**所有的内存在 Go 中都是经过初始化的**

变量的命名规则遵循骆驼命名法，即首个单词小写，每个新单词的首字母大写，例如：numShips 和 startDate 。

使用关键字var和括号，可以将一组变量定义放在一起。

```
var (
    a int
    b string
    c []float32
    d func() bool
    e struct {
        x int
    }
)
```

除 var 关键字外，还可使用更加简短的变量定义和初始化语法。

`名字 := 表达式`

需要注意的是，简短模式（short variable declaration）有以下限制：

- 定义变量，同时显式初始化。
- 不能提供数据类型。
- 只能用在函数内部。

和 var 形式声明语句一样，简短变量声明语句也可以用来声明和初始化一组变量：

`i, j := 0, 1`

```
func main() {
   x:=100
   a,s:=1, "abc"
}
```

## 1.2 变量的初始化

标准写法 : `var 变量名 类型 = 表达式`

例如 `var hp int = 100`	

编译器推导类型的格式 ：  `var hp = 100`

默认0.17为高精度 float64



短变量声明并初始化 ：`hp :=100`

  这是 Go 语言的推导声明写法，编译器会自动根据右值类型推断出左值的对应类型。

注意：由于使用了`:=`，而不是赋值的`=`，因此推导声明写法的左值变量必须是没有定义过的变量。若定义过，将会发生编译错误。  

注意：在多个短变量声明和赋值中，至少有一个新声明的变量出现在左值中，即便其他变量名可能是重复声明的，编译器也不会报错，代码如下：

```
conn, err := net.Dial("tcp", "127.0.0.1:8080")
conn2, err := net.Dial("tcp", "127.0.0.1:8080")
```

上面的代码片段，编译器不会报err重复定义。



多重赋值

```
var a int = 100
var b int = 200
b, a = a, b
fmt.Println(a, b)


200 100
```

 多重赋值时，变量的左值和右值按从左到右的顺序赋值。

多重赋值在 Go 语言的错误处理和函数返回值中会大量地使用。  

## 2.3 匿名变量

匿名变量特定是一个下划线"_" ,任何值都可以被赋给他，但是会被抛弃，不会再后续代码中使用该变量

```go
func GetData() (int, int) {
    return 100, 200
}

a, _ := GetData()
_, b := GetData()

fmt.Println(a, b)

代码说明如下：
第 5 行只需要获取第一个返回值，所以将第二个返回值的变量设为下画线。
第 6 行将第一个返回值的变量设为匿名。
```

**匿名变量不占用命名空间，不会分配内存。匿名变量与匿名变量之间也不会因为多次声明而无法使用。**

## 2.4 变量的作用域

  一个变量（常量、类型或函数）在程序中都有一定的作用范围，称之为作用域。如果一个变量在函数体外声明，则被认为是全局变量，可以在整个包甚至外部包（被导出后）使用，不管你声明在哪个源文件里或在哪个源文件里调用该变量。

在函数体内声明的变量称之为局部变量，它们的作用域只在函数体内，参数和返回值变量也是局部变量。在今后的学习中我们将会学习到像 if 和 for 这些控制结构，而在这些结构中声明的变量的作用域只在相应的代码块内。一般情况下，局部变量的作用域可以通过代码块（用大括号括起来的部分）判断。  

不要将作用域和生命周期混为一谈。声明语句的作用域对应的是一个源代码的文本区域；它是一个编译时的属性。一个变量的生命周期是指程序运行时变量存在的有效时间段，在此时间区域内它可以被程序的其他部分引用；是一个运行时的概念。



和 for 循环类似，if 和 switch 语句也会在条件部分创建隐式词法域，还有它们对应的执行体词法域。下面的 if-else 测试链演示了 x 和 y 的有效作用域范围：

```go
if x := f(); x == 0 {
    fmt.Println(x)
} else if y := g(x); x == y {
    fmt.Println(x, y)
} else {
    fmt.Println(x, y)
}
fmt.Println(x, y) // 编译错误: x 和 y 未定义

第二个 if 语句嵌套在第一个内部，因此第一个 if 语句条件初始化词法域声明的变量在第二个 if 中也可以访问。switch 语句的每个分支也有类似的词法域规则：条件部分为一个隐式词法域，然后每个是每个分支的词法域。
```

```go
if f, err := os.Open(fname); err != nil { // 编译错误: unused: f
    return err
}
f.ReadByte() // 编译错误: undefined f
f.Close()    // 编译错误: undefined f
```

【例子 】获取当前的工作目录然后保存到一个包级的变量中。这可以本来通过直接调用 os.Getwd 完成

如果要正确更新包级变量，最好不用 `:=` 

```go
var cwd string

func init() {
    var err error    //通过单独声明 err 变量，来避免使用 := 的简短声明方式：
    cwd, err = os.Getwd()
    if err != nil {
        log.Fatalf("os.Getwd failed: %v", err)
    }
}
```

在二进制传输、读写文件的结构描述时，为了保持文件的结构不会受到不同编译目标平台字节长度的影响，不要使用 int 和 uint。

小数点前面或后面的数字都可能被省略（例如 .707 或 1.）。很小或很大的数最好用科学计数法书写，通过 e 或 E 来指定指数部分：

```
const Avogadro = 6.02214129e23  // 阿伏伽德罗常数
const Planck   = 6.62606957e-34 // 普朗克常数
```

用 Printf 函数的 %g 参数打印浮点数，将采用更紧凑的表示形式打印，并提供足够的精度，但是对应表格的数据，使用 %e（带指数）或 %f 的形式打印可能更合适。所有的这三个打印形式都可以指定打印的宽度和控制打印精度。

内置的 complex 函数用于构建复数，内建的 real 和 imag 函数分别返回复数的实部和虚部：

```go
var x complex128 = complex(1, 2) // 1+2i
var y complex128 = complex(3, 4) // 3+4i
fmt.Println(x*y)                 // "(-5+10i)"
fmt.Println(real(x*y))           // "-5"
fmt.Println(imag(x*y))           // "10"
```

  函数 real(c) 和 imag(c) 可以分别获得相应的实数和虚数部分。

在使用格式化说明符时，可以使用 %v 来表示复数，但当你希望只表示其中的一个部分的时候需要使用 %f。

复数支持和其它数字类型一样的运算。当你使用等号 == 或者不等号 != 对复数进行比较运算时，注意对精确度的把握。cmath 包中包含了一些操作复数的公共方法。如果你对内存的要求不是特别高，最好使用 complex128 作为计算类型，因为相关函数都使用这个类型的参数。  

## 2.5 sin图像

  在 Go 语言中，正弦函数由 math 包提供，函数入口为 math.Sin。正弦函数的参数为 float64，返回值也是 float64。在使用正弦函数时，根据实际精度可以进行转换。

Go 语言的标准库支持对图片像素进行访问，并且支持输出各种图片格式，如 JPEG、PNG、GIF 等。  

```go
package main

import (
    "image"
    "image/color"
    "image/png"
    "log"
    "math"
    "os"
)

func main() {
	
	//  设置图片背景色
	
	
    // 图片大小
    const size = 300
    // 根据给定大小创建灰度图
    pic := image.NewGray(image.Rect(0, 0, size, size))

    // 遍历每个像素
    for x := 0; x < size; x++ {
        for y := 0; y < size; y++ {
            // 填充为白色
            pic.SetGray(x, y, color.Gray{255})
        }
    }
    /*
    代码说明如下：
    第 2 行，声明一个 size 常量，值为 300。
    第 5 行，使用 image 包的 NewGray() 函数创建一个图片对象，使用区域由 image.Rect 结构提供。image.Rect 描述一个方形的两个定位点 (x1,y1) 和 (x2,y2)。image.Rect(0,0,size,size) 表示使用完整灰度图像素，尺寸为宽 300，长 300。
    第 8 行和第 9 行，遍历灰度图的所有像素。
    第 11 行，将每一个像素的灰度设为 255，也就是白色。
    */
    
    //   绘制正弦函数轨迹
    
    
    // 从0到最大像素生成x坐标
    for x := 0; x < size; x++ {

        // 让sin的值的范围在0~2Pi之间
        s := float64(x) * 2 * math.Pi / size

        // sin的幅度为一半的像素。向下偏移一半像素并翻转
        y := size/2 - math.Sin(s)*size/2

        // 用黑色绘制sin轨迹
        pic.SetGray(x, int(y), color.Gray{0})
    }

	//  写入图片文件
	
	
    // 创建文件
    file, err := os.Create("sin.png")

    if err != nil {
        log.Fatal(err)
    }
    // 使用png格式将数据写入文件
    png.Encode(file, pic) //将image信息写入文件中

    // 关闭文件
    file.Close()
    /*
    第 2 行，创建 sin.png 的文件。
    第 4 行，如果创建文件失败，返回错误，打印错误并终止。
    第 8 行，使用 PNG 包，将图形对象写入文件中。
    第 11 行，关闭文件。
    */
}
```

![](./goproject/src/go_code/project01/main/sin.png)



## 2.6 bool 类型

布尔值可以和 &&（AND）和 ||（OR）操作符结合，并且有短路行为：如果运算符左边值已经可以确定整个布尔表达式的值，那么运算符右边的值将不再被求值，因此下面的表达式总是安全的：

`s != "" && s[0] == 'x'`

布尔值并不会隐式转换为数字值 0 或 1，反之亦然。不允许布尔类型强制转换int

``` go
flag:=true
fmt.Println((flag)*1) // 会报错
```



如果需要经常做类似的转换, 包装成一个函数会更方便:

```go
// 如果b为真，btoi返回1；如果为假，btoi返回0
func btoi(b bool) int {
    if b {
        return 1
    }
    return 0
}
```

数字到布尔型的逆转换则非常简单, 不过为了保持对称, 我们也可以包装一个函数:

```go
// itob报告是否为非零。
func itob(i int) bool { return i != 0 }
```



