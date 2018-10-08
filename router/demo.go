package router

import (
  "web/context"
  "web/controller"

  "github.com/gin-gonic/gin"
)

func registerDemoRouter(g *gin.RouterGroup) {

  // 这是一个测试请求
  g.GET("", /*middleware.Jwt,*/ context.CreateCtx(controller.Get))

}
