package process

import (
	"fmt"
	"message"
	"util"
	"json/encoding"
)

type SmsProcess struct{}

func (sp *SmsProcess) SendGroupMes(content string) (err error) {
	// 1 创建一个Mes
	var mes message.Message
	mes.Type = message.SmsMesType

	//2 创建一个SmsMes实例
	var smsMes message.SmsMes
	smsMes.Content = content
	smsMes.UserId = CurUser.UserId
	smsMes.UserStatus = CurUser.userStatus

	// 3 序列化 smsmes
	data, err := json.Marshal(smsMes)
	if err != nil {
		fmt.Println("SendGroupMes err = ", err)
		return
	}
	mes.Data = string(data)

	// 4 对mes再次序列化
	data, err = json.Marshal(mes)
	if err != nil {
		fmt.Println("SendGroupMes err = ", err)
		return
	}

	//5 将mes发送给服务器
	tf := &util.Transfer{
		Conn: CurUser.Conn,
	}

	// 6 发送
	err = tf.WritePkg(data)
	if err != nil {
		fmt.Println("WritePkg err = ", err)
	}
	return

}