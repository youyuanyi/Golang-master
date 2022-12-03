# Golang Redis

## Nosql

### 什么是Nosql

- 不仅仅是数据
- 没有固定的查询语言
- 键值对存储(redis)、列存储、文档存储(MongoDB)、图形数据库
- 最终一致性

## Redis基础

### Redis是什么

**Redis(Remote Dictionary Server)**

### 用途

- 数据库
- 缓存
- 消息中间件MQ

### Redis命令

#### 切换数据库

```bash
select [number]  // 切换到第number号是数据库,例如select 3
```

#### 查看当前数据库下所有的key

```bash
keys *
```

#### 清空当前数据库

```bash
flushdb
```

#### 清空所有数据库

```
flushall
```



### Redis是单线程

Redis是单线程模型指的是执行Redis命令的核心模块是单线程的，而不是整个Redis实例就一个线程

Redis是基于内存操作，其瓶颈是机器的内存和网络的带宽



### Redis是单线程的为什么还这么快且能支撑高并发

1. Redis是基于内存操作，内存的读写速度非常快
2. 单线程避免了多线程的频繁上下文切换、加锁、CPU消耗等问题
3. 使用了IO多路复用技术，可以处理并发的连接



## Redis数据类型

### 基础类型

#### String

```bash
127.0.0.1:6379[1]> select 0    # 选择0号数据库
OK
127.0.0.1:6379> clear
127.0.0.1:6379> set key1 v1   # 设置key为key1的value为v1
OK
127.0.0.1:6379> get key1      # 获取key1的value
"v1"
127.0.0.1:6379> keys *			# 获取当前数据库下所有的key
1) "key1"
127.0.0.1:6379> EXISTS key1    # 是否存在key1
(integer) 1
127.0.0.1:6379> EXISTS key2
(integer) 0
127.0.0.1:6379> APPEND key1 hello   # 往key1中添加hello字符串
(integer) 7
127.0.0.1:6379> get key1
"v1hello"
127.0.0.1:6379> STRLEN key1			# key为key1的value的长度
(integer) 7
127.0.0.1:6379> set key1 hello,yjc
OK
127.0.0.1:6379> get key1
"hello,yjc"
127.0.0.1:6379> GETRANGE key1 1 5   # 截取字符串的部分，左闭右闭
"ello,"
127.0.0.1:6379> GETRANGE key1 0 -1  # 获取所有的值
"hello,yjc"
127.0.0.1:6379> SETRANGE key1 1 xx  # 替换字符串指定位置开始的内容
(integer) 9
127.0.0.1:6379> get key1
"hxxlo,yjc"
127.0.0.1:6379> setex key3 30 hello  # 设置key3时指定过期时间为30s
OK
127.0.0.1:6379> ttl key3
(integer) 28
127.0.0.1:6379> ttl key3
(integer) 26
127.0.0.1:6379> ttl key3
(integer) 25
127.0.0.1:6379> ttl key3
(integer) -2
127.0.0.1:6379> set mykey 123
OK
127.0.0.1:6379> setnx mykey 456  # 如果key存在则不设置，常常用在分布式锁中
(integer) 0
127.0.0.1:6379> get mykey
"123"
127.0.0.1:6379> mset k1 v1 k2 v2 k3 v3  # 批量设置key
OK
127.0.0.1:6379> keys *
1) "k3"
2) "k1"
3) "k2"
127.0.0.1:6379> mget k1 k2 k3
1) "v1"
2) "v2"
3) "v3"

```



#### List

所有的List命令都以`L`或者`R`开头

