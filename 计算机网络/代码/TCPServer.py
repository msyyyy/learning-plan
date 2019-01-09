from socket import *
serverPort = 12000
serverSocket = socket(AF_INET,SOCK_STREAM)
serverSocket.bind(('',serverPort)) # 绑定套接字和端口号,serverSocket是我们的欢迎套接字
serverSocket.listen(1)  # 让服务器聆听来自客户的TCP连接请求，定义请求连接的最大数为1
print("The server is ready to receive")
while True:
    connectionSocket, addr = serverSocket.accept() 
    # 当客户敲门时，程序为serverSocket调用accept()方法,创建了新的连接套接字，由这个特定用户专用
    #　此时客户与服务器完成握手并connectionSocket建立TCP连接
    sentence = connectionSocket.recv(1024).decode()
    capitalizedSentence = sentence.upper() # 大写
    connectionSocket.send(capitalizedSentence.encode())
    connectionSocket.close()# 关闭connectionSocket套接字
