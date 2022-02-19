package util

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"message"
	"net"
)

type Transfer struct {
	Conn net.Conn
	Buf  [4096]byte
}

func (this *Transfer) Readpkg() (mes message.Message, err error) {

	n, err := this.Conn.Read(this.Buf[:4])
	if n != 4 || err != nil {
		// fmt.Println("conn.Read err = ", err)
		return
	}
	//根据buf[:4]转成uint32
	pkglen := binary.BigEndian.Uint32(this.Buf[:4])

	//根据pkglen读取消息内容
	n, err = this.Conn.Read(this.Buf[:pkglen])
	if n != int(pkglen) || err != nil {
		// fmt.Println("conn.Read fail err = ", err)
		return
	}

	//把pkglen 反序列化 -> message.Message
	err = json.Unmarshal(this.Buf[:pkglen], &mes)
	if err != nil {
		fmt.Println("json unmatshal err", err)
		return
	}
	return
}

func (this *Transfer) WritePkg(data []byte) (err error) {
	// 先把data的长度发送给服务器
	pklen := uint32(len(data))
	var buf [4]byte
	binary.BigEndian.PutUint32(buf[:4], pklen)

	n, err := this.Conn.Write(buf[:4])
	if n != 4 || err != nil {
		fmt.Println("conn.Write(bytes) fail", err)
		return err
	}
	//fmt.Printf("客户端发送 %d %s\n", len(data), string(data))

	_, err = this.Conn.Write(data)
	if err != nil {
		return
	}
	return
}
