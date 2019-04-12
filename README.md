## 操作系统

1. **进程的状态**

   - 可执行状态   R

   ​	这些进程的PCB被放入可执行队列中

   - 可中断的睡眠状态  S

     等待某某事件发生，PCB在对应事件的等待队列中

   - 不可中断的睡眠状态   D

     内核的某些处理流程是原子操作，是不能中断的

   - 暂停状态  T

     进程收到SIGSTOP，SIGSTP ,SIGTIN, SIGTOU信号后终止

   - 僵尸状态  Z

     进程占用资源被回收，除了PCB，PCB中保存进程的退出码和一些信息

2. 


## 网络

1. **TCP三次握手**

   假设A为客服端，B为服务器端

   - B处于listen状态
   - A向B发送一个请求报文，SYN置为1，ACK为0，序号为x，A进入SYN_SENT状态
   - B向A发送确认报文，SYN置为1，ACK为1，序号为y，确认号为x+1，B进入SYN_RECVD状态
   - A向B发送确认报文，ACK置为1，序号为x+1，确认号y+1,A结束SYN_SENT状态，B结束SYN_RECVD状态
   - 只有第三次握手才能传输数据

2. **为啥需要三次握手**

   防止已失效的请求报文再次到达服务器端，造成服务器端的资源浪费

3. **TCP四次挥手**

   假设A为客户端，B为服务器端

   - A向B发送一个请求断开报文，FIN置为1，序号为 x，A进入FIN_WAIT_1状态
   - B向A发送确认报文，ACK为1，序号为y，确认号为x+1，B进入CLOSE_WAIT状态
   - B向A发送请求断开报文，FIN置为1，序号为z，确认号为x+1,B进入LAST_ACK状态，A进入FIN_WAIT_2状态
   - A向B发送确认报文，序号为x+1，确认号为z+1，A进入TIME_WAIT状态，持续2MSL，B接收LAST_ACK状态

4. **为什么要四次挥手**

   因为TCP是全双工的，客户端和服务器端都要发送请求断开报文和确认报文，当服务器收到请求断开的报文时，服务器可能还有未传输完的数据，等数据传输完后才会发送请求断开报文


5. **TIME_WAIT的作用**

   - 确保客户端发送的确认报文能到达客户端
   - 确保本次TCP连接中产生的有效报文段都从网络中消失

6. **为什么要2MSL**

   - ACK从A到B最多1MSL，超过这个时间，B会重发FIN
   - B重发的FIN报文最多1MSL到达A

   所以如果B重发FIN报文，在2MSL时间内一定能被A收到

7. **UDP和TCP的区别**

   - UDP是无连接的，支持 广播和多播，TCP是全双工通信的的
   - UDP是尽最大努力交付，但不保证什么，TCP保证数据按序，完整到达
   - UDP首部8字节，TCP首部20字节(不考虑偏移量)
   - UDP面向报文，TCP面向字节流，(数据块数量可以不相等，但是字节数一样)
   - UDP有拥塞控制，TCP没有拥塞控制

8. **TCP如何保证可靠性**

   - 建立连接(标志位)   ：通信前确认通信实体存在
   - 序号和确认号 ：保证数据按序，完整到达
   - 检验和： 保证数据正确
   - 超时重传： 保证丢失的数据能重发
   - 滑动窗口：减轻接收方接收数据的压力大小，流量控制
   - 拥塞控制：减轻网络传输数据的压力大小

9. 

10. 

## C++  

1. **右值引用**

   左值： 地址值，可以被赋值

   右值：数据值，不可被赋值，不能被访问地址，即将消亡的数据

   如 int a,b   为左值   a+b为右值

   右值引用可以延长右值的生存周期，窃取其资源而不需要进行拷贝

   例如有个函数返回一个类的对象，那么他先对返回的那个值进行了构造函数，然后接收返回值的那个值调用拷贝构造函数

   如果使用右值引用，就能使用移动构造函数，窃取函数中的返回值对象资源，不需要再调用拷贝构造函数

   之前不能被访问的右值地址现在为接收值的地址了，接收值是左值能被访问

   move可以将左值变为一个右值，move指针，接收值指向原内存地址，原指针失效，move基础值无意义

