package message

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
)

const (
	LoginSuccess    = 0
	LoginFail       = 1
	RegisterSuccess = 10
	RegisterFail    = 11
)

const (
	LoginType       = "LoginMessage"
	LoginRspType    = "LoginRspMesssage"
	RegisterType    = "RegisterMessage"
	RegisterRspType = "RegisterRspMessage"
	NotifyLoginType = "NotifyLoginMessage"
)

type Message struct {
	Type string    `json:"type"`
	Content string `json:"content"`
}

func (m *Message)Send(conn net.Conn) (err error) {
	msgSlic, err := json.Marshal(*m)
	if err != nil {
		return err
	}

	pio := PackIo{
		Conn: conn,
	}
	err =pio.SendPack(msgSlic)
	return
}

func (m *Message)Parse() (err error) {
	switch m.Type {
		case NotifyLoginType :
			err = m.processNotify()
	}
	return
}

func (m *Message)processNotify() (err error) {
	var notifyMsg NotifyOnlineMessage
	err = json.Unmarshal([]byte(m.Content), &notifyMsg)
	if err != nil {
		return
	}

	if(notifyMsg.User.Status == OnLine) {
		fmt.Println(notifyMsg.User.Account, "刚刚上线了，快去找他聊聊吧")
	}

	return
}

type LoginMessage struct {
	Account string  `json:"account"`
	Password string `json:"password"`
}

type LoginRspMessage struct {
	Code int             `json:"code"`
	Status string        `json:"status"`
	OnlineUsers []string `json:"online_user"`
}

type RegisterMessage struct {
	User User `json:"user"`
}

type RegisterRspMessage struct {
	Code int      `json:"code"`
	Status string `json:"status"`
}

type NotifyOnlineMessage struct {
	User User `json:"user"`
}

// 将该消息绑定到Message类型中并返回Message
func (n *NotifyOnlineMessage)Bind() (msg *Message, err error) {
	x, err := json.Marshal(*n)
	if err != nil {
		return
	}
	msg.Type = NotifyLoginType
	msg.Content = string(x)
	return 
}


type PackIo struct {
	Conn net.Conn
	buf [2048]byte
}

func (this *PackIo)SendPack(data []byte) (err error) {
	// 为了避免粘包，在发送数据时先发送数据的长度
	msgLen := uint32(len(data))
	var byteslice [4]byte 
	binary.BigEndian.PutUint32(byteslice[:], msgLen)

	_, err = this.Conn.Write(byteslice[:])
	if err != nil {
		err = errors.New("发送报文长度时出错")
		return
	}

	_, err = this.Conn.Write(data)
	if err != nil {
		err = errors.New("发送报文内容时出错")
		return
	}
	return
}

func (this *PackIo)RecvPack() (msg Message, err error){
	_, err = this.Conn.Read(this.buf[:4])
	if err != nil {
		if err == io.EOF {
			return
		}
		err = errors.New("读取报文长度出错")
		return
	}
	numLen := binary.BigEndian.Uint32(this.buf[:4])

	_, err = this.Conn.Read(this.buf[:numLen])
	err = json.Unmarshal(this.buf[:numLen], &msg)
	if err != nil {
		err = errors.New("接受报文反序列化过程出错")
		return
	}
	return
}