- [c++基础](#1)
    - [变量和基本类型](#1-1)
        - [处理类型](#1-1-5)
        





<h1 id='1'>C++基础</h1>


<h2 id='1-1'>变量和基本类型</h2>

<h3 id='1-1-5'> 处理类型 <h3>

#### 类别别名

 what： 某种类型的同义词

 why ：让复杂的代码更容易理解
 ```
如 typedef char * pstring;
那么 pstring是 char *的别名 如果用到这个别名 他的基础类型就是指针
const pstring cstr=0; //指向char的常指针
const pstring *ps;  // 指向常指针的 一个指针
 ```
 how ： 2种声明方法
 ```
1. typdef double base,*p;
// base是 double的同义词 p是double *的同义词

2. using base = double; 
 ```
 
 #### auto 类型说明符

 what : 让编译器替我们分析表达式所属类型

 why ：我们在声明过程中可能不是很清楚了解表达式类型

 how：
 ```
1. 一条声明语句只能有一个基本数据类型
auto sz = 0,pi = 3.14//错误,sz和pi类型不一致

2. auto一般会忽视顶层const 会保留底层const，auto会找到被引用变量
const int ci, &cr=ci;
auto b = ci ;// 整数 const属性被忽略
auto c = cr;// 整数 auto找到被引用变量
auto e= &ci; // 指向整数常量的指针

3. 希望得到const 类型要明确指出
const auto f = ci // 整数常量

PS.常量引用可以绑定字面值(如 42 )

 ```

 #### decltype提示符

what：选择并返回操作数（函数或者变量）的数据类型

why: 希望从表达式获得对应数据类型 但不想用auto 用这个表达式来初始化变量

how：
```
1.使用对象为变量 返回该变量类型 ,可以为引用类型
const int ci=0, &cj = ci;
decltype(ci) x=0; // 常整数
decltype(cj) y=x;// y为常整数引用 ,y必须绑定到x
decltype(cj) z; // 错误 ,z 为 const int& 必须初始化

2. 对象为表达式,返回表示式结果类型
int i =42, *p = &i ,&r = i ;
decltype(r + 0) b; // 正确 r为引用 r+0为int
decltype(*p) c;// 错误 c为int& 必须初始化
（解引用得到的是引用类型）

3. decltype((任何东西)) 得到的永远是引用,双层括号的结果是引用

```


