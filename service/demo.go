package service

import (
  `web/db`
  `web/model`

  `gopkg.in/mgo.v2/bson`
)

// demo server
func InsertMockDataToDB(foo string, bar string) (error) {

  // 克隆一个db
  db, close, err := db.CloneDB()

  // 如果没克隆成，报错，否则在函数退出时关闭他
  if err != nil {
    return err
  } else {
    defer close()
  }

  // 创建数据
  data := model.DemoModel{
    Id:  bson.NewObjectId(),
    Foo: foo,
    Bar: bar,
  }

  // 返回
  return db.C(model.DemoCollection).Insert(data)
}