```bash
######################################################################################
127.0.0.1:6379> LPUSH list one  # 将一个值或多个值插入列表的头部
(integer) 1
127.0.0.1:6379> LPUSH list two
(integer) 2
127.0.0.1:6379> LPUSH list three
(integer) 3
127.0.0.1:6379> LRANGE list 0 -1
1) "three"
2) "two"
3) "one"
127.0.0.1:6379> RPUSH list four
(integer) 4
127.0.0.1:6379> RPUSH list five
(integer) 5
127.0.0.1:6379> LRANGE list 0 -1
1) "three"
2) "two"
3) "one"
4) "four"
5) "five"
127.0.0.1:6379> LPOP list  # 弹出左侧第一个
"three"
127.0.0.1:6379> RPOP list  # 弹出右侧第一个
"five"
127.0.0.1:6379> LRANGE list 0 -1
1) "two"
2) "one"
3) "four"
127.0.0.1:6379> LINDEX list 2  # 获取某个List的某个Index的值
"four"
127.0.0.1:6379> LINDEX list 1
"one"
127.0.0.1:6379> LREM list 1 two  # 移除指定的值
(integer) 1
127.0.0.1:6379> LRANGE list 0 -1
1) "one"
2) "four"
127.0.0.1:6379> lset list 1 123  # 设置List中的某个Index的值
OK
127.0.0.1:6379> LRANGE list 0 -1
1) "one"
2) "123"
######################################################################################
127.0.0.1:6379> RPUSH mylist hello
(integer) 1
127.0.0.1:6379> RPUSH mylist world
(integer) 2
127.0.0.1:6379> LINSERT mylist before "world" "other"  # 在world前添加other
(integer) 3
127.0.0.1:6379> LRANGE mylist 0 -1
1) "hello"
2) "other"
3) "world"
127.0.0.1:6379> LINSERT mylist after "world" "new"   # 在world后添加new
(integer) 4
127.0.0.1:6379> LRANGE mylist 0 -1
1) "hello"
2) "other"
3) "world"
4) "new"

```



#### Set

set中的值不能重复，命令以`s`开头,是无序不重复集合

```bash
127.0.0.1:6379> SADD myset hello  # 往myset中添加元素
(integer) 1
127.0.0.1:6379> SADD myset yjc
(integer) 1
127.0.0.1:6379> SADD myset fd
(integer) 1
127.0.0.1:6379> SMEMBERS myset   # 查看myset所有元素
1) "yjc"
2) "fd"
3) "hello"
127.0.0.1:6379> SISMEMBER myset yjc  # 判断yjc是否在myset中存在
(integer) 1
127.0.0.1:6379> scard myset			# 获取myset集合中的元素个数
(integer) 3
127.0.0.1:6379> SREM myset hello	# 删除myset中的hello元素
(integer) 1
127.0.0.1:6379> SMEMBERS myset
1) "yjc"
2) "fd"
127.0.0.1:6379> SRANDMEMBER myset  # 随机选出一个元素
"yjc"
127.0.0.1:6379> SRANDMEMBER myset
"fd"
######################################################################################
127.0.0.1:6379> sadd myset hello
(integer) 1
127.0.0.1:6379> sadd myset world
(integer) 1
127.0.0.1:6379> sadd myset yjc
(integer) 1
127.0.0.1:6379> sadd myset2 set2
(integer) 1
127.0.0.1:6379> smove myset myset2 yjc  # 将myset中的yjc移动到myset2
(integer) 1
127.0.0.1:6379> SMEMBERS myset
1) "world"
2) "hello"
127.0.0.1:6379> SMEMBERS myset2
1) "yjc"
2) "set2"
###################################################################################
127.0.0.1:6379> sadd key1 a
(integer) 1
127.0.0.1:6379> sadd key1 b
(integer) 1
127.0.0.1:6379> sadd key1 c
(integer) 1
127.0.0.1:6379> sadd key2 b
(integer) 1
127.0.0.1:6379> sadd key2 c
(integer) 1
127.0.0.1:6379> SDIFF key1 key2   # 差集
1) "a"
127.0.0.1:6379> SINTER key1 key2  # 交集
1) "c"
2) "b"
127.0.0.1:6379> sadd key2 d
(integer) 1
127.0.0.1:6379> SUNION key1 key2  # 并集
1) "a"
2) "c"
3) "d"
4) "b"

```



