package db

import (
  "github.com/gomodule/redigo/redis"
  "time"
  "web/conf"
)

var (
  RedisPool *redis.Pool
)

// 初始化redis连接池
func InitRedisPool() {
  if RedisPool == nil {
    RedisPool = &redis.Pool{
      MaxIdle:     3,
      IdleTimeout: 240 * time.Second,
      Dial: func() (redis.Conn, error) {
        return redis.Dial("tcp", conf.RedisDBUrl)
      },
    }
  }
}
