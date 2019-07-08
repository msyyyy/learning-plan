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