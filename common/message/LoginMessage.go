package message

import (
	"encoding/json"
	"fmt"
)

type LoginMessage struct {
	Account  string `json:"account"`
	Password string `json:"password"`
}

func (l *LoginMessage) Bind() (msg *Message, err error) {
	msg = &Message{}
	data, err := json.Marshal(*l)
	if err != nil {
		fmt.Println("序列化注册信息失败")
		return
	}
	msg.Type = LoginType
	msg.Content = string(data)
	return
}