package main

import (
	"fmt"
	"model"
	"net"
	"time"
)

func process_conn(conn net.Conn) {
	//读取客户端发送的信息
	defer conn.Close()
	processor := Mainprocess{
		Conn: conn,
	}
	err := processor.processmain()
	if err != nil {
		fmt.Println("客户端与服务器通讯协程错误", err)
	}
}

func initUserDao() {
	//这里的pool 本身就是一个全局变量
	model.MyUserDao = model.NewUserDao(Pool)
}

func main() {
	//初始化连接池
	initPool("8.142.31.201:6379", "myRedis", 16, 0, 300*time.Second)
	initUserDao()

	//提示信息
	fmt.Println("服务器在8889端口监听")
	listen, err := net.Listen("tcp", "localhost:8889")
	if err != nil {
		fmt.Printf("监听失败 %v\n", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept error", err)
		} else {
			//一旦连接成功，则启动一个协程与客户通讯
			go process_conn(conn)
		}
	}
}
