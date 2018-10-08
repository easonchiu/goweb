package model

import (
  "gopkg.in/mgo.v2/bson"
)

// 文档名
const DemoCollection = "demos"

// 文档的字段结构
type DemoModel struct {
  Id  bson.ObjectId `bson:"_id,omitempty"` // id自动生成
  Foo string        `bson:"foo"`
}

// 实现接口
func (d DemoModel) Parse() bson.M {
  return bson.M{
    "id":  d.Id,
    "foo": d.Foo,
    "len": len(d.Foo),
  }
}
