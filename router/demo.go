package router

import (
  "web/controller"

  "github.com/gin-gonic/gin"
)

func registerDemoRouter(g *gin.RouterGroup) {

  g.GET("", /*middleware.Jwt,*/ controller.DemoControl)

}
