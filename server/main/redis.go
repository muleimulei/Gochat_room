package main

import (
	"time"

	"github.com/gomodule/redigo/redis"
)

//定义一个全局的pool

var Pool *redis.Pool

func initPool(address, pwd string, maxIdle, maxActive int, idleTimeout time.Duration) {
	Pool = &redis.Pool{
		MaxIdle:     maxIdle,     //最大空闲连接数
		MaxActive:   maxActive,   //表示和数据库的最大连接数,0 没有限制
		IdleTimeout: idleTimeout, //最大空闲时间
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", address,
				redis.DialPassword(pwd))
		},
	}
}
