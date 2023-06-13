package utils

import "errors"

var (
	ERROR_USER_EXIST     = errors.New("注册的用户名已经存在")
	ERROR_USER_NOT_FOUND = errors.New("没有找到对应的用户")
	ERROR_PASSWORD_WRONG = errors.New("用户登录密码错误")
)