package main

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

// 声明一个全局的rdb变量
var rdb2 *redis.Client

// initial
func initClient2() (err error) {
	rdb2 = redis.NewClient(&redis.Options{
		Addr:     "1.15.56.246:6379",
		Password: "",
		DB:       0,
		PoolSize: 100, // 连接池大小
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb2.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil

}

func main() {
	ctx := context.Background()
	_ = initClient2()
	rdb2.Publish(ctx, "channel1", "hello")
}
