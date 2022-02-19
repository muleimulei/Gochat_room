package main

import (
	"fmt"
	"process"
)

var userId string
var usrPwd string
var usrName string

func main() {
	//接收用户选择
	var key int
	//判断是否继续显示菜单
	var loop = true

	for loop {
		fmt.Println("-------------欢迎登录多人聊天系统-------------")
		fmt.Println("             1 登录聊天室")
		fmt.Println("             2 注册用户")
		fmt.Println("             3 退出系统")
		fmt.Println("             4 请选择(1-3)")

		fmt.Scanln(&key)
		switch key {
		case 1:

			fmt.Println("登录聊天室")
			//用户要登录
			fmt.Println("请输入用户的id")
			fmt.Scanln(&userId)
			fmt.Println("请输入用户的密码")
			fmt.Scanln(&usrPwd)

			up := process.UserProcess{}
			err := up.Login(userId, usrPwd)
			if err != nil {
				fmt.Println("登录失败")
			}
			loop = false
		case 2:
			fmt.Println("注册用户")
			fmt.Println("请输入用户id：")
			fmt.Scanln(&userId)
			fmt.Println("请输入用户密码：")
			fmt.Scanln(&usrPwd)
			fmt.Println("请输入用户名称：")
			fmt.Scanln(&usrName)

			up := process.UserProcess{}
			err := up.Register(userId, usrPwd, usrName)
			if err != nil {
				fmt.Println("注册失败")
			}
			loop = false
		case 3:
			fmt.Println("退出系统")
			loop = false
		default:
			fmt.Println("你的输入有误,请重新输入")
		}
	}

}