#### Hash

key-{k/v,k/v...}，以`h`开头

```bash
127.0.0.1:6379> hset myhash field1 yjc
(integer) 1
127.0.0.1:6379> hget myhash field1
"yjc"
127.0.0.1:6379> hmset myhash field1 hello field2 world
OK
127.0.0.1:6379> hmget myhash field1 field2
1) "hello"
2) "world"
127.0.0.1:6379> hgetall myhash   # 获取hash中全部的k/v
1) "field1"
2) "hello"
3) "field2"
4) "world"
127.0.0.1:6379> hdel myhash field1  # 删除hash指定的key字段和value
(integer) 1
127.0.0.1:6379> hgetall myhash
1) "field2"
2) "world"
```



#### Zset

有序集合

```bash
127.0.0.1:6379> zadd myset 1 one
(integer) 1
127.0.0.1:6379> zadd myset 2 two 3 three
(integer) 2
127.0.0.1:6379> zrange myset 0 -1
1) "one"
2) "two"
3) "three"
127.0.0.1:6379> zadd salary 2500 xiaohong
(integer) 1
127.0.0.1:6379> zadd salary 500 zhangsan
(integer) 1
127.0.0.1:6379> zadd salary 5000 lis
(integer) 1
127.0.0.1:6379> ZRANGEBYSCORE salary -inf +inf   # 排序
1) "zhangsan"
2) "xiaohong"
3) "lis"
127.0.0.1:6379> ZRANGEBYSCORE salary -inf +inf  withscores
1) "zhangsan"
2) "500"
3) "xiaohong"
4) "2500"
5) "lis"
6) "5000"

```



### 特殊类型

#### geospatial

地理位置

朋友的定位，附近的人，打车距离计算

```bash
127.0.0.1:6379> geoadd china:city 116.40 39.90 beijing   # 添加地理位置(经度、纬度)
(integer) 1
127.0.0.1:6379> geoadd china:city 121.47 31.23 shanghai
(integer) 1
127.0.0.1:6379> geoadd china:city 106.50 29.53 chongqing
(integer) 1

```

```
georadius：以给定的经纬度为中心，找出某一半径内的元素
```



#### hyperloglog

基数统计：统计不重复的元素，可以接受误差

优点：占用的内存是固定的：12KB

```bash
127.0.0.1:6379> PFADD mykey a b c d e
(integer) 1
127.0.0.1:6379> PFCOUNT mykey
(integer) 5
127.0.0.1:6379> pfadd mykey2 f g h j k
(integer) 1
127.0.0.1:6379> pfcount mykey2
(integer) 5
127.0.0.1:6379> PFMERGE mykey3 mykey mykey2
OK
127.0.0.1:6379> PFCOUNT mykey3
(integer) 10

```



#### Bitmap

操作二进制位来进行记录，就只有0和1两个状态

统计用户信息：活跃/不活跃，登录/未登录等两个状态



## Redis基本事务操作

### Redis事务本质

一组命令的集合，一个事务中的所有命令都会被序列化，在执行过程中按照顺序执行

### Redis事务特点

- Redis的单条命令可以保证原子性，但是**事务不保证原子性**
- **Redis事务没有隔离级别的概念**，即所有的命令在事务中并没有被直接执行，只有发起执行命令的时候才会执行

### Redis事务操作

#### 开启事务

`Multi`

#### 命令入队



#### 执行事务

`exec`



#### 取消事务

`DISCARD`

#### 例子

