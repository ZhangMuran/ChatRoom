package packio

import (
	"chatroom/common/message"
	"encoding/binary"
	"encoding/json"
	"errors"
	"io"
	"net"
)

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

func (this *PackIo)RecvPack() (msg message.Message, err error){
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