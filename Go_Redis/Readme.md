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



#### Set

#### Hash

#### Zset



### 扩展类型



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