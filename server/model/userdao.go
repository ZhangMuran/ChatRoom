package model

import "github.com/gomodule/redigo/redis"

type UserDao struct {
	pool *redis.Pool
}

func (this *UserDao)CheckLogin(account string, password string) {
	conn := this.pool.Get()
	defer conn.Close()
}