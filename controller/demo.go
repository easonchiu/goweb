package controller

import (
  "web/context"
  "web/model"
  "web/service"

  "github.com/gin-gonic/gin"
  "gopkg.in/mgo.v2/bson"
)

func Get(ctx *context.New) {

  foo := ctx.GetQueryDefault("foo", "default foo")

  // 存储数据
  err := service.Insert(ctx, foo)

  // 如果有报错
  if err != nil {
    ctx.Error(err)
    return
  }

  // 返回成功数据
  ctx.Success(gin.H{
    "data": "success",
  })
}

func Post(ctx *context.New) {
  foo, _ := ctx.GetRaw("foo")

  res := model.DemoModel{
    Id:  bson.NewObjectId(),
    Foo: foo,
  }

  ctx.Success(gin.H{
    "data": res.GetMap(),
  })
}

func Put(ctx *context.New) {

  // 更新数据
  err := service.Update(ctx, "err id", bson.M{
    "a": 1,
    "b": 2,
  })

  // 如果有报错
  if err != nil {
    ctx.Error(err)
    return
  }

  // 返回成功数据
  ctx.Success(gin.H{
    "data": "success",
  })
}
