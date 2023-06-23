package message

import (
	"encoding/json"
	"fmt"
)

type SmsMessage struct {
	Content string
	Users   []string
}

func (s *SmsMessage) Bind() (msg *Message, err error) {
	msg = &Message{}
	data, err := json.Marshal(*s)
	if err != nil {
		fmt.Println("序列化注册信息失败")
		return
	}
	msg.Type = SmsSendType
	msg.Content = string(data)
	return
}