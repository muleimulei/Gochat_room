package process

import (
	"fmt"
	"message"
	"model"
)

//客户端维护的map

var onlineUsers map[int] *message.User = make(map[int] *message.User, 10)

var CurUser model.CurUsr //我们在用户登录成功后，完成初始化

//在客户端显示当前在线的用户
func outputOnlineUser() {
	fmt.Println("当前在线用户列表")
	for id, _ := range onlineUsers {
		fmt.Println("用户id: ", id )
	}
}


// 编写一个方法， 处理返回的NotifyUserStatusMes
func updateUserStatus(notifyUserStatusMes *message.NotifyUserStatusMes) {
	user, ok := onlineUsers[notifyUserStatusMes.UserId]

	if ok == false {
		onlineUsers[notifyUserStatusMes.UserId] = &message.User{
			UserId: notifyUserStatusMes.UserId
			UserStatus: notifyUserStatusMes.UserStatus
		}
		return
	}
	
	user.Status = notifyUserStatusMes.UserStatus

	outputOnlineUser()
}

