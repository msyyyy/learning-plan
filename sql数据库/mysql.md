linux安装mysql [博客](<https://blog.csdn.net/zhoukikoo/article/details/78982281>)

`$ mysql -u root -p` 
访问数据库 

### mysql常用命令

```
1. 查看当前所有数据库
show databases;
2. 打开指定的库
use 库名;
3. 查看当前库的所有表
show tables;
4. 查看其他库的所有表
show tables from 库名;
5. 创建库
CREATE DATABASE 数据库名;
6. 创建表
create table 表名(
	列名  列类型,
	列名  列类型,
	........
);
7. 查看表结构
desc 表名;

8. 查看服务器版本
方法一 登陆到mysql服务端
select version();
方法二：没有登陆到mysql服务端
mysql --V
```

### mysql的语法规范

1. 基础查询

```

select  查询列表 from 表名；

特点 
1. 查询结果是一个虚拟的表格
2. 查询列表可以是: 表中的字段 常量值 表达式 函数

2. 起别名 AS  或 空格
① 便于理解
② 如果查询的字段有重名，使用别名可以区分

如果要起的别名中有特殊字符，例如空格 最好加双引号区分"out put"

3. 去重  DISTINCT

4. +号的作用
mysql只有运算符功能
select 100+90;  190
select '123'+90; 其中一方为字符型，试图将字符型转换为数值型
				转换成功 ，继续做加法运算 213
select 'ab'+90; 转换失败 ，字符型转换为0，  90
select null+90; 有一方为null 则为null

5. 用 CONCAT 函数进行拼接
select CONCAT('a','b','c') as pp;

6. ifnull
ifnull(a,b)
如果a为null 输出b 否则输出a本身
```

2. 条件查询

```;
语法：
 select 
 		查询列表
 from
 		表名
 where
 		筛选条件；
 

1. like 模糊查询
/*
 ①一般和通配符搭配使用
 		通配符
		% 任意多个字符
		- 任意单个字符
		\ 转义
*/

#案例一  查询员工中包含字符a的员工信息
SELECT *
FROM 表
WHERE
	last_name LIKE '%a%';

#案例二 查询员工名中第二个字符为 _ 的员工名
select last_name
from 表
where 
	last_name like '_$_%' ESCAPE '$'; 

ESCAPE代表这个字符是转义字符

2. between and
/*
① 提高简洁度
② 完全等价于 大于等于前一个值，小于等于后一个值
*/
案例1 查询员工编号在100到200之间的员工信息
select *
from 表
where 
	id between 100 and 200；
	
3. in
/*
 ① in列表的值类型必须兼容
 
*/
案例 查询员工的工种编号是 IT_PRGC,AD_PRES,AD_VP 中的一个员工名和工种编号
select
	last_name,
	job_id
from
	表
where
	job_id in('IT_PRGC','AD_PRES','AD_VP');
	
4. is null 
 is not null 
 
案例 查询没有奖金的员工名和奖金
select 
	last_name,
	momey
from
	表
where
	momey is null;
4. 安全等于 <=>
案例 查询没有奖金的员工名和奖金
select 
	last_name,
	momey
from
	表
where
	momey <=> null;
	

```

3. 排序查询

```
语法：
select 查询列表					 3
from 表 						   1
【where 筛选条件】				2
order by 排序列表  【asc|desc】	4	

案例1 查询部门编号>=90 的员工信息 ，按入职时间的先后排序 
select *
from 表
where id>=90
order by time asc；

案例二  按年薪的高低显示员工的信息和年薪 【按表达式排序】
select *,salary*12*(1+ifnull(奖金率,0)) as 年薪
from 表
order by salary*12*(1+ifnull(奖金率,0)) desc;

案例三  按年薪的高低显示员工的信息和年薪 【按别名排序】
select *,salary*12*(1+ifnull(奖金率,0)) as 年薪
from 表
order by 年薪 desc;

案例4 按姓名的长度显示员工姓名和工资 【按函数排序】
length()


select last_name, gongzi
from biao
order by length(last_name) desc;

案例5 查询员工信息，要求先按工资升序 ，再按员工标号降序

select *
from biao
order by gongzi asc ,id desc;


```

4. 常见函数

```

分类:1. 单行函数  ，concat ,length ,ifnull  (字符函数， 数字函数， 日期函数， 其他函数，流程控制函数)
  	2. 分组函数 即聚合函数  统计函数
  	
  	
一. 字符函数

1. length() 获取参数值的字节个数

show variables like '%char%' 显示字符集

select length('join');

2. concat 拼接字符串
select concat (lats_name,'-',first_name) 姓名 from 表；

3. upper,lower  变大写，变小写

示例 将姓变大写，名变小写然后拼接
select concat(upper(last_name),lower(first_name)) 姓名 from 表;

4. substr  截取字符
/*
   mysql索引从 1 开始
   两个参数
   
*/
select substr('abcd',2);  

> bcd

select substr('abcd',1,2);

> ab

5. instr 子串在索引串的第一次出现的索引，如果找不到返回0

select instr('abcdcd','cd');

> 3

6. trim 

去前后空格  
select trim('    abc   ');

> abc 

去前后指定字符
select trim('a' from 'aaabcdaaaefaa') as pp;

> bcdaaaef

select trim('aa' from 'aaabcdaaaefaa') as pp;

> abcdaaaef

7. lpad  用指定字符实现左填充 指定长度

select lpad('abc',10,'*') as test;

> *******abc

select lpad('abc',2,'*') as test;

> ab

8.rpad  用指定字符实现右填充 指定长度

9. replace 替换

select replace ('abcdab','ab','ef') as testl;

> efcdef


二.数学函数

1. round 四舍五入

四舍五入，保留整数
select round(-1.555) as test;

> -2

保留两位
select round(1.567,2) as test;

> 1.57

2. ceil 向上取整，返回>=该参数的最小整数

3. floor 向下取整，返回<=该参数的最大整数

4. truncate 截断

截取2位小数
select truncate(1.6999,2) as test;

> 1.69

5. mod 取余

三 日期函数

1. now 返回当期系统日期

返回当前系统日期+时间
select now(); 
> 2019-06-14 11:49:20 

返回当前系统日期不加时间
select curdate();
> 2019-06-14

返回当前时间
select curtime();
> 11:51:24 

获取指定部分，年 月 日 小时 分钟 秒
year ，month ， monthname


2. str_to_date 将日期格式的字符转换为指定格式的日期

select str_to_date('9-13-1999','%m-%d-%Y');

> 1999-09-13

3. date_format 将日期转换成字符

select date_format(now(),'%y年%m月%d日') as test;
> 19年06月15日

4. datediff() 两个日期相差天数

四 其他函数

版本号
select version() ;

select database();

select user();

五 流程控制函数

1. if函数,类似三目运算符
select if(10<5,'a','b');
> b

2. case函数

使用一，类似于switch
语法:
case 要判断的字段或表达式
when 常量1 then 要显示的值1或语句1
when ..2..
..
else 要显示的值n或语句n;
end

案例： 查询员工的工资，要求
部门号=30 ，显示工资为1.1倍
部门号=40,显示工资为1.2倍
否则为原工资

select  员工名，工资,部门号，
case 部门号
when 30 then 1.1*工资
when 40 then 1.2*工资
else 工资
end as 新工资
from 表

使用二 类似于多重if
case 
when 条件1 then 要显示的值1或语句1
when 条件2 then 要显示的值2或语句2
else 要显示的值n或语句n
end

案例  查询员工工资情况
如果工资>20000 ，显示a级别
如果工资>15000,显示b级别
如果工资>10000,显示c级别
否则显示d级别

select 工资,
case
when 工资>20000 then 'A'
when 工资>15000 then 'B'
when 工资>10000 then 'C'
else 'D' 
end as 工资级别
from 表





```

```
分组函数

sum 求和
avg 平均值
max 最大值
min 最小值
count 计算不为空个数

sum，avg 一般用于处理数值型
max,min,count 可以处理任何类型
所有分组函数都忽略null值

可以和distance搭配，实现去重

count函数
select count(*)  from 表; 统计行数

select count(常量值) from 表; 也相当于统计行数，因为多加了一列全为该常量的列

```

```
分组查询 group by

分组前筛选	原始表			group by 前	where
分组后筛选	分组后结果集	  group by 后  having

语法
select 分组函数，列 (要求出现在group by 后面)
from 
[where ] 
[group by 分组的列表]
[having ]
[order by ]

案例1 查询每个工种的最高工资

select max(工资),工种
from 表
group by 工种

案例2 查询那个部门的员工个数大于2
having可以进行分组后筛选

select 部门
from 表
group by 部门
having count(*)>2;

案例3 查询每个工种有奖金的员工的最高工资>12000的工种编号和最高工资
select 工种编号, max(工资)
from 表
where 奖金 is not null
group by 工种编号
having max(工资)>12000;

案例4 查询领导编号>102的每个领导手下最低工资>5000的领导编号是那个，以及最低工资
select 领导编号 ,min(工资)
from 表
where 领导编号>102
group by 领导编号
having min(工资)>5000;


#按表达式分组
having 支持别名

案例 按员工姓名长度分组，查询每一组的员工个数，筛选员工个数>5的有哪些

select count(*) c
from 表
group by length(姓名)
having c>5;

# 按多个字段分组

案例 查询每个部门每个工种的员工的平均工资

select avg(工资)，部门，工种
form 表
group by 部门，工种;

# 添加排序


```

```
连接查询

又称多表查询，当查询字段来自多个表时，会用到连接查询
sql92标准，仅支持内连接
sql99标准，全支持

按功能分类
	内连接：
		等值连接
		非等值连接	
		自连接
	外连接：
		左外连接
		右外连接
		全外连接(sql99不支持)
	交叉连接

#1. 等值连接
案例1 查询女神名和对应的男神名
select name,boyname
from boys,beauty
where beauty.boyfriend_id =boy.id;

可以为表起别名，提高语句简洁度，区分多个重名的字段
如果为表起了别名，那么查询字段不能用原来的表名去限定

案例2 查询每个城市的部门个数

select count(*) 个数，city
from d,l
d.id=l.id
group by city;

案例3 查询每个工种的工种名和员工的个数，并且按员工个数降序

select 工种名，count(*)
from 工种,员工
where 工种id = 员工的工种id
group by 工种
order by count(*) desc;

#2.非等值连接

案例1. 查询员工的工资和工资级别
员工表  工资  
工资级别表  等级   最低工资  最高工资

select 工资，工资级别
from 员工，级别表
where 工资 between 级别表.最低工资 and  级别表.最高工资

#3. 自连接

案例  查询 员工名和上级的名称

员工表 编号  名字 上级编号

select a.名字，b.名字
from 员工表 a ,from 员工表 b
where a.上级编号 = b.编号
```

```
sql99语法
select 查询列表
from 表1 别名 【连接类型】
join 表2 别名 
on 连接条件
【where 筛选条件】
【group by 分组】
【having 筛选条件】
【order by 排序列表】

分类
内连接 inner，可以省略
外连接
	左外 left  【outer】
	右外 right 【outer】
	全外 full  【outher】
交叉连接 cross


内连接
select 查询列表
from 表1 别名
inner join 表2 别名
on 连接条件

#1. 等值连接

案例1 查询名字包含e的员工名和工种名
select 员工名,工种名
from 员工 a
inner join 工种 b
on a.工种 = b.工种
where a.名字 like '%e%';

案例2 查询部门个数> 3 的城市名和部门个数
select 城市名,count(*)
from 城市 a
inner join 部门 b
on a.城市 = b.城市
group by 城市
having count(*) >3;

#2. 非等值连接

#3. 自连接

```

```
外连接
分主从表，主表中有，从表没有，用null填充主表，主表信息都会显示
/*

应用场景：用于查询一个表有，另一个表没有的记录
*/
```



