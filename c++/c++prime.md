
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

