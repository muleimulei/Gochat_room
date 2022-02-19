package process

import (
	"encoding/json"
	"fmt"
	"message"
	"model"
	"net"
	"util"
)

type UserProcessor struct {
	Conn net.Conn
	UserId string
}

// 这里我们编写通知所有在线的用户的方法
// userId 要通知其它在线用户，我上线

func (up *UserProcessor) NotifyOtherOnlineUser(userId string){

	// 遍历 onlineUsers ， 然后一个一个发送NotifyUserStatusMes
	for id, up := range userMgr.onlineUsers {
		if userId == id {
			continue
		}
		// 开始通知
		up.NotifyMeOnline(userId)
	}
}

func (up *UserProcessor) NotifyMeOnline(userId int) {

	//组装我们的NotifyUserStatusMes
	var mes message.Message
	mes.Type = message.NotifyUserStatusMesType

	var notifyUserStatusMes message.NotifyUserStatusMes
	notifyUserStatusMes.UserId = userId
	notifyUserStatusMes.Status = message.UserOnline

	//将notifyUserStatusMes序列化
	data, err := json.Marshal(notifyUserStatusMes)
	if err != nil {
		fmt.Println("json.Marshal err = ", err)
		return
	}

	//将序列化后的notifyUserStatusMes赋值给mes.Data
	mes.Data = string(data)

	//对mes序列化
	data, err = json.Marshal(mes)

	//6. 发送数据
	tf := &util.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("NotifyMeOnline err", err)
	}
}


func (this *UserProcessor) ServerProcessLogin(mes *message.Message) (err error) {
	// 1.先从mes中取出 mes.Data ，并直接反序列化成LoginMes
	var loginMes message.LoginMes
	err = json.Unmarshal([]byte(mes.Data), &loginMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err ", err)
		return
	}

	//声明一个 resMes
	var resMes message.Message
	resMes.Type = message.LoginResMesType

	// //声明一个 LoginResMes
	var loginResMes message.LoginResMes

	user, err := model.MyUserDao.Login(loginMes.UserId, loginMes.UserPwd)
	if err != nil {
		if err == model.ERROR_USER_NOTEXIST {
			loginResMes.Code = 500
			loginResMes.Error = err.Error()
		} else if err == model.ERROR_USER_PWD {
			loginResMes.Code = 403
			loginResMes.Error = err.Error()
		} else {
			loginResMes.Code = 505
			loginResMes.Error = "服务器内部错误..."
		}
	} else {
		loginResMes.Code = 200

		//用户登录成功，我们就把该登录成功的用户放入usermgr中
		//将登录成功的用户id赋值给this
		this.UserId = loginMes.UserId
		userMgr.AddOnlineUser(this)

		//通知其它在线用户，我上线了
		this.NotifyOtherOnlineUser(loginMes.UserId)

		//将当前在线用户id 放入到loginResMes.UserIds中
		for id, _ := range userMgr.onlineUsers {
			loginResMes.UserIds = append(loginResMes.UserIds, id)
		}

		fmt.Println(user, "登录成功")
	}

	//将 loginResMes系列化
	data, err := json.Marshal(loginResMes)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}

	// 将data赋值给resMes
	resMes.Data = string(data)

	// 5. 对resMes进行序列化，准备发送
	data, err = json.Marshal(&resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	//6. 发送数据
	tf := &util.Transfer{
		Conn: this.Conn,
	}
	err = tf.WritePkg(data)
	return
}


func (up *UserProcessor) ServerProcessRegister(mes *message.Message)(err error){
	// 1.先从mes中取出 mes.Data ，并直接反序列化成LoginMes
	var registerMes message.RegisterMes
	err = json.Unmarshal([]byte(mes.Data), &registerMes)
	if err != nil {
		fmt.Println("json.Unmarshal fail err ", err)
		return
	}

	//声明一个 resMes
	var resMes message.Message
	resMes.Type = message.RegisterResMesType
	// //声明一个 registerResMes
	var registerResMes message.RegisterResMes

	//我们需要到redis数据库完成注册
	//1. 使用model.MyUserDao到 redis去验证
	err = model.MyUserDao.Register(&registerMes.User)
	if err != nil {
		if err == model.ERROR_USER_EXIST {
			registerResMes.Code = 505
			registerResMes.Error = err.Error()
		} else {
			registerResMes.Code = 506
			registerResMes.Error = "注册未知错误"
		}
	}else{
		registerResMes.Code = 200
		registerResMes.Error = "注册成功"
	}

	//将 registerResMes系列化
	data, err := json.Marshal(registerResMes)
	if err != nil {
		fmt.Println("序列化失败", err)
		return
	}

	// 将data赋值给resMes
	resMes.Data = string(data)

	// 5. 对resMes进行序列化，准备发送
	data, err = json.Marshal(&resMes)
	if err != nil {
		fmt.Println("json.Marshal fail", err)
		return
	}

	//6. 发送数据
	tf := &util.Transfer{
		Conn: up.Conn,
	}
	err = tf.WritePkg(data)
	return

}