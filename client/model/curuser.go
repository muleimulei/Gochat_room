package model

import (
	"fmt"
	"message"
	"net"
)

type CurUsr struct {
	Conn net.Conn
	message.User
}