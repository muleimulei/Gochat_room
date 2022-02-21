package process

import (
	"encoding/json"
	"fmt"
	"message"
	"net"
	"util"
)

type UserProcess struct {
}

func (u *UserProcess) Register(userId string, usrPwd string, userName string) (err error){
	//1. 连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return err
	}

	//2. 准备通过conn发送消息
	var mes message.Message
	mes.Type = message.RegisterMesType

	// 3. 创建一个LoginMes结构体
	var registerMes message.RegisterMes
	registerMes.User.UserId = userId
	registerMes.User.UserPwd = usrPwd
	registerMes.User.UserName = userName

	data, err := json.Marshal(registerMes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}
	//5. 把data赋值给mes.data字段
	mes.Data = string(data)

	//6. 将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}

	tf := &util.Transfer{
		Conn: conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("发送注册消息失败")
		return
	}

	//查看服务器返回的消息
	mes, err = tf.Readpkg()

	if err != nil {
		fmt.Println("readPkg error", err)
		return
	}

	//将mes的Data部分反序列化
	var registerResMes message.RegisterResMes
	
	err = json.Unmarshal([]byte(mes.Data), &registerResMes)
	
	if registerResMes.Code == 200 {
		fmt.Println("注册成功，请登录")
	} else {
		fmt.Println(registerResMes.Error)
	}
	return nil
}



func (u *UserProcess) Login(userId string, usrPwd string) error {
	//1. 连接到服务器
	conn, err := net.Dial("tcp", "localhost:8889")
	if err != nil {
		fmt.Println("net.Dial err=", err)
		return err
	}

	//2. 准备通过conn发送消息
	var mes message.Message
	mes.Type = message.LoginMesType

	// 3. 创建一个LoginMes结构体
	var loginMes message.LoginMes
	loginMes.UserId = userId
	loginMes.UserPwd = usrPwd

	// 4. 将loginMes序列化
	data, err := json.Marshal(loginMes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return err
	}

	//5. 把data赋值给mes.data字段
	mes.Data = string(data)

	//6. 将mes进行序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return err
	}

	tf := &util.Transfer{
		Conn: conn,
	}

	err = tf.WritePkg(data)
	if err != nil {
		return err
	}

	//查看服务器返回的消息

	mes, err = tf.Readpkg()

	if err != nil {
		fmt.Println("readPkg error", err)
		return err
	}

	//将mes的Data部分反序列化
	var loginResMes message.LoginResMes

	err = json.Unmarshal([]byte(mes.Data), &loginResMes)

	if loginResMes.Code == 200 {

		//初始化 CurUser
		CurUser.Conn = conn
		CurUser.UserId = userId
		CurUser.UserStatus = message.UserOnline

		//显示当前在线用户列表，遍历loginResMes.UserId
		fmt.Println("当前在线用户列表如下:")
		for _, v := range loginResMes.UserIds {
			fmt.Println("用户 id: ", v)
			if v == userId {
				continue
			}

			onlineUsers[v] = &message.User{
				UserId: v,
				UserStatus: message.UserOnline
			}
		}

		go serverProcessMes(conn)
		for {
			ShowMenu()
		}
	} else {
		fmt.Println(loginResMes.Error)
	}

	return err
}
