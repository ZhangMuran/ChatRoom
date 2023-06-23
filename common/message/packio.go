package message

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
)

type PackIo struct {
	Conn net.Conn
	buf  [2048]byte
}

func (p *PackIo) SendPack(data []byte) (err error) {
	
	// 为了避免粘包，在发送数据时先发送数据的长度
	msgLen := uint32(len(data))
	var byteslice [4]byte
	binary.BigEndian.PutUint32(byteslice[:], msgLen)

	_, err = p.Conn.Write(byteslice[:])
	if err != nil {
		err = errors.New("发送报文长度时出错")
		return
	}
	_, err = p.Conn.Write(data)
	if err != nil {
		err = errors.New("发送报文内容时出错")
		return
	}
	return
}

func (p *PackIo) RecvPack() (msg Message, err error) {
	_, err = p.Conn.Read(p.buf[:4])
	if err != nil {
		if err != io.EOF {
			fmt.Println("客户端断开连接没有正确收到io.EOF, err =", err)
			err = io.EOF
		}
		fmt.Println("读取报文长度出错")
		return
	}
	numLen := binary.BigEndian.Uint32(p.buf[:4])

	_, err = p.Conn.Read(p.buf[:numLen])
	if err != nil {
		return
	}
	err = json.Unmarshal(p.buf[:numLen], &msg)
	if err != nil {
		err = errors.New("接受报文反序列化过程出错")
		return
	}
	return
}