package process

import "fmt"

type onlineUserManage struct{
	onlineMap map[string]*UserProcess
}

var onlineUser *onlineUserManage

func init() {
	onlineUser = &onlineUserManage{
		onlineMap: make(map[string]*UserProcess),
	}
}

func (o *onlineUserManage)addOnlineUser(up *UserProcess) {
	o.onlineMap[up.account] = up
}

func (o *onlineUserManage)delOnlineUser(account string) {
	delete(o.onlineMap, account)
}

func (o *onlineUserManage)GetOnlineMap() map[string]*UserProcess {
	return o.onlineMap
}

func (o *onlineUserManage)getInfoByAccount(account string) (up *UserProcess, err error) {
	up, ok := o.onlineMap[account]
	if !ok {
		err = fmt.Errorf("用户 %v 目前离线或不存在", account)
		return
	}
	return
}