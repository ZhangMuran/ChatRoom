package packio

import (
	"chatroom/common/message"
	"encoding/json"
	"errors"
	"io"
	"net"
	"strconv"
)

func SendPack(conn net.Conn, data []byte) (err error) {
	// 为了避免丢包，在发送数据时先发送数据的长度
	msgLen := len(data)
	lenStr := strconv.Itoa(msgLen)
	// 再将数据发送过去
	data = append([]byte(lenStr), data...)
	_, err = conn.Write(data)
	if err != nil {
		err = errors.New("发送报文内容时出错")
		return
	}
	return
}

func RecvPack(conn net.Conn) (msg message.Message, err error){
	var buf [1024]byte
	lenBuf, err := conn.Read(buf[:])
	if err != nil {
		if err == io.EOF {
			return
		}
		err = errors.New("读取报文过程出错")
		return
	}
	var numLen int
	var msgStr []byte
	for i := 0; i < lenBuf; i++ {
		if buf[i] >= '0' && buf[i] <= '9' {
			numLen = numLen * 10 + int(buf[i] - '0')
		} else {
			msgStr = buf[i:lenBuf]
			break
		}
	}
	if numLen != len(msgStr) {
		err = errors.New("发生了丢包")
		return
	}
	err = json.Unmarshal(msgStr[:], &msg)
	if err != nil {
		err = errors.New("接受报文反序列化过程出错")
		return
	}
	return
}