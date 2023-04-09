package redis

import (
	"context"

	"github.com/go-redis/redis/v8"
)

type Redis struct {
	Ip string
	Port int
	Password string
	Db int
}

type DisLockRedis struct {
	key       string             //锁名称
	ttl       int64              //锁超时时间
	isLocked  bool               //上锁成功标识
	cancelFun context.CancelFunc //用于取消自动续租协程

	redis *redis.Client
	debug bool
}


func NewRedisStd() *Redis  {
	rdb := Redis{Ip: "localhost",Port: 6379,Password: "1qaz2wsx",Db: 0}
	return &rdb
}