```bash
127.0.0.1:6379> MULTI      # 开启事务
OK
#  命令入队
127.0.0.1:6379(TX)> set k1 v1
QUEUED
127.0.0.1:6379(TX)> set k2 v2
QUEUED
127.0.0.1:6379(TX)> get k2
QUEUED
127.0.0.1:6379(TX)> set k3 v3
QUEUED
# 执行事务
127.0.0.1:6379(TX)> exec
1) OK
2) OK
3) "v2"
4) OK
127.0.0.1:6379> MULTI
OK
127.0.0.1:6379(TX)> set k1 v1
QUEUED
127.0.0.1:6379(TX)> set k2 v2
QUEUED
127.0.0.1:6379(TX)> set k4 v4
QUEUED
127.0.0.1:6379(TX)> DISCARD   # 取消事务
OK
127.0.0.1:6379> get k4
(nil)

```



### Redis事务异常处理

#### 编译型异常

命令语法有问题，则Redis事务中的**所有命令都不执行**

```bash
127.0.0.1:6379> multi
OK
127.0.0.1:6379(TX)> set k1 v1
QUEUED
127.0.0.1:6379(TX)> set k2 v2
QUEUED
127.0.0.1:6379(TX)> set k3 v4
QUEUED
127.0.0.1:6379(TX)> getset k3   # 语法错误
(error) ERR wrong number of arguments for 'getset' command
127.0.0.1:6379(TX)> set k4 v4
QUEUED
127.0.0.1:6379(TX)> exec   # 检查出编译型异常，所有命令都不执行
(error) EXECABORT Transaction discarded because of previous errors.
127.0.0.1:6379> get k1
(nil)
127.0.0.1:6379> get k4
(nil)

```



#### 运行时异常

除数为0，字符串错误操作等运行时异常错误，那么**执行事务时，其他没有错误的命令仍然会执行**

```bash
127.0.0.1:6379> multi
OK
127.0.0.1:6379(TX)> set k1 "v1"  
QUEUED 
127.0.0.1:6379(TX)> incr k1   # 对字符串+1，是无法执行的
QUEUED
127.0.0.1:6379(TX)> set k2 v2
QUEUED
127.0.0.1:6379(TX)> set k3 v3
QUEUED
127.0.0.1:6379(TX)> get k3
QUEUED
127.0.0.1:6379(TX)> exec
1) OK
2) (error) ERR value is not an integer or out of range
3) OK
4) OK
5) "v3"
```



## Redis乐观锁

### 乐观锁

- 认为什么时候都不会出问题，所以不会上锁，在更新数据的时候去判断一下在此期间是否有人修改过数据
- 获取version
- 更新时比较version

### Redis监视测试

```bash
127.0.0.1:6379> set money 100
OK
127.0.0.1:6379> set out 0
OK
127.0.0.1:6379> WATCH money  # 监视money
OK
127.0.0.1:6379> multi   # 事务正常结束，数据期间没有发生变动
OK
127.0.0.1:6379(TX)> DECRBY money 20
QUEUED
127.0.0.1:6379(TX)> INCRBY out 20
QUEUED
127.0.0.1:6379(TX)> exec
1) (integer) 80
2) (integer) 20

```

### 测试多线程

```bash
### 线程1:
127.0.0.1:6379> watch money
OK
127.0.0.1:6379> multi
OK
127.0.0.1:6379(TX)> DECRBY money 10
QUEUED
127.0.0.1:6379(TX)> INCRBY out 10
QUEUED

#### 此时线程2突然执行
127.0.0.1:6379> get money
"80"
127.0.0.1:6379> set money 1000
OK

### 回到线程1,提交事务，此时watch检测到money已经被修改，事务取消
127.0.0.1:6379(TX)> exec
(nil)
## 线程1恢复事务正常
127.0.0.1:6379> unwatch  # 解锁，不过Redis不管事务是否执行成功，都会对watch对象解锁
OK
127.0.0.1:6379> watch money
OK
127.0.0.1:6379> DECRBY money 10
(integer) 990
127.0.0.1:6379> INCRBY out 10
(integer) 30
127.0.0.1:6379> INCRBY money 10
(integer) 1000
127.0.0.1:6379> DECRBY out 10
(integer) 20


```



