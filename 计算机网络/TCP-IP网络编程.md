- [开始网络编程](#1)
    - [理解网络编程和套接字](#1-1)
    - [套接字类型与协议设置](#1-2)
    - [地址族与数据序列](#1-3)


<h1 id='1'>开始网络编程</h1>

<h2 id='1-1'>理解网络编程和套接字</h2>

```
linux 头文件 #include <sys/socket.y>
windows 头文件 #include <winsock2.h>
```

### `基于linux平台的实现`

网络编程结束连接请求的套接字创建过程为
```
1. 调用socket函数创建套接字

int socket(int domain,int type ,int protocol);

2. 调用bind函数分配IP地址和端口号

int bind(int sockfd, struct sockaddr *myaddr, socklen_t addrlen);

3. 调用listen函数转化为可接收请求状态

int listen(int sockfd, int backlog);

4. 调用accept函数受理连接请求

int accept(int sockfd, struct sockaddr *addr , socklen_t *addrlen);

```
linux不区分文件和套接字
```
打开文件 
#include <sys/types.h>
#include <sys/stat.h>
#include <fcntl.h>
open(const char *path , int flag);// path为文件地址， flag为文件开始模式，可能有多个，由|连接
例如 fd = open("data.txt",O_CREAT|O_WRONLY|O_TRUNC)

O_CREAT     必要时创建文件
O_TRUNC     删除全部现有数据
O_APPEND    维持现有数据，保存到后面
O_RDONLY    只读打开
O_WRONLY    只写打开
O_RDWR      读写打开

关闭文件
#include <unistd.h>
int close(int fd);// fd为文件描述符

将数据写入文件
#include <unistd.h>
ssize_t write(int fd,const void * buf ,size_t nbytes)

size_t为无符号整形(unsigned int)的别名， ssize_t是signed int 类型

读取文件中数据
#include <unistd.h>

ssize_t read(int fd,void *buf,size_t nbytes);
// fd 文件描述符 ，buf 保存接收数据缓冲地址值 nbytes 接收数据最大字节数
```

### `基于Windows平台的实现`

进行 Winsock编程时，首先调用WSAStartup函数
```
#include <winsock2.h>

int WSAStartup(WORD wVersionRequested , LPWSAData lpWSAData);
程序员要用的winsock版本信息 和 WSADATA结构体变量的地址值
```
Winsock编程的基础公式,初始化Winsock库
```
int main(int argc,char* argv[])
{
    WSADATA wsaData;
    ....
    if(WSAStartup(MAKEWORD(2,2),&wsaData) != 0)// MAKEWORD(1,2) 主版本号为1，副版本号为2，返回0x0201
        ErrorHandling("WSAStartup() error!");
    ....
    return 0;
}
```
`注销库，int WSACleanup(void);  成功时返回0，失败时返回SOCKET_ERROR`

### `基于Windows的套接字相关函数及展示`

```
SOCKET socket(int af,int type,int protocol)

int bind(SOCKET s, const struct sockaddr *name , int namelen);

int listen(SOCKET s, int backlog)

SOCKET accept(SOCKET s, struct sockaddr *addr , int * addrlen) 成功时返回套接字句柄

int connect (SOCKET s, const struct sockaddr *name ,int namelen)

关闭套接字函数，在linux中关闭文件和关闭套接字都会调用close函数，而windows中有专门关闭套接字的函数

int closesocket(SOCKET s)

```
winsock数据传输函数
```
int send(SOCKET s, const char *buf, int len ,int flags); 成功返回传输字节数
s 套接字句柄值  buf 保存待传输数据的缓冲地址值， len 传输字节数，flags 多项选项信息

和linux的 send函数相比，只多了flags参数

和send对应的 recv函数 ，接收数据
int recv(SOCKET s, const char *buf ,int len , int flags); 成功返回接收的字节数
```

<h2 id='1-2'>套接字类型与协议设置</h2>

```
int socket(int domain, int type ,int protocol)

domain : 套接字中使用的协议族信息

type: 套接字数据传输类型信息

protocol: 计算机间通信使用的协议信息
```
协议族 : 协议分类信息
```
PF_INET         IPv4互联网协议族
PF_INET6        IPv6
PF_LOCOL        本地通信的UNIX协议族
PF_PACKET       底层套接字的协议族
PF_IPX          IPX Novell协议族
```
套接字类型(type)：套接字的数据传输方式

1. 面向连接的套接字(SOCK_STREAM)

特征：可靠，按序基于字节的面向连接(一对一)的数据传输方式的套接字 

2. 面向消息的的套接字(SOCK_DGRAM)

特征: 不可靠，不按序，以数据的高速传输为目的的套接字

具体指定协议信息(protocol)

为啥需要第三个参数: 同一协议族中存在多个数据传输方式相同的协议

TCP套接字(IPPROTO_TCP) ， write函数调用次数可以和不同于read函数调用次数

<h2 id='1-3'>地址族与数据序列</h2>

### **分配给套接字的IP地址与端口号**

IP是为收发网络数据而分配给计算机的值，端口号是为区分程序中创建的套接字而分配给套接字的序号

IPv4： 4字节地址族  IPv6 ： 16字节地址族

IPv4标准的4字节IP地址分为网络地址和主机地址，且根据网络ID和主机ID所占字节的不同，分为A(0-127)，B(128-191)，C(192-223)，D，E

主机传输数据是先根据网络ID发送到相应路由器或交换机然后在根据主机ID向目标主机传递数据


