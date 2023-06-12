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

func (this *UserProcess)Login(account string, password string) (err error) {
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
	err = packio.SendPack(conn, data)
	if err != nil {
		err = errors.New("发送loginMessage包出错")
		return
	}
	rspMsg, err := packio.RecvPack(conn)
	if err != nil {
		err = errors.New("接收loginRspMessage出错")
		return
	}
	var loginRspMsg message.LoginRspMesssage
	err = json.Unmarshal([]byte(rspMsg.Content), &loginRspMsg)
	if err != nil {
		err = errors.New("反序列化loginRspMessage出错")
		return
	}
	if loginRspMsg.Status == "OK" {
		fmt.Println("登陆成功，欢迎您admin")
	} else {
		fmt.Println("登陆失败，用户不存在")
	}

	return nil
}