## Redis.conf详解

### Redis对大小写不敏感

```bash
# Note on units: when memory size is needed, it is possible to specify
# it in the usual form of 1k 5GB 4M and so forth:
#
# 1k => 1000 bytes
# 1kb => 1024 bytes
# 1m => 1000000 bytes
# 1mb => 1024*1024 bytes
# 1g => 1000000000 bytes
# 1gb => 1024*1024*1024 bytes
#
# units are case insensitive so 1GB 1Gb 1gB are all the same.

```

### Redis可以include其他配置文件

```bash
# If instead you are interested in using includes to override configuration
# options, it is better to use include as the last line.
#
# include /path/to/local.conf
# include /path/to/other.conf

```

### Redis网络

```bash
# bind 127.0.0.1 ::1  默认绑定127.0.0.1，如果需要外部访问，则注释掉
port 6379
tcp-keepalive 300
protected-mode yes # 保护模式
```



### 通用

```bash
daemonize yes # 以守护进程进行
pidfile /var/run/redis/redis-server.pid  # 如果以后台方式运行，就需要制定pid文件
# Specify the server verbosity level.
# This can be one of:
# debug (a lot of information, useful for development/testing)
# verbose (many rarely useful info, but not a mess like the debug level)
# notice (moderately verbose, what you want in production probably)
# warning (only very important / critical messages are logged)
loglevel notice  # 日志级别
logfile /var/log/redis/redis-server.log  # 日志文件
databases 16  # 数据库数量

```

### 快照

```bash
appendfsync everysecsave 900 1  # 如果900s (15min) 内至少修改了1个key，则进行持久化操作
save 300 10  # 如果300s内，至少10个key进行了操作
save 60 10000  # 如果60s内，至少10000个key进行了修改
stop-writes-on-bgsave-error no  # 持久化如果出错，是否继续工作
rdbcompression yes # 是否压缩rdb文件，需要消耗一些CPU资源
rdbchecksum yes # 保存rdb文件时，校验rdb文件，如果出错，会自动去修复
dbfilename dump.rdb  # rdb持久化后保存的文件
dir /var/lib/redis  # rdb持久化后保存文件路径
appendonly no # 默认不开启AOF模式，默认使用RDB
appendfilename "appendonly.aof"  # 持久化的AOF文件
appendfsync everysec  # 每秒执行一次sync，可能会丢失这一秒的数据
# appendfsync always  # 每次修改，都会同步
```



### REPLICATION-复制



### 设置密码

```bash
# requirepass 123456
```



### 限制CLIENTS

```bash
# maxclients 10000

```



### redis内存

```bash
# maxmemory <bytes>  # redis配置最大的内存容量
# maxmemory-policy noeviction  # 内存达到上限之后的处理策略
1.volatile-lru：只对设置了过期时间的key进行LRU （默认）
2.allkeys-lru：删除lru算法的key
3.volatile-random：随机删除即将过期的key
4.allkeys-random：随机删除
5.volatile-ttl：删除即将过期的
6.noeviction：永不过期，返回错误
```





## Redis数据持久化

### RDB

**根据配置的规则定时将内存中的数据持久化到硬盘上**

RDB的持久化是通过**快照**的方式完成的。**当符合某种规则时，会将内存中的数据全量生成一份副本存储到硬盘上。**

RDB执行快照的几种情况：

- 根据配置规则进行自动快照
- 用户执行`SAVE`，`BGSAVE`命令
- 执行`FLUSHALL`命令
- 执行复制(replication)时

### AOF

在每次**执行写命令之后将命令记录下来**，保存到AOF文件中。

AOF是为了弥补RDB会发生数据不一致性的问题，所以采用日志的形式来记录每个写操作，并保存到文件中

Redis重启时会根据AOF文件中的内容将写指令从前到后执行一次



## Redis发布订阅

