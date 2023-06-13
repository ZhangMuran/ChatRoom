package model

import (
	"chatroom/common/message"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gomodule/redigo/redis"
)

type userDao struct {
	Pool *redis.Pool
}

var userDaoInstance *userDao

func GetUserDao() *userDao {
	if userDaoInstance == nil {
		userDaoInstance = &userDao{}
	}
	return userDaoInstance
}

func (this *userDao)getUserByAccount(conn redis.Conn, account string) (userInfo message.User, err error) {
	uInfo, err := redis.String(conn.Do("hget", "users", account))
	if err != nil {
		if err == redis.ErrNil {
			err = errors.New("该用户名不存在，请检查后重试")
		}
		return
	}
	err = json.Unmarshal([]byte(uInfo), &userInfo)
	if err != nil {
		return
	}
	return
}

func (this *userDao)CheckLogin(account string, password string) (err error) {
	conn := this.Pool.Get()
	defer conn.Close()

	userInfo, err := this.getUserByAccount(conn, account)
	if err != nil {
		return
	}

	if userInfo.Password == password {
		fmt.Println("登陆成功")
	} else {
		err = errors.New("密码错误，请重新输入")
	}

	return
}