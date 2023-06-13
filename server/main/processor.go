package main

import (
	"chatroom/common/message"
	"chatroom/common/packio"
	"chatroom/server/process"
	"errors"
	"fmt"
	"io"
	"net"
)

type processor struct {
	conn net.Conn
}

func (this *processor)clientConn() {
	defer this.conn.Close()
	
	pio := &packio.PackIo{
		Conn: this.conn,
	}
	for {
		msg, err := pio.RecvPack()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端断开了连接")
				return
			}
			fmt.Println("接收报文过程出错! err =", err)
			return
		}

		err = this.processMsg(&msg)
		if err != nil {
			fmt.Println("服务器处理报文信息出错! err =", err)
			return
		}
	}
}

func (this *processor)processMsg(msg *message.Message) (err error) {
	pro := &process.UserProcess{
		Conn: this.conn,
	}
	switch msg.Type {
		case message.LoginType :
			err = pro.ProcessLogin(msg)
		case message.RegisterType :
			err = pro.ProcessRegister(msg)
		default :
			err = errors.New("不存在的msg类型")	
	}
	return
}