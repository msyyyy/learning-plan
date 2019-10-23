# 36讲(22) panic 下

<a name="In1WY"></a>
## 22 | panic函数、recover函数以及defer语句（下）

<a name="bPZFv"></a>
### 怎样让 panic 包含一个值，以及应该让它包含什么样的值？
可以用panic函数包裹一个值将其返回，可以包裹interface{}，最好包裹error类型，定义Error方法

<a name="MkThR"></a>
### 怎样施加应对 panic 的保护措施，从而避免程序崩溃？
Go 语言的内建函数`recover`专用于恢复 panic，或者说平息运行时恐慌。`recover`函数无需任何参数，并且会返回一个空接口类型的值。<br />如果用法正确，这个值实际上就是即将恢复的 panic 包含的值。并且，如果这个 panic 是因我们调用`panic`函数而引发的，那么该值同时也会是我们此次调用`panic`函数时，传入的参数值副本

需要调用defer语句   `defer`语句就是被用来延迟执行代码的。延迟到什么时候呢？这要延迟到该语句所在的函数即将执行结束的那一刻，无论结束执行的原因是什么。被延迟执行的是`defer`函数

```go
package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("Enter function main.")
	defer func() {
		fmt.Println("Enter defer function.")
		if p := recover(); p != nil {
			fmt.Printf("panic: %s\n", p)
		}
		fmt.Println("Exit defer function.")
	}()
	// 引发 panic。
	panic(errors.New("something wrong"))
	fmt.Println("Exit function main.")
}
```

<a name="PuD77"></a>
### 如果一个函数中有多条`defer`语句，那么那几个`defer`函数调用的执行顺序是怎样的？
如果只用一句话回答的话，那就是：在同一个函数中，`defer`函数调用的执行顺序与它们分别所属的`defer`语句的出现顺序（更严谨地说，是执行顺序）完全相反。

当一个函数即将结束执行时，其中的写在最下边的`defer`函数调用会最先执行，其次是写在它上边、与它的距离最近的那个`defer`函数调用，以此类推，最上边的`defer`函数调用会最后一个执行。<br />**相当于会把defer函数都放到一个栈中，最后函数结束时执行**

**
<a name="lEpFY"></a>
### 我们可以在`defer`函数中恢复 panic，那么可以在其中引发 panic 吗
可以

```go
package main

import (
	"errors"
	"fmt"
)

func main() {
	defer func() {
		panic(errors.New("some error"))
	}()
	fmt.Println("a")
}

```

