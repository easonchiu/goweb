package main

import (
  "flag"
  "web/db"
  "web/router"

  "github.com/gin-gonic/gin"
)

func init() {
  db.ConnectMgoDB()
  db.InitRedisPool()
}

func main() {

  // 退出时关闭数据库连接
  defer db.CloseMgoDB()
  defer db.RedisPool.Close()

  // 初始化gin框架
  // Default With the Logger and Recovery middleware already attached
  g := gin.Default() // gin.New()

  // 注册路由
  router.Register(g)

  // 获取端口号
  // e.g.  go run main.go --port=8080
  port := ""
  flag.StringVar(&port, "port", "8080", "port addr")
  flag.Parse()

  // 启动
  g.Run(":" + port)
}
