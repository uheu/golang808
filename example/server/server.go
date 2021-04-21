package main

import (
	"fmt"
	"net"
	"golang808"
)

func main(){
	//1.监听端口:9500
	fmt.Println("启动服务端:9500")
	//net为网络I/O提供了一个便携式接口
	listen,err:=net.Listen("tcp",":9500")
	if err != nil{
		fmt.Println("[GPS Server] 启动GPS服务器监听失败")
		return
	}

	for{
		//2.接收客户向服务端建立连接
		//可以与客户端建立连接.如果没有连接挂起阻塞状态
		conn,err:=listen.Accept()
		if err != nil{
			fmt.Println("[GPS Server] 新的终端链接失败")
			return
		}
		//3.处理用户的连接信息
		go gpsHander(conn)
	}
}

func gpsHander(conn net.Conn){
	//关闭链接
	defer conn.Close()
	// fmt.Println("GPS终端已链接",conn.RemoteAddr().String())
	for{
		golang808.Handle(conn)
	}
}