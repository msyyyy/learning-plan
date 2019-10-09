**个人理解，并不保证正确，有大佬发现偏差欢迎指教**

## linux
1. **查看端口号**

lsof -i:端口号

2. **top**

CPU、内存、运行时间、交换分区(类似于Windows的虚拟内存，就是当内存不足的时候，把一部分硬盘空间虚拟成内存使用,从而解决内存容量不足的情况)、执行的线程

```
PID — 进程id
USER — 进程所有者
PR — 进程优先级
NI — nice值。负值表示高优先级，正值表示低优先级
VIRT — 进程使用的虚拟内存总量，单位kb。VIRT=SWAP+RES
RES — 进程使用的、未被换出的物理内存大小，单位kb。RES=CODE+DATA
SHR — 共享内存大小，单位kb
S — 进程状态。D=不可中断的睡眠状态 R=运行 S=睡眠 T=跟踪/停止 Z=僵尸进程
%CPU — 上次更新到现在的CPU时间占用百分比
%MEM — 进程使用的物理内存百分比
TIME+ — 进程使用的CPU时间总计，单位1/100秒
COMMAND — 进程名称（命令名/命令行）
```
   

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

2. **进程和线程**

   进程是系统进行资源分配的基本单位，有独立的内存地址空间

   线程是程序执行时的最小单位，是进程的执行流，与其他同一进程下的线程共享数据区，堆区，但是有自己独立的栈区

   从性能上比较，线程具有如下优点：

   1. 在一个已有进程中创建一个新线程比创建一个全新进程所需的时间要少许多
   2. 终止一个线程比终止一个进程花费的时间少
   3. 同一进程内线程间切换比进程间切换花费的时间少
   4. 线程提高了不同的执行程序间通信的效率（在大多数操作系统中，独立进程间的通信需要内核的介入，以提供保护和通信所需要的机制。但是，由于在同一个进程中的线程共享内存和文件，它们无须调用内核就可以互相通信）
   
   **协程**
   
   
- 进程和线程是os通过调度算法，保存当前的上下文，然后从上次暂停的地方再次开始计算，重新开始的地方不可预期，每次CPU计算的指令数量和代码跑过的CPU时间是相关的，跑到os分配的cpu时间到达后就会被os强制挂起。

- 协程是编译器的魔术，通过插入相关的代码使得代码段能够实现分段式的执行，重新开始的地方是yield关键字指定的，一次一定会跑到一个yield对应的地方。

**IO密集型为什么用线程不好**

当涉及到大规模的并发连接时，例如10K连接。以线程作为处理单元，系统调度的开销还是过大。当连接数很多 —> 需要大量的线程来干活 —> 可能大部分的线程处于ready状态 —> 系统会不断地进行上下文切换。 可见性能瓶颈在上下文切换。

协程的出现出现为克服同步模型和异步模型的缺点，并结合他们的优点提供了可能： 现在假设我们有3个协程A,B,C分别要进行数次IO操作。

这3个协程运行在同一个调度器或者说线程的上下文中，并依次使用CPU。 调度器在其内部维护了一个多路复用器（epoll/select/poll）。

协程A首先运行，当它执行到一个IO操作，但该IO操作并没有立即就绪时，A将该IO事件注册到调度器中，并主动放弃CPU。
这时调度器将B切换到CPU上开始执行，同样，当它碰到一个IO操作的时候将IO事件注册到调度器中，并主动放弃CPU。
调度器将C切换到cpu上开始执行。
当所有协程都被“阻塞”后，调度器检查注册的IO事件是否发生或就绪。**假设此时协程B注册的IO时间已经就绪，调度器将恢复B的执行，B将从上次放弃CPU的地方接着向下运行。**A和C同理。
这样，对于每一个协程来说，它是同步的模型；但是对于整个应用程序来说，它是异步的模型。 意味着

协程访问全局变量是不需要加锁的，因为一个个协程轮流运行
在协程轮流执行的过程中如果一个协程进入死循环，或者陷入内核态，那么其他的协程也会被堵死。



3. **进程的地址空间布局**

   从上到下是栈区，堆区，未初始化数据区，初始化数据区，代码区

