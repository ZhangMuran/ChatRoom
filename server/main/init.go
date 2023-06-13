package main

import (
	"chatroom/server/model"
	"time"

	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool

func initRedisPool(address string, maxIdle, maxActive int, idleTimeout time.Duration) {
	pool = &redis.Pool {
		MaxIdle: maxIdle, // 最大空闲连接数
		MaxActive: maxActive, // 和数据库的最多连接数，0表示不限制
		IdleTimeout: idleTimeout, // 最大空闲时间
		Dial: func()(redis.Conn, error) { //连接数据库方法
			return redis.Dial("tcp", address)
		},
	}
}

func initUserDao(pool *redis.Pool) {
	userDao := model.GetUserDao()
	userDao.Pool = pool
}