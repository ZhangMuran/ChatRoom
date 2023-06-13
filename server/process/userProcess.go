package process

import (
	"chatroom/common/message"
	"chatroom/common/packio"
	"chatroom/server/model"
	"encoding/json"
	"errors"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

func (this *UserProcess)ProcessLogin(msg *message.Message) (err error) {
	var loginMsg message.LoginMessage
	err = json.Unmarshal([]byte(msg.Content), &loginMsg)
	if err != nil {
		err = errors.New("反序列化LoginMessage失败")
		return
	}

	var loginRsp message.LoginRspMesssage
	loginRsp.Code = message.LoginSuccess

	dao := model.GetUserDao()
	err = dao.CheckLogin(loginMsg.Account, loginMsg.Password)

	if err == nil {
		loginRsp.Status = "OK"
	} else {
		loginRsp.Status = err.Error()
	}

	loginRspslice, err := json.Marshal(loginRsp)
	if err != nil {
		err =  errors.New("序列化LoginRspMessage失败")
		return
	}
	sendMsg := message.Message{Type: message.LoginRspType, Content: string(loginRspslice)}
	sendMsgSlice, err := json.Marshal(sendMsg)

	pio := &packio.PackIo{
		Conn: this.Conn,
	}

	err = pio.SendPack(sendMsgSlice)
	if err != nil {
		err = errors.New("发送LoginRspMessage失败")
		return
	}

	return
}