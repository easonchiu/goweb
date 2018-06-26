package main

import (
  `flag`
  `web/db`
  `web/middleware`
  `web/router`

  `github.com/gin-gonic/gin`
)

func init() {
  db.ConnectDB()
}

func main() {

  // close db before unmount
  defer db.CloseDB()

  // initialization
  // Default With the Logger and Recovery middleware already attached
  g := gin.Default() // gin.New()

  // register middleware
  middleware.Register(g)

  // register router
  router.Register(g)

  // get port args
  // e.g.  go run main.go --port=8080
  port := ""
  flag.StringVar(&port, "port", "8080", "port addr")
  flag.Parse()

  // start
  g.Run(":" + port)

  // 启动之后浏览器访问
  // http://localhost:8080/demo?foo=1&bar=2
  // 就可以存一条数据至mongodb

}
