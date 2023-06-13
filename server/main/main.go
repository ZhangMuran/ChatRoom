package main

import (
	"fmt"
	"net"
	"time"
)

func init() {
	initRedisPool("localhost:6379", 16, 0, 300 * time.Second)
	initUserDao(pool)
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
			fmt.Println("与客户端连接出错! err =", err)
			continue
		}
		connect := &processor{
			conn: conn,
		}
		go connect.clientConn()
	}
}
