package main

import (
	"chatroom/common/message"
	"chatroom/common/packio"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
)

func processLogin(conn net.Conn, msg *message.Message) (err error) {
	var loginMsg message.LoginMessage
	err = json.Unmarshal([]byte(msg.Content), &loginMsg)
	if err != nil {
		err = errors.New("反序列化LoginMessage失败")
		return
	}

	var loginRsp message.LoginRspMesssage
	if loginMsg.Account == "admin" && loginMsg.Password == "password" {
		loginRsp.Code = message.LoginSuccess
		loginRsp.Status = "OK"
	} else {
		loginRsp.Code = message.LoginUserNotExist
		loginRsp.Status = "ERROR"
	}

	loginRspslice, err := json.Marshal(loginRsp)
	if err != nil {
		err =  errors.New("序列化LoginRspMessage失败")
		return
	}
	sendMsg := message.Message{Type: message.LoginRspType, Content: string(loginRspslice)}
	sendMsgSlice, err := json.Marshal(sendMsg)
	err = packio.SendPack(conn, sendMsgSlice)
	if err != nil {
		err = errors.New("发送LoginRspMessage失败")
		return
	}

	return
}

func processMsg(conn net.Conn, msg *message.Message) (err error) {
	switch msg.Type {
		case message.LoginType :
			err = processLogin(conn, msg)
		default :
			err = errors.New("不存在的msg类型")	
	}
	return
}

func clientConn(conn net.Conn) {
	defer conn.Close()

	for {
		msg, err := packio.RecvPack(conn)
		if err != nil {
			if err == io.EOF {
				return
			}
			fmt.Println("接收报文过程出错！err =", err)
		}
		err = processMsg(conn, &msg)
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:10000")
	if err != nil {
		fmt.Println("listen error, err =", err)
		return
	}
	defer listen.Close()

	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept err, err =", err)
			return
		}
		go clientConn(conn)
	}
}
