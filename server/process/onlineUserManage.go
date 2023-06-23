package process

import (
	"fmt"
)

type onlineUserManage struct{
	onlineMap map[string]*Processor
}

var OnlineUser *onlineUserManage

func init() {
	OnlineUser = &onlineUserManage{
		onlineMap: make(map[string]*Processor),
	}
}

func (o *onlineUserManage)addOnlineUser(up *Processor) {
	o.onlineMap[up.Account] = up
}

func (o *onlineUserManage)DelOnlineUser(account string) {
	delete(o.onlineMap, account)
}

func (o *onlineUserManage)GetOnlineMap() map[string]*Processor {
	return o.onlineMap
}

func (o *onlineUserManage)GetInfoByAccount(account string) (up *Processor, err error) {
	up, ok := o.onlineMap[account]
	if !ok {
		err = fmt.Errorf("用户 %v 目前离线或不存在", account)
		return
	}
	return
}