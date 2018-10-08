package controller

import (
  "web/context"

  "github.com/gin-gonic/gin"
)

func Get(ctx *context.New) {
  ctx.Success(gin.H{
    "data": "OK",
  })
}
