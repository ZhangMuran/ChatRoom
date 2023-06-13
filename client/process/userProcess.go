package process

import (
	"chatroom/common/message"
	"chatroom/common/packio"
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type UserProcess struct {

}

func (u *UserProcess)Login(account string, password string) (err error) {
	conn, err := net.Dial("tcp", "127.0.0.1:10000")
	if err != nil {
		return
	}
	defer conn.Close()

	userInfo := message.LoginMessage{Account:account, Password:password}
	data, err := json.Marshal(userInfo)
	if err != nil {
		err = errors.New("序列化LoginMessage失败")
		return
	}
	var msg message.Message
	msg.Type = message.LoginType
	msg.Content = string(data)

	data, err = json.Marshal(msg)
	if err != nil {
		err = errors.New("发送登陆消息过程中，序列化message失败")
		return
	}

	pio := packio.PackIo{
		Conn: conn,
	}
	err = pio.SendPack(data)
	if err != nil {
		err = errors.New("发送loginMessage包出错")
		return
	}
	rspMsg, err := pio.RecvPack()
	if err != nil {
		err = errors.New("接收loginRspMessage出错")
		return
	}

	var loginRspMsg message.LoginRspMessage
	err = json.Unmarshal([]byte(rspMsg.Content), &loginRspMsg)
	if err != nil {
		err = errors.New("反序列化loginRspMessage出错")
		return
	}
	if loginRspMsg.Status == "OK" {
		fmt.Println("登陆成功，欢迎您", account)
	} else {
		fmt.Println(loginRspMsg.Status)
	}

	return nil
}

func (u *UserProcess)Register(account string, password string) (err error) {
	conn, err := net.Dial("tcp", "127.0.0.1:10000")
	if err != nil {
		return
	}
	defer conn.Close()
	userInfo := message.RegisterMessage{User: message.User{Account:account, Password:password}}

	data, err := json.Marshal(userInfo)
	if err != nil {
		err = errors.New("客户端序列化RegisterMessage失败")
		return
	}
	var msg message.Message
	msg.Type = message.RegisterType
	msg.Content = string(data)

	data, err = json.Marshal(msg)
	if err != nil {
		err = errors.New("发送注册消息过程中，序列化message失败")
		return
	}

	pio := packio.PackIo{
		Conn: conn,
	}
	err = pio.SendPack(data)
	if err != nil {
		err = errors.New("发送RegisterMessage包出错")
		return
	}
	rspMsg, err := pio.RecvPack()
	if err != nil {
		err = errors.New("接收loginRspMessage出错")
		return
	}

	var loginRspMsg message.RegisterRspMessage
	err = json.Unmarshal([]byte(rspMsg.Content), &loginRspMsg)
	if err != nil {
		err = errors.New("反序列化RegisterRspMessage出错")
		return
	}
	if loginRspMsg.Status == "OK" {
		fmt.Println("注册成功，您可以登陆了")
	} else {
		fmt.Println(loginRspMsg.Status)
	}

	return nil
}