package model

import (
  "gopkg.in/mgo.v2/bson"
)

// collection name
const DemoCollection = "demos"

// collection model
type DemoModel struct {
  Id  bson.ObjectId `json:"id" bson:"_id"`
  Foo string        `json:"foo"`
  Bar string        `json:"bar"`
}
