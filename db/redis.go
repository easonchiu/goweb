package db

import (
  "time"
  "web/conf"

  "github.com/gomodule/redigo/redis"
)

var (
  RedisPool *redis.Pool
)

// 初始化redis连接池
func InitRedisPool() {
  if conf.RedisDisabled {
    return
  }

  if RedisPool == nil {
    RedisPool = &redis.Pool{
      MaxIdle:     3,
      IdleTimeout: 240 * time.Second,
      Dial: func() (redis.Conn, error) {
        return redis.Dial("tcp", conf.RedisdbUrl)
      },
    }
  }
}

func CloseRedisPool() {
  if RedisPool != nil {
    RedisPool.Close()
  }
}

func GetRedis() redis.Conn {
  if conf.RedisDisabled {
    return nil
  }
  return RedisPool.Get()
}