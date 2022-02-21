package process

import (
	"fmt"
	"net"
	"util"
)

type SmsProcess struct {}

func (sp *SmsProcess) SendGroupMes(mes *message.Message){
	//遍历 服务器 端的 onlineUsers map[int] *UserProcessor
	// 将消息转发取出

	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("json. Unmashal err = ", err)
		return
	}

	data, err := json.Marshal(mes)
	if err != nil {
		fmt.Println("json Marshal err = ", err)
		return
	}

	for id, up := range UserMgr.onlineUsers{
		if id == smsMes.UserId {
			continue
		}

		this.SendMesToEachOnlineUser(data, up.Conn)
	}
}

func (sp *SmsProcess)SendMesToEachOnlineUser(data []byte, conn net.Conn){
	tf := &util.Transfer{
		Conn: conn,
	}
	err := tf.WritePkg(data)
	if err != nil {
		fmt.Println("转发失败 err = ", err)
	}
}