4. **进程间通信的方法**

   - 管道

     半双工通信，如果要同时读写，要创建两个管道，分为有名管道和无名管道，无名管道只用于亲缘进程间

   - 信号量

     它是一个计数器，用于为多个进程提供对共享数据对象的访问。

   - 消息队列

   - 共享内存

     映射一段能被其他进程访问的内存，优点：无需复制  缺点：需要解决进程间的同步问题

   - 套接字

     可用于不同机器间的进程通信

5. **进程的创建**

   - 创建一个PCB
   - 分配内存
   - 初始化PCB
   - 将PCB放入相应的队列中( 进程有执行，阻塞，就绪三个状态)

6. **进程的终止**

   进程不再执行，但是PCB会保留一个记录(状态码和一些计时统计数据供其他进程收集，一旦收集完成，PCB完全清空)

7. **进程调度算法**

   要根据不同的系统选择合适的算法

   如**批处理系统**没有太多的用户操作，在该系统中，调度算法目标是保证吞吐量和周转时间

   - 先来先服务，按请求顺序进行服务，可能会导致短作业等待时间过长
   - 短作业优先，可能会导致长作业一直得不到调度
   - 最短剩余时间优先

   如**交互式系统**有大量的用户交互操作，在该系统中调度算法的目标是快速地进行响应

   - 时间片轮转
   - 优先级调度
   - 多级反馈队列(时间片轮转+优先级调度)

8. **死锁**

   各进程互相等待对方手里的资源，造成无法向下推进的情况

   **产生原因**： ①：系统资源不足  ②：资源分配不当  ③：进程推进顺序不合理

   **产生条件**:        ①：互斥 ，某资源只能同时被一个进程占用

   ​			②：请求和保持，进程已经保存了至少一个资源，又请求新的资源

   ​			③： 资源不可剥夺，被进程占用的资源只能等待进程自己释放

   ​			④ ：循环等待，有一个资源等待循环链，比如a等待b具有的资源，b等待a具有的资源

   **预防**:    ①：打破不可剥夺条件，当进程已占用资源又申请资源而无法满足时，则退出原占有的资源

   ​	     ②：打破请求和保持条件，采用资源预分配策略，即程序运行前申请全部所需资源，如果不满足则等待

   ​	   ③：打破循环等待条件,对所有资源编号，所有进程只有按序号递增申请资源，如果申请序号小的资源，需要把序号大的资源都释放

9. **自旋锁和互斥锁**

   互斥锁：用于保护临界区，确保同一时间只有一个线程访问数据。对共享资源的访问，先对互斥量进行加锁，如果互斥量已经上锁，调用线程会阻塞，直到互斥量被解锁。在完成了对共享资源的访问后，要对互斥量进行解锁。

   自旋锁：与互斥量类似，它不是通过休眠使进程阻塞，而是在获取锁之前一直处于忙等(自旋)阻塞状态，会占用cpu。用在以下情况：锁持有的时间短，而且线程并不希望在重新调度上花太多的成本。"原地打转"。

   自旋锁与互斥锁的区别：线程在申请自旋锁的时候，线程不会被挂起，而是处于忙等的状态。

10. **虚拟内存**

   虚拟内存的目的是为了让物理内存扩充成更大的逻辑内存，从而让程序获得更多的可用内存。

   为了更好的管理内存，操作系统将内存抽象成地址空间。每个程序拥有自己的地址空间，这个地址空间被分割成多个块，每一块称为一页。这些页被映射到物理内存，但不需要映射到连续的物理内存，也不需要所有页都必须在物理内存中。当程序引用到不在物理内存中的页时，由硬件执行必要的映射，将缺失的部分装入物理内存并重新执行失败的指令。

11. **页面置换算法**

    - 最佳 OPT

      所选择的被换出的页面将是最长时间内不再被访问，通常可以保证获得最低的缺页率。

      是一种理论上的算法，因为无法知道一个页面多长时间不再被访问。

    - 最近最久未使用 LRU

      LRU 将最近最久未使用的页面换出。

      为了实现 LRU，需要在内存中维护一个所有页面的链表。当一个页面被访问时，将这个页面移到链表表头。这样就能保证链表表尾的页面是最近最久未访问的。

    - 时钟

      第二次机会算法需要在链表中移动页面，降低了效率。时钟算法使用环形链表将页面连接起来，再使用一个指针指向最老的页面。有个标志位，如果为1代表最近有访问过，不需要置换这个

