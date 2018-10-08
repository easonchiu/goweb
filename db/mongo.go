package db

import (
  "errors"
  "net/http"
  "web/conf"
  "web/util"

  "gopkg.in/mgo.v2"
)

var (
  Session    *mgo.Session
  Mongo      *mgo.DialInfo
  Connecting = false
)

func ConnectMgoDB() {
  if conf.MongodbDisabled {
    return
  }

  mongo, err := mgo.ParseURL(conf.MongodbUrl)

  s, err := mgo.Dial(conf.MongodbUrl)

  if err != nil {
    panic(err)
  }

  s.SetSafe(&mgo.Safe{})

  util.Println("Connect database successed.")

  Session = s
  Mongo = mongo
  Connecting = true
}

// 克隆一个mongodb的session
// 使用完成后需要关闭session
//   e.g.  defer session.close()
func CloneMgoDB() (*mgo.Database, func(), error) {
  if conf.MongodbDisabled {
    return nil, nil, nil
  }

  if Connecting {
    session := Session.Clone()
    closeFn := func() {
      session.Close()
    }
    return session.DB(Mongo.Database), closeFn, nil
  }

  return nil, nil, errors.New(http.StatusText(http.StatusBadGateway))
}

// 关闭mongodb数据库
func CloseMgoDB() {
  if conf.MongodbDisabled {
    return
  }

  if Connecting {
    Session.Close()
    Connecting = false
    util.Println("Database is closed.")
  } else {
    panic(errors.New("Database is not connected."))
  }
}
