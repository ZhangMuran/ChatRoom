package process

import (
	"chatroom/common/message"
	"net"
)

type SmsProcess struct {
	data string
	Conn net.Conn
}

func (s *SmsProcess) SendSms() (err error) {
	smsMessage := message.SmsMessage{
		Content: s.data,
	}
	for _, x := range MyInfo.Friend {
		smsMessage.Users = append(smsMessage.Users, x)
	}

	msg, err := smsMessage.Bind()
	if err != nil {
		return
	}
	msg.Send(MyInfo.Conn)

	return
}

func (s *SmsProcess) ReceiveSms() (err error) {
	return
}