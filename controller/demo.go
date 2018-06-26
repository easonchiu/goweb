package controller

import (
  "web/service"

  "github.com/gin-gonic/gin"
)

func DemoControl(g *gin.Context) {
  ctx := CreateCtx(g)
  defer ctx.handleErrorIfPanic()

  foo := ctx.getQuery("foo")
  bar := ctx.getQuery("bar")

  // 存储数据
  err := service.InsertMockDataToDB(foo, bar)

  // 如果有报错就panic，在handleErrorIfPanic中会处理错误
  if err != nil {
    panic(err)
    return
  }

  // 返回成功数据
  ctx.Success(gin.H{
    "data": "success",
  })
}
