package process

import (
	"chatroom/common/message"
	"chatroom/server/model"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
)

type Processor struct {
	Conn    net.Conn
	Account string
}

func (p *Processor)ClientConn() {
	defer p.Conn.Close()
	
	pio := &message.PackIo{
		Conn: p.Conn,
	}
	for {
		msg, err := pio.RecvPack()
		if err != nil {
			if err == io.EOF {
				fmt.Println("客户端断开了连接")
				OnlineUser.DelOnlineUser(p.Account)
				fmt.Println("delete user", p.Account)
				fmt.Println(OnlineUser.onlineMap)
				return
			}
			fmt.Println("接收报文过程出错! err =", err)
			return
		}

		err = p.processMsg(&msg)
		if err != nil {
			fmt.Println("服务器处理报文信息出错! err =", err)
			return
		}
	}
}

func (p *Processor)processMsg(msg *message.Message) (err error) {
	switch msg.Type {
		case message.LoginType :
			err = p.ProcessLogin(msg.Content)
		case message.RegisterType :
			err = p.ProcessRegister(msg.Content)
		default :
			err = errors.New("不存在的msg类型")	
	}
	return
}

/*
 * 服务器处理用户的登陆操作，并根据处理结果回复消息给客户端
 * @params: json序列化成string格式的LoginMessage
 */
 func (p *Processor) ProcessLogin(msg string) (err error) {
	var loginMsg message.LoginMessage
	err = json.Unmarshal([]byte(msg), &loginMsg)
	if err != nil {
		err = errors.New("反序列化LoginMessage失败")
		return
	}

	var loginRsp message.LoginRspMessage
	loginRsp.Code = message.LoginSuccess

	dao := model.GetUserDao()
	err = dao.CheckLogin(loginMsg.Account, loginMsg.Password)

	// 成功登录，需要展示所有在线成员，并提醒所有人该用户上线
	if err == nil {
		loginRsp.Status = "OK"

		p.Account = loginMsg.Account
		notifymsg := message.NotifyOnlineMessage{
			User: message.User{
				Account: p.Account,
				Status:  message.OnLine,
			},
		}
		var nMsg *message.Message
		nMsg, err = notifymsg.Bind()
		if err != nil {
			return
		}

		for id, up := range OnlineUser.onlineMap {
			// 在loginRsp中增加在线成员的切片，使客户端可以看到在线列表
			loginRsp.OnlineUsers = append(loginRsp.OnlineUsers, id)
			// 向所有在线成员发送notify消息，通知新用户上线
			nMsg.Send(up.Conn)
		}
		fmt.Println("add user", p.Account)
		OnlineUser.addOnlineUser(p)
		fmt.Println(OnlineUser.onlineMap)
	} else {
		loginRsp.Code = message.LoginFail
		loginRsp.Status = err.Error()
	}

	loginRspslice, err := json.Marshal(loginRsp)
	if err != nil {
		err = errors.New("序列化LoginRspMessage失败")
		return
	}

	sendMsg := message.Message{Type: message.LoginRspType, Content: string(loginRspslice)}
	sendMsgSlice, err := json.Marshal(sendMsg)
	if err != nil {
		fmt.Println("login类型的message包序列化错误， err =", err)
		return
	}

	pio := &message.PackIo{
		Conn: p.Conn,
	}

	err = pio.SendPack(sendMsgSlice)
	if err != nil {
		fmt.Println("发送LoginRspMessage失败，err =", err)
		return
	}

	return
}

/*
 * 服务器处理用户的注册操作，并根据处理结果回复消息给客户端
 * @params: json序列化成string格式的RegisterMessage
 */
func (p *Processor) ProcessRegister(msg string) (err error) {
	var registerMsg message.RegisterMessage
	err = json.Unmarshal([]byte(msg), &registerMsg)
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

	pio := &message.PackIo{
		Conn: p.Conn,
	}

	err = pio.SendPack(sendMsgSlice)
	if err != nil {
		fmt.Println("发送registerRspMessage失败，err =", err)
		return
	}

	return
}
