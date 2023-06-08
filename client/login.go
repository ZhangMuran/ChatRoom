package main

import (
	"chatroom/common/message"
	"chatroom/common/packio"
	"encoding/json"
	"fmt"
	"net"
	"time"
)

func login(account string, password string) error {

	conn, err := net.Dial("tcp", "127.0.0.1:10000")
	if err != nil {
		fmt.Println("dial err, err =", err)
		return err
	}
	defer conn.Close()

	var msg message.Message
	msg.Type = message.LoginType

	userInfo := message.LoginMessage{account, password}
	data, err := json.Marshal(userInfo)
	if err != nil {
		fmt.Println("json Mashal userInfo err, err =", err)
		return err
	}
	msg.Content = string(data)

	data, err = json.Marshal(msg)
	if err != nil {
		fmt.Println("json Mashal msg err, err =", err)
		return err
	}

	packio.SendPack(conn, data)

	time.Sleep(time.Second*10)
	return nil
}