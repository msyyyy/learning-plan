from socket import *
serverName = '118.24.197.134' # 服务器端主机号
serverPort = 12000 # 服务器端 端口号
clientSocket = socket(AF_INET,SOCK_DGRAM) # 创建客户的套接字，称为clientSocket，
# AF_INET 为一个地址簇 ，指示底层网络使用 IPv4  SOCK_DGRAM 代表 他是一个UDP套接字
# 我们通过系统为我们指定客户套接字的端口号
message = input('Input lowercase srntence: ')  # 用户输入一行 与提示信息
clientSocket.sendto(message.encode(),(serverName,serverPort)) 
# encode()把字符串变成字节   sendto发送到指定主机的端口号
modifiedMessage, serverAddress = clientSocket.recvfrom(2048)
# recvfrom 从服务器接收返回的数据 ，2048为缓存长度 前一个存数据 后一个存服务器源地址
print(modifiedMessage.decode()) #decode 将字节转换为字符串，然后输出
clientSocket.close() # 关闭
