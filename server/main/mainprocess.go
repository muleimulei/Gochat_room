package main

import (
	"fmt"
	"message"
	"net"
	"process"
	"util"
)

type Mainprocess struct {
	Conn net.Conn
}

// 编写一个serverProcessMes函数
//功能：根据客户端发送消息种类不同，决定调用哪个函数来处理
func (this *Mainprocess) serverProcessMes(mes *message.Message) (err error) {
	switch mes.Type {
	case message.LoginMesType:
		{
			//处理登录
			up := &process.UserProcessor{
				Conn: this.Conn,
			}
			err = up.ServerProcessLogin(mes)
		}
	case message.RegisterMesType:
		{
			//处理登录
			up := &process.UserProcessor{
				Conn: this.Conn,
			}
			err = up.ServerProcessRegister(mes)
		}
	}
	return
}

func (this *Mainprocess) processmain() (err error) {
	for {
		//这里读取数据包
		tf := &util.Transfer{
			Conn: this.Conn,
		}

		mes, err := tf.Readpkg()
		if err != nil {
			// fmt.Println("readpkg error", err)
			fmt.Println(this.Conn.RemoteAddr().String(), "退出")
			return err
		} else {
			err = this.serverProcessMes(&mes)
			if err != nil {
				fmt.Println("serverProcessMes error", err)
				return err
			}
		}
	}
}
