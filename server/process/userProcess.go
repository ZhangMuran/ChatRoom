package process

import (
	"chatroom/common/message"
	"chatroom/common/packio"
	"chatroom/server/model"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type UserProcess struct {
	Conn net.Conn
}

func (u *UserProcess)ProcessLogin(msg *message.Message) (err error) {
	var loginMsg message.LoginMessage
	err = json.Unmarshal([]byte(msg.Content), &loginMsg)
	if err != nil {
		err = errors.New("反序列化LoginMessage失败")
		return
	}

	var loginRsp message.LoginRspMessage
	loginRsp.Code = message.LoginSuccess

	dao := model.GetUserDao()
	err = dao.CheckLogin(loginMsg.Account, loginMsg.Password)

	if err == nil {
		loginRsp.Status = "OK"
	} else {
		loginRsp.Code = message.LoginFail
		loginRsp.Status = err.Error()
	}

	loginRspslice, err := json.Marshal(loginRsp)
	if err != nil {
		err =  errors.New("序列化LoginRspMessage失败")
		return
	}

	sendMsg := message.Message{Type: message.LoginRspType, Content: string(loginRspslice)}
	sendMsgSlice, err := json.Marshal(sendMsg)
	if err != nil {
		fmt.Println("login类型的message包序列化错误， err =", err)
		return
	}

	pio := &packio.PackIo{
		Conn: u.Conn,
	}

	err = pio.SendPack(sendMsgSlice)
	if err != nil {
		fmt.Println("发送LoginRspMessage失败，err =", err)
		return
	}

	return
}

func (u *UserProcess)ProcessRegister(msg *message.Message) (err error) {
	var registerMsg message.RegisterMessage
	err = json.Unmarshal([]byte(msg.Content), &registerMsg)
	if err != nil {
		err = errors.New("反序列化RegisterMessage失败")
		return
	}
	
	var registerRsp message.RegisterRspMessage
	registerRsp.Code = message.RegisterSuccess

	dao := model.GetUserDao()
	err = dao.UserRegister(registerMsg.User)

	if err == nil {
		registerRsp.Status = "OK"
	} else {
		registerRsp.Code = message.RegisterFail
		registerRsp.Status = err.Error()
	}

	registerRspslice, err := json.Marshal(registerRsp)
	if err != nil {
		fmt.Println("序列化registerRspMessage失败，err =", err)
		return
	}

	sendMsg := message.Message{Type: message.RegisterRspType, Content: string(registerRspslice)}
	sendMsgSlice, err := json.Marshal(sendMsg)
	if err != nil {
		fmt.Println("register类型的message包序列化错误， err =", err)
		return
	}

	pio := &packio.PackIo{
		Conn: u.Conn,
	}

	err = pio.SendPack(sendMsgSlice)
	if err != nil {
		fmt.Println("发送registerRspMessage失败，err =", err)
		return
	}

	return
}