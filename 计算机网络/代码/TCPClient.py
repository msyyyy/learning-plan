from socket import *
serverName = '118.24.197.134'
serverPort = 12000
clientSocket = socket(AF_INET,SOCK_STREAM) # SOCK_STREAM代表通过TCP
# 我们仍是通过系统为我们指定客户套接字的端口号
clientSocket.connect((serverName,serverPort)) # 这条代码执行完后，执行三次握手并创建TCP连接
sentence = input('Input lowercase srntence: ')
clientSocket.send(sentence.encode()) # 只需要send 不需要目的地址
modifiedSentence = clientSocket.recv(1024) # 接收
print('From Server: ',modifiedSentence.decode())
clientSocket.close()
