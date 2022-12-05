package goredis

import (
	"context"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

var (
	rdb *redis.Client
	ctx = context.Background()
	mutex sync.Mutex
)


//InitRedis 初始化Redis连接
func InitRedis() *redis.Client{
	r := NewRedisStd()
	val := []string{r.Ip, strconv.Itoa(r.Port)}
	rdb = redis.NewClient(&redis.Options{
		Addr:   strings.Join(val , ":"),
		Password: r.Password, // no password set
		DB:       0,  // use default DB
	})
	return rdb
}

func Lock(key string) bool {
	mutex.Lock()
	defer mutex.Unlock()
	result, err := rdb.SetNX(ctx, key, 1, 10*time.Second).Result()
	if err != nil {
		log.Println(err.Error())
	}
	return result
}
func UnLock(key string) int64 {
	nums, err := rdb.Del(ctx, key).Result()
	if err != nil {
		log.Println(err.Error())
		return 0
	}
	return nums
}
