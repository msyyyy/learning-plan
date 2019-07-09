package main

import (
    "net/http"
)

func main() {
    http.Handle("/", http.FileServer(http.Dir(".")))
    http.ListenAndServe(":8080", nil)
}
/*
HTTP 文件服务器是常见的 Web 服务之一。
开发阶段为了测试，需要自行安装 Apache 或 Nginx 服务器，下载安装配置需要大量的时间。
使用 Go语言实现一个简单的 HTTP 服务器只需要几行代码，如下所示。
*/
/*
下面是代码说明：
第 1 行，标记当前文件为 main 包，main 包也是 Go 程序的入口包。
第 3~5 行，导入 net/http 包，这个包的作用是 HTTP 的基础封装和访问。
第 7 行，程序执行的入口函数 main()。
第 8 行，使用 http.FileServer 文件服务器将当前目录作为根目录（/目录）的处理器，访问根目录，就会进入当前目录。
第 9 行，默认的 HTTP 服务侦听在本机 8080 端口。
*/