12. **分段和分页**

    段页式：

    程序的地址空间划分成多个拥有独立地址空间的段，每个段上的地址空间划分成大小相同的页。这样既拥有分段系统的共享和保护，又拥有分页系统的虚拟内存功能。

    - 对程序员的透明性：分页透明，但是分段需要程序员显式划分每个段。
    - 地址空间的维度：分页是一维地址空间，分段是二维的。
    - 大小是否可以改变：页的大小不可变，段的大小可以动态改变。
    - 出现的原因：分页主要用于实现虚拟内存，从而获得更大的地址空间；分段主要是为了使程序和数据可以被划分为逻辑上独立的地址空间并且有助于共享和保护。

13. **编译链接过程**

    ![img](https://github.com/CyC2018/CS-Notes/raw/master/notes/pics/b396d726-b75f-4a32-89a2-03a7b6e19f6f.jpg)

    - 预处理阶段：处理以 # 开头的预处理命令；
    - 编译阶段：翻译成汇编文件；
    - 汇编阶段：将汇编文件翻译成可重定位目标文件；
    - 链接阶段：将可重定位目标文件和 printf.o 等单独预编译好的目标文件进行合并，得到最终的可执行目标文件。

    

14. **静态链接**

    静态链接器以一组可重定位目标文件为输入，生成一个完全链接的可执行目标文件作为输出。链接器主要完成以下两个任务：

    - 符号解析：每个符号对应于一个函数、一个全局变量或一个静态变量，符号解析的目的是将每个符号引用与一个符号定义关联起来。
    - 重定位：链接器通过把每个符号定义与一个内存位置关联起来，然后修改所有对这些符号的引用，使得它们指向这个内存位置。

    ![img](https://github.com/CyC2018/CS-Notes/raw/master/notes/pics/47d98583-8bb0-45cc-812d-47eefa0a4a40.jpg)

    

15. **动态链接**

    静态库有以下两个问题：

    - 当静态库更新时那么整个程序都要重新进行链接；
    - 对于 printf 这种标准函数库，如果每个程序都要有代码，这会极大浪费资源。

    共享库是为了解决静态库的这两个问题而设计的，在 Linux 系统中通常用 .so 后缀来表示，Windows 系统上它们被称为 DLL。它具有以下特点：

    - 在给定的文件系统中一个库只有一个文件，所有引用该库的可执行目标文件都共享这个文件，它不会被复制到引用它的可执行文件中；
    - 在内存中，一个共享库的 .text 节（已编译程序的机器代码）的一个副本可以被不同的正在运行的进程共享。

    ![img](https://github.com/CyC2018/CS-Notes/raw/master/notes/pics/76dc7769-1aac-4888-9bea-064f1caa8e77.jpg)

16. 











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

   - UDP是无连接的，支持 广播和多播，TCP是有连接的，全双工通信的的
   - UDP是尽最大努力交付，但不保证什么，TCP保证数据按序，完整到达
   - UDP首部8字节，TCP首部20字节(不考虑偏移量)
   - UDP面向报文，TCP面向字节流，(数据块数量可以不相等，但是字节数一样)
   - UDP没有拥塞控制，TCP有拥塞控制

8. **TCP如何保证可靠性**

   - 建立连接(标志位) ：通信前确认通信实体存在
   - 序号和确认号 ：保证数据按序，完整到达
   - 检验和： 保证数据正确
   - 超时重传： 保证丢失的数据能重发
   - 滑动窗口：减轻接收方接收数据的压力大小，流量控制
   - 拥塞控制：减轻网络传输数据的压力大小

9. **TCP滑动窗口**

   TCP滑动窗口主要是为了进行流量控制，让接收端口来得及接收。

   因为TCP是全双工的，发送端和接收端都要维护一个发送窗口和接收窗口

   发送窗口是用来现在发送方可以发送的数据大小，数值依赖于接收端发来的TCP报文中的窗口字段大小

   接收窗口是标记可接收的数据大小

   发送数据分为  1.已发送且被确认 2. 已发送但未被确认 3. 未发送但可发送 4. 不可发送

   发送窗口指的是 2 ，3 部分

   接收数据分为 1.已接收 2. 未接收但准备接收 3. 未接收但不准备接收

   接收窗口为 2 部分

   发送窗口只有收到某个未确认数据的ACK响应时，才移动窗口，因为TCP时采用累积确认的方法，接收到ACK时，代表该确认号之前的数据都被正确接收

   接收窗口只有当接收到数据且最左侧连续时才移动

10. **当接收窗口为0时怎么办**

  客户端会启动零窗口探测定时器，定期向接收端发送探测报文

11. **滑动窗口的缺点**

    如果发送方网络状态不好或者接收方接收的比较慢，这样发现的都是很小的数据包，造成网络资源的浪费，可以使用nagle算法。让一个TCP连接中同时只能有一个未被确认的小分组

12. **TCP拥塞控制**

    TCP拥塞控制分为3个部分 ①. 慢启动 ②. 拥塞避免 ③. 快速恢复

    慢启动是将拥塞窗口cwnd设置为 1MSS (最大报文段长度)，然后每收到一个确认报文，cwnd增长1MSS，所以慢启动是2倍增长的

    **慢启动**有3个停止的方法 ①. 由超时指示的丢包，这时cwnd重置为1，重新开始慢启动  ② cwnd超过慢启动阈值，进入拥塞避免状态 ③ 由三个冗余ACK指示的丢包，先快速重传，然后进入快速恢复状态

    **拥塞避免**，是线性增长的，每个连接返回时间(RTT) 增长1MSS ，防止增长速率过快，导致丢包

    **快速恢复**，阈值变为原cwnd的一半，cwnd也变为原cwnd的一半，进入拥塞避免状态。让拥塞窗口不需要从1MSS开始增长，加快恢复速度

13. **ARP的工作原理**

    主机向自己所在的网络广播一个ARP请求，请求包含目标机器的网络地址，此网络的其他机器都会收到这个请求，不过只有被请求的目标机器会回应一个ARP应答，其中包含自己的物理地址

    通常维护一个高速ARP缓存，其中包含经常访问或是最近访问的目标机器IP地址到物理地址的映射，避免重复请求

14. **I/O复用 select和epoll的区别**

    IO复用是为了提高效率，让一个进程能同时监听多个文件描述符

    **select** ①每次都要将整个fd_set数组从应用进程缓冲区复制到内核缓冲区，这是很慢的(  CPU要在用户态和核心态之间切换，堆栈需要恢复上下文)

    ② select返回的是整个fd_set数组，需要通过轮询每一位才知道哪些描述符变化

    不过select也有优点，在多个操作系统中均有实现，可移植性强

    **epoll**是通过维护一个内核事件表，每个的文件描述符只需要向内核注册一次，已注册的描述符维护在红黑树上，通过回调函数，内核将发生变化的描述符 加到链表上，进程通过epoll_wait得到所有发生变化的描述符

    所有说epoll只需要将描述符从进程缓冲区向内核缓冲区复制一次，而且不需要通过轮询获得变化的描述符

    select只支持水平触发，而epoll可以设置为边缘触发

15. **水平触发和边缘触发**

    水平触发就是一个事件，你如果不处理他，他会重复触发，所以当事件发生时，你可以不立即处理他或不处理完，下次还是会触发该事件

    边缘触发，一个事件只会触发一次，所有当事件发生时，必须立刻处理完

    边缘触发减少同一事件的重复触发次数，因此效率比水平触发高

16. **同步IO和异步IO**

    IO有两个阶段

    ① 数据准备阶段   ( 非阻塞可以在此阶段做别的事)

    ② 从内核向进程复制数据 ( 异步可以在此阶段做别的事)

    同步IO是程序进行IO操作，在内核向进程复制数据这一阶段，程序被阻塞

    异步IO是内核进行IO操作，当IO完成时，通知程序数据的存储地址，所以在内核向进程复制数据这一阶段，程序可以完成别的操作

17. **键入一个url到返回网址会发生什么**

    - 先解析url是否正确

    - 在浏览器的HSTS表中查询该url是否存在，存在的话默认访问方式是https

    - 在主机host 查询是否有对应IP

    - 访问本地dns服务器

    - 如果本地dns服务器和主机在同一个子网，那么ARP请求本地dns服务器，如果不在同一个子网，那么先ARP请求默认网关，在ARP请求本地dns服务器

    - 通过dns找到目标主机ip

    - 如果是https，会先进行ssl

    - http是依赖于tcp的，所以进行tcp三次握手

    - http访问，得到请求资源

    - 在本地解析并渲染

      

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

   
