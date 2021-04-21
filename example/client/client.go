package main

import (
	"fmt"
	"net"
)

//启动gps服务
func main(){
	//1.创建建立连接
	conn,_:=net.Dail("tcp",":9500")
	defer conn.Close()
	fmt.Println(111)
}