package process

import (
	"chatroom/common/message"
	"fmt"
	"net"
)

type UserMgr struct {
	Friend map[string]string
	message.User
	Conn  net.Conn
}

var MyInfo UserMgr

func init() {
	MyInfo.Friend = make(map[string]string)
}

func (u *UserMgr)InitFriend(onlineUsers *[]string) {
	for _, id := range *onlineUsers {
		u.Friend[id] = id
	}
}

func (u *UserMgr)ShowOnlineFriend() {
	if len(u.Friend) == 0 {
		fmt.Println("暂时还没有其他用户在线哦")
		return
	}
	fmt.Println("目前在线的用户数量：", len(u.Friend))
	for _, id := range u.Friend {
		fmt.Println(id)
	}
	fmt.Println("快去找他们聊聊吧！")
}