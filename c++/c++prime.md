
- [c++基础](#1)
    - [变量和基本类型](#1-1)
        - [处理类型](#1-1-5)
        - [自定义数据类型](#1-1-6)
    - [字符串，向量和数组](#1-2)
        - [命名空间的using声明](#1-2-1)
        - [标准库类型string](#1-2-2)
        - [标准库类型vector](#1-2-3)
        - [迭代器](#1-2-4)
        - [数组](#1-2-5)
        - [多维数组](#1-2-6)
    - [表达式](#1-3)
        - [基础](#1-3-1)
        - [算术运算符](#1-3-2)
        - [逻辑和关系运算符](#1-3-3)
        - [赋值运算符](#1-3-4)
        - [递增和递减运算符](#1-3-5)
        - [成员访问运算符](#1-3-6)
        - [条件运算符](#1-3-7)
        - [位运算符](#1-3-8)
        - [sizeof运算符](#1-3-9)
        - [逗号运算符](#1-3-10)
        - [类型转换](#1-3-11)
    - [语句](#1-4)
        - [try语句块和异常处理](#1-4-6)
    - [函数](#1-5)
        - [函数基础](#1-5-1)
        - [参数传递](#1-5-2)
        - [返回类型和return语句](#1-5-3)
        - [函数重载](#1-5-4)
        - [特殊用途语言特性](#1-5-5)
        - [函数匹配](#1-5-6)
        - [函数指针](#1-5-7)
    - [类](#1-6)
        - [定义抽象数据类型](#1-6-1)
        - [访问控制和封装](#1-6-2)
        - [类的其他特性](#1-6-3)
        - [类的作用域](#1-6-4)
        - [构造函数再探](#1-6-5)
        - [类的静态成员](#1-6-6)
- [c++标准库](#2)
    - [IO库](#2-8)
        - [IO类](#2-8-1)
        - [文件输入输出](#2-8-2)
        - [string流](#2-8-3)
    - [顺序容器](#2-9)
        - [顺序容器概述](#2-9-1)
        - [容器库概览](#2-9-2)
        - [顺序容器操作](#2-9-3)
        - [vector对象是如何增长的](#2-9-4)
        - [额外的string操作](#2-9-5)
        - [容器适配器](#2-9-6)
    - [泛型算法](#2-10)
        - [概述](#2-10-1)
        - [初识泛型算法](#2-10-2)
        - [定制操作](#2-10-3)
        - [再探迭代器](#2-10-4)
        - [泛型算法结构](#2-10-5)
        - [特定容器算法](#2-10-6)
    - [关联容器](#2-11)
        - [使用关联容器](#2-11-1)
        - [关联容器概述](#2-11-2)
        - [关联容器操作](#2-11-3)
        - [无序容器](#2-11-4)
    - [动态内存](#2-12)
        - [动态内存与智能指针](#2-12-1)
        - [动态数组](#2-12-2)
        - [使用标准库：文本查询程序](#2-12-3)


<h1 id='1'>C++基础</h1>


<h2 id='1-1'>变量和基本类型</h2>

<h3 id='1-1-5'> 处理类型 <h3>

#### 类别别名

 what： 某种类型的同义词

 why ：让复杂的代码更容易理解
 ```c++
如 typedef char * pstring;
那么 pstring是 char *的别名 如果用到这个别名 他的基础类型就是指针
const pstring cstr=0; //指向char的常指针
const pstring *ps;  // 指向常指针的 一个指针
 ```
 how ： 2种声明方法
 ```c++
1. typdef double base,*p;
// base是 double的同义词 p是double *的同义词

2. using base = double; 
 ```
 
 #### auto 类型说明符

 what : 让编译器替我们分析表达式所属类型

 why ：我们在声明过程中可能不是很清楚了解表达式类型

 how：
 ```c++
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
```c++
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

Ps. 赋值是会产生引用的一类典型表达式
int a = 3, b = 4;
decltype(a) c = a; // c为int 
decltype(a = b) d = a; // d 为int &
```

<h3 id='1-1-6'> 自定义数据类型 <h3>

#### 定义 

what: 把数据元素组织起来 用户能直接访问其中元素 也能实现基本操作

how ：
```c++
struct node
{

};
1. 创建对象时类内初始值将用于初始化数据成员 ，如果没有初始值将被默认初始化

2. 也能用class 创建

```

#### 编写头文件

what: 在函数体外币定义

why:想在不同文件中用同一个类

how：
```c++
1. 类所在头文件应与类名字一样

2. 预处理器 ，编译之前执行能部分改变所写程序，如
#include<stdio.h> 用 stdio.h 有文件的内容来代替#include<stdio.h>

3. 头文件保护符，防止重复包含的发生
#define 把一个名字设定为预处理变量

#ifdef  当且仅当变量已定义时为真
#ifndef 当且仅当变量未定义时为真

这两个一旦为真一直执行后续操作直到遇到 #endif  
为假 跳过 到 #endif为止的部分

通常做法是基于头头文件中类的名字来构建保护符的名字
```

<h2 id='1-2'>字符串，向量和数组 </h2>

<h3 id='1-2-1'> 命名空间的using声明 </h3>

```c++
1. 
using std::cin; // 以后不用 std::cin 了直接 cin就行

2.  头文件不应该包含using 申明 ,不然每个使用该头文件的都会有申明 可能会有名字冲突

```

<h3 id ='1-2-2'> 标准库类型string </h3>

what : 可变长字符序列

#### 定义和初始化string 对象

直接初始化和拷贝初始化 

拷贝初始化用了 = 
如 string s2 = s1 
直接  string s2(s1)
string s(4,'c') // s为"cccc"

#### string对象上的操作

1. cin 和 getline 
```c++
string s;
getline(cin,s); //读取一整行，可读取空格 不包含换行符

s.empty() // 为空为真
s.size() // 返回长度,返回的是 string::size_type 类型的值,是个无符号整数,而且能存下任何string对象的大小
所以混用 size和 int 可能会存在问题 比如 s.size() < n , n如果为一个负数，会转变为较大的无符号值

2. string 比较
根据字典序比较

3. 两个 string 相加
类如 a ="aa"  b= "bb"
 a + b = "aabb"

 a + = b = "aabb"

4. string 可以和字面值相加  如 "o" 或 '/n' 都算字面值 ,字面值和string 是不同的类型
但是字面值不能和字面值相加

```

#### 处理string对象中的字符

1. 用cctype头文件的函数处理单个字符

isalnum() 如果参数是字母数字，即字母或数字，该函数返回true 

isalpha() 如果参数是字母，该函数返回true 

isblank() 如果参数是空格或水平制表符，该函数返回true 

iscntrl() 如果参数是控制字符，该函数返回true

isdigit() 如果参数是数字（0～9），该函数返回true 

isgraph() 如果参数是除空格之外的打印字符，该函数返回true 

islower() 如果参数是小写字母，该函数返回true

isprint() 如果参数是打印字符（包括空格），该函数返回true 

ispunct() 如果参数是标点符号，该函数返回true 

isspace() 如果参数是标准空白字符，如空格、进纸、换行符、回车、水平制表符或者垂直制表符，该函数返回true 

isupper() 如果参数是大写字母，该函数返回true 

isxdigit() 如果参数是十六进制的数字，即0～9、a~f、A~F，该函数返回true 

tolower() 如果参数是大写字符，则返回其小写，否则返回该参数 

toupper() 如果参数是小写字母，则返回其大写，否则返回该参数

tolower()——toupper() 大写转为小写——小写转为大写

isupper()——islower() 判断是否为大写——判断是否为小写，若是返回true，否则返回该参数

isalnum()——isalpha() 判断是否为字母或数字，若是返回true，否则返回该参数——判断是否为字母，若是大写字母返回1，若是小写字母返回2，若不是字母返回0


2. 基于范围的for 例如
```c++
// 输出字符串每个字符
string str;
for (auto c :str)
    cout<< c <<endl;

// 想要改变字符 如全变成大写
string str;
for (auto &c :str) // c变成了引用
    c = toupper(c);//通过函数让每个c变成大写
```

3. 运用下标时要确定那个位置上有值而且下标类型
为 string::size_type 即无符号整数

<h3 id="1-2-3">标准库类型vector</h3>

what : 表示对象的集合 ，其中所以对象的类型都相同 类模板 

vector<vector<int> > 需要有空格不过 c++11 没有空格也合法

#### 定义和初始化vector对象

how：
```c++
vector<T> v1;
vector<T> v2(v1);// v2包含v1所有元素的副本,v2和v1的元素类型要一样
vector<T> v3(n,val); // 含有n个重复元素 ，每个元素值为val
vector<T> v4｛a，b，c，...｝
vector<int> v5(10) // 有10个元素，每个元素都默认初始化为0

```
圆括号  构造 
花括号 初始化
但是如果使用花括号但是给予的值不能用来初始化那么会变成用来构造
如 vector<string> v{10}   // v有10个默认初始化的元素
   vector<string> v("hi") // 错误 ，不能用字符串字面值构建vector对象
   vector<string> v{10，"hi"} // v有10个初始化为"hi"的元素

#### 向vector 对象中添加元素

push_back

如果循环体内含有向vector对象添加元素的语句，则不能使用范围for循环

#### 其他vector操作

支持关系运算符 按字典序比较

不允许通过下标来添加新元素 ，只能队确知已存在元素进行下标操作

<h3 id='1-2-4'>迭代器</h3>

what: 对对象的间接访问，其对象是容器的元素或者string对象中的字符

#### 使用迭代器

begin  指向第一个元素

end 指向尾元素的下一位置

如果容器为空 begin 和 end 返回元素相同

```c++
*iter // 返回迭代器iter所指元素，解引用必须保证确实指着某个元素
iter->mem  // 解引用iter并获取该元素的名为mem的成员，等价于 (*iter).mem
++iter
--iter

```
iterator 能读能写

const_iterator 只能读

begin 和end 返回的值由对象是否为常量觉得 为常量返回const_iterator 否则 iterator  

cbegin 和cend 返回为 const_iterator  不过也可以进行 ++  --运算，只是不行修改值

#### 迭代器运算

两个迭代器指向同一容器的元素或者尾元素的下一位置，就能相减 得到的是 名为`difference_type`的带符号整形数

二分时 只能 mid = beg + ( end - beg )/2,不能 mid = ( beg + end )/2  ，因为迭代器没有 + 功能

<h3 id='1-2-5'>数组</h3>

what: 存放类型相同对象的容器 但数组大小确定不变

#### 定义和初始化内置数组

what： 数组声明 a[d] a为数组名  d为维度 ，d必须为一个常量表达式(整数字面值 或const及constexpr类型  )

how：
```c++
1. 显示初始化数组元素
int a[] = {1,2} // 维度为2 ，元素 为 ｛1,2｝

2. 字符数组的特殊性
char a1[] = {'c','+','+'}; // 列表初始化 没有空字符 ，输出时可能有问题
char a2[] = {'c','+','+','\0'}; //列表初始化 含有显式的空字符，输出时不存在空字符
char a3[] = "c++"; // 自动添加表示字符串结束的空字符
char a4[6] = "Daniel" ; // 错误 没有空间可存放空字符

3. 不允许拷贝和赋值

4. 数组声明
int arr[10];
int *ptrs[10]; // 含有10个整型指针的数组
int &refs[10] = /* ? */; // 不存在引用数组
int (*Parray)[10] =&arr; // Parray 指向一个含有10个整数的数组
int (&arrRef)[10] = arr; // arrRef引用一个含有10个整数的数组
int * (&arry)[10]=ptrs; // arry是数组的引用 该数组有10个指针

ptrs[0]=&arr[0];

下面几个值都相同 因为数组的地址和数组第一个元素地址相同
printf("%d\n",&arr[0]);
printf("%d\n",Parray);
printf("%d\n",ptrs[0]);
printf("%d\n",arry[0]);

下面几个值都相同
printf("%d\n",arrRef[0]);
printf("%d\n",*ptrs[0]);
printf("%d\n",*arry[0]);

5. 定义在函数外的数组会默认初始化   定义在函数内的数组未定义

```

#### 访问数组元素

可以使用范围for或下标(使用下标时，为size_t类型 无符号类型)

#### 指针和数组

what：`使用数组是编译器一般会把他转换成指针`

how：
```c++
1. 很多用到数组名字的地方 编译器会自动将其替换成一个指向数组首元素的指针
int ia[]= {0,1};
auto ia2(ia) // ia2 是一个整型指针 指向ia的第一个元素,因为这个式子等价于 auto ia2(&ia[0])
当使用 decltype时转换不会发生
decltype(ia) ia3 ={1,2};
ia3[0] = 4;

2. 数组使用迭代器
int  arr[]= {0,1};
int *e = &arr[2]; // 指向arr尾元素的下一位置的指针 相当于 end

不过这种方法容易出错 ,c++11提供begin和end方法
string nums[] ={"1","2"};

string *beg=begin(nums);
string *last=end(nums);
for (auto i=beg;i!=last;++i)
    cout<<*i<<endl;

3. 迭代器范围不要超过尾元素的下一位置
两个指针相减得到`ptrdiff_t ` 带符号类型

4. 下标和指针运算, 内置的下标索引值不是无符号类型

int ia []={0,1,2,3,4,5,6};
int i=ia[2]; // ia转换成指向数组首元素的指针 ia[2]代表 得到 (ia+2) 所指的元素 
int *p = &ia[2]; // p指向索引为2的元素
int j = p[1] // p [1] 等价于 *(p+1) 即得到的是 ia[3] 表示的元素
int k = p[-2] // 得到ia[0] 表示的元素

```

#### C风格字符串

what ： 存放在字符数组中并以`空字符结束`

使用 strlen 返回长度 strcmp 比较  strcat 附加 strcpy  拷贝   ！！必须以空字符结尾

可以用安全版本如 strcpy_s(str, strlen(str1)+1, str1);

> 当使用数组时真正用的是指向数组首元素的指针

使用c风格字符串时 我们得保证数组足够容纳字符串 (字符串后面有空字符)

将两个指针相减可以表示两个指针(在同一数组中)相距的距离，将指针加上一个整数也可以表示移动这个指针到某一位置。但是两个指针相加并没有逻辑上的意义，因此两个指针不能相加。

#### 与旧代码的接口

what: 提供c++与数据

how：
```c++
1. 可以用 以空字符结束的字符数组来初始化string对象或为其赋值

2. 可以用 string的 c_str()属性来返回一个c风格字符串，也就是说函数返回结果是一个指针 ，该指针指向一个以空字符结尾的字符数组 ，指针类型是 const char *

const char *str = s.s_str()

但是不保证 s.s_str函数返回数组一直有效 如果我们改变了s的值可能会让之前返回数组失效 所以如果想一直使用返回的数组 最好将该数组重新拷贝一份

char str[10];
string s("1234");
int main()
{
   const char * p = s.c_str(); //一定姚 const char *
   strcpy(str,p); //拷贝一份
   cout<<p[1]<<endl;
   s.clear(); //修改后 之前的p指针可能会有问题
    cout<<str<<endl;
}

3. 可以用数组初始化vector
int a[4]={0,1,2,3};
vector<int> aa(a+1,a+4);// 数据为 1， 2 ，3 

4. 尽量使用标准库而非数组
```

<h3 id='1-2-6'>多维数组</h3>

what: 数组的组数

1. 多维数组想要用范围for处理 ，除了最内层循环，其他所有循环的控制变量都为引用，例如
```c++
int ia[3][4];
 for(auto &row:ia) 
    for(auto col:row)
        cout<<col;
因为 auto row :ia 为指针类型
    auto &row:ia 为数组类型

int (*p)[4] = ia; // p 指向含有4个整数的数组
p = &ia[2] ; // p指向ia的尾元素
```
2. 也能通过 迭代器 begin和end处理
```c++
int ia[3][4];
for(auto p = begin( ia);p != end( ia ); ++p ) // p为迭代器
    for(auto q = begin( *p ); q != end( *p ); q++ ) //q为迭代器
        cout << *q ;
```
3. 可以用类型别名简化多维数组的指针
```c++
int ia[3][4];
using int_array = int[4] // 将类型"四个整数组成的数组"命名为int_array
for(int_array *p = ia ; p != ia + 3 ; ++p )
    for( int *q = *p ; q != *p + 4 ; ++q )
        cout << *q ;
```
例如
```
int ia[3][4] =
{
    { 0, 1, 2, 3 },
    { 4, 5, 6, 7 },
    { 8, 9, 10, 11 }
};

for(const int(&row)[4] :ia)
{
    for(auto col:row)
    cout<<col<<"  ";
}

cout<<endl;

for(size_t i = 0 ; i != 3 ; ++i )
    for(size_t j = 0 ; j != 4 ; ++j )
        cout<<ia[i][j]<<"  ";

cout<<endl;
for(int (*p)[4] = ia ; p != ia + 3 ; ++p )
    for( int *q = *p ; q != *p + 4 ; ++q )
        cout << *q <<"  " ;


for (auto& p : ia)
    for (auto q : p)
        cout << q << " ";
cout << endl;

for (auto i = 0; i != 3; ++i)
    for (auto j = 0; j != 4; ++j)
        cout << ia[i][j] << " ";
cout << endl;

for (auto p = ia; p != ia + 3; ++p)
    for (auto q = *p; q != *p + 4; ++q)
        cout << *q << " ";
cout << endl;
```

<h2 id='1-3'> 表达式 </h2>

what: 由一个或多个运算读写组成 ，对表达式式求值将得到一个结果 运算符和多个运算对象组合起来可以生成较复杂的表达式

<h3 id='1-3-1'>基础</h3>

#### 基本概念

一元运算符:作用于一个运算对象  取地址符(&)  解引用符(*)

二元运算符 : 作用于二个运算对象  相等运算符( == )  乘法运算符 ( * )

三元

当对象被用作右值时，用的是对象的值(内容),当对象被用作左值时，用的时对象的身份(在内存中的位置)

例如 取地址符作用于一个左值运算对象，返回一个指向该运算对象的指针，这个指针是个右值， 

解引用，下标运算符，递增递减运算符，赋值运算符 获得左值

如果表达式的求值结果是左值,decltype作用于该表达式(不是变量)得到一个引用类型,例如decltype(*p) 结果为int &

#### 优先级和结合律

括号无视优先级和结合律

#### 求值顺序

优先级规定了运算对象的组合类型，但是没有说明运算对象安照什么顺序求值，例如

int i = f1() * f2()

我们知道f1和f2会在乘法之前调用 ,但我们不知道f1和f2哪个会先调用，对于函数的调用顺序没有明确规定
```
int a=0;
cout<<a<< ++a<<endl；//类似于这样的行为是错误的，是未定义的，我们不能保证最后输出是什么
可能是 1 1 也可能是 0 1

类似 * ++iter 则不会和改变iter的值的函数产生冲突 ，因为他解引用的是iter下个元素 
```

<h3 id='1-3-2'>算术运算符</h3>

优先级顺序 (一元正号  一元负号) (乘  除  取余) (加 减)

算术运算符的运算对象和求值结果都是右值，在求值表达式求值之前，所有运算对象最终会转换成同一类型

bool不应该参与运算，例如
```
bool b= true;
bool b2 = -b;//b2为true ，因为-b转换为int为-1，-1!= 0
```
m % (-n) = m%n        (-m) % n = -( m%n )

当计算的结果超出该类型所能表示的范围时就会产生溢出。

<h3 id='1-3-3'>逻辑和关系运算符</h3>

关系运算符作用于算术类型或指针类型，逻辑运算符作用于任意能转换成布尔值的类型,这两类运算对象和求值结果都为右值
```c++
string text;

for (const auto &s : text)
{
    cout<< s;
}
因为text是string类型可能会很大将s申明为引用，可以避免对元素进行拷贝，
又因为不需要对string进行写操作所以被申明为常量引用
```
不能连立逻辑运算符例如 if(i< j < k )  ,因为 i< j 返回的是bool类型 ,bool类型与k进行比较

得 if(i< j && j < k ) 

进行比较运算时，除非比较对象时bool 否则不要使用bool字面值false和true进行运算对象

关系运算符优先级高于逻辑运算符 例如 
```c++
int a,b,c;
a=1,b=2,c=3;
if(a!= b<c )
printf("--");
else
printf("++"); //输出为++
```

<h3 id='1-3-4'>赋值运算符</h3>

赋值运算符左侧运算对象必须是一个可修改的左值

无论左侧运算对象类型是什么，初始值类别都可以为空。此时 编译器创建一个值初始化的临时量并将其赋给左侧运算对象
如
```c++
int main()
{
  int k;
  k={};
  cout<<k ; // k 为 0 
}
```
#### 赋值运算符满足右结合律

```c++
int ival, *pval;
ival = pval = 0 // 错误 ，虽然0能赋给任何对象但是 先 pval = 0 然后 ival = pval 时
// pval 是 (int *)  无法转换为 int
```

#### 赋值运算优先级较低

因为赋值运算符优先级低于关系运算符 所以在条件语句中 赋值部分一般会加上括号

#### 复合赋值运算符

+= -=之类的,区别在于 复合运算符只求值一次 ，普通运算符求值两次

<h3 id='1-3-5'>递增和递减运算符</h3>

迭代器不支持算术运算但支持递增和递减

前置版本和后置版本，建议使用前置版本，提高性能
```c++
int i=0,j=0;
j = ++i; // i=1 ,j=1  前置 先+1将改变后的对象作为求值结果 ，得左值
j = i++; // i=2, j=1  后置 也+1但求值结果是运算对象改变之前那个值的副本，得右值

但是如果我们想既将变量+1或-1 又能使用它原来的值，我们就可以使用后置版本，如
auto pbeg = v.begin(); 
//输出元素知道遇到第一个负值为止
while (pbeg != v.end() && *pbeg >= 0)
    cout << *pbeg++ <<endl;   

*pbeg++ 什么意思，因为后置递增运算符的优先级高于解引用运算符 相当于 *(pbeg++) ,pbeg++ 把 pbeg值 加1然后返回 pbeg 原值的副本
```

#### 运算对象可按任意顺序求值

所以例如 
```
vector v;
int b;
v[b++] <= v[b] //错误 ，c++未规定 关系运算符两边的求值顺序
```

<h3 id='1-3-6'>成员访问运算符</h3>

点运算符和箭头运算符都可以访问成员

ptr为指针类型  ptr -> mem 相当于 (*ptr).mem  {需要加括号，因为解引用的优先级低于点运算符}

箭头运算符作用于一个指针的运算对象，结果为左值 ，点运算符结果依靠成员所属的对象，对象是左值，结果为左值，对象为右值 结果为右值

<h3 id='1-3-7'>条件运算符</h3>

cond ? expr1 : expr2 

comd 是判断条件的表达式 ,为真 对 expre1 求值并返回 为假，对 expre2 求值并返回(条件运算符只对expre1和expre2 中的一个求值)

当条件运算符两个表达式 都是左值或者都能转换成同一种左值类型时，运算结果是左值；否则运算结果是右值

#### 嵌套条件运算符

cond1 ? expr1 : cond2 ? exprr2 : exprr3

#### 在输出表达式中使用条件运算符

条件运算符的优先级非常低 

```c++
cout << (( grade < 60) ? "fail" : "pass");//正确，输出fail或pass

cout << ( grade < 60) ? "fail" : "pass"; // 输入 0 或 1
//相当于 ( cout << ( grade < 60 ) ) ? "fail" : "pass";

cout << grade < 60 ? "fail" : "pass"; //错误，试图比较cout和60
// 相当于 （ cout << grade ） < 60 ? "fail" : "pass";
```

<h3 id='1-3-8'> 位运算符 </h3>

what ：位运算符作用于整数类型的运算对象，并把运算对象看成是二进制位的集合

关于符号位如何处理没有明确规定 ，强烈建议仅仅将位运算用于处理无符号类型

#### 移位运算符 <<  或 >> ,移出边界之外的位被舍弃掉了

#### 位求反

char类型运算对象首先提升为int ，往高位添 0 ，然后在逐位取反

#### 移位运算符(又叫IO运算符) 满足左结合律

移位运算符优先级 比算术运算符低， 比关系运算符 赋值运算符 和条件运算符 高 反正一次多个运算符 建议使用括号

<h3 id='1-3-9'> sizeof运算符 </h3>

what: 返回一条表达式或一个类型名字所占的字节数

how ： 1. 满足右结合律 2. 得到的是size_t类型(无符号)
```
int data,*p;
sizeof data; // data 的类型的大小 即 sizeof (int);
sizeof (int); // 存储int类型的对象所占的空间大小
sizeof p; // 指针所占空间大小
sizeof *p;// p所指类型的空间大小，即 sizeof (int); 因为是满足右结合律 相当于 sizeof (*p) 
//  sizeof不会实际求运算对象的值    所以 p为未初始化指针也没关系

对引用类型执行sizeof运算得到指针本身所占空间的大小
对数组 ： 对数组中所有的元素各执行一次sizeof运算并将所得结果求和，sizeof不会把数组当成指针
int a[10];
sizeof( *a ) 返回a的元素数量
sizeof(a)/sizeof(*a) // 得到 10
对 string或vector 只会返回该类型固定部分的大小，不计算长度

sizeof不会实际求运算对象的值，sizeof 返回是一个常量表达式 
```
sizeof 优先级高于算术运算符 ，低于点运算符和箭头访问运算符 

<h3 id= '1-3-10'> 逗号运算符</h3>

waht: 有两个运算对象，按从左到右顺序

how： 先对左侧表达式求值，然后将求值结果丢弃。运算符真正结果是右侧表达式的值，如果右侧表达式为左值，那么最终结果也为左值

优先级最低

<h3 id='1-3-11'>类型转换</h3>

隐式转换

#### 算术转换

what: `运算符的运算对象将转换成最宽的类型`

整数提升: 小整数提升为大整数

如果一个运算对象时无符号类型，另一个带符号 

1. 其中无符号类型不小于带符号类型,那么带符号会转换为无符号
```
比如
int a=-3 ;
unsigned int b=2;
cout<<a+b<<endl; //输出为 4294967295,a+b 过程中 a 转变为unsigned int类型
```
2. 如果无符号所有值都能存在于该带符号类型中 ，转换为带符号,所以可能取决于运行该机器的机器每个类型所占内存大小
```
long long  a=-3 ;
unsigned int b=2;
cout<<a+b<<endl; // 输出为-1
```

#### 其他隐式类型的转换

1. 数组转换为指针 : 数组自动转换为数组首元素指针

被用作decltype关键字参数  取地址符(&)  sizeof 和 typeid 等运算符的运算对象时，转换不会发生

2. 指针转换 : 常量整数值0或者字面值nullptr 转换成任意指针类型 ； 指向任意非常量指针能转换为 void *;
指向任意对象的指针能转换为 const void *;

3. 转换成bool类型 : 如果指针或算数类型值为0，转换结果为false

4. 转换成常量
```c++
  int i;
  const int &j = i;// 非常量转换成const int 的引用
   int &r = j; // 错误，不允许const转换成非常量 j是常量 如果赋值成功 就能修改j值了
   //如果 const int &r = j 可以赋值成功

  const int *p = &i; // 非常量的地址转换成const 的地址
  int *q = p; //错误 
```

5. 类类型定义的转换

例如 :  while( cin>>s ) while 把 cin转换成布尔值

如果同时提出多个转换请求，这些请求将被拒绝

#### 显式转换

强制转换:  cast-name < type > (expression);

cast-name 种类 1. static_cast  2. dynamic_cast 3. const_cast 4. reinterpret_cast  

<h2 id='1-4'>语句</h2>

<h3 id='1-4-6'> try语句块和异常处理</h3>

异常处理： 1. throw表达式(引发(raise)异常) 2 2. try语句块  (try语句中抛出异常会被某个catch子句处理，异常处理代码) 3. 异常类 (用于在throw表达式和相关catch子句间传递异常具体信息)

#### throw表达式

throw表达式包含关键字 throw 和紧跟着的一个表达式
```c++
int a=1,b=2;
if(a==b)
    printf("a=b");
else
    throw runtime_error("a!=b");
```

#### try语句

try语句块后面有一个或多个catch子句，catch子句块也无法访问 try 语句块声明的变量
```c++
int a,b;
while(cin>>a>>b)
{
 try {
   if(a==b)
     cout<< "a=b" <<endl;
   else  throw runtime_error("a!=b"); // 寻找最近的runtime_error函数 处理异常
 } catch(runtime_error err){
    cout<< err.what() // what的信息为 "a!=b"
        << "\n Try Again? Enter y or n" << endl;
    char c;
    cin >> c;
    if(!cin||c=='n')
        break; //跳出while循环
 }
}

在多层 try 中 如果一层没找到相应处理函数 会返回上层寻找  ，一直没找到 会调用 terminate 库函数，将导致异常程序退出
```

#### 标准异常

头文件 : exception (最通用异常类exception,只负责报告异常的发生，不提供信息)    stdexcept (常用类)  
new (bad_alloc)  type_info (bad_cast)

exception 最常见问题

`runtime_error` 只有在运行时能检测出的问题

range_error 运行时出错 : 生成的结果超出了有意义的值域范围

overflow_error 运行时出错 : 计算上溢

underflow_error 运行时出错 : 计算下溢

`logic_error 程序逻辑错误`

domain_error 逻辑错误 : 参数对应的结果值不存在

invalid_argument 逻辑错误 ： 无效参数

length_error  逻辑错误 ：试图创建一个超过该类型最长长度的对象

out_of_range 逻辑错误： 使用一个超出有效范围的值

只能默认初始化 exception bad_alloc bad_cast  其他异常类型都不能默认初始化，需提供错误信息 例如字符串

异常类型定义what成员函数  ，返回值是一个指向C风格字符串 的 const char *

<h2 id='1-5'> 函数</h2>

<h3 id='1-5-1'>函数基础</h3>

what调用函数：1. 用实参初始化函数的形参 2. 将控制权转移给被调用函数，此时 主调函数的执行被暂时中断

遇到 return ：1. 返回return语句的值(如果有的话) 2. 将控制权转移给主调函数

#### 局部对象

名字有作用域 ， 对象有生命周期(程序执行过程中该对象存在的一段时间)

形参和函数内部定义的变量称之为 局部变量 ，仅在函数作用域内可见 ，同时局部变量会隐藏在外层作用域的同名其他所有声明中 。 局部变量的生命周期依赖于定义方式

自动对象：只存在于块执行期的对象，当块的执行结束后，快内创建的自动对象就变成未定义的了

形参是一种自动对象

局部静态对象 (static 类型) :在程序的执行路径`第一次`经过对象定义语句时初始化，直到程序终止才被销毁
```c++
int aa()
{
   static int a =0; // 仅第一次会初始化
   return ++a;
}
int main()
{
    for(int i=1;i<=10;i++)
        cout<< aa()<<endl; // 输出为 1 到 10
}
```

#### 函数声明

函数三要素 (返回类型 函数名 形参类型)

建议在头文件中声明函数和变量， 在源文件中定义

#### 分离式编译
`Abss.h`
```c++
#ifndef MEMORY_CELL_H
#define MEMORY_CELL_H
int abss(int a,int b);

#endif
```
`fact.cpp`
```c++
#include "Abss.h"
#include <iostream>

int abss(int a,int b)
{
    return (a>b?a:b);
}
```
`main.cpp`
```c++
#include "fact.cpp"
#include <iostream>

using namespace std;
int main()
{
	cout<<abss(3,5)<<endl;
}
```

<h3 id='1-5-2'>参数传递</h3>

如果形参式引用那么会绑定到实参，否则将实参的值赋给形参

#### 引用形参

引用形参可以避免拷贝 ，提高效率 ，如果无需改变值可以用 常量引用，而且可以用引用参数来返回值

#### const 形参和实参

顶层const 例如 允许int 拷贝 const int类型

允许若干个相同名字的函数 但是其形参列表必须不同

把函数值不会改变的形参定义为常量引用，且如果定义为普通引用， const对象 ，字面值和需要类型转换的的对象就不能传递进来

#### 数组形参

int *  a 等价于 int a[] 

```c++
void aa(int* a,const int *beg ,const int *endd)
{
    a[0]=3;
}
int a[2]={1,2};
int main()
{
	aa(a,begin(a),end(a)); // a转换为 int * 并指向 a[0]
    cout<<a[0];
}
```
数组引用形参：  int ( &a )[2]
```c++
void aa(int (&a)[2])
{
    a[0]=3;
}
int a[2]={1,2};
int main()
{
	aa(a);
    cout<<a[0];
}
```

传递多维数组 : int ( *a )[2]  指向含有2个整数的数组的指针，等价于 int a[][2]

#### 含有可变形参的函数

initializer_list 形参 ： 表示某种特定类型的值的数组 
```
initializer_list< T > lst; 默认初始化 ，T类型元素的空列表
initializer_list< T > lst{a,b,c...};
last2(lst)
last2=lst

lst.size()
lst.begin()
lst.end()
```

省略符形参 (便于c++程序访问某些特殊的c代码，只用于与c的接口)

<h3 id='1-5-3'>返回类型和return语句</h3>

#### 有返回值

不要返回局部对象的引用或指针 ： 函数完成后，局部对象被销毁 ，内存被释放 要想返回引用 ，我们得知道我们所引的是函数之前已经存在的哪个对象

函数返回值为 引用 则得到左值，其他返回类型得到右值，我们可以为返回引用的函数赋值
```c++
int &aa(int &a)
{
   a=2;
   return a;
}

int main()
{
	int a=1;
	aa(a)=3;
	cout<<a<<endl;
}
```
列表初始化返回值
```c++
vector<string> aa(int a,int b)
{
    if(a==0&&b==0)
    return {};
    if(a==b)
        return {"a=b","OK"};
    else
        return {"a!=b","a","b"};
}

int main()
{
	int a,b;
	vector<string> v;
	while(~scanf("%d%d",&a,&b))
    {
        v=aa(a,b);
        for(auto s:v)
            cout<<s<<endl;
    }
}
```
主函数main 返回值:默认是在最后 返回return 0 (表示成功),如果想显式返回   返回失败  return EXIT_FAILURE     ，返回成功  return EXIT_SUCCESS    其中EXIT_FAILURE和EXIT_SUCCESS 都是预处理变量

#### 返回数组指针

1. 使用别名  ： 
```c++
using arrt = int[10];
arrt* func(int i); // func返回一个指向含有10个整数的数组的指针
```
2. 记住维度
```c++
int ( *func( int i ) )[10] // func返回一个指向含有10个整数的数组的指针

func(int i) 表示func函数需要一个int类型实参
( *func( int i ) ) 表示 返回的是指针
( *func( int i ) )[10] 表示指针指向大小为 10的数组
int ( *func( int i ) )[10] 表示 数组类型为int
```
3. 使用尾置返回类型
```c++
auto func(int i) - > int (*)[10] // func 接受一个int类型实参 返回一个指针，该指针指向含有10个整数的数组
```
4. 使用declty
```c++
int odd[]={0,1,2,3,4,5,6,7,8,9};
decltype(odd) *func(i) // 返回一个指针，该指针指向含有10个整数的数组
{

}
```

<h3 id='1-5-4'>函数重载</h3>

重载函数： 同一作用域下几个函数名字相同但是 **形参列表**(形参数量或类型) 不同，(main函数不能重载)

错误重载
```c++
1. 仅仅改变返回值
Record lookup(const Account&)
bool lookup(const Account&) //错误，只有返回值不同
2. 类型相同
Record lookup(const Account &acct)
Record lookup(const Account&) //错误，只是忽略了形参

Record lookup(const  Phone&) 
Record lookup(const  Telno&) //错误，类型和数量相同

3. 顶层const 形参 和 没有顶层const的形参 无法区分开来
Record lookup( Phone ) 
Record lookup( const Phone ) //错误 ，是顶层const

Record lookup( Phone*) 
Record lookup( Phone* const ) //错误 ，是顶层const
```

正确重载
```c++
Record lookup( Phone& ) 
Record lookup( const Phone& ) // 正确，作用于常量引用，是底层const

Record lookup( Phone* ) 
Record lookup( const Phone* ) //正确 ，作用于指向常量的指针，是底层const

PS. const对象只能传递给const 形参 ， const形参能接收非const对象 ，但是如果传递的是非const对象，那么优先选用非const函数版本
```

区别底层和顶层const
```c++
int const *a; //底层const
int *const b; //顶层const
```

const_cost 和 重载 : 使用const 函数 创建 非const版本函数
```c++
// 比较两个string 的长度 返回较短的那个
const string &shorterString(const string &s1,const string &s2)
{
    cout<<"const string&"<<endl;
    return s1.size()<= s2.size() ? s1 : s2;
}
string &shorterString(string &s1, string &s2)
{
    auto &r = shorterString(const_cast<const string&>(s1),
                            const_cast<const string&>(s2));
    cout<<"string&"<<endl;
    return const_cast<string&>(r);
}
int main()
{
    string s1 ="11",s2="1";
    const string s3 ="11",s4="1";
    shorterString(s1,s2);
    cout<<endl;
    shorterString(s3,s4);
}

结果
const string&
string&

const string&
```

<h3 id='1-5-5'>特殊用途语言特性</h3>

默认实参 内联函数  constexpr函数

#### 默认实参
一个形参添加默认值后 ，默认值后的形参也得有默认值
```c++
int aa(const int longg=1,const int wight=2 ,const int hign=3)
{
    return longg*wight*hign;
}
int main()
{
    cout<< aa() <<endl;  //6
    cout<<aa(2,3,4) <<endl; // 24
    cout<< aa(2,3) <<endl; // 18 
}
```
我们不能修改一个已经存在的默认值
```c++
int aa(const int longg,const int wight ,const int hign=3)；

int aa(const int longg,const int wight ,const int hign=4)；//错误

int aa(const int longg=1,const int wight=2 ,const int hign) //正确
```
局部变量不能作为默认参数

#### 内联函数和constexpr函数

what内联函数： 将内联函数在每个调用点上"内联地"展开

why内联函数： 调用函数一般比求等价表达式慢一些  函数( 1. 调用前要先保存寄存器，在返回时恢复 2. 可能要拷贝实参 3. 程序转向一个新的位置继续执行)
```c++
inline int
aa(const int longg=1,const int wight=2 ,const int hign=3)
{
  return longg*wight*hign;
}
int main()
{
    cout<< aa() <<endl; //6
    cout<<aa(2,3) <<endl; // 24
}

编译过程中展开类似于 cout<< longg*wight*hign <<endl， 消除函数运行时开销
```
内联函数用于优化规模较小，流程直接 频繁调用的函数

`constexpr函数`  what : 能用于常量表达式的函数(1. 函数的返回类型和返回值都得是字面值类型 2. 函数体中必须有且只有一条return语句 )
```c++
constexpr int new_sz() {return 4;}
constexpr size_t scale(size_t cnt) { return new_sz() * cnt;} // 相乘
int main()
{
  cout<<scale(2)<<endl;
  int a[scale(2)] ; // scale(2) 是常量表达式
  int v = 2;
  int aa[scale(v)] ;  // 错误 scale(a) 不是常量表达式
}

所以 用非常量表达式调用constexpr函数 得到非常量表达式
```

#### 调试帮助

assert 预处理宏 ： assert (expr); expr是表达式 ，如果expr为真 ，什么也不做，否则输出信息并结束程序执行

assert(word.size()) > threshole) ; // 要求给定单词都大于某个阈值

NDEBUG: 定义了NDEBUG ，assert就不运作，例如可以在开头定义 #define NDEBUG
```c++
int aa(const int &a,const int &b)
{
    #ifndef NDEBUG
        cerr << __func__<<endl; // __func__是编译器定义的局部静态变量 用于存放函数名字
    #endif
    assert(a > b);
    return a;
}
int main()
{
   cout<<aa(1,3)<<endl; 
   //输出
   aa
   Assertion failed: a > b, file C:\Users\??????\Desktop\c++\main.cpp, line 8
}
```
```c++
#define NDEBUG#include <assert.h> // 在assert.h添加了 NDEBUG 使 assert无效
#include<bits/stdc++.h>
using namespace std;
int aa(const int &a,const int &b)
{
    #ifndef NDEBUG
        cerr << __func__<<endl;
    #endif
    assert(a > b);
    return a;
}
int main()
{
   cout<<aa(1,3)<<endl; 
   //输出 1 
}

```
```
 4 int aa(const int &a,const int &b)
 5 {
 6    cerr << __FILE__<<endl;
 7 }
```
预处理器定义的对程序调试有帮助的名字 1.  __func__(存放函数名字)< aa > 2. __FILE__(存放文件名的字符串字面值) < C:\Users\姚杨伟\Desktop\c++\main.cpp >  3. __LINE__(存放当前行号的整形字面值)< 6 >  4. __TIME__ (存放文件编译时间的字符串字面值)< 14:58:07 > 5. __DATE__ (存放文件编译日期的字符串字面值)<Jan  7 2019>

<h3 id='1-5-6'>函数匹配</h3>

调用函数
```c++
void f();
void f(int);
void f(int ,int);
void f(double ,double=3.14);

f(5,6);
```
1. 确认候选函数

候选函数: how (1. 与被调用函数同名 2. 其声明在调用点可见) ,例子中 有 4个

2. 确认可行函数

可行函数: how(1. 其形参数量与本次调用提供的实参数相等{ 也可形参比较多 但是 形参减去有默认值的形参<=实参数}  2. 每个实参类型与对应的形参类型相同，或者能转换成形参类型 ) ，例子由 2个  void f(int); 和 void f(double ,double=3.14);

3. 寻找最佳匹配

实参和形参类型最接近，匹配的越好，找到 void f(double ,double=3.14);

```c++
void f(const int ,const int ){}
void f (const double ,const double ){}
int main()
{
    f(1,2.2); // 错误 存在二义性 ，1 更符合 int ，2.2 更符合 double
}
```
```c++
void f(){cout<<1<<endl;}
void f(const int){cout<<2<<endl;}
void f(const int ,const int ){cout<<3<<endl;}
void f (const double ,const double ){cout<<4<<endl;}
int main()
{
    f();
    f(1);
    f(1,1);
    f(1.1,1.1);
}

1 2 3 4
```
#### 实参类型转换

转换排序 1. 精准匹配 2. const转换 3. 类型提升 4. 匹配类型转换或指针转换 5. 类类型转换

所有算术类型转换级别都一样  char如果没有精准匹配会优先转换int而不是 short

能用非常量函数 会先匹配 非常量的函数

<h3 id ='1-5-7'>函数指针</h3>

函数指针: 函数的类型由它的返回类型和形参类型共同决定
```c++
void aa(const int a){cout<< a <<endl;} 
void (*p)(const int); //函数指针
int main()
{

   auto pf = &aa;  // 函数指针 pf
   p = aa;  //让 p 指向 aa
   p(0);  // 能输出 0
   (*p)(0); // 能输出 0
   p =0 // 让 p 不指向任何一个函数
}
```

如果定义了指向重载函数的指针 ，那么指针类型必须与重载函数的某一个精确匹配

与数组类似，如果实参为函数类型 会自动转换为该函数类型的指针，如果形参是函数类型，会自动转换为函数指针类型

返回指向函数的指针, 返回类型不会自动把函数类型转换为指针类型,我们必须显式转换
```c++
using F = int(int *,int); // F 是函数类型
using PF = int(*)(int *,int);// F 是指针类型
int main()
{
    PE f1(int); // 正确 ,f1返回指向函数的指针
    F * f1(int); //正确 ,f1返回指向函数的指针，显式指定
}
```
decltype作用于某个函数时，返回函数类型而不是指针类型所以我们要显式指定，

<h2 id='1-6'>类</h2>

####定义改进的类

类的基本思想是数据抽象和封装。

封装实现了类的接口和实现分离 。类的接口包含用户能实现操作 ，类的实现包括类的数据成员，负责接口实现的函数体以及定义类所需的各种私有函数

<h3 id='1-6-1'>定义抽象数据类型</h3>

成员函数的声明必须在类的内部，他的定义则可以在类的内部也可以在类的外部，作为接口组成部分的非成员函数，他们的定义和声明在类的外部

定义在类内部的函数是隐式的inline 函数
```c++
struct  node
{
   string isbn() const {return aa;} // 成员函数通过 this 的额外的隐式参数来访问调用他的那个对象,实际为 node::isbn( &a )
    string aa="1"; // 之所以aa定义在后面上面的函数还能使用是因为，编译器首先编译成员的声明，然后才轮到成员函数体
};
int main()
{
  node a;
  cout<<a.isbn(); // 输出 1 
}
```
```c++
struct  node
{
    string isbn() {return aa;} //这边应该修改为string isbn() const {return aa;} 才能正确访问 
    string aa="1";
};
int main()
{
  const node a;
  cout<<a.isbn(); //报错 ，因为this的类型是指向类 类型非常量版本的常量指针，所以 这边a为 const node 类型 不能通过非常量类型访问
}
```
```c++
在内部声明 外部定义
struct  node
{
      int uu()const;
      int u=6;
};
int node::uu() const {return u;}
int main()
{
  const node a;
  cout<<a.uu(); //输出 6 
}
```

return *this  返回调用该函数的对象

#### 定义类相关的非成员函数

如果非成员函数是类接口的组成部分，在这些函数的声明一个和类声明在同一个头文件内

IO类属于不能拷贝的类型，只能通过普通引用来传递
```c++
struct  node
{
      int u;
};
istream &readd(istream &is,node &item) // readd会改变对象内容
{
    is >> item.u;
    return is;
}
ostream &out(ostream &os,const node &item) // out不会改变所以能设置为const
{
    os << item.u;
    return os;
}
int main()
{
  node a;
  readd(cin,a); // 键入
  out(cout,a); // 输出
}
```

#### 构造函数

what：初始化对象的数据成员，构造函数不能被声明为const，因为直到构造函数完成初始化工作，对象才能真正取得常量属性(构造函数在const 对象构造过程中可以向里面写值)

合成的默认构造函数：没有显示的定义构造函数，编译器就会隐式的定义一个构造函数 (1. 如果存在类内初始值，用它来初始化 成员 2. 否则，默认初始化该成员) 

合成的默认构造函数  缺点 ：1. 只有在不包含任何构造函数下才会生成  2. 可能会执行错误操作，定义在块内的内置类型或复合类型(比如 数组和指针)的对象默认初始化，他的值时未定义的  3. 类中包含其他类成员且该成员未初始化，编译器不能初始化该成员


`= default (需要默认的行为 ，生成默认的构造函数)`
```c++
struct  node
{

    int u;
    node () = default; // 这个函数作用完全等价于合成的默认构造函数，因为我们需要其他形式的构造函数也想保留合成的默认构造函数
};
int main()
{
  node a;
  cout<<a.u<<endl;
}
```

不过有些编译器不支持类内初始化 ，`需要用 构造函数初始化列表`
```c++
struct  node
{
    node () = default; // 默认构造函数
    node (const int &a):a(a) {} // 构造函数初始化列表
    node (const int &a,const int &b): a(a) ,b(b) {}
    int a;
    int b;
};
int main()
{
  node a(1);
  cout<<a.a<<" "<<a.b<<endl; // 支持类内初始值 输出 1  0        不支持 输出 1   7012224(未定义的值)
  node b(2,3);
  cout<<b.a<<" "<<b.b<<endl; // 输出 2 3
}


如果编译器支持类内初始值 ，那么 node (const int &a):a(a) {}  相当于  node (const int &a):a(a)  b(0) {} 
如果不支持，需要显式的初始化每个内置类型的成员
```
`在类外部定义构造函数`：编写接口
```c++
struct  node
{
    node () = default;
    node (const int &a):a(a) {}
    node (const int &a,const int &b): a(a) ,b(b) {}
    node (istream &);
    int a;
    int b;
};
istream &readd(istream &is,node &item)
{
    is >> item.a >> item.b;
    return is;
}
node::node(istream &is)
{
    readd(is,*this); // 从is中读取一条信息后存入this中
}
int main()
{
   node a(cin);
   cout<<a.a<<" "<<a.b<<endl; //输入 1 2    输出 1 2 
}

```
将接口定义到内部
```c++
struct  node; // 先声明
istream &readd(istream &is,node &item); // 先声明
struct  node
{
    node () = default;
    node (const int &a):a(a) {}
    node (const int &a,const int &b): a(a) ,b(b) {}
    node (istream &is)
    {
         readd(is,*this);
    }
    int a;
    int b;
};
istream &readd(istream &is,node &item)
{
    is >> item.a >> item.b;
    return is;
}
int main()
{
   node a(cin);
   cout<<a.a<<" "<<a.b<<endl;
}
```

#### 拷贝 ，赋值和析构(销毁)

拷贝：初始化变量以及以值的方式传递或返回一个对象

赋值：让我们使用赋值运算符时

销毁：局部对象在创建他的块结束时销毁，vector(数组)销毁时，里面的内容也会销毁

不主动定义这些类，编译器会主动合成，但是有些特殊类会有问题(比如需要动态内存的类)，对于需要动态内存的类，我们最好使用vector或string

<h3 id='1-6-2'>访问控制和封装</h3>

访问说明符 ： 1. public (定义类的接口(构造函数和部分成员函数) ，在整个程序都能访问)  2. private (封装了类的实现细节(数据成员和实现部分的函数)   能被类的成员函数访问，但不能被使用该类的代码访问 )

```c++
class aa {
public:
    aa():a(0),b(0) {}
    aa(const double a,const double b):a(a),b(b) {}
    istream &readd(istream &is,aa &item)
    {
        is >> item.a >> item.b;
       return is;
    }
    aa(istream& is)
    {
         readd(is,*this);
    }
    double bb()const {return a+b; }
private:
    double a=0.0;
    double b=0.0;
};
int main()
{
  aa  p(cin);
  cout << p.bb() <<endl;
   cout<< p.a; //会报错 因为 a.a是类的私有成员
   //如果将readd函数接口放在外部也就不能访问私有成员
}
```

class和struct的区别在于默认的访问权限，claa是 private ，struct 是public

#### 友元


友元：what:类允许其他类或者非成员函数访问他的非公有成员，友元只能定义在类的内部(最后定义在开头)

```c++
class aa;
istream &readd(istream &is,aa &item);//为友元函数提供单独的声明
class aa {
    // 为aa的非成员函数所做的友元声明
    friend istream &readd(istream &is,aa &item);
public:
    aa():a(0),b(0) {}
    aa(const double a,const double b):a(a),b(b) {}
    aa(istream& );
    double bb()const {return a+b; }
private:
    double a=0.0;
    double b=0.0;
};
istream &readd(istream &is,aa &item) //readd为友元 定义在外部也能访问类的私有成员
{
    is >> item.a >> item.b;
    return is;
}
aa::aa(istream &is)
{
    readd(is,*this);
}
int main()
{
  aa  p(cin);
  cout << p.bb() <<endl;

}
```

why封装: 1. 确保用户代码不会无意破坏封装的对象 2. 被封装的类的具体实现细节可以随时改变，而无需调整用户级别的代码(使用该类的源文件需要重新编译) 

友元的声明仅仅指定了访问权限，所以我们姚专门对函数进行声明，为了使友元对类的用户可见，我们最好将友元的声明与类本身放在同一头文件，为友元函数提供单独的声明

<h3 id='1-6-3'>类的其他特性</h3>

`类型成员` ：某种类型在类中的别名.必须先定义后使用所以一般放在类开始的地方
```c++
class Screen{ // 显示器的一个窗口
    public:
        typedef std::string::size_type pos;
    private:
        pos cursor=0; //光标位置
        pos height = 0,width = 0; // 屏幕宽和高
        std::string contents;
};
```
`内联函数`： **定义**在类内部的成员函数默认为内联函数，类外部定义可以用inline关键字修饰

可变数据成员 :mutable 即使他是const对象的成员也能改变

```c++
class Screen{
public:
    int vv(){return a;}
    void aa() const {++a;} //即使是const对象也能改变 可变数据成员
private:
    mutable int a=0;
};
int main()
{
  Screen p;
  p.aa();
  cout<<p.vv()<<endl; //结果为1
}
```

类内初始值必须使用=的初始化形式或者花括号括起来的直接初始化形式

```c++
class Screen{
public:
    Screen():a(0),b(0) {}
    Screen(int a,int b): a(a),b(b) {}
    int aa(){return a;}
    int bb(){return b;}
private:
     int a;
     int b;
};
class Window_mgr {
public:
    vector<Screen> Screens{Screen(4,9)};
};
int main()
{
  Window_mgr p;
  cout<< p.Screens[0].aa()<<" "<< p.Screens[0].bb()<<endl; // 输出 4 9
}

```

#### 返回*this的成员函数

```c++
class Screen{
public:
    Screen():a(0),b(0) {}
    Screen(int a,int b): a(a),b(b) {}
    int aa(){return a;}
    int bb(){return b;}
    Screen &upa(int); // 返回的是引用，是左值，意味着返回的是对象本身而不是副本
private:
     int a;
     int b;
};
inline Screen &Screen::upa(int c) // (内联函数 返回左值，意味着可以继续运算)
{
    a+=c;   // 将c的值加到Screen对象的a值上
    return *this;
}
int main()
{
  Screen p;
  p.upa(10).upa(9);
  cout<<p.aa()<<" "<<p.bb()<<endl; // 输出 19 0
}


如果 返回的是副本  

Screen upa(int);

inline Screen Screen::upa(int c) 
{
    a+=c;   // 将c的值加到Screen对象的a值上
    return *this;
}
int main()
{
  Screen p;
  p.upa(10).upa(9); // 只有前一个有效 ，upa(10)后返回的是右值，只是副本不影响原值
  cout<<p.aa()<<" "<<p.bb()<<endl; //返回的是 10 0
}

```

##### 从const成员函数中返回 *this

一个const对象成员函数如果以引用的形式返回*this ，那么他的返回类型将是个常量引用。
```c++
class Screen{
public:
    Screen():a(0),b(0) {}
    Screen(int a,int b): a(a),b(b) {}
    int aa(){return a;}
    int bb(){return b;}
    // 根据对象是否是const重载了display函数
    const Screen &display() const {return *this;} //返回个常量引用,它引用的也是常量
    Screen &display() {return *this;}

    Screen &upa(int);
private:
     int a;
     int b;
};
inline Screen &Screen::upa(int c)
{
    a+=c;
    return *this;
}
int main()
{
  Screen p;
   p.display().upa(9);
  //p.display().upa(9); 在仅有常量版本时不工作  因为常量引用不能继续进行upa(9)操作
  cout<<p.aa()<<" "<<p.bb()<<endl;
}
```

#### 类类型

即使两个类的成员列表完全一致，他们也是不同类型
```c++
class x; // 声明 x类
class y; // 声明 y 类
class x {
y* a =nullptr; // x中包含一个指向y类的指针
};
class y {
 x b;  // y中包含一个x类
};
```

#### 友元再探

类可以让类成为自己友元 ，也可以让一个类的某个成员函数成为自己友元

类友元,aa类成员函数都能访问Screen的私有部分，友元关系不具有传递性(a的友元b，b的友元c ，c不是a的友元)
```c++
class aa;
class Screen{
    friend class aa;
private:
     int a;
};
class aa{
public:
    void f()
    {
        Screen p;
        p.a=1;
        cout<<p.a<<endl;
    }
};
int main()
{
    aa u;
    u.f();
}
```
成员函数友元，仅仅类的该成员函数能访问类的私有部分
```c++
class aa{   
public:
    void f();  //1. 先定义aa并声明 f
};
class Screen{
    friend void aa::f();  // 2. 声明并定义 Screen 包括对 f 的友元声明
private:
     int a;
};
void aa::f()   // 3. 定义 f 此时可以使用Screen的成员
{
    Screen p;
    p.a=1;
    cout<<p.a<<endl;
}
int main()
{
    aa u;
    u.f();
}
```

尽管重载函数名字相同，但是仍是不同函数，如果想对一组重载函数都声明友元，得每一个都声明

<h3 id='1-6-4'>类的作用域</h3>

一个类就是一个作用域

在我们将函数定义在类外时需要提供类名，在遇到类名后，定义的剩余部分就在类的作用域之内。如果我们要返回的类型也是类内命名的数据类型，那么之前也要提供类名 比如 aa中 定义了A类型数据  aa::A  aa:: f() {  } ，返回类型就是A类型

#### 名字查找和类的作用域

编译器处理完类中的全部声明之后才会处理成员函数的定义，如果在类外，就会只考虑在名字的使用前出现的声明

类内找不到声明，会去外层作用域中寻找。在类中，如果成员使用了外层作用域的某个名字，那么这个名字不能被重新定义

成员函数查找规则：1. 在成员函数中查找 2. 在类内查找， 3. 在成员函数之前的作用域找

我们可以通过类的名字或显式的使用this指针来强制访问成员 this ->a 或 Screen::a

当成员定义在类的外部时，名字查找的第三步不仅要考虑类定义之前的全局作用域中的声明，也要考虑成员函数定义之前的全局作用域中的声明

<h3 id='1-6-5'>构造函数再探</h3>

#### 构造函数初始值列表

如果成员是const ，引用，或者属于某种为提供默认构造函数的类类型，我们必须通过构造函数列表为这些成员提供初始值

构造函数初始值列表的初始化顺序和在类定义的出现顺序一致
```c++
class x{
    int i;
    int j;
public:
    x(int val): j(val),i(j){} //因为定义时i在前，所以i是先初始化，此时j未定义
    void aa(){cout<<i<<" "<<j<<endl;} 
};
int main()
{
    x p(2);
    p.aa(); // 结果为 4309678(未定义的) 2
}
```

#### 委托构造函数

```c++
class aa{
public:
    aa(int i,int j,int k):i(i),j(j),k(k) {}

    // 委托构造函数(其余构造函数都委托给另一个构造函数)
    aa(): aa(0,0,0) {} 
    aa(int i):aa(i,0,0) {}

    void printff() {cout<<i<<' '<<j<<' '<<k<<endl;}
private:
    int i,j,k;
};
int main()
{
     aa p(2);
     p.printff(); // 2 0  0
}
```

#### 默认构造函数的作用

如果定义了其他构造函数，最好也提供一个默认构造函数  = default;

只有当一个类没有定义任何构造函数的时候，编译器才会自动生成一个默认构造函数。

假定有一个名为 NoDefault 的类，它有一个接受 int 的构造函数，但是没有默认构造函数。定义类 C，C 有一个 NoDefault 类型的成员，定义C 的默认构造函数。
```c++
class NoDefault {
public:
    NoDefault(int i) { }
};

class C {
public:
    C() : def(0) { } 
private:
    NoDefault def;
};
```

#### 隐式的类类型转换

能通过一个实参调用的构造函数定义了一条从构造函数的参数类型向类类型隐式转换的规则
```c++
class aa{
public:
    aa(string i,int j,int k):i(i),j(j),k(k) {}
      aa(string i):aa(i,0,0) {}
    void printff() {cout<<i<<' '<<j<<' '<<k<<endl;}
private:
    string i;
    int j,k;
};
int main()
{
     string s="we";
     aa string1 = s; // 这样是可以的，通过隐式转换成了aa类型
     // aa string2 = "we" 这样是不行的，因为编译器只会自动执行一步类型转换，
                        //他要先从"we"转到string 再转成aa，有问题
     string1.printff();
}
```
通过 explicit抑制隐式转换，explicit只对一个实参的构造函数有用。因为需要多个实参的不能进行隐式转换
```c++
class aa{
public:
    aa(string i,int j,int k):i(i),j(j),k(k) {}
     explicit aa(string i):aa(i,0,0) {} // 抑制隐式操作
    void printff() {cout<<i<<' '<<j<<' '<<k<<endl;}
private:
    string i;
    int j,k;
};
int main()
{
     string s="we";
      aa string1 =s; //错误操作，因为没有支持隐式转换的构造函数了
     string1.printff();
}
```

当我们用，explicit关键字声明构造函数时，以后该函数只能以直接初始化的形式使用，不过我们可以显式强制转换
```c++
string s="we";
aa string1 = (aa(s)); // 强制转换
string1.printff();
```

#### 聚合类

what:1. 所以成员都是public 2. 没有任何构造函数 3. 没有类内初始值 4. 没有基类也没有virtual函数

我们可以用花括号括起来的成员初始值列表用来初始化集合类的数据成员:1. 初始值顺序必须与声明的顺序一致 2. 如果列表中元素小于类成员数，后面的成员被值初始化，个数不能超过成员数
```c++
struct node
{
    int a,b;
};
int main()
{
     node s={1,2};
     cout<<s.a<<' '<<s.b<<endl; // 1 2
}
```

#### 字面值常量类

字面值常量类： 数据成员都是字面值类型的聚合类 或满足以下要求:1. 数据成员都是字面值类型 2. 类必须至少含有一个constexpr的构造函数 3. 如果一个数据成员有类内初始值，那么这个初始值必须是一条常量表达式 (或者成员属于某种类类型，初始值必须使用成员自己的constexpr构造函数) 4.  类必须使用析构函数的默认定义 ,该成员负责销毁类的对象 

constexpr： 参数和返回值必须是字面值类型

constexpr构造函数可以声明成=default 形式或者是删除函数的形式，必须初始化所以数据成员，初始值是constexpr构造函数或者是常量表达式

<h3 id='1-6-6'>类的静态成员</h3>

通过static关键字使得类的静态成员和类关联在一起(一旦改变，所以该类的对象获取到的数据都会改变)

类的静态成员存在于任何对象之外，对象中不包含任何与静态成员有关的数据，不与任何对象绑在一块，不包含this指针

访问： 我们可以作用域运算符直接访问静态成员，虽然静态成员不属于类的某个对象，但我们仍然可以用类的对象，引用或指针来访问静态成员

定义： static关键字只出现在类内部的声明语句中

静态成员的类内初始化：初始值必须是常量表达式。

静态成员可以是不完全类型，也能使用静态成员作为默认实参

<h1 id='2'>c++标准库</h1>

<h2 id='2-8'>IO库</h2>

<h3 id='2-8-1'>IO类</h3>

```c++
`iostream`： 定义读写流的基本类型 `fstream`：定义读写命名文件的类型  `sstream`:定义读写内存string对象的类型

#include <iostream>
istream//从流中读取数据
ostream//从流中写入数据
iostream//读写流
 
#include <fstream>//文件
ifstream
ofstream
fstream
 
#include <sstream>//字符串
istringstream
ostringstream
iostringstream
```
cerr: 输出程序错误信息，写入到标准错误

宽字符：wchar_t

标准库通过`继承机制`使我们可以忽略不同类型流之间的差别。因为一个派生类(继承类)可以当做基类(所继承的类)使用,fstream和sstream都继承于iostream,所以iostream有的操作他们都有

#### `IO对象无拷贝或赋值`

由于不能拷贝IO对象，我们不能将形参或返回类型设置为IO对象，一般都是以**引用**方式传递和返回流
，读写IO对象会改变其状态，所以不能加const

#### `条件状态`

```c++
strm代表一种IO类型,流处于好的状态时这些bit位全部处于0。
 
strm::iostate   iostate一种机器相关类型，提供了表达条件状态的完整功能
strm::badbit    指出流已经崩溃,系统级错误，不可恢复
strm::failbit   指出IO操作失败，比如int确读入char类型，可恢复
strm::eofbit    指出流达到了文件的末尾
strm::goodbit   指出流处于未错误状态

```
```c++
s.eof()          若流处于eofbit,返回true
s.fail()         若流处于failbit,返回true
s.bad()          若流处于badbit,返回true
s.good()         若流处于有效状态,返回true
s.clear()        将流所有状态恢复，设置为有效，返回void
s.clear(flag)    将指定bit位恢复，设置为有效，返回void。flag类型为strm::iostate
s.setstate(flag) 根据给定的bit标志位，将流s中对应条件状态位置位。flag类型为strm::iostate
s.rdstate()      读取当前的状态
```

```c++

int 读入char后，得先将条件状态位复原，在清除缓存区中的错误缓存

 int p;
    for(int i=1;i<=5;i++)
    {
        cin>>p;
        if(cin.fail()) 
        {
            cin.clear(); //复原
            cout<<cin.get()<<endl; // 清除缓存
        }
        cout<<p<<endl;
    }
```

####　管理输出缓存

每个输出流有个缓存区，用来保存程序读写的数据。　有了缓存，操作系统可以将多个输出操作组合成单一的系统级写操作(提高性能)

`缓存刷新` : 1. 程序正常结束 ，mian函数的return 操作会刷新缓存 2. 缓冲区慢  3. 操纵符 endl(输出换行，刷新)，flush(刷新)和ends (输出空字符，刷新)   4. 通过 unitbuf设置实时刷新 (nounitbuf恢复到正常) 5. 一个输出流关联到另一输出流(例如 cin和cerr和cout关联，读cin或写cerr会更新cout的缓冲区)

```c++
cout<< unitbuf; // 以后所有输出操作都会立刻刷新缓存区
cout << nounitbuf; // 恢复正常
```

如果程序崩溃，输出缓冲区是不会刷新的

关联输出和输入流

tie ：1. 不带参数 ，返回指向输出流的指针(如果本对象关联到一个输出流，那就指向他，否则返回空指针) 2. 带参数， 接收一个指向ostream的指针，返回指向输出流的指针，再将自己关联到此ostream  (例如 x.tie(&o) 将流x关联到输出流o )
```c++
    cin.tie(&cout); // 将cin和cout关联，虽然cin和cout本来就关联
    cout<<cin.tie()<<endl; //0x489940，cout地址
    ostream *old_tie = cin.tie(nullptr);// old_tie值为cout地址，然后cin和空指针关联
    cout<<old_tie<<endl;//0x489940
    cin.tie(&cerr);  // cin和cerr关联，不推荐
    cout<<cin.tie()<<endl;//0x489880
    cin.tie(old_tie); // cin和cout关联
    cout<<cin.tie()<<endl;//0x489940
```

</h3 id='2-8-2'>文件输入输出</h3>

fstrem头文件定义了 ：ifstream从文件中读取数据  ，ofstream向一个文件写数据  ，fstream可以读写

在使用基类型对象的地方，我们可以用继承类型的对象代替，(如可以用fstream代替 iostream& )

```c++
ifstream in("D:\\123.txt"); // 打开文件123.txt ，然后就和cin差不多，不过是从文件中读入
ofstream out;     // 创建文件输出流
out.open("D:\\3.txt"); // 绑定到D:\\3.txt，可以不存在，不存在自己会创建
if (out) // 如果打开成功。流用open和文件关联起来，如果open失败，failbit会被置位。因为open可能会失败，所以检查流是个好习惯
{
string s;
while(in>>s)
{
    out<<s<<endl;
}
out.close();当将流文件关联到另外一个文件时，必须用close关闭当前的文件。
}

当然一个frstream对象销毁时，会自动close

123.txt ： 
1234 3534 534dfgdf

3.txt：
1234
3534
534dfgdf
```

#### `文件模式`
```
fstream fs("文件名", mode); //mode就是文件模式

in         以读的方式打开
out        以写的方式打开     
app        每次写操作前均定位到末尾(不加这个写的话是创建一个新文件写，会丢弃原文件)
ate        打开文件后立刻定位到文件末尾
trunc      截断文件(丢弃)
binary     以二进制方式IO
```
我们要保留原文件记得加app
```c++
ofstream out;
out.open("D:\\3.txt",ofstream::app); // 写 和 追加
out<<"asda34d"<<endl;
```

<h3 id='2-8-3'>string流</h3>

```c++
三个IO
istringstream
ostringstream
stringstream
特有操作
sstream strm;       strm是一个未绑定的stringstream对象
sstream strm(s);    strm是一个sstream对象，保存s的一个拷贝,此构造函数是explicit的
strm.str();         返回strm所保存的string的拷贝
strm.str(s);        将string s拷贝到strm中，返回void
```

```c++
struct node
{
    string name;// 名字
    vector<string> p;//电话号码，可能有多个
};
using namespace std;
int main()
{
    ifstream in("D:\\123.txt");
    string line,word;
    vector<node>kk; // 所有人信息的总和
    while(getline(in,line)) // 从文件中一行行读入
    {
        node info;
        istringstream record(line); // 对line中的每个单词进行处理，绑定到string读入流
        record >>info.name; //第一个单词为人名
        while(record>>word) // 后面单词都是手机号
            info.p.push_back(word);
        kk.push_back(info); 
    }
}

```
<h2 id='2-9'>顺序容器</h2>

顺序容器： 提供控制元素存储和访问顺序的能力

<h3 id='2-9-1'>顺序容器概述</h3>

```c++

#include <iterator>       //迭代器，包含c++11的begin() 和end()函数
#include <array>          //c++11 数组类型，长度固定，提供了更好、更安全的接口，执行效率和内置数组相同，可以有效替代内置数组
#include <valarray>       //c++11 值类型的数组类型，针对值类型的数组，有更多的操作，比如求和，最大最小数等。
#include <list>           //双向链表,插入删除速度快,不支持随机访问
#include <forward_list>   //c++11 单向链表,单向访问,插入删除速度快，不支持随机访问,没有size操作
#include <deque>          //双端队列,支持快速随机访问 ，在中间插入很慢，但是在头尾插入很快
#include <string>         //string类型,插入删除耗时。随机访问快，在尾部插入很快
#include <vector>         //迭代器类型,插入删除耗时。随机访问快
```
string和vector为连续储存，随机访问快,但是插入数据时，会把后面的元素全部移动来保证连续储存

list和forward_list 在哪里插入都很快，不支持随机访问，为了寻找一个元素我们得遍历一遍，有很多空间的额外开销

通常，我们一般都选择vector

<h3 id='2-9-2'>容器库概览</h3>

类型别名 
```
iterator            迭代器
const_iterator      const迭代器
size_type           无符号整数类型(迭代器大小)
different_type      带符号整数类型(迭代器距离)
value_type          元素类型
reference           元素左值类型，相当于 value_type&
const_reference     元素const左值类型
```
构造函数
```
C a
C a1(a2)
C a(b,e)            构造a，将迭代器 b和e指定的范围内元素拷贝到a(array不支持)
C a{u,i,o...}       列表初始化
```
赋值与swap (swap得容器重定义过`=`才行，因为swap用到了=)
```
c1=c2
c1 = {a,b,c...}     c1元素装换为列表元素(array不支持)
a.swap(b)
swap(a,b) 
```
大小
```
c.size()            c中元素数
c.max_size()        c可保存的最大元素数目
c.empty()
```
添加/删除元素(不适应array)，不同容器中操作接口不同
```
c.insert( args )    将args中的元素拷贝进c
c.emplace( inits )  使用inits构造c中一个元素
c.erase( args )     删除args指定元素
c.claer()           删除所以元素

关系运算符 
==   !=
< ,<= ,>,  >=       (无序关联容器不支持)
```
获取迭代器
```
c.begin()
c.end()
c.cbegin()
c.cend()
```
反向容器成员(不支持 forword_list)
```
reverse_iterator      按逆序寻址元素的迭代器
const_reverse_iterator  
c.rbegin() , c.rend()   返回指向c的尾元素和首元素之前位置的迭代器
c.crbegin(), c.crend()
```

#### `迭代器`

如果一个迭代器提供某个操作，那么所以提供相同操作的迭代器对这个操作的实现方法相同

算术迭代器只适用于 vector ， string ， array ， deque 

迭代器范围：[begin, end ) ,左闭右开区间
```
while(begin!=end)
{

}
```
#### `容器类型成员`

需要元素类型  value_type ,需要元素类型的引用:reference或const_reference
```c++
list<deque<int>>a;
list<deque<int>>::value_type p;
list<deque<int>>::iterator it;
auto k=a.begin();
```
为了读取string 的list 中的元素，应该使用什么类型？如果写入list，又应该使用什么类型？
```c++
list<string>::const_iterator // 读
list<string>::iterator // 写
```

#### `容器定义和初始化`

```c++
vector<int> c1;  // 默认初始化
vector<int> c2(c1);  // 拷贝初始化
vector<int> c3{1,2,3}; // 列表初始化
vector<int> c4(c1.begin(),c1.end()); // 迭代器初始化
vector<int> c5(10);  // 10个元素初始化
vector<int> c6(10,8); // 10个元素初始化为8
```
将一个容器初始化为另一个容器
```c++
list<string> lis1{"11","22","33"};
vector<const char*> article={"aa","bb","cc"};

list<string> lis2(lis1);

//vector<string> v1(article); 错误，拷贝初始化，容器类型和元素类型都得相同

deque<string> v1(article.begin(),article.end());  迭代器初始化，只要元素类型相同或能转换即可

list<string>::iterator  it;
it=lis1.begin();
it++;
vector<string>s(lis1.begin(),it); // 11
```

列表初始化 ，元素类型是内置类型或具有默认构造函数类型，可以只提供容器大小，否则得指定显示的元素初始值
```
vector<int> s(10,1)  提供了容器大小和类型初始值
```

> 标准array具有固定大小
```c++
array<int ,10> s; // 创建array数组，需要元素类型和容器大小
```
array内元素都被默认初始化。array可以进行拷贝
```c++
array<int,5>p={1,23,4};
array<int,5> o= p;
```

#### `赋值和swap`

array不支持assign,也不允许用花括号包围的值列表进行赋值(codeblock好像可以赋值。。)

赋值相关操作会导致指向左边容器内部的迭代器，引用和指针实现，而swap不会使其失效(除了array和string)

赋值
```
c1=c2

c={a,b,c..}         替换为初始化列表中元素的拷贝(array不适用)

swap(c1,c2)
c1.swap(c2)

assign操作不适用于关联容器和array
seq.assign(b,e)     seq元素替换为b,e迭代器范围内的元素，但是b，e不能指向seq的元素

seq.assign(il)      seq转换为初始化列表il中的元素   p.assign({123,3,5});

seq.assign(n,t)     seq元素替换为n个值为t的元素
```

assgin支持类型的转化。但是传递给assgin的迭代器不能指向本身其中的元素

容器 swap操作很快，因为元素本身为交换，只是交换了两个容器的内部数据结构。元素本身不会移动，

对string进行swap会使指向容器的迭代器，引用，指针失效

对array进行swap会真正交换元素，所以时间和元素个数成正比，因此指针，引用，迭代器所绑定元素不变，不过元素值改变了

#### `容器大小操作`

forward_list 支持max_size和empty但不支持size

#### `关系运算符`

每个容器都支持 == 和!= ,但是比较运算符< ,>等运算符两边都得是相同类型的容器，比较规则和string相似

相等运算符=通过 == 运算符实现，比较运算符通过 < 运算符实现

<h3 id='2-9-3'>顺序容器操作</h3>

顺序容器和管理容器的不同之处在于两者组织元素的方式

向一个vector和string和deque插入元素会使所有指向容器的迭代器，引用和指针失效
```c++
deque<int>v1{1,2,3};
deque<int>v2{7,8,9};
v1.push_back(4);             // 尾部创建,返回void
v1.emplace_back(5);          //速度比push_back快
v2.push_front(6);            // 头部，vector不支持,返回void
v2.emplace_front(5);
deque<int>::iterator it=v1.begin()+2;       //z指向3的迭代器
deque<int>::iterator it2=v1.insert(it,6);   // 在迭代器it指向的元素之前插入6，返回指向6的迭代器
deque<int>::iterator it3=v1.emplace(it2,7); //返回指向7的迭代器
deque<int>::iterator it4=v1.insert(it3,5,1);// 在it3之前插入5个1，返回指向第一个新添加的元素
deque<int>::iterator it5=v1.insert(it4,v2.begin(),v2.end());// 将迭代器范围内的元素插入到it4之前，
                                        //迭代器范围不能是指向v1的元素,返回指向第一个新添加的元素
v1.insert(it5,{9,9,9});//插入元素值列表，返回指向第一个新添加的元素

for(auto k:v1)
    cout<<k<<' ';//  1 2 9 9 9 5 6 7 8 9 1 1 1 1 1 7 6 3 4 5
    
// 向一个vector和string和deque插入元素会使所有指向容器的迭代器，引用和指针失效
```

向`vector和string除尾部`之外的位置或者`deque除首尾`之外的任何位置添加元素，都需要`移动元素`。
因为vector和string的分配内存是2倍增加的，超过一定内存需要扩大两倍，需要重新分配空间然后将元素从旧空间移动到新空间

我们用一个对象初始化容器或者将对象插入到容器，其实放入的是对象值的一个拷贝

insert操作除了vector和string尾部，或deque首位都很耗时

`emplace操作`:push或insert是将元素类型对象传递给他们，这些对象被拷贝到容器中，而emplace是将参数传递给`元素类型的构造函数`,在容器管理内存中直接构造，速度快

emplace函数在容器中直接构造元素，传递给emplace的函数的参数必须与元素类型的构造函数匹配
```c++
struct node
{
    node(int a,int b): a(a),b(b){}
    int a, b;
};
int main()
{
    vector<node>v; 
    v.emplace_back(1,2); // 可以传递参数给元素的构造函数，在容器中直接构造
    // v.push_back(5,4); 不能这么操作，因为是传递对象
    v.push_back(node(5,4));  
}
```

#### `访问元素`

在调用front,back(或解引用迭代器之前)要确保容器非空(!c.empty())

访问成员函数返回的都是引用，我们可以改变元素的值
```c++
vector<int>v{1,2,3,4,5,6,7};
cout<< v.back() <<endl;     // 返回v中尾元素的引用,7
cout<< v.front() <<endl;    // 返回v中首元素的引用,1
cout<< v[3] <<endl;         // 返回v中下标为n的引用,4, 如果越界，行为未定义
cout<< v.at(3) <<endl;      //返回v中下标为n的引用,4, 如果越界，抛出 out_of_range异常
cout<< v.at[7] <<endl;      // 抛出异常
```

#### `删除元素`

删除deque中除首尾之外的元素会使其任何迭代器，引用和指针失效,指向vector或string`删除点之后`任何位置的迭代器,引用和指针失效(因为他们都是`顺序容器`)
```c++
//array无这些操作，因为这些操作会改变容器大小

// forward_list 不支持pop_back
//  vector 和string 不支持pop_front
deque<int>v{0,1,2,3,4,5,6,7,8,9,10};
v.pop_back();//删除v中为元素，返回void
// 0 1 2 3 4 5 6 7 8 9
v.pop_front();// 删除v中首元素,返回void
//1 2 3 4 5 6 7 8 9
deque<int>::iterator it=v.begin()+2;
deque<int>::iterator it1 = v.erase(it); // 删除迭代器it指定元素，返回一个被删元素之后的元素迭代器
//1 2 4 5 6 7 8 9,    it1指向元素4
deque<int>::iterator it2=v.begin()+4;   //指向元素6
deque<int>::iterator it3=v.erase(it1,it2);//删除范围中的元素(左闭右开),返回最后一个被删除元素之后的元素迭代器，即指向it2原本所指元素
// 1 2 6 7 8 9
v.clear();//删除v中所有元素，返回void
```
```c++
int ia[]={0,1,2,3,4,5,6,7,8,9,10};
  vector<int>v(begin(ia),end(ia));
```

`特殊的forward_list操作`

forward_list:单向链表，因为单向链表不能很好的访问前驱，只能访问后驱，所以我们添加或删除元素都是通过改变给定元素之后的元素来进行操作

```c++
forward_list<int> v={0,1,2,3};
vector<int> p = {7,7,7};
auto it = v.before_begin();// 返回指向首元素之前不存在的元素迭代器，不能解引用
// auto it = v.cbefore_begin();

auto it1 = v.insert_after(it,9); //在it之后的位置插入9，返回指向最后一个插入元素的迭代器，返回指向9的迭代器
// 9 0 1 2 3
auto it2 = v.insert_after(it1,2,8); // 返回指向最后一个8的迭代器
// 9 8 8 0 1 2 3
auto it3 = v.insert_after(it2,p.begin(),p.end());
// 9 8 8 7 7 7 0 1 2 3
auto it4 = v.insert_after(it3,{6,6});
// 9 8 8 7 7 7 6 6 0 1 2 3
auto it5 = v.emplace_after(it4,5); // 返回指向5的迭代器
// 9 8 8 7 7 7 6 6 5 0 1 2 3
auto it6 = v.erase_after(it5); // 删除it5指向元素的后一位元素，返回指向1的迭代器
// 9 8 8 7 7 7 6 6 5 1 2 3
auto it7 = v.before_begin();  // 指向首元素之前
    it7++;           //指向首元素
auto it8 = v.erase_after(it7,it6); // it7指向 9 ，it6 指向1 ，删除it7之后和不包括it6 范围内的元素,返回指向1的迭代器
// 9 1 2 3
```

#### `改变容器大小`

resize ，如果当前大小大于所要求大小，容器后部元素会被删除，如果当前大小小于新大小，会将新元素添加到后面
```
vector<int>v(5,42);
v.resize(2);            // 42 42
v.resize(6,-1);         // 42 42 -1 -1 -1 -1
v.resize(7);            // 42 42 -1 -1 -1 -1 0
```

#### `容器操作可能是迭代器失效`

对应list和forward_list ，添加或删除元素，指向容器的迭代器，指针，引用都仍有效

deque 在除首尾的其他位置添加或删除，全失效， 添加首尾，迭代器失效，指针和引用有用，删除首，全部无影响，删除尾，尾后迭代器失效，指针和引用有效

string和vector ,插入如果存储空间不重新分配，那么和删除一样，插入之前的迭代器，指针，引用有效，之后的都无效，重新分配，都无效

每次添加或删除元素时，记得更新迭代器，引用或指针，我们不需要保存end返回的迭代器，因为他添加或删除元素后都会失效

</h3 id='2-9-4'>vector对象是如何增长的</h3>

管理容量的成员函数:
```c++
vector<int>v(5,42);
v.push_back(9);
cout<<v.capacity()<<endl; // 计算v不重新分配空间，可以保存多少元素，10
v.reserve(50);            // 分配至少能容纳50个元素的空间
cout<<v.capacity()<<endl; // 50
v.shrink_to_fit();        // 将capacity()减少为size()相同大小
cout<<v.capacity()<<endl; // 6
v.reserve(1);             // 分配的空间不会少于当前元素大小
cout<<v.capacity()<<endl; // 10
```

当size大于capacity时，容器就会重新分配空间，这时候所以元素都会移动

<h3 id='2-9-5'>额外的string操作</h3>

构造string
```c++
const char *cp = "hello world!!!"; // 以空字符结束的数组
char nonull[] = {'h','i'}; // 不以空字符结束
string s1(cp);// 拷贝cp直到遇到空字符
// hello world!!!
string s2(nonull,2);// 从nonull中拷贝两个字符，hi
string s3(nonull); // 错误， 因为nonull结尾无空字符
string s4(cp+6,5); // 从cp[6]开始拷贝，拷贝5个字符,world
string s5(s1,6,5); // 从s[6]开始拷贝，拷贝5个字符,world
string s6(s1,6);// 从s[6]开始拷贝，拷贝到末尾  world!!!
string s7(s1,6,100);// 从s[6]开始拷贝，只拷贝到末尾  world!!!
string s8(s1,100);// 抛出out_of_range异常
```
`substr`操作返回一个string,传递一个开始位置和可选的长度(不加长度就到末尾为止)

`append`:在string末尾进行插入操作的简写

`replace`:替换 s.replace(11,3,"5th")；删除s[11]开始的3个字符，然后插入"5th"

args形式
```
str
str,pos,len
cp,len  cp指向的字符数组，长度为len
cp
n,c     n个c
b,e     迭代器
初始化列表
```

#### `string搜索操作`

搜索返回size_type值，如果搜索失败，返回npos，初始化为-1，因为npos为无符号整形，-1其实为任何string最大的可能大小

```
s.find(args)                    查找s中args第一次出现位置
s.rfind(args)                   查找s中args最后一次出现位置
s.find_first_of(args)           在s中查找args任何一个字符第一次出现的位置
s.find_last_of(args)            在s中查找args任何一个字符最后一次出现的位置
s.find_first_not_of(args)       在s中查找第一个不在args中的字符的位置
s.find_last_not_of(args)        在s中查找最后一个不在args中的字符的位置

args形式
c,pos           s中位置pos开始查找字符c，pos默认0
s2,pos          从s中位置pos查找字符串s2，pos默认0
cp,pos          从s中位置pos查找c风格字符串，pos默认0
cp,pos,n        从s中位置pos查找cp所指数组的前n个字符，pos默认0
```

#### `compare函数`

比较函数
s.cpmpare(args)
```
args形式

s2
pos1,n1,s2   
pos1,n1,s2,pos2,n2   s中从pos1开始的n1个字符与s2从pos2开始的n2个字符比较

cp
pos1,n1,cp
pos1,n1,cp,n2
```

#### `数值转换`

<h3 id='2-9-6'>容器适配器</h3>

stack栈,queue队列和priority_queue优先队列

<h2 id='2-10'>泛型算法</h2>

泛型算法:可以作用于不同类型的与那是和多种容器类型

<h3 id='2-10-1'>概述</h3>

算法本身不会执行容器的操作，他们只会运行于迭代器之上

<h3 id='2-10-2'>初识泛型算法</h3>

#### `只读算法`

1. accmulate，接收三个函数，前两个是求和的范围，后一个是初值
```
vector<int>v{1,2,3};
string s({'1','2','3'});
cout<< accumulate(v.begin(),v.end(),0)<<endl;               // 6
cout<< accumulate(s.begin(),s.end(),string(""))<<endl;      // 123
cout<< accumulate(s.begin(),s.end(),""))<<endl; //报错，因为""为const char*，没有+号运算符
```
注意 给定初值的类型
```c++
vector<double>v{1.2,2,3};
cout<< accumulate(v.begin(),v.end(),0)<<endl;// 会输出6
```

2. equal比较两个序列是否保存相同的值,给定三个值，前两个是第一个元素范围，最后是第二个元素列表的首元素，长度和第一个范围一样长

```c++
vector<int>v1{1,2,3};
vector<int>v2{1,2,3,4,5};
cout<<equal(v1.cbegin(),v1.cend(),v1.cbegin())<<endl; // 输出为1

vector<string>v1={"1","2"};
vector<const char*>v2={"1","2","3"};
cout<<equal(v1.begin(),v1.end(),v2.begin())<<endl;// 1
```

#### `写容器的算法`

向目的位置写入数据的算法要保证能容纳要写入的元素

file_n(dest,n,val)，dest为迭代器，向dest开始的序列至少要包含n个元素，这些元素都被重置为val

back_inserter接受一个指向容器的引用，返回一个与该容器绑定的插入迭代器。当我们对这个迭代器赋值时，赋值运算符会调用push_back将一个具有定值的元素添加到元素

```c++
vector<int>v{1,2,3};
auto it=back_inserter(v);
*it=42;// 1 2 3 42
fill_n(back_inserter(v),3,5);//1 2 3 42 5 5 5
```

copy,接收三个迭代器，前两个拷贝范围，后一个目的起始位置copy返回拷贝到目的的元素的尾元素后面的元素
```c++
vector<int>v{1,2,3};
vector<int>v1{5,6,7,8,9};
auto ret = copy(v.begin(),v.end(),v1.begin());//1 2 3 8 9,ret是指向8的迭代器
```
replace读入一个序列，将其中的等于给定值的元素都改为另一个值
```c++
vector<int>v{1,2,3,1};
vector<int>v1{5,5,5,8,9};
replace(v.begin(),v.end(),1,9);// 9 2 3 9，改变了v的值
replace_copy(v1.begin(),v1.end(),back_inserter(v),5,4);
// v的值 9 2 3 9 4 4 4 8 9
// v1的值 5 5 5 8 9
```

#### `重排容器元素的算法`

sort通过< 运算符实现

排序去重
```c++
vector<int>v{9,9,8,8,5,3,2,2};
sort(v.begin(),v.end());//2 2 3 5 8 8 9 9
auto end_unique = unique(v.begin(),v.end());// 2 3 5 8 9 8 9 9，
//返回指向第一个重复元素的指针，指向后面的8
v.erase(end_unique,v.end());// 2 3 5 8 9
```

<h3 id='2-10-3'>定制操作</h3>

#### `向算法传递函数`

谓词：可调用的表达式，其返回结果是一个能用作条件的值

```c++
bool cmp(const string &s1,const string &s2)
{
    if(s1.size()==s2.size())
        return s1<s2;
    return s1.size()<s2.size();
}
int main()
{
   vector<string>v{"asd","gd","fhe","rgd","asdf"};
   sort(v.begin(),v.end(),cmp);
   // gd asd fhe rgd asdf
}
```

#### `lambda表达式`

lambda(匿名函数)，有 一个捕获列表，一个返回类型，一个参数列表，一个函数体(一定要有 捕获列表和函数体)

一个lambda没有默认参数，调用的实参数一定和形参数相等
```c++
vector<string>v{"asd","gd","rhe","bgd","asdf"};
sort(v.begin(),v.end(),[](const string &s1,const string &s2){return s1.size()<s2.size();});//按长度排序
int sz=3;
auto wc = find_if(v.begin(),v.end(),[sz](const string &s1){return s1.size()>=sz;});// 返回第一个满足长度大于sz的元素的迭代器
for(auto it=wc;it!=v.end();++it)
cout<<*it<<' '; // asd rhe bgd asdf
```

#### `lambda捕获和返回`

定义lambda时，编译器定义了一个新类型并创建该类型的一个对象

值捕获:
```c++
int a=42;
auto f = [a]{return a;};
a=0;
cout<<f()<<endl;//42
```
引用捕获,必须确保在lambda执行时变量是存在的
```c++
int a=42;
auto f = [&a]{return a;};
a=0;
cout<<f()<<endl;//0
```
隐式捕获,让编译器根据lambda中的代码推断我们要使用哪些变量，&表示采用捕获引用方式，=表示采用值捕获
```c++
int a=42,b=42;
auto f1 = [&]{return a;};
auto f2 = [=]{return b;};
a=0,b=0;
cout<<f1()<<' '<<f2()<<endl;//0 42

混合
int a=42,b=42;
auto f1 = [&,b]{cout<<a<<' '<<b<<endl;};
a=0,b=0;
f1();// 0 42
```
可变lambad，我们想要改变拷贝得到的值.
```c++
int a=42;
auto f1 = [=]()mutable{a++;cout<<a<<endl;};// 43
f1();
cout<<a<<' '<<endl; // 42
```

指定返回类型
```c++
[](int i)-> int { return i;};
```
#### `参数绑定`

通过绑定来代替lambda
```c++
using namespace std;
using namespace std::placeholders; // _n类型的占位符都在这个命名空间中
bool check_size(const string &s,string::size_type sz)
{
    return s.size()>=sz;
}
int main()
{
    vector<string>v{"asd","gd","rhe","bgd","asdf"};
    sort(v.begin(),v.end(),[](const string &s1,const string &s2){return s1.size()<s2.size();});
    int sz=3;
    //auto wc = find_if(v.begin(),v.end(),[sz](const string &s1){return s1.size()>=sz;});
    auto check6 = bind(check_size, _1,sz ); // 
    auto wc = find_if(v.begin(),v.end(),bind(check_size, _1,sz ));
    //_1为占位符，绑定到check_size函数，第二个形参被绑定为sz
    for(auto it=wc;it!=v.end();++it)
    cout<<*it<<' ';
}
```
bind参数
```c++
using namespace std::placeholders;
void confun(int a,int b,int c)
{
	cout << "a=" << a << ends << "b=" << b << ends << "c=" << c << endl;
}
int main(void)
{
	auto con = bind(confun, 3, _2, _1);// 第二形参是第二实参，第三形参是第一实参
	con(5,4); //a=3 b=4 c=5 
}


vector<string>v{"asd","gd","rhe","bgd","asdf"};
// 由短到长排序 gd asd rhe bgd asdf
sort(v.begin(),v.end(),cmp);
// 由长到短排序
sort(v.begin(),v.end(),bind(cmp,_2,_1));
```

如果我们想bind传递引用，使用ref函数，或cref返回const 引用
```C++
void confun(int &a,int b,int c)
{
	cout << "a=" << a++ << ends << "b=" << b << ends << "c=" << c << endl;
}
int main(void)
{
    int sz=3;
	auto con = bind(confun, ref(sz), _2, _1);// 第二形参是第二实参，第三形参是第一实参
	con(5,4); //a=3 b=4 c=5
	cout<<sz<<endl;// 4
}
```

<h3 id='2-10-4'>再探迭代器</h3>

插入迭代器，流迭代器，反向迭代器，移动迭代器

####  `插入迭代器`

```
it = t;  在it指定的当前位置插入值t

*it，++it,it++ ;   操作不会做任何事，每个操作都返回it

```

back_inserter : 创建使用push_back的迭代器

front_inserter: 创建使用push_front的迭代器

inserter: 创建使用insert的迭代器，接受第二个参数(指向给定容器的迭代器)，元素插入到给定迭代器之前,返回的还是原本的迭代器

```c++
vector<int>v{1,2,3,4};
auto it=v.begin()+3;
auto it1=inserter(v,it);
it1 = 6;//1 2 3 6 4
it1 = 7;//1 2 3 6 7 4
```

#### `流迭代器`

istream_iterator 读入流 , ostream_iterator 向输出流读写数据

```c++
vector<string>v;
// in(cin)表示关联到cin，定义时读一个
// eof是可以当尾后迭代器使用
istream_iterator<string> in(cin),eof;
// out和cout关联，每次遇上他就输出，顺便输出\n
cout<<*in<<endl; // 输出一开始读入的
ostream_iterator<string>out(cout,"\n");
// 将刚刚输出的in值全部插入到v后面，eof是尾后迭代器
v.insert(v.end(),in,eof);
// 将v中的值都复制到out中，因为调用out每个值都会被输出而且每个值后面会有\n
copy(v.begin(),v.end(),out);

//windouw : cotl + z 表示输入结尾

istream_iterator<string> in(cin),eof; //会先读入一个
for(int i=0;i<5;i++)
{
    cout<<(*in)<<' ';//输出
    ++in; // 在读入一个，不过in不会自动升序，我们得++
}
读入 1 2 3 4 5 6
输出 1 2 3 4 5

ostream_iterator<string>out(cout,"\n");
for(int i=0;i<5;i++)
{
    string s="";
    s+=i+'0';
    out=s; //赋值然后输出,每输出一个out会自动升序，指向下一个out
}
输出
0
1
2
3
4
```

istream_iterator操作
```
istream_iterator<T> in(is);         in从输入流is读取类型为T的值
istream_iterator<T> eof;            读取类型为T的尾后迭代器
*in                                 返回流中读取的值
in1 == in2
in1 != in2
in ->mem(等于 (*in).mem )
++in
in++           
```

```c++
istream_iterator<string> in(cin),eof;
for(int i=0;i<5;i++)
{
    cout<<*in++;// 读取下一条记录并将第一个记录输出s
}
```

```c++
ifstream in3("D:\\3.txt");
istream_iterator<string> in(in3),eof;
vector<string>v(in,eof);
for(auto k:v)
    cout<<k<<' ';

3.txt

asda34d
a234da34d

输出
asda34d a234da34d
```

#### `反向迭代器`

```c++
vector<string>v{"abcd","efg","hig"};
for(auto it=v.rbegin();it!=v.rend();it++)
    cout<<*it<<' '; 
// hig efg abcd
```

<h3 id='2-10-5'>泛型算法结构</h3>

迭代器类型: 1.输入迭代器 2. 输出迭代器 3.前向迭代器 4.双向迭代器 5. 随机访问迭代器

```
输入迭代器,顺序访问
1. == 和！=
2. ++
3. * (解引用只出现在赋值运算符的右侧)
* ->

find和accumulate 要求输入迭代器
istream_iterator 是输入迭代器

输出迭代器，单遍扫描
1. ++
2. * (解引用只出现在赋值运算符的左侧)

copy的第三个参数和ostream_iterator 是输出迭代器

前向迭代器，可以读写，不过只能沿一个方向

replace迭代器和forward_list迭代器

双向迭代器

除了forward_list以外的标准库

随机访问迭代器
1.双向迭代器所有功能
2.<, > ,<= ,>=
3.+,+=,-,-=
4.iter[n]

sort,array ,deque,string,vector,用于访问数组元素的指针

```

为防止重载歧义，一般能接收谓词的都有_if版本，例如find_if()

拷贝元素版本提供_copy，在提供一个目的位置迭代器表示拷贝的目的位置,例如reverse_copy(beg,end,dest)

_copy和_if可以连用 remove_copy_if(beg,end,dest,pred)

<h3 id='2-10-6'>特定容器算法</h3>

对于list和forward_list，可以通过改变元素链接实现`交换`，而不需要真的交换，提高效率

```c++
bool cmp(const int &a)
{
    if(a==6) return 1;
    return 0;
}
bool cmp2(const int &a,const int &b)
{
    return a>b;
}
int main(void)
{
   list<int>lst1{2,1,3},lst2{6,7,8};
   lst1.merge(lst2);//将lst2合并入lst1,lst2元素删除
   // lst1      2 1 3 6 7 8
   // lst2
   lst1.remove(2);// 去掉所有值为2的元素
   // lst1       1 3 6 7 8
   lst1.remove_if(cmp);// 去掉所有另cmp为真的元素
    // lst1      1 3 7 8
    lst1.reverse();// 反转
    // lst1      8 7 3 1
    lst1.sort();//使用<排序
    // lst1      1 3 7 8
    lst1.sort(cmp2);// 根据cmp2排序
    // lst1      8 7 3 1
    lst1.push_back(1);
    // lst1      8 7 3 1 1
    lst1.unique();//删除同一个值的连续拷贝
    // lst1      8 7 3 1
}

```

<h2 id='2-11'>关联容器</h2>

按`关键词`来保存和访问的,map和set

```
按关键字保存有序元素

map             关联数组：保存 关键字-值 对
set             关键字即值，只保存关键字的容器
multimap        关键字可重复出现
multiset        关键字可重复出现

无序集合

unordered_map           用哈希函数组织的map 
unordered_set           用哈希函数组织的set
unordered_multimap      哈希组织的map，关键字可以重复出现
unordered_multiset      哈希组织的set，关键字可以重复出现
```

<h3 id='2-11-1'>使用关联容器</h3>

map可以使用 e.first  ,e.second 

<h3 id='2-11-2'>关联容器概述</h3>

关联容器不支持顺序容器的位置相关操作，例如push_front和push_back，也不支持构造函数和插入操作

关联容器的迭代器都是双向的

#### `定义关联容器`

可以列表初始化，但是map每一对得用{ }
```c++
map<string,int>m{ {"a",1},{"bb",2} };
set<string>s{"a","bb","cc"};
for(auto k:m)
    cout<<k.first<<' '<<k.second<<' ';//a 1 bb 2
for(auto k:s)
    cout<<k<<' ';// a bb cc
```

map和set关键字唯一，multiset和multimap则可以出现多个相同关键字

#### `关键字类型的要求`

```c++
struct node
{
    string s;
    int t;
    node(const string &s,const int &t): s(s),t(t) {}
};
bool cmp(const node &n1,const node &n2)
{
    if(n1.s.size()==n2.s.size())
        return n1.t<n2.t;
    return n1.s.size()<n2.s.size();
}
int main()
{
    //为了使用自己比较的操作，定义时提供两个参数：关键字参数类型node和比较操作类型(函数指针类型，指向cmp)
    multiset<node,decltype(cmp)*>sett(cmp);
    sett.insert(node("a123da",3234));
    sett.insert(node("asda",32));
    sett.insert(node("asddda",32));
    for(auto k:sett)
        cout<<k.s<<' '<<k.t<<' ';
    //asda 32 asddda 32 a123da 3234
}
```

#### `pair类型`

定义在头文件utility中，一个pair保存两个数据成员，是用来生成特定类型的模版

pair的数据成员是public类型的，两个成员被命名为first和second
```
pair<int,int>p1;   //  默认初始化
pair<int,int>p2(3,4);
pair<int,int>p4={7,8};
auto p3 = make_pair(2,5);
```
pair对象的函数
```c++
pair<string,int> aa(vector<string> &v)
{
    if(!v.empty())
        return {v.back(),v.back().size()};// 尾元素和尾元素长度
        // return make_pair(v.back(),v.back().size()); 也行
    else
        return pair<string,int> ();// 隐式构造返回值
}
```

<h3 id='2-11-3'>关联容器操作</h3>

关联容器额外的类型别名
```
key_type        此容器的关键字类型

mapped_type     map容器才有，值类型

value_type      set的话和key_type相同， map为pair<key_type,mapped_type>类型
```

#### `关联容器迭代器`

当解引用关联容器迭代器时，得到value_type类型，但是，我们可以修改值类型，但是不能修改关键字类型(const )

当使用迭代器遍历map,multimap,set和multiset时，迭代器按关键字升序遍历元素
```c++
map<int,int>m;
m[2]=2;
m[1]=1;
auto it=m.cbegin();
while(it!=m.cend())
{
    cout<<it->first<<' '<<it->second<<' ';
    ++it;
}
// 1 1 2 2
```
#### `添加元素`

关联容器insert成员向容器中添加一个元素或一个元素范围，添加已存在元素对容器无影响

```c++
map<string,int>m;
m.insert({"word",1});
m.insert(make_pair("asd",2));
auto p=m.emplace(make_pair("adsd",3));
cout<<p.first->first<<' '<<p.second<<endl;// first是指向具有指定关键字的指针，second是指示是否插入成功的bool值
// adsd 1

vector<pair<string,int>> v{{"d",2},{"as",4}};
m.insert(v.begin(),v.end());

m.insert({{"34d",2},{"a6s",4}});

m.insert(p.first,{"dad",4});

m.emplace(p.first,{"dad",4});
```

#### `删除元素`

```
c.erase(k)    从c中删除每个关键字为k的元素，返回一个size_type指出删除元素的数量
c.erase(p)     p为迭代器，且p不能为c.end，删除p所指元素，返回指向p之后的迭代器
c.erase(b,e)    删除迭代器b,e所表示范围内的元素，返回e
```

####   `map的下标操作`

```c++
map<string,int>m;
m["a"]=1;
m["b"];//访问关键字为“b”的元素,默认初始化
m.at("c");// 访问关键字为“c”的元素，如果不存在返回out_of_range异常
```

#### `访问元素`

```
c.find(k);        返回一个迭代器，指向第一个关键字为k的元素，若不存在，返回尾后迭代器

c.count(k);       返回关键字等于k的元素的数量

c.lower_bound(k); 返回一个迭代器，指向第一个关键字不小于k的元素

c.upper_bound(k); 返回一个迭代器，指向第一个关键字大于k的元素

c.equal_range(k); 返回一个迭代器pair，表示关键字等于等于k的元素范围，如果不存在，两个成员都是c.end()

```

如果multimap或multiset中有多个元素具有给定关键字，那么这些元素在容器中会相邻存储

```c++
vector<pair<int,int>>v{{1,6},{2,1},{1,3},{4,3},{1,9}};
multimap<int,int>m;
m.insert(v.begin(),v.end());
auto it=m.find(1);
auto num=m.count(1);
while(num)
{
    cout<< it->second <<' ';//6 3 9
    ++it;
    --num;
}

使用equal_range

vector<pair<int,int>>v{{1,6},{2,1},{1,3},{4,3},{1,9}};
multimap<int,int>m;
m.insert(v.begin(),v.end());
for(auto p=m.equal_range(1);p.first!=p.second;++p.first)
{
    cout<<p.first->second<<' ';// 6 3 9
}
```

#### `一个单词转换的map`

```c++
map<string,string> buildmap(ifstream &map_file)
{
    map<string,string>m;
    string key;
    string value;
    while(map_file>>key&&getline(map_file,value))// 先读入一个字符串，然后把这一行剩下的都读入
    {
        if(value.size()>1)// 检查是否有转换规则
            m[key]=value.substr(1);// 跳过前导空格
        else
            throw runtime_error("no rule for"+key);// 引发一个错误
    }
    return m;
}
const string & transformm(const string &s,const map<string,string> &m)
{
    auto map_it=m.find(s);// 查找s是否存在于map中
    if(map_it!=m.cend())
        return map_it->second;// 存在，返回转换后结果
    else
        return s;// 不存在,返回原值
}
void word_transform(ifstream &map_file,ifstream &input,ofstream &output)
{
    auto m=buildmap(map_file);
    string text;
    while(getline(input,text))// 读入input的每一行
    {
        istringstream stream(text);//将这一行转化为string输入流
        string word;
        bool firstword = true;// 除了第一个单词以外之前都要加空格
        while(stream>>word)// 读入这一行的每一个单词
        {
            if(firstword)
                firstword=false;
            else
                output<<' ';
            output<< transformm(word,m);
        }
        output<< endl;// 一行结束换行

    }
}
int main()
{
   ifstream map_file("D:\\123.txt");
   ifstream input("D:\\3.txt");
   ofstream output("D:\\12.txt");
   word_transform(map_file,input,output);
}


123.txt

brb be right back
k obky?
y why
r are
u you
pic picture
thk thanks!
18r later

3.txt

where r u
y dont u send me a pic
k thk 18r

12.txt

where are you
why dont you send me a picture
obky? thanks! later
```

<h3 id='2-11-4'>无序容器</h3>

无序关联容器不是使用比较运算符来组织容器，而是通过使用一个`哈希函数`和`关键字类型的==运算符`

无序容器在存储上组织为一组桶，每个桶·保存零个或多个元素。无序容器使用哈希函数将元素映射到桶中，为了访问一个元素，容器回先计算元素的哈希值，他将指出搜索哪个桶。所有有相同哈希值的所有元素都会保存在同一个桶中。

无序容器性能依赖于哈希函数的质量和桶的数量和大小

桶接口
```c++
unordered_map<int,int>m;
m.insert({{1,2},{2,3},{5,6}});
cout<< m.bucket_count() <<endl;// 正在使用的桶的数目 ， 11
cout<< m.max_bucket_count() <<endl;// 容器能容纳的最多的桶的数目， 357913941
cout<< m.bucket_size(1) <<endl; // 第1个桶中有多少个元素， 1
cout<< m.bucket(5) <<endl;//关键字为5的元素在哪个桶中,  5
```
桶迭代
```
local_iterator              可以用来访问桶中元素的迭代器类型
const_local_iterator    
c.begin(n),    c.end(n)     桶n的首元素迭代器和尾后迭代器
c.cbegin(n),    c.cend(n)
```
哈希策略
```c++
    unordered_map<int,int>m;
    m.insert({{1,2},{2,3},{5,6}});
    cout<< m.bucket_count() <<endl;// 正在使用的桶的数目 ， 11
    cout<< m.load_factor() <<endl;// 返回每个桶的平均数量，0.272727
    cout<< m.max_load_factor() <<endl;// 最大桶存储元素数量
    m.rehash(20);  //重组存储，使使用桶数>=20，且桶数至少大于(size/max_load_factor)
    cout<< m.bucket_count() <<endl;// 23
    m.reserve(30); // 重组存储，使m可以保存20个元素，而不用rehash
    cout<< m.bucket_count() <<endl;// 31
```

无序容器使用关键字类型的==来比较元素，他们还使用一个hash<key_type> 类型的对象来生成每个元素的哈希值

我们可以直接定义关键字是内置类型的无序容器，但是不能直接定义关键字为自定义类类型的无序容器

```c++
struct node
{
    int a,b;
};
int hasher(const node &sd)
{
    return hash<int>()(sd.a);// 重载哈希函数
}
bool eqop(const node &k1,const node &k2)
{
    return k1.a==k2.a;// 重载==运算符
}
int main()
{

    using nodd = unordered_multiset<node,decltype(hasher)*,decltype(eqop)*>;

    // 参数是桶大小，哈希函数指针，和相等性判断运算符指针
    nodd aa(42,hasher,eqop);
}
```
如果我们的类定义了==运算符，我们也可以只重载哈希函数

<h2 id='2-12'>动态内存</h2>

`静态内存`:保存static对象，类static数据成员,定义在任何函数之外的变量。(在使用之前分配，在程序结束时销毁)

`栈内存`: 保存定义在函数内的非static对象，(仅在程序块运行时才存在)

程序拥有`自由空间(堆)`，用来存储`动态内存`，动态内存生存期由程序控制，当动态对象不在使用时，我们必须显式的销毁他们

<h3 id='2-12-1'>动态内存与智能指针</h3>

头文件为`memory`

动态内存的管理通过一对运算符：`new`，在动态内存中为对象分配空间并返回一个指向该对象的指针，`delete`,接收一个动态对象的指针，销毁该对象，并释放与之关联的内存

为了更安全的使用动态内存，提供智能指针，可以自动释放所指向的对象。种类:1. `shared_ptr`，允许多个指针指向同一个对象，2.`unique_ptr`，独占所指向的对象,3.`week_ptr`,一种弱引用，指向shared_ptr所管理的对象

shared_ptr和unique_ptr都有的操作
```
shared_ptr < T > sp     空智能指针，指向类型为T的对象
unique_ptr < T > up
p                       若p不为空，则为true
*p                      解引用p，获得他指向的对象
p->mem                  
p.get()                 返回p中保存的指针
swap(p,q)
p.swap(q)
```
#### `shared_ptr`

当指向一个对象的最后一个shared_ptr被销毁时，shared_ptr类会通过析构函数自动销毁此对象

```c++
int y=4;
shared_ptr<int> p1=make_shared<int>();//为默认初始化
shared_ptr<int> p2=make_shared<int>(42);
shared_ptr<string> p3=make_shared<string>(10,'9');
cout<<p1<<endl;//0x86104c
cout<<p1.get()<<endl;//0x86104c
cout<<*p3<<endl;// 9999999999

// 指向42元素的计数器变为2，指向0元素的计数器变为0，这时，0元素被销毁，内存被释放
p1=p2;
    
```
当动态对象不在被使用时，shared_ptr类会自动释放动态对象，但是仅当被指向的内存没有其他引用时，指向的内存被销毁
```c++
shared_ptr<int> aa(int a)
{
    shared_ptr<int> p = make_shared<int>(a);
    return p;// 当我们返回p时，引用技术进行递增操作
}//所以当p离开作用域时，p指向的内存不会被释放
```

使用动态内存的原因：1. 程序不知道自己需要使用多少对象 2. 程序不知道所需对象的准确类型 3. 程序需要在多个对象间共享数据

动态内存允许多个对象共享相同的状态。

`exercise.h`
```c++
#include<memory>
#include<vector>
#include<string>
#include <initializer_list>
#include <exception>
using namespace std;
class StrBlob {
public:
    using size_type=vector<string>::size_type;// 别名
    StrBlob();// 不接受参数的构造函数
    StrBlob(initializer_list<string> il);//接收参数列表的构造函数
    size_type size() const {return data->size();};// 返回大小
    bool empty() const {return data->empty();};// 是否为空
    void push_back(const string &t) {data->push_back(t);};// 添加元素
    void pop_back();// 删除元素

    string& front();// 访问首元素
    string& back();// 访问尾元素
private:
    shared_ptr<vector<string>> data;
    void check(size_type i,const string &msg) const;// 检查访问的元素是否超范围
};
StrBlob::StrBlob(): data(make_shared<vector<string>>()) {}
StrBlob::StrBlob(initializer_list<string> il): data(make_shared<vector<string>>(il)) {}
void StrBlob::check(size_type i,const string &msg) const
{
    if (i>=data->size())
        throw out_of_range(msg);
}
string& StrBlob::front()
{
    // 如果vector为空，check返回一个异常
    check(0,"front on empty StrBlob");
    return data->front();
}
string& StrBlob::back()
{
    check(0,"back on empty StrBlob");
    return data->back();
}
void StrBlob::pop_back()
{
    check(0,"pop_back on empty StrBlob");
    data->pop_back();
}
```
`main.cpp`
```c++
#include "exercise.h"
#include<iostream>
int main()
{
  StrBlob b1;
  {
      StrBlob b2 ={"a","an","the"};
      b1=b2;
      b2.push_back("about");
  }
  // 它们实际操作的是同一个vector，都包含4个元素。在代码的结尾，b2 被析构了，不影响 b1 的元素。
}
```

#### `直接管理内存`

默认情况下，动态分配是默认初始化的，即内置类型或组合类型的值时未定义的

```c++
int *pi1 = new int;     // 默认初始化,*pi1的值未定义
int *pi2 = new int();   // 值初始化为0，*pi2的值为0

auto p1 = new auto(obj);// 若poj为int，那么p1为int*
```
一般如果动态内存不足，new表达式就会失败，返回bad_alloc,我们也可以向new传递额外参数使得new不抛出异常而返回空指针
```c++
int *p1 = new int;// 分配失败，抛出 bad_alloc异常
int *p2 = new (nothrow) int;// 分配失败,new返回空指针
```

delete表达式释放内存，但是传递给delete的指针必须是指向动态分配的指针，或者是空指针

由内置指针(而不是智能指针)管理的动态内存在显式释放之前一直都会存在

当我们释放一个指针后他就变为空悬指针，要记得将得指向nullptr，但是这样还是可能存在问题，例如：所以最好使用智能指针
```c++
int *p(new int(42));// p指向动态内存
auto q = p; // q和p指向相同内存
delete p; // p和q同时无效，所以我们得将p和q都绑定到空指针上
```

####  `shared_ptr和new结合使用`

接受指针参数的智能指针构造函数是explicit的，因此我们要将内置函数转换为智能指针，必须使用直接初始化形式 
```
 shared_ptr<int> p2(new int(1024));
```
用来初始化智能指针的普通指针必须指向动态内存
```
shared_ptr<T> p(q)      p管理内置指针q所指向对象

shared_ptr<T> p(q,d)    p管理内置指针q所指向对象且p使用可调用对象d(例如lambda)来代替delete

shared_ptr<T> p(u)      p从unique_ptr u那边接管对象的所有权，u置为空

shared_ptr<T> p(q2,d)   p是shared_ptr p2的拷贝,但是用d代替delete

p.resert()              若p为唯一指向其对象的shared_ptr，resert会释放对象，将p置为空

p.resert(q)             若p为唯一指向其对象的shared_ptr,释放对象,将p指向q

p.resert(q,d)           若p为唯一指向其对象的shared_ptr,释放对象,将p指向q, 调用d而不是delete来释放q
```

不要混用普通指针和智能指针
```c++
void aa(shared_ptr<int> ptr)
{

}
int main()
{
   int *x(new int(1024));// x为普通指针
   // aa(x); 错误，不能将int *隐式转换为shared_ptr
   aa(shared_ptr<int>(x));
// 合法，但是有问题，因为临时shared_ptr对象的作用域在这个表达式结束后就结束了
// shared_ptr被销毁那么x指向的内存被释放，以后我们就不能使用x了
    int j = *x;// 错误，x是空悬指针
}
```
shared_ptr的get操作用于将指针的访问权限传递给代码。不要用get初始化或赋值给另一个智能指针
```c++
shared_ptr<int> p(new int(42));
int *q =p.get();// 现在q也指向p的对象
{//新程序块
    shared_ptr<int>(q);
}// 程序块结束，q被销毁，指向的内存被释放
int f = *p;// 未定义，因为p指向的内存已经被释放了
```
一般reset和unique一起使用，因为修改对象之前，要确认自己是否是唯一用户
```c++
shared_ptr<int> p(new int(42));
p.reset(new int(1024));// p指向一个新对象，旧对象被释放

if(!p.unique())
    p.reset(new int(*p));// 不是唯一用户，分配新的拷贝
```

#### `智能指针和异常`

在new之后对应的delete之前发送异常，且异常未被函数捕获，那么内存不会释放

 假设智能指针管理的资源不是new分配的的内存，记得传递一个删除器。当我们创建一个shared_ptr时，可以传递一个(可选的)指向删除器函数的参数，用我们传递的来代替delete,例如
```c++
struct node
{
    int a,b;
    node(){a=4,b=5;}
};
void end_node(node *p) {p->a=1,p->b=2;}
void f(node &c)
{
    shared_ptr<node> p(&c,end_node);
}
int main()
{
  node c;
  cout<<c.a<<c.b<<endl;//45
  f(c);
  cout<<c.a<<c.b<<endl;//12
}
```
#### `unique_ptr`

初始化unique_ptr只有直接初始化， unique_ptr<int> p2(new int(42));

```c++
 unique_ptr<int> p2(new int(42));
 //　p2放弃对指针的控制，返回指针，将p2置为空
 unique_ptr <int> p1(p2.release());// p1所指对象为42
 int a=7;
 int *b = &a;
 // 释放p1所指对象，提供指针b，p1指向b所指对象
 p1.reset(b);// p1所指对象为7
 
 //release一般用于初始化或赋值给另一个智能指针,不包括释放对象
```

提供删除器和shared_ptr有所不同
```c++
struct node
{
    int a,b;
    node(){a=4,b=5;}
};
void end_node(node *p) {p->a=1,p->b=2;}
void f(node &c)
{
    unique_ptr<node,decltype(end_node)*> p(&c,end_node);
    // 需要提供删除器的类型，这里为函数指针类型，decltype返回函数类型，*表示为是该类型的指针
}
int main()
{
  node c;
  cout<<c.a<<c.b<<endl;//45
  f(c);
  cout<<c.a<<c.b<<endl;//12
}
```

#### `weak_ptr`

weak_ptr是不受控制所指对象生存期的智能指针，指向由shared_ptr管理的对象，将weak_ptr绑定到一个shared_ptr时不会改变shared_ptr的引用计数,所以当最后的shared_ptr被销毁时，即使存在weak_ptr指向对象，对象还是会被释放

```c++
auto p=make_shared<int>(42);
auto pp=p;
weak_ptr<int> w(p);// 绑定shared_ptr
w=pp;
// use_count返回与w共享对象的shared_ptr数
cout<< w.use_count() <<endl;// 2
// expired，如果use_count为0,返回true，否则返回false
cout<< w.expired() <<endl;// 0
// 因为不确定所指对象是否存在，我们得调用lock，而不能直接访问
// lock 如果他共享对象的shared_ptr数为0，返回空shared_ptr，
// 否则返回指向共享对象的shared_ptr
if(auto np=w.lock())
{
    cout<<*np;// 35
}
w.reset();// 将w置为空
```

weak_ptr可以检查类中的shared_ptr对象是否还存在，p421页

<h3 id='2-12-2'>动态数组</h3>

#### `new和数组`

分配动态数组得到的是元素类型的指针，所以对于动态数组不能使用begin和end也不能用范围for

```c++
int *p = new int[10];// 10个未分配的int
int *p1 = new int[10] ();// 10个默认初始值为0的int
int *p2 = new int [10] {0,1,2};// 列表初始化，后几个为默认初始化
// 如果初始化器数目小于元素数目，new表达式会失败，不会分配任何内存
```
动态分配一个空数组是合法的，差不多相当于尾后迭代器
```
char *cp = new char[0], 正确但是不能解引用
```

释放动态数组得加方括号    delete [] p; 元素逆序销毁，最后一个元素最先销毁

`智能指针和动态数组`

unique_ptr支持管理动态数组，shared_ptr不直接支持，想用得提供自己定义的删除器

unique_ptr支持下标运算
```c++
// up指向一个包含10个未初始化int的数组
unique_ptr<int[]> up(new int[10]);
for(size_t i=0;i!=10;++i)
    up[i]=i;// 通过下标运算符为每个元素赋予新值
up.release();// 自动用delete[]销毁其指针
```
shared_ptr不支持下标运算符，也需要自己提供删除器
```c++
  shared_ptr<int>sp (new int[10],[](int *p){ delete [] p; } );
  sp.reset();// 使用我们提供的lambda释放数组，即使用delete[]

  // 不支持下标运算符
  for (size_t i=0;i!=10;++i)
    *(sp.get()+i) = i;// 使用get获取内置指针
```

#### `allocator类`

定义于memory头文件中，可以让内存分配和对象构造分离开来

```c++
  // 定义名为a的allocator对象，他可以为类型为string型的对象分配内存
  allocator<string> a;
  // 分配一段原始的，未构造的内存，保存20个类型为string的对象
  auto const p = a.allocate(20);
  auto q = p;
  // construct接收一个指针和零个或多个其他参数，其他参数用来初始化构造对象
  a.construct(q,3,'c');
  cout<<*q;// ccc
  ++q;
  a.construct(q); // *q为空字符串
  ++q;
  //还为构造就使用会报错

  // 通过对每个构造的元素调用destroy来销毁，执行析构函数
  while(q!=p)
    a.destroy(--q);// 只能对真正构造的元素进行destroy

  // 通过deallocate释放内存
  a.deallocate(p,20);
```

```c++
  vector<string> v{"a","b","c"};
  // 定义名为a的allocator对象，他可以为类型为string型的对象分配内存
  allocator<string> a;
  // 分配一段原始的，未构造的内存，保存20个类型为string的对象
  auto p = a.allocate(20);
  auto q = p;
  // uninitialized_copy返回一个指针，指向最后构造元素之后的位置
  q = uninitialized_copy(v.begin(),v.end(),q); //a b c

  q = uninitialized_copy_n(v.begin(),1,q);//a b c a

  //  uninitialized_fill(b,e,t) 迭代器b和e指定范围内创建对象，对象值为t的拷贝

  // 迭代器b指向的内存地址开始创建2个对象，对象值为do
  uninitialized_fill_n(q,2,"do");// a b c a do do
  q+=2;
  while(q!=p)
    cout<<*(p++)<<' ';

  // 通过对每个构造的元素调用destroy来销毁，执行析构函数
  while(q!=p)
    a.destroy(--q);// 只能对真正构造的元素进行destroy

  // 通过deallocate释放内存
  a.deallocate(p,20);
```

<h3 id='2-12-3'>使用标准库：文本查询程序</h3>

