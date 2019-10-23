# 36讲(18) -if语句、for语句和switch语句

<a name="4NZoo"></a>
# if语句、for语句和switch语句

<a name="VoBC7"></a>
### 使用携带`range`子句的`for`语句时需要注意哪些细节？

- 只有1个变量时，该变量代表的是索引

```go
func main() {
    // 注意 [...]类型为数组，是值类型
	numbers2 := [...]int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(numbers2) - 1
	for i, e := range numbers2 {
		if i == maxIndex2 {
			numbers2[0] += e
		} else {
			numbers2[i+1] += e
		}
	}
	fmt.Println(numbers2)

}

[7 3 5 7 9 11]
```

<a name="haTL7"></a>
### **`switch`语句中的`switch`表达式和`case`表达式之间有着怎样的联系？**

由于**需要进行判等操**作，所以**switch**和** case **中的子表达式的结果类型需要相同。

<a name="3MJYw"></a>
### **`switch`语句对它的`case`表达式有哪些约束？**
**`switch`语句在`case`子句的选择上是具有唯一性的。不允许`case`表达式中的子表达式结果值存在相等的情况

**不过可以绕过,方法如下**

```go
value5 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch value5[4] {
	case value5[0], value5[1], value5[2]:
		fmt.Println("0 or 1 or 2")
	case value5[2], value5[3], value5[4]:
		fmt.Println("2 or 3 or 4")
	case value5[4], value5[5], value5[6]:
		fmt.Println("4 or 5 or 6")
	}
```

这种绕过方式对用于类型判断的`switch`语句（以下简称为类型`switch`语句）就无效了。因为类型`switch`语句中的`case`表达式的子表达式，都必须直接由类型字面量表示，而无法通过间接的方式表示

普通`case`子句的编写顺序很重要，最上边的`case`子句中的子表达式总是会被最先求值，在判等的时候顺序也是这样。因此，如果某些子表达式的结果值有重复并且它们与`switch`表达式的结果值相等，那么位置靠上的`case`子句总会被选中。

<a name="VKRlA"></a>
### 在类型`switch`语句中，我们怎样对被判断类型的那个值做相应的类型转换

```go
value1 := [...]int8{0, 1, 2, 3, 4, 5, 6}
	switch 1+3 {
	case int(value1[0]), int(value1[1]):
		fmt.Println("0 or 1")
	case int(value1[2]), int(value1[3]):
		fmt.Println("2 or 3")
	case int(value1[4]), int(value1[5]), int(value1[6]):
		fmt.Println("4 or 5 or 6")
	}
类型转换 ，如果需要使用类型判断，得先转成 interface{}

```

