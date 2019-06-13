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

```
1. 基础查询
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
select null+90; 有一方为null 则，为null
```