2. **智能指针**

   c++11后只有shared_ptr、unique_ptr、weak_ptr

   智能指针主要是为了解决内存泄漏的问题，智能指针把普通指针封装成类，然后使用**引用计数**，在引用数为0时，自动释放占用的内存。普通指针可能会有忘记释放或多次释放的问题，智能指针不会

   shared_ptr，多个指针能指向相同的对象，不过可能出现循环引用的问题，循环引用可以由week_ptr解决

   week_ptr，弱指针，指针指向对象，但是对应引用计数不增加，观测资源。**循环引用**是a类有指针指向b，b类有指针指向a（现在新建2个指针  a* p1, b* p2   ,先将p1=nullptr，但是指向a的引用计数为2，现在为1，不会将a析构，p2也一样，所以最后a和b都不会被析构），**解决**：类里的指针用week_ptr

   unique_ptr, 同一时刻只能有一个unique_ptr指向给定对象（通过禁止拷贝语义、只有移动语义来实现）

3. **多态**

   多态分为动态时多态和静态时多态

   静态多态是**函数重载**，在编译时确定调用哪个函数。函数重载是函数相同，形参列表不同( 形参类型 和数量 ) 和返回值类型无关

   动态多态是**虚函数**，在运行时确定调用哪个类的虚函数。一个基类指针指向派生类对象，当对这个指针调用虚函数时，调用的虚函数为派生类的虚函数，这就是动态绑定

   **如何实现**: 每个有虚函数的类都有一个虚函数表指针，指向虚函数表。虚函数表存储着类的所有虚函数位置，派生类继承基类时会先继承基类的虚表然后再对其进行修改

   每个类对象再创建时，编译器会自动加上那个虚表指针，所以使用该对象时，调用的就是该类虚表中的虚函数

4. **数组名和指针的区别**

   数组名相当于含有数组大小的指针常量

   数组名不能++，--

   sizeof(指针) ，显示指针大小，sizeof(数组名) ,显示数组大小

## 数据库

1. **三大范式**

   - 第一范式： 属性不可分割
   - 第二范式： 非主属性完全函数依赖于候选码
   - 第三范式： 非主属性不传递函数依赖于候选码
   - BCNF范式： 主属性完全函数依赖于候选码且不传递函数依赖于候选码

   在一张表中，若 X → Y，且对于 X 的任何一个真子集（假如属性组 X 包含超过一个属性的话），X ' → Y 不成立，那么我们称 Y 对于 X **完全函数依赖**

   在『Y 不包含于 X，且 X 不函数依赖于 Y』这个前提, 假如 Z 函数依赖于 Y，且 Y 函数依赖于 X，那么我们就称 Z **传递函数依赖**于 X

   设 K 为某表中的一个属性或属性组，若除 K 之外的所有属性都完全函数依赖于 K（这个“完全”不要漏了），那么我们称 K 为**候选码**，简称为**码**。在实际中我们通常可以理解为：**假如当 K 确定的情况下，该表除 K 之外的所有属性的值也就随之确定，那么 K 就是码。**一张表中可以有超过一个码。（实际应用中为了方便，通常选择其中的一个码作为**主码**）

2. **事务**

   满足ACID的一组操作

   - A  原子性

     事务被视为不可分割的最小单元，事务的所有操作要么都成功要么都失败回滚

     回滚可以用回滚日志来实现，回滚日志记录着事务所执行的修改操作，在回滚时反向执行这些修改操作即可。

   - C  一致性

     数据库在事务执行前后都保持一致性状态。在一致性状态下，所有事务对一个数据的读取结果都是相同的。

   - I  隔离性

     一个事务所做的修改在最终提交以前，对其它事务是不可见的。

   - D 持久性

     一旦事务提交，则其所做的修改将会永远保存到数据库中。即使系统发生崩溃，事务执行的结果也不能丢失。使用重做日志来保证持久性。

3. 

