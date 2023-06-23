package message

import (
	"encoding/json"
	"net"
)

const (
	LoginSuccess    = 0
	LoginFail       = 1
	RegisterSuccess = 10
	RegisterFail    = 11
)

const (
	LoginType       = "LoginMessage"
	LoginRspType    = "LoginRspMesssage"
	RegisterType    = "RegisterMessage"
	RegisterRspType = "RegisterRspMessage"
	NotifyLoginType = "NotifyLoginMessage"
	SmsSendType     = "SmsMessage"
)

type Message struct {
	Type string    `json:"type"`
	Content string `json:"content"`
}

func (m *Message)Send(conn net.Conn) (err error) {
	msgSlic, err := json.Marshal(*m)
	if err != nil {
		return err
	}

	pio := PackIo{
		Conn: conn,
	}
	err =pio.SendPack(msgSlic)
	return
}

// func (m *Message)Receive(conn net.Conn) (err error) {

// }

type LoginRspMessage struct {
	Code int             `json:"code"`
	Status string        `json:"status"`
	OnlineUsers []string `json:"online_user"`
}

type RegisterMessage struct {
	User User `json:"user"`
}

type RegisterRspMessage struct {
	Code int      `json:"code"`
	Status string `json:"status"`
}

type NotifyOnlineMessage struct {
	User User `json:"user"`
}

// 将该消息绑定到Message类型中并返回Message
func (n *NotifyOnlineMessage)Bind() (msg *Message, err error) {
	msg = &Message{}
	x, err := json.Marshal(*n)
	if err != nil {
		return
	}
	msg.Type = NotifyLoginType
	msg.Content = string(x)
	return 
}
