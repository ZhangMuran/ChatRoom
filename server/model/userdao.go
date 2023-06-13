package model

import (
	"chatroom/common/message"
	"chatroom/server/utils"
	"encoding/json"
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

func (u *userDao)isUserExist(conn redis.Conn, account string) (exist bool, userInfo string, err error) {
	userInfo, err = redis.String(conn.Do("hget", "users", account))
	if err != nil {
		if err == redis.ErrNil {
			err = nil
		} else {
			fmt.Println("redis寻找用户过程出错了，err =", err)
		}
		return
	}
	exist = true
	return
}

func (u *userDao)getUserByAccount(conn redis.Conn, account string) (userInfo message.User, err error) {
	exist, uInfo, err := u.isUserExist(conn, account)
	if err != nil {
		return
	} else if !exist {
		err = utils.ERROR_USER_NOT_FOUND
		return
	}

	err = json.Unmarshal([]byte(uInfo), &userInfo)
	if err != nil {
		return
	}
	return
}

func (u *userDao)CheckLogin(account string, password string) (err error) {
	conn := u.Pool.Get()
	defer conn.Close()

	userInfo, err := u.getUserByAccount(conn, account)
	if err != nil {
		return
	}

	if userInfo.Password == password {
		fmt.Println("登陆成功")
	} else {
		err = utils.ERROR_PASSWORD_WRONG
	}

	return
}

func (u *userDao)UserRegister(user message.User) (err error) {
	conn := u.Pool.Get()
	defer conn.Close()

	exist, _, err := u.isUserExist(conn, user.Account)
	if err != nil {
		return
	} else if exist {
		err = utils.ERROR_USER_EXIST
		return
	}

	userInfo, err := json.Marshal(user)
	if err != nil {
		return
	}

	_, err = conn.Do("hset", "users", user.Account, string(userInfo))
	if err != nil {
		fmt.Println("保存注册用户数据过程出错，err =", err)
		return
	}

	return
}