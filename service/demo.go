package service

import (
  "errors"
  "web/context"
  "web/errgo"
  "web/model"
  "web/util"

  "gopkg.in/mgo.v2/bson"
)

// demo server
func Insert(ctx *context.New, data model.DemoModel) error {
  // 存
  err := ctx.MgoDB.C(model.DemoCollection).Insert(data)

  // 返回结果
  if err != nil {
    return errors.New(errgo.ErrServerError)
  }
  return nil
}

func Update(ctx *context.New, id string, data bson.M) error {
  // 限制更新的字段及类型
  util.Only(
    data,
    util.Keys{
      "foo": util.TypeString,
    },
  )

  // 验证id
  ctx.Errgo.ErrorIfStringNotObjectId(id, errgo.ErrIdError)

  if err := ctx.Errgo.PopError(); err != nil {
    return err
  }

  // 更新
  err := ctx.MgoDB.C(model.DemoCollection).UpdateId(bson.ObjectIdHex(id), bson.M{
    "$set": data,
  })

  // 返回结果
  if err != nil {
    return errors.New(errgo.ErrServerError)
  }

  return nil
}
