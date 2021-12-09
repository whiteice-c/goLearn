package main

import "goLearn/httpServer"

//编写一个HTTP服务器
//1.接受客户端的request，并将request中带的header写入response header
//2.读取当前系统的环境变量中的VERSION配置，并写入response header
//3.Server端记录访问日志包括客户端ip，HTTP返回码，输出到server端的标准输出
//4.当访问localhost/healthz时，应返回200

func main() {
	httpServer.ServerStart()
}
