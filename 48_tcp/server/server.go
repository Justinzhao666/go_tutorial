package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("server begin to listen...")
	// 开启监听 ：参数1 表示使用tcp协议，参数2 表示监听的ip和端口
	listener, e := net.Listen("tcp", "0.0.0.0:8888") //0.0.0.0 ipv4和v6都支持，127那个只支持ipv4
	if e != nil {
		fmt.Println("listen error:", e)
		return
	}
	defer listener.Close()
	for {
		//等待客户端连接,获取连接
		fmt.Println("wait to accept...")
		conn, e := listener.Accept() // 这里会阻塞等待着,可以使用telnet进行测试
		if e != nil {
			fmt.Println("fail to get conn")
		}
		fmt.Println(conn)

	}

	fmt.Println(listener)

}
