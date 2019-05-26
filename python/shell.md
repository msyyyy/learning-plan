# 第一章 小试牛刀

执行一个脚本  bash  a.sh

1. `echo`

终端打印

默认情况下， echo 会在输出文本的尾部追加一个换行符。可以使用选项 -n 来禁止这种行为。
echo 同样接受双包含转义序列的双引号字符串作为参数。在使用转义序列时，需要使用 echo -e
" 包含转义序列的字符串 " 这种形式

2. `查看环境变量`

   ```
   要查看其他进程的环境变量，可以使用如下命令：
   cat /proc/$PID/environ
   其中， PID 是相关进程的进程ID（ PID 是一个整数）。
   假设有一个叫作gedit的应用程序正在运行。我们可以使用 pgrep 命令获得gedit的进程ID：
   $ pgrep gedit
   12501
   那么，你就可以执行以下命令来查看与该进程相关的环境变量：
   $ cat /proc/12501/environ
   ```

3. 

   ```
   注意， var = value 不同于 var=value 。把 var=value 写成 var = value
   是一个常见的错误。两边没有空格的等号是赋值操作符，加上空格的等号表示的
   是等量关系测试。
   在变量名之前加上美元符号（$）就可以访问变量的内容。
   var="value" # 将"value" 赋给变量var
   echo $var
   也可以这样写：
   echo ${var}
   输出如下：
   value
   ```

4. 可以用下面的方法获得变量值的长度：
   length=${#var}

   root用户的 UID 是 0 

5. 使用shell进行数学运算

   let ,  [] 和 （（））

   

   ```
   #!/bin/bash
   # 文件名:a.sh
   no1=4;
   no2=5;
   let a1=no1+no2
   echo $a1
   a2=$[no2 - no1]
   echo $a2 
   ```

6. 文件描述符与重定向

   ```
   0 代表标准输入  1 代表标准输出  2 代表标准错误
   (1) 使用大于号将文本保存到文件中：
   $ echo "This is a sample text 1" > temp.txt
   该命令会将输出的文本保存在temp.txt中。如果temp.txt已经存在，大于号会清空该文件中
   先前的内容。
   (2) 使用双大于号将文本追加到文件中：
   $ echo "This is sample text 2" >> temp.txt
   
   你可以将 stderr 和 stdout 分别重定向到不同的文件中：
   $ cmd 2>stderr.txt 1>stdout.txt
   
   也可以使得 stderr 和 stdout 都被重定向到同一个文件中：
   $ cmd 2>&1 alloutput.txt
   或者这样
   $ cmd &> output.txt
   ```

7. 别名

   ```
   1) 创建别名。
   $ alias new_command='command sequence'
   下面的命令为 apt-get install 创建了一个别名：
   $ alias install='sudo apt-get install'
   定义好别名之后，我们就可以用 install 来代替 sudo apt-get install 了。
   (2) alias 命令的效果只是暂时的。一旦关闭当前终端，所有设置过的别名就失效了。为了
   使别名在所有的shell中都可用，可以将其定义放入~/.bashrc文件中。每当一个新的交互式
   shell进程生成时，都会执行 ~/.bashrc中的命令。
   $ echo 'alias cmd="command seq"' >> ~/.bashrc
   (3) 如果需要删除别名，只需将其对应的定义（如果有的话）从~/.bashrc中删除，或者使用
   unalias 命令。也可以使用 alias example= ，这会取消别名 example 。
   (4) 我们可以创建一个别名 rm ，它能够删除原始文件，同时在backup目录中保留副本。
   alias rm='cp $@ ~/backup && rm $@'
   
   
   对命令进行转义可以忽略别名
   $ \command
   字符 \ 可以转义命令，从而执行原本的命令
   
   alias 命令可以列出当前定义的所有别名：
   ```

   

   



