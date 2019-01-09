from socket import *
serverPort = 12000
serverSocket = socket(AF_INET,SOCK_DGRAM)
serverSocket.bind(('',serverPort)) # 将端口号12000与该服务器的套接字serverSocket相连
print("The server is ready to receive")
while True: # 一直保持 等待连接状态
    message,clientAddress = serverSocket.recvfrom(2048) # 接收数据 和客户源地址
    modifiedMessage = message.decode().upper() # 转换为字符串后转换为大写
    serverSocket.sendto(modifiedMessage.encode(),clientAddress) # 发送回去
