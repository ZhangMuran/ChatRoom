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

func processLogin(msg *message.Message) (err error) {
	var loginMsg message.LoginMessage
	err = json.Unmarshal([]byte(msg.Content), &loginMsg)
	if err != nil {
		err = errors.New("反序列化LoginMessage失败")
	}

	// var loginRsp message.LoginRspMesssage
	if loginMsg.Account == "admin" && loginMsg.Password == "password" {


	} else {
		fmt.Println("登陆失败！")
	}
	return
}

func processMsg(msg *message.Message) (err error) {
	switch msg.Type {
		case message.LoginRspType :
			err = processLogin(msg)
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
		// fmt.Println(msg)
		err = processMsg(&msg)
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
