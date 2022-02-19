package process

import (
	"fmt"
	"os"
	"util"
	"message"
)

//显示登录成功后的界面
func ShowMenu() {
	fmt.Println("----------恭喜xxx登录成功--")
	fmt.Println("---1 显示在线用户列表---")
	fmt.Println("---2 发送消息-----------")
	fmt.Println("---3 信息列表--")
	fmt.Println("---4 退出系统--")
	fmt.Println("-请选择(1 - 4)-")

	var key string
	fmt.Scanln(&key)

	switch key {
	case "1":
		{
			fmt.Println("显示在线用户列表")
			outputOnlineUser()
		}
	case "2":
		{
			fmt.Println("发送消息")
		}
	case "3":
		{
			fmt.Println("信息列表")
		}
	case "4":
		{
			fmt.Println("你选择了退出了系统...")
			os.Exit(0)
		}
	default:
		fmt.Println("重新输入")
	}

}

func serverProcessMes(conn net.Conn) {
	//创建一个transfer实例，不停的读取服务器发送的消息
	tf := &util.Transfer{
		Conn: conn,
	}
	for {
		fmt.Println("客户端正在等待读取服务器发送的消息")
		mes, err := tf.Readpkg()
		if err != nil {
			fmt.Println("tf.ReadPkg err = ", err)
			return
		}
		//如果读取到消息，又是下一步处理逻辑
		fmt.Println("mes = %v\n", mes)

		switch mes.Type {
		case message.NotifyUserStatusMesType:
			{
				var notifyUserStatusMes message.NotifyUserStatusMes
				err := json.Unmarshal([]byte(mes.Data), &notifyUserStatusMes)
				if err != nil {
					fmt.Println("unmarshal notifyUserStatusMes err = ", err)
					return
				}

				updateUserStatus(notifyUserStatusMes)
			}
		default:
			{
				fmt.Println("未知类型...")
			}
		}
	}
}