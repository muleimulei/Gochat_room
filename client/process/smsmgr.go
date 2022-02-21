package process

import (
	"fmt"
	"encoding/json"
	"message"
)

func OutPutGroupMes (mes *message.Message){
	
	var smsMes message.SmsMes
	err := json.Unmarshal([]byte(mes.Data), &smsMes)
	if err != nil {
		fmt.Println("Unmarshal err = ", err)
		return
	}

	//显示信息
	info := fmt.Sprintf("用户id: %s 对大家说: %s\n", smsMes.UserId, smsMes.Content)

	fmt.Println(info)
}
