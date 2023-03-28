package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

type RedisMutex struct {
	client *redis.Client
	key    string
	value  string
	expiry time.Duration
}

// NewRedisMutex 给key设置过期时间
func NewRedisMutex(client *redis.Client, key string, value string, expiry time.Duration) *RedisMutex {
	return &RedisMutex{
		client: client,
		key:    key,
		value:  value,
		expiry: expiry,
	}
}

func (m *RedisMutex) Lock(ctx context.Context) error {
	for {
		ok, err := m.client.SetNX(ctx, m.key, m.value, m.expiry).Result()
		if err != nil {
			return err
		}
		if ok {
			return nil
		}
		select {
		case <-ctx.Done():
			return ctx.Err()
		case <-time.After(time.Millisecond * 100):
			continue
		}
	}
}

func (m *RedisMutex) Unlock(ctx context.Context) error {
	script := `
		if redis.call("get", KEYS[1]) == ARGV[1] then
			return redis.call("del", KEYS[1])
		else
			return 0
		end
	`
	result, err := m.client.Eval(ctx, script, []string{m.key}, m.value).Result()
	if err != nil {
		return err
	}
	if result.(int64) == 0 {
		return fmt.Errorf("unlock failed: key=%s, value=%s", m.key, m.value)
	}
	return nil
}
func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "1qaz2wsx", // no password set
		DB:       0,          // use default DB
	})

	m := NewRedisMutex(client, "myLock", "myMutex", time.Second*5)

	ctx := context.Background()

	if err := m.Lock(ctx); err!= nil {
		fmt.Println("lock error:", err)
		return
	}
	defer m.Unlock(ctx)

	fmt.Println("lock acquired")

	time.Sleep(time.Second * 3)

	fmt.Println("lock released")
}
