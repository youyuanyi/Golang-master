package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

// 声明一个全局的rdb变量
var rdb *redis.Client

// initial
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "1.15.56.246:6379",
		Password: "",
		DB:       0,
		PoolSize: 100, // 连接池大小
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil

}

// 设置String
func redisString() {
	ctx := context.Background()
	err := rdb.Set(ctx, "yjc", "123", 0).Err()
	if err != nil {
		panic(err)
	}
	result, err := rdb.Get(ctx, "yjc").Result()
	if err == redis.Nil {
		fmt.Println("key yjc does not exits")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Printf("result: %v\n", result)
	}

}

func redisHash() {
	ctx := context.Background()
	// user_1是hash key, username是field,yjc是value
	err := rdb.HSet(ctx, "user_1", "username", "yjc").Err()
	if err != nil {
		panic(err)
	}
	// 获取指定key的指定field
	username, err := rdb.HGet(ctx, "user_1", "username").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(username)
	// 获取指定key的所有field
	data, err := rdb.HGetAll(ctx, "user_1").Result()
	if err != nil {
		panic(err)
	}
	// data 是一个map类型，需要用range来遍历
	for field, val := range data {
		fmt.Println(field, ": ", val)
	}

	err = rdb.HSet(ctx, "user_1", "password", "123").Err()
	if err != nil {
		panic(err)
	}
	// 根据key返回所有字段名
	keys, err := rdb.Keys(ctx, "user_1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(keys)
}

func redisList() {
	ctx := context.Background()
	// 从列表左边插入一个数据
	rdb.LPush(ctx, "l1", "data1")
	err := rdb.LPush(ctx, "l1", 1, 2, 3, 4, 5).Err()
	if err != nil {
		panic(err)
	}
	// 返回list的一个范围内的数据
	vals, err := rdb.LRange(ctx, "l1", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(vals)
	// 从列表右侧删除数据
	val, err := rdb.RPop(ctx, "l1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)
	// 返回list的一个范围内的数据
	vals, err = rdb.LRange(ctx, "l1", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(vals)
}

// 设置set类型
func redisSet() {
	ctx := context.Background()
	// 添加元素
	err := rdb.SAdd(ctx, "s1", 100).Err()
	if err != nil {
		panic(err)
	}
	rdb.SAdd(ctx, "s1", 100, 200, 300, 400)
	// 获取集合元素个数
	size, err := rdb.SCard(ctx, "s1").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(size) // 4
	// 判断元素是否在集合中
	ok, _ := rdb.SIsMember(ctx, "s1", 100).Result()
	if ok {
		fmt.Println("ok")
	} else {
		fmt.Println("No")
	}

	// 获取集合中所有的元素
	es, _ := rdb.SMembers(ctx, "s1").Result()
	fmt.Println(es) // [100 200 300 400]

	// 删除set中的元素
	rdb.SRem(ctx, "s1", 100)
	es, _ = rdb.SMembers(ctx, "s1").Result()
	fmt.Println(es) // [200 300 400]
}

func redisZset() {
	key := "language_rank"

	ctx := context.Background()
	// ZADD
	rdb.ZAdd(ctx, key, &redis.Z{
		Score:  90,
		Member: "Golang",
	})
	rdb.ZAdd(ctx, key, &redis.Z{
		Score:  95,
		Member: "Java",
	})
	rdb.ZAdd(ctx, key, &redis.Z{
		Score:  91,
		Member: "Rust",
	})
	rdb.ZAdd(ctx, key, &redis.Z{
		Score:  94,
		Member: "Python",
	})
	rdb.ZAdd(ctx, key, &redis.Z{
		Score:  99,
		Member: "C/C++",
	})

	// IncrBy
	newScore, err := rdb.ZIncrBy(ctx, key, 10.0, "Golang").Result()
	if err != nil {
		fmt.Printf("zincrby failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	// 取数据
	ret, err := rdb.ZRange(ctx, key, 0, -1).Result()
	if err != nil {
		panic(err)
	}
	// 返回了一个slice
	for _, val := range ret {
		fmt.Println(val)
	}

}

// 连接redis集群
func initClinetCluster() (err error) {
	crdb := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs: []string{"1.15.56.246:6379", "1.15.56.246:6380", "1.15.56.246:6381"},
	})
	ctx := context.Background()
	_, err = crdb.Ping(ctx).Result()
	if err != nil {
		return err
	}
	return nil
}

func redisSubscribe() {
	ctx := context.Background()
	// 订阅一个channel
	sub := rdb.Subscribe(ctx, "channel1")
	//for ch := range sub.Channel() {
	//	fmt.Println(ch.Channel) // 频道名称
	//	fmt.Println(ch.Payload) // 内容
	//}
	// 方式二
	for {
		message, err := sub.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}
		fmt.Println(message.Channel)
		fmt.Println(message.Payload)
	}

	// 查询指定的channel有多少个订阅者
	chs, _ := rdb.PubSubNumSub(ctx, "channel1").Result()
	for ch, count := range chs {
		fmt.Println(ch)    // channel名字
		fmt.Println(count) // channel的订阅者数量
	}
	//取消订阅
	sub.Unsubscribe(ctx, "channel1")

}

// redis处理事务
func RedisWork() {
	// 开启一个TxPipelin事务
	pipe := rdb.TxPipeline()
	// 执行事务操作，可以通过pipe读写redis
	ctx := context.Background()
	incr := pipe.Incr(ctx, "tx_pipeline_counter")
	pipe.Expire(ctx, "tx_pipeline_counter", time.Hour)
	_, err := pipe.Exec(ctx)
	//以上操作等同于执行了
	// MULTI
	// INCR pipeline_counter
	// EXPIRE pipeline_counts 3600
	// EXEC
	// 查询结果
	fmt.Println(incr.Val(), err)
}

func redisWatch() {
	// redis乐观锁
	ctx := context.Background()
	// 定义一个callback，用于处理事务
	fn := func(tx *redis.Tx) error {
		// 先查询下当前watch监听的key的值
		v, err := tx.Get(ctx, "key").Int()
		if err != nil && err != redis.Nil {
			return err
		}
		//处理业务
		v++
		// 如果key的值没有改变，Pipelined函数才会调用成功
		tx.Pipelined(ctx, func(pipe redis.Pipeliner) error {
			// 在这里给key设置最新值
			pipe.Set(ctx, "key", v, 0)
			return nil
		})
		return err
	}
	rdb.Watch(ctx, fn, "key")
}
func main() {
	_ = initClient()
	redisWatch()
	//err := initClinetCluster()
	//if err != nil {
	//	fmt.Println("connect cluster success")
	//}
